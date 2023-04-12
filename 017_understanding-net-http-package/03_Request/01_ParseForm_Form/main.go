package main

import (
	"html/template"
	"log"
	"net/http"
)

type customHttpHandler int

// ServeHTTP implement the Handler interface, thus, it can be input into the http.ListenAndServe
func (m customHttpHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	// here calling the temple -> which basically return an html return to the request
	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d customHttpHandler
	http.ListenAndServe(":8080", d)
}
