package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := makeRouter()

	http.ListenAndServe(":8080", router)
}

func makeRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", handler).Methods("GET")

	//file server
	staticFileDir := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDir))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

	return r
}


func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
