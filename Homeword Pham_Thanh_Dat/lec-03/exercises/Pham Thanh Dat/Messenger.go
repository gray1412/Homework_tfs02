package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var id = 0
var data []string

type Messenger struct {
	Id   int
	Name string
	Mess string
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
func set(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	id++
	name := r.URL.Query().Get("Name")
	messenger := r.URL.Query().Get("Messenger")
	bb := &Messenger{Id: id, Name: name, Mess: messenger}
	re, err := json.Marshal(bb)
	if err != nil {
		fmt.Println(err)
		return
	}
	if id == 1 {
		data = append(data, string(re))
	} else {
		data = append(data, ","+string(re))
	}
	fmt.Fprintf(w, "%v", data)

}
func getAll(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	fmt.Fprintf(w, "%v", data)
}
func loadMessenger(w http.ResponseWriter, r *http.Request) {
	var numberMessenger = r.URL.Query().Get("numberMessenger")
	n := 0
	n, _ = strconv.Atoi(numberMessenger)
	enableCors(&w)
	bb := &Messenger{Id: id, Name: "", Mess: ""}
	re, err := json.Marshal(bb)
	if err != nil {
		fmt.Println(err)
		return
	}
	var newMessenger []string
	newMessenger = append(newMessenger, string(re))
	for i := len(data) - n; i < len(data); i++ {
		newMessenger = append(newMessenger, data[i])
	}
	fmt.Fprintf(w, "%v", newMessenger)
}

func main() {
	http.HandleFunc("/set", set)
	http.HandleFunc("/getAll", getAll)
	http.HandleFunc("/loadMessenger", loadMessenger)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
