package handlers

import (
	"encoding/json"
	"net/http"

	"NDT/connectDb"
	"NDT/memory"
)

func CreateData(w http.ResponseWriter, r *http.Request) {
	db := *connectDb.ConnectDB()
	defer db.Close()
	if (!db.HasTable(&memory.Data{})) {
		db.CreateTable(&memory.Data{})
	}

	var newData memory.Data
	if err := json.NewDecoder(r.Body).Decode(&newData); err != nil {
		connectDb.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}

	db.Create(newData)
	connectDb.ResponseWithJson(w, http.StatusCreated, newData)

}
