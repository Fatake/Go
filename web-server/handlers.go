package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// HandleHome HandleHome
func HandleHome(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second * 20)
	fmt.Fprintf(w, "Hello from the very best server!")
}

// HandleRoot HandleRoot
func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API Running")
}

// PostRequest PostRequest
func PostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var metadata MetaData
	if err := decoder.Decode(&metadata); err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	fmt.Fprintf(w, "Request Payload: %v\n", metadata)
}

// CreateRequest CreateRequest
func CreateRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Everything seems to be fine!")
}

// MetaData MetaData
type MetaData interface{}
