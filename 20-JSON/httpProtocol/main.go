package main

import (
	"log"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Página raiz!"))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá mundo!"))
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Users!"))
}

func main() {

	http.HandleFunc("/", root)
	http.HandleFunc("/home", home)
	http.HandleFunc("/usuarios", usuarios)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
