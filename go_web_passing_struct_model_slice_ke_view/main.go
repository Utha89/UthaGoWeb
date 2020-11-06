package main

import (
	"go_web_passing_struct_model_slice_ke_view/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	//mux.HandleFunc
	// func(pattern string, handler func(ResponseWriter, *Request))
	//pattern string isi /hello
	mux.HandleFunc("/", handler.HomeHandler)
	mux.HandleFunc("/hello", handler.HelloHandler)
	mux.HandleFunc("/mario", handler.MarioHandler)
	mux.HandleFunc("/product", handler.ProductHandler)

	log.Println("Starting web on port 8080")

	//jalanin servernya
	//LListenAndServe func(addr string, handler Handler) error
	//parameter pertamanya addr string
	//keduanya si handle mux itu sendiri
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
