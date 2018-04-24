// https://golang.org/doc/effective_go.html#web_server
// $ go run web-server.go -addr=:1718
package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":1718", "http service address") // Q=17, R=18

var tmpl = template.Must(template.New("qr").Parse(templateStr))

func main() {
	flag.Parse()
	http.Handle("/", http.HandlerFunc(QR))
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func QR(w http.ResponseWriter, req *http.Request) {
	err := tmpl.Execute(w, req.FormValue("s"))
	if err != nil {
		fmt.Println(err)
	}
}

const templateStr = `
<html>
<head>
<title>QR Link Generator</title>
</head>
<body>
{{if .}}
<img src="https://chart.apis.google.com/chart?chs=300x300&cht=qr&choe=UTF-8&chl={{.}}" />
<br>
{{.}}
<br>
<br>
{{end}}
<form action="/" name="f" method="GET">
	<input maxLength="1024" size="70" name="s" value="" title="Text to QR Encode">
	<input type="submit" value="Show QR" name="qr">
</form>
</body>
</html>
`
