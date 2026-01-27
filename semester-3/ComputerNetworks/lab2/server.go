package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	log "github.com/mgutz/logxi/v1"
	"golang.org/x/net/html"
)

var ITEMS []*Item

func getAttr(node *html.Node, key string) string {
	for _, attr := range node.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}

func isElem(node *html.Node, tag string) bool {
	return node != nil && node.Type == html.ElementNode && node.Data == tag
}

type Item struct {
	Date, USD, EUR, CNY string
}

func getText(node *html.Node) string {
	if node == nil {
		return ""
	}
	if node.Type == html.TextNode {
		return strings.TrimSpace(node.Data)
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if txt := getText(c); txt != "" {
			return txt
		}
	}
	return ""
}

func readItem(tr *html.Node) *Item {
	var cols []string
	for c := tr.FirstChild; c != nil; c = c.NextSibling {
		if isElem(c, "td") {
			cols = append(cols, getText(c))
		}
	}
	if len(cols) < 4 {
		return nil
	}
	return &Item{
		Date: cols[0],
		USD:  cols[1],
		EUR:  cols[2],
		CNY:  cols[3],
	}
}

func search(node *html.Node) []*Item {
	if isElem(node, "tbody") {
		var items []*Item
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			if isElem(c, "tr") {
				if item := readItem(c); item != nil {
					items = append(items, item)
				}
			}
		}
		ITEMS = items
		return items
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if items := search(c); items != nil {
			return items
		}
	}
	return nil
}

func downloadRates() []*Item {
	log.Info("sending request to sberometer.ru")
	if response, err := http.Get("https://www.sberometer.ru/cbr/"); err != nil {
		log.Error("request failed", "error", err)
	} else {
		defer response.Body.Close()
		status := response.StatusCode
		log.Info("got response", "status", status)
		if status == http.StatusOK {
			if doc, err := html.Parse(response.Body); err != nil {
				log.Error("invalid HTML", "error", err)
			} else {
				log.Info("HTML parsed successfully")
				return search(doc)
			}
		}
	}
	return nil
}

const INDEX_HTML = `
<!doctype html>
<html lang="ru">
<head>
    <meta charset="utf-8">
    <title>Курсы валют ЦБ РФ</title>
    <style>
        body { font-family: sans-serif; margin: 20px; }
        table { border-collapse: collapse; width: 60%; }
        th, td { border: 1px solid #ccc; padding: 6px 12px; text-align: center; }
        th { background-color: #f0f0f0; }
        .info-block { background-color: #f8f9fa; padding: 15px; margin-bottom: 20px; border-left: 4px solid #007bff; }
        .info-block h3 { margin-top: 0; color: #333; }
        .info-block p { margin: 5px 0; }
    </style>
</head>
<body>
    <h2>Курсы валют ЦБ РФ</h2>
    {{if .Items}}
    <table>
        <tr>
            <th>Дата</th>
            <th>USD</th>
            <th>EUR</th>
            <th>CNY</th>
        </tr>
        {{range .Items}}
        <tr>
            <td>{{.Date}}</td>
            <td>{{.USD}}</td>
            <td>{{.EUR}}</td>
            <td>{{.CNY}}</td>
        </tr>
        {{end}}
    </table>
    {{else}}
        <p>Не удалось загрузить данные!</p>
    {{end}}
</body>
</html>
`

const USD_HTML = `
<!doctype html>
<html lang="ru">
<head>
    <meta charset="utf-8">
    <title>Курсы валют ЦБ РФ</title>
    <style>
        body { font-family: sans-serif; margin: 20px; }
        table { border-collapse: collapse; width: 60%; }
        th, td { border: 1px solid #ccc; padding: 6px 12px; text-align: center; }
        th { background-color: #f0f0f0; }
    </style>
</head>
<body>
    <h2>Курсы валют ЦБ РФ</h2>
    {{if .}}
    <table>
        <tr>
            <th>Дата</th>
            <th>USD</th>
        </tr>
        {{range .}}
        <tr>
            <td>{{.Date}}</td>
            <td>{{.USD}}</td>
        </tr>
        {{end}}
    </table>
    {{else}}
        <p>Не удалось загрузить данные!</p>
    {{end}}
</body>
</html>
`

