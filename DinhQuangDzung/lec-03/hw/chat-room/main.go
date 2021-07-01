package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Message slice to hold messages
var messages []map[string]string

// Handle POST request: Read from request body
func handlePost(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal []byte data from request body to m map
	m := make(map[string]string)
	json.Unmarshal(reqBody, &m)

	messages = append(messages, m)    // Append unmarshalled message to slice
	data, _ := json.Marshal(messages) // Marshal slice to write to output file

	_ = ioutil.WriteFile("output.json", data, 0644) // Write to output file
}

// Handle GET request: Read from output file and print to screen
func handleGet(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadFile("output.json")
	fmt.Fprintln(w, string(data))
}

func main() {
	log.Println("Server started on port 8080")
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/chat", handlePost)
	http.HandleFunc("/message", handleGet)
	http.ListenAndServe(":8080", nil)
}
