package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080", nil)
	if err != nil {
		log.Println("Erro ao fazer a request pelo client")
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Erro ao obter a response pelo client")
		panic(err)
	}
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
	log.Println("Request concluída com sucesso no lado do client")
}
