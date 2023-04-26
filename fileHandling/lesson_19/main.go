package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	c := http.Client{}
	//readers sempre trabalham com buffers
	jsonVar := bytes.NewBuffer([]byte(`{"name": "wesley"}`))
	resp, err := c.Post("https://www.google.com", "application/json", jsonVar)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	//bypassing para stdout diretamente.
	io.CopyBuffer(os.Stdout, resp.Body, nil)
}
