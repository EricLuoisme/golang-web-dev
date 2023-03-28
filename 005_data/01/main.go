package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	// here pass the data into the {{.}} into html file
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", "data")
	if err != nil {
		log.Fatalln(err)
	}
}
