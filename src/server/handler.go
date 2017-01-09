package server

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"

	"io/ioutil"

	"github.com/gorilla/mux"
)

// ArithURL for the subscriber service
const ArithURL = "http://192.168.50.5:8001/arithsubscriber"

// Handlers its all here...
func handlers() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/ping", handlerPing)
	r.HandleFunc("/arith", handlerArith)
	return r
}

// handles arithmetic requests
func handlerArith(w http.ResponseWriter, r *http.Request) {

	// Parses the raw query from the POST form received
	if err := r.ParseForm(); err != nil {
		log.Print("rest api: failed to parse form - ", err)
	}

	// Grab the task and wait for a result
	result := arithSender(r.PostFormValue("task"))

	// Read the result
	resultData, err := ioutil.ReadAll(result)
	if err != nil {
		log.Fatal("rest api: failed to read result - ", err)
	}
	defer result.Close()

	// Send it back to the producer
	w.Header().Set("X-Custom-Header", "arithsubcriber result")
	w.Header().Set("Content-Type", "application/json")
	w.Write(resultData)
}

// just a handler for testing
func handlerPing(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

// Send the task to arithsubcriber
func arithSender(task string) io.ReadCloser {

	// Make a POST reques to the subscriber with the task
	req, err := http.NewRequest("POST", ArithURL, bytes.NewBuffer([]byte(task)))

	// Set header
	req.Header.Set("X-Custom-Header", "arithproducer task")
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	return resp.Body
}
