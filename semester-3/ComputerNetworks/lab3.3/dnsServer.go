package main

import (
	"log"
	"net"
	"strings"

	"github.com/miekg/dns"
)

type DNSHandler struct{}

func (h *DNSHandler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	msg := new(dns.Msg)
	msg.SetReply(r)
	msg.Authoritative = true

	for _, question := range r.Question {
		domain := strings.ToLower(question.Name)
		if strings.Contains(domain, "veterok.everywhere") {
			switch question.Qtype {
			case dns.TypeA:
				rr := &dns.A{
					Hdr: dns.RR_Header{
						Name:   question.Name,
						Rrtype: dns.TypeA,
						Class:  dns.ClassINET,
						Ttl:    60,
					},
					A: net.ParseIP("185.102.139.168"), // Использую ваш IP
				}
				msg.Answer = append(msg.Answer, rr)
				log.Printf("Resolved %s to 185.102.139.168", question.Name)
			}
		}
	}

	w.WriteMsg(msg)
}

func main() {
	handler := &DNSHandler{}

	serverUDP := &dns.Server{
		Addr:    "127.0.0.1:53",
		Net:     "udp",
		Handler: handler,
	}

	log.Printf("Starting DNS servers...")
	log.Printf("UDP: 127.0.0.1:53")
	log.Printf("veterok.everywhere -> 185.102.139.168")

	if err := serverUDP.ListenAndServe(); err != nil {
		log.Printf("UDP server failed: %v", err)
	}
}
