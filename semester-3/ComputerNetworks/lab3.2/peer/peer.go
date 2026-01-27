package peer

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type Peer struct {
	PeerName    string
	IPadress    string
	IPParent    string
	ParentConn  *websocket.Conn
	ChildrenIPs []string
}

type RequestJSON struct {
	From      string `json:"from"`
	Msg       string `json:"msg"`
	NewParent string `json:"newParent"`
}

type ResponseParentName struct {
	ParentIP   string
	ParentName string
}

type TotalNameJSONResp struct {
	Receiver string
	Names    []string
}

func CreatePeer(name, ipAdress, ParentIP string) *Peer {
	return &Peer{
		PeerName:    name,
		IPadress:    ipAdress,
		IPParent:    ParentIP,
		ChildrenIPs: make([]string, 0),
	}
}

func StartPeer(peer *Peer) {
	http.HandleFunc("/", peer.handleClient)
	go func() {
		log.Println("Starting WebSocket server...")
		addr := fmt.Sprintf(":%s", peer.IPadress)
		if err := http.ListenAndServe(addr, nil); err != nil {
			log.Fatalf("ListenAndServe error: %v", err)
		}
	}()
	if peer.IPParent != "" {
		log.Println("send hello")
		go peer.sendHelloToParent()
	}
}

func (peer *Peer) sendHelloToParent() {
	ip := fmt.Sprintf("ws://localhost:%s/ws", peer.IPParent)
	conn, _, err := websocket.DefaultDialer.Dial(ip, nil)
	if err != nil {
		return
	}
	jsonReq := &RequestJSON{peer.PeerName, "hello", peer.IPadress}
	req, _ := json.Marshal(&jsonReq)
	if err := conn.WriteMessage(websocket.TextMessage, req); err != nil {
		log.Println("Write error:", err)
		return
	}
	_, reply, _ := conn.ReadMessage()
	if string(reply) == "ok" {
		log.Println("Succesfully found parent")
	}
	peer.ParentConn = conn
}

func (peer *Peer) handleClient(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil)
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		var jsonReq RequestJSON
		err = json.Unmarshal(msg, &jsonReq)
		if err != nil {
			log.Println("Fail with json Unmarshal")
		}
		log.Printf("Get message from %s: %s\n", jsonReq.From, jsonReq.Msg)
		switch jsonReq.Msg {
		case "getName":
			jsonResp := &ResponseParentName{peer.IPadress, peer.PeerName}
			resp, _ := json.Marshal(jsonResp)
			conn.WriteMessage(websocket.TextMessage, resp)
		case "getTotalName":
			jsonResp := &TotalNameJSONResp{peer.PeerName, make([]string, 0)}
			jsonResp.Names = append(jsonResp.Names, peer.PeerName)
			if peer.IPParent != "" {
				coonection := peer.ParentConn
				jsonReqForParent := &RequestJSON{peer.PeerName, "getTotalName", ""}
				req, _ := json.Marshal(jsonReqForParent)
				coonection.WriteMessage(websocket.TextMessage, req)
				_, reply, _ := coonection.ReadMessage()
				var jsonRespFromParent TotalNameJSONResp
				json.Unmarshal(reply, &jsonRespFromParent)
				jsonResp.Names = append(jsonResp.Names, jsonRespFromParent.Names...)
			}
			resp, _ := json.Marshal(jsonResp)
			conn.WriteMessage(websocket.TextMessage, resp)
		case "NewParent":
			if peer.ParentConn != nil {
				peer.ParentConn.Close()
				peer.ParentConn = nil
			}
			peer.IPParent = jsonReq.NewParent
			if peer.IPParent != "" {
				go peer.sendHelloToParent()
			}
			conn.WriteMessage(websocket.TextMessage, []byte("ok"))
			log.Printf("New parent %s\n", jsonReq.NewParent)
		case "hello":
			peer.ChildrenIPs = append(peer.ChildrenIPs, jsonReq.NewParent)
			conn.WriteMessage(websocket.TextMessage, []byte("ok"))
			log.Printf("Save children ip: %s", jsonReq.NewParent)
		}
	}
}

func (peer *Peer) GetNameFromParent() string {
	conn := peer.ParentConn
	req := &RequestJSON{peer.IPadress, "getName", ""}
	msg, _ := json.Marshal(req)
	if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
		log.Println("Write error:", err)
	}
	_, reply, _ := conn.ReadMessage()
	var jsonResp ResponseParentName
	err := json.Unmarshal(reply, &jsonResp)
	if err != nil {
		log.Println("Fail with json Unmarshal")
	}
	return jsonResp.ParentName
}

func (peer *Peer) GetTotalName() []string {
	conn := peer.ParentConn
	req := &RequestJSON{peer.PeerName, "getTotalName", ""}
	msg, _ := json.Marshal(req)
	if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
		log.Println("Write error:", err)
	}
	_, reply, _ := conn.ReadMessage()
	var jsonResp TotalNameJSONResp
	err := json.Unmarshal(reply, &jsonResp)
	if err != nil {
		log.Println("Fail with json Unmarshal")
	}
	return jsonResp.Names
}

func (peer *Peer) Exit() {
	var wg sync.WaitGroup
	for i, ip := range peer.ChildrenIPs {
		wg.Add(1)
		go func(i int, ip string) {
			defer wg.Done()
			ipAddr := fmt.Sprintf("ws://localhost:%s/ws", ip)
			conn, _, err := websocket.DefaultDialer.Dial(ipAddr, nil)
			if err != nil {
				return
			}
			req := &RequestJSON{peer.PeerName, "NewParent", peer.IPParent}
			msg, _ := json.Marshal(req)
			if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				log.Println("Write error:", err)
				return
			}
			conn.SetReadDeadline(time.Now().Add(3 * time.Second))
			_, reply, err := conn.ReadMessage()
			if err != nil {
				log.Println("Read error:", err)
				return
			}
			if string(reply) == "ok" {
				log.Printf("Succesfully changed parent for child #%d\n", i)
			}
		}(i, ip)
	}
	wg.Wait()
	log.Println("Succesfully changed parent for all")
}
