package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Handlers ...
func handlers() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/ping", handlerPing)
	r.HandleFunc("/arith", handlerArith)
	return r
}

func handlerArith(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for _, v := range r.Form {
		fmt.Println(v[0])
	}
}

func handlerPing(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}
