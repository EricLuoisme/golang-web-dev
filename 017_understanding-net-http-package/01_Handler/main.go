package main

import (
	"fmt"
	"net/http"
)

type hotdog int

// any function implement ServeHTTP(ResponseWriter, *Request) would a kind of Http-Handler
func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Any code you want in this func")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
