package main

import (
	"fmt"
	"net/http"
)

func Server() {
	http.HandleFunc("/", ResponseHome)
	fmt.Println("http://localhost:3001")
	fmt.Println(http.ListenAndServe(":3001", nil))

}
func ResponseHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong"))
}
