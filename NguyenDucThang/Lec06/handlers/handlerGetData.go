package handlers

import (
	"NDT/connectDb"
	"NDT/memory"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetData(w http.ResponseWriter, r *http.Request) {
	db := *connectDb.ConnectDB()
	defer db.Close()

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		connectDb.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Invalid data id"})
		return
	}

	var data []memory.Data
	db.Find(&data, uint(id))

	for _, data := range data {
		connectDb.ResponseWithJson(w, http.StatusOK, data)
		return
	}
	connectDb.ResponseWithJson(w, http.StatusNotFound, map[string]string{"message": "Data not found"})
}
