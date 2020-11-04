package main

import (
	"fmt"
	"strconv"
	"log"
	"net/http"
)

func main()  {
	mux:=http.NewServeMux()

	

	//mux.HandleFunc
	// func(pattern string, handler func(ResponseWriter, *Request))
	//pattern string isi /hello
	mux.HandleFunc("/",homeHandler)
	mux.HandleFunc("/hello",helloHandler)
	mux.HandleFunc("/mario",marioHandler)
	mux.HandleFunc("/product",productHandler)
	
	
	log.Println("Starting web on port 8080")

	//jalanin servernya 
	//LListenAndServe func(addr string, handler Handler) error
	//parameter pertamanya addr string
	//keduanya si handle mux itu sendiri
	err:=http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
//funtion baru yg membutuhkan 2 paramater yaitu 
//1. w http.ResponseWriter
//2. pointer nya yaitu r*http.Request
func helloHandler(w http.ResponseWriter, r*http.Request)  {
	//body dr fnctionnya 
	w.Write([]byte("Lagi Belajar Go Lang guys"))
}

func marioHandler(w http.ResponseWriter, r*http.Request)  {
	//body dr fnctionnya 
	w.Write([]byte("Mario Balotelli"))
}

func homeHandler(w http.ResponseWriter, r*http.Request)  {
	//body dr fnctionnya 
	
	//ngehandle 4040
	log.Println(r.URL.Path)
	if r.URL.Path !="/"{
		http.NotFound(w ,r)
		return
	}
	w.Write([]byte("Home Page"))
}

//function product handler
func productHandler(w http.ResponseWriter, r*http.Request)  {
	//body dr fnctionnya 
	//menangkap query berdasrkan id 
	 id := r.URL.Query().Get("id")

	 //convery string ke integer pake nya Atoi
	 //balikannya dia angka dan error
	 idNumber, err := strconv.Atoi(id)

	 if(err != nil || idNumber < 1 ){
		http.NotFound(w ,r)
		return
	 }
	 //w.Write([]byte("Product Page"))

	 
	 //cetaknya pake Fprinf(w writer, "string", object bebas)
	 fmt.Fprintf(w, "Product Page : %d", idNumber)
} 