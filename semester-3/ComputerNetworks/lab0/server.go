package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/SlyMarbo/rss"
)

func getCurrentDomain(req *http.Request) string {
	headers := []string{
		"X-Forwarded-Host",
		"X-Original-Host",
		"Host",
	}

	for _, header := range headers {
		if host := req.Header.Get(header); host != "" {
			return strings.Split(host, ":")[0]
		}
	}

	return strings.Split(req.Host, ":")[0]
}

func getFullURL(req *http.Request) string {
	scheme := "http"
	if req.TLS != nil {
		scheme = "https"
	}

	if proto := req.Header.Get("X-Forwarded-Proto"); proto != "" {
		scheme = proto
	} else if req.Header.Get("X-Forwarded-Scheme") != "" {
		scheme = req.Header.Get("X-Forwarded-Scheme")
	}

	host := req.Host
	if forwardedHost := req.Header.Get("X-Forwarded-Host"); forwardedHost != "" {
		host = forwardedHost
	}
	fullURL := fmt.Sprintf("%s://%s%s", scheme, host, req.URL.RequestURI())

	return fullURL
}

func HomeRouterHandler(w http.ResponseWriter, req *http.Request) {
	currentDomain := getCurrentDomain(req)
	fullURL := getFullURL(req)
	fmt.Println("=== ЗАГОЛОВКИ ЗАПРОСА ===")
	fmt.Printf("Host: %s\n", req.Host)
	fmt.Printf("URL: %s\n", req.URL.String())
	fmt.Printf("Current Domain: %s\n", currentDomain)
	fmt.Printf("Full URL: %s\n", fullURL)
	fmt.Println("==========================")

	q := req.URL.Query()
	data := q.Get("data")
	fmt.Println(data)

	feed, err := rss.Fetch("https://tvsamara.ru/rss/")
	if err != nil {
		log.Fatalf("Error fetching RSS: %v", err)
	}

	html := "<html><body>"
	html += ` <div style='background: #f0f0f0; padding: 10px; margin-bottom: 20px;'>
            <p><strong>Текущий домен:</strong> ` + currentDomain + `</p>
            <p><strong>Полный URL:</strong> ` + fullURL + `</p>
        </div>`
	html += fmt.Sprintf("<h1>%s</h1>", feed.Title)
	html += fmt.Sprintf("<p>%s</p>", feed.Description)
	html += fmt.Sprintf("<p>Количество новостей: %d</p>", len(feed.Items))
	html += "<hr>"

	for i, item := range feed.Items {
		html += fmt.Sprintf("<div style='margin-bottom: 20px; padding: 10px; border: 1px solid #ccc;'>")
		html += fmt.Sprintf("<h3>Новость %d</h3>", i+1)
		html += fmt.Sprintf("<p><strong>Заголовок:</strong> %s</p>", item.Title)
		html += fmt.Sprintf("<p><strong>Ссылка:</strong> <a href='%s'>%s</a></p>", item.Link, item.Link)
		html += "</div>"
	}

	html += "<br><a href='http://localhost:3000/another_page'>Перейти на другую страницу</a>"
	html += "</body></html>"

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(html))
}

func AnotherRouterHandler(w http.ResponseWriter, req *http.Request) {
	html := `<html><body>
		<h1>Другая страница</h1>
		<a href='http://localhost:3000/'>Вернуться назад</a>
		</body></html>`

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(html))
}

func main() {
	http.HandleFunc("/", HomeRouterHandler)
	http.HandleFunc("/another_page", AnotherRouterHandler)

	err := http.ListenAndServe(":4040", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
