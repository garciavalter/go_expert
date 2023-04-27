package main

import (
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	port := ":8080"
	l, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	log.Printf("Server Starting at port: %s", port)
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	err = http.Serve(l, mux)
	if err != nil {
		panic(err)
	}

}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request Inicial")
	defer log.Println("Request finalizada")
	select {
	case <-time.After(5 * time.Second):
		log.Println("Request processada com sucesso")
		w.Write([]byte("Request processada com sucesso"))
	case <-ctx.Done():
		log.Println("Request cancelada pelo cliente")
		http.Error(w, "Request cancelada pelo cliente", http.StatusRequestTimeout)
	}

}
