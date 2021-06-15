package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Start....")
	http.HandleFunc("/chat", chat)
	http.ListenAndServe(":8080", nil)
	defer func() {
		fmt.Println("end")
	}()

}

type DataChat struct {
	Name     string `json: "name"`
	Time     string `json: "time"`
	DataChat string `json: "datachat"`
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

var DataSlice []string

func chat(w http.ResponseWriter, rep *http.Request) {
	enableCors(&w)
	nameChat := rep.URL.Query().Get("name")
	time := rep.URL.Query().Get("time")
	data := rep.URL.Query().Get("chat")
	i := DataChat{
		Name:     nameChat,
		Time:     time,
		DataChat: data,
	}
	bb, _ := json.Marshal(&i)
	if len(DataSlice) == 0 {
		DataSlice = append(DataSlice, string(bb))
	} else {
		DataSlice = append(DataSlice, ","+string(bb))
	}
	fmt.Fprintln(w, DataSlice)

}
