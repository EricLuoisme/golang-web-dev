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

	// no matter the size of your slice
	sages := []string{"Gandhi", "MLK", "Buddha", "Jesus", "Muhammad", "One More", "Two More"}

	// html -> {{range.}} can hold all data size you input
	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
