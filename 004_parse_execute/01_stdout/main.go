package main

import (
	"log"
	"os"
	"text/template" // make sure use text template -> parsing the file into text
)

func main() {
	// tpl would be a container for holding all files input
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// execute would exactly read the file data to writer (e.g. os print or other buffer writer)
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

// Do not use the above code in production
// We will learn about efficiency improvements soon!