const EUR_HTML = `
<!doctype html>
<html lang="ru">
<head>
    <meta charset="utf-8">
    <title>Курсы валют ЦБ РФ</title>
    <style>
        body { font-family: sans-serif; margin: 20px; }
        table { border-collapse: collapse; width: 60%; }
        th, td { border: 1px solid #ccc; padding: 6px 12px; text-align: center; }
        th { background-color: #f0f0f0; }
    </style>
</head>
<body>
    <h2>Курсы валют ЦБ РФ</h2>
    {{if .}}
    <table>
        <tr>
            <th>Дата</th>
            <th>EUR</th>
        </tr>
        {{range .}}
        <tr>
            <td>{{.Date}}</td>
            <td>{{.EUR}}</td>
        </tr>
        {{end}}
    </table>
    {{else}}
        <p>Не удалось загрузить данные!</p>
    {{end}}
</body>
</html>
`

const CNY_HTML = `
<!doctype html>
<html lang="ru">
<head>
    <meta charset="utf-8">
    <title>Курсы валют ЦБ РФ</title>
    <style>
        body { font-family: sans-serif; margin: 20px; }
        table { border-collapse: collapse; width: 60%; }
        th, td { border: 1px solid #ccc; padding: 6px 12px; text-align: center; }
        th { background-color: #f0f0f0; }
    </style>
</head>
<body>
    <h2>Курсы валют ЦБ РФ</h2>
    {{if .}}
    <table>
        <tr>
            <th>Дата</th>
            <th>CNY</th>
        </tr>
        {{range .}}
        <tr>
            <td>{{.Date}}</td>
            <td>{{.CNY}}</td>
        </tr>
        {{end}}
    </table>
    {{else}}
        <p>Не удалось загрузить данные!</p>
    {{end}}
</body>
</html>
`

var indexHtml = template.Must(template.New("index").Parse(INDEX_HTML))

func serveClient(response http.ResponseWriter, request *http.Request) {

	log.Info("=== ИНФОРМАЦИЯ О ЗАПРОСЕ ===")
	log.Info("Host", "value", request.Host)
	log.Info("URL", "value", request.URL.String())
	log.Info("Path", "value", request.URL.Path)
	log.Info("Method", "value", request.Method)
	log.Info("============================")

	path := request.URL.Path
	log.Info("got request", "Method", request.Method, "Path", path)
	if path != "/" && path != "/index.html" {
		response.WriteHeader(http.StatusNotFound)
		return
	}
	items := downloadRates()
	data := struct {
		Items   []*Item
		Domain  string
		FullURL string
		Path    string
	}{
		Items: items,
		Path:  request.URL.Path,
	}

	if err := indexHtml.Execute(response, data); err != nil {
		log.Error("HTML creation failed", "error", err)
	} else {
		log.Info("response sent successfully")
	}
}

func serveUSD(response http.ResponseWriter, request *http.Request) {
	var indexHtml2 = template.Must(template.New("index").Parse(USD_HTML))
	if err := indexHtml2.Execute(response, ITEMS); err != nil {
		log.Error("HTML creation failed", "error", err)
	} else {
		log.Info("response sent successfully")
	}
}

func serveEUR(response http.ResponseWriter, request *http.Request) {
	var indexHtml2 = template.Must(template.New("index").Parse(EUR_HTML))
	if err := indexHtml2.Execute(response, ITEMS); err != nil {
		log.Error("HTML creation failed", "error", err)
	} else {
		log.Info("response sent successfully")
	}
}

func serveCNY(response http.ResponseWriter, request *http.Request) {
	var indexHtml2 = template.Must(template.New("index").Parse(CNY_HTML))
	if err := indexHtml2.Execute(response, ITEMS); err != nil {
		log.Error("HTML creation failed", "error", err)
	} else {
		log.Info("response sent successfully")
	}
}

func main() {
	http.HandleFunc("/", serveClient)
	http.HandleFunc("/usd", serveUSD)
	http.HandleFunc("/eur", serveEUR)
	http.HandleFunc("/cny", serveCNY)
	log.Info("starting server on :4060")
	fmt.Println("listener failed", "error", http.ListenAndServe(":4060", nil))
}
