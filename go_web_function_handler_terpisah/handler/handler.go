package handler

import (
	"log"
	"strconv"
	"fmt"
	"net/http"
)

//funtion baru yg membutuhkan 2 paramater yaitu 
//1. w http.ResponseWriter
//2. pointer nya yaitu r*http.Request
func HelloHandler(w http.ResponseWriter, r*http.Request)  {
	//body dr fnctionnya 
	w.Write([]byte("Lagi Belajar Go Lang guys"))
}

func MarioHandler(w http.ResponseWriter, r*http.Request)  {
	//body dr fnctionnya 
	w.Write([]byte("Mario Balotelli"))
}

func HomeHandler(w http.ResponseWriter, r*http.Request)  {
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
func ProductHandler(w http.ResponseWriter, r*http.Request)  {
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