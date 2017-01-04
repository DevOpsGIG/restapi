package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	server   *httptest.Server
	reader   io.Reader
	arithURL string
)

func init() {
	server = httptest.NewServer(handlers())
	arithURL = fmt.Sprintf("%s/arith", server.URL)
}

func TestHandlerArith(t *testing.T) {
	defer server.Close()
	// Test data
	testJSON := `{"Task": "test"}`
	// Parse string to reader
	reader = strings.NewReader(testJSON)
	// Create request with JSON body
	req, err := http.NewRequest("POST", arithURL, reader)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if resp.StatusCode != 200 {
		t.Errorf("Arith handler test failed: status code %d", resp.StatusCode)
	}
}
