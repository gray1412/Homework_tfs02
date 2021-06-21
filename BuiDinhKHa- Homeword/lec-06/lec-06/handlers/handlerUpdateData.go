package handlers

import (
	"NDT/connectDb"
	"NDT/memory"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func UpdateData(w http.ResponseWriter, r *http.Request) {
	db := *connectDb.ConnectDB() //ket noi DB
	defer db.Close()

	params := mux.Vars(r)
	_, err := strconv.Atoi(params["id"])

	if err != nil {
		connectDb.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Invalid id"})
		return
	}

	var data memory.Data
	json.NewDecoder(r.Body).Decode(&data)
	db.Save(&data)

}
