package main

// call this via HTTP GET with a URL like:
//     http://localhost:9999/get-spanish-greeting?name=Bob
//     http://localhost:9999/get-spanish-farewell?name=Bob

import (
	"fmt"
	"net/http"
)

func spanishGreetingHandler(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["name"]
	if ok {
		name := keys[0]
		translation := fmt.Sprintf("¡Hola, %s!", name)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, translation)
	} else {
		http.Error(w, "Missing required 'name' parameter.", http.StatusBadRequest)
	}
}

func spanishFarewellHandler(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["name"]
	if ok {
		name := keys[0]
		translation := fmt.Sprintf("¡Adios, %s!", name)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, translation)
	} else {
		http.Error(w, "Missing required 'name' parameter.", http.StatusBadRequest)
	}
}

func main() {
	http.HandleFunc("/get-spanish-greeting", spanishGreetingHandler)
	http.HandleFunc("/get-spanish-farewell", spanishFarewellHandler)
	http.ListenAndServe(":9999", nil)
}
