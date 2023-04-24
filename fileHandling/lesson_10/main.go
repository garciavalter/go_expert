package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)
	mux.Handle("/blog", blog{tittle: "My Blog"})
	http.ListenAndServe(":8080", mux)

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Wesley"))
	})
	http.ListenAndServe(":8081", mux2)

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

type blog struct {
	tittle string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.tittle))
}
