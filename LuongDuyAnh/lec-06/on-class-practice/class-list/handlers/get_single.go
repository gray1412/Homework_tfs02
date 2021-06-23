package handlers

import (
	"net/http"
	"strconv"
	"tfs-02/lec-06/on-class-practice/class-list/pkg"
	"tfs-02/lec-06/on-class-practice/class-list/storage"

	"github.com/gorilla/mux"
)

func GetData(w http.ResponseWriter, r *http.Request) {
	db := *pkg.ConnectDB()
	defer db.Close()

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		pkg.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Invalid data id"})
		return
	}

	var data []storage.Data
	db.Find(&data, uint(id))

	for _, data := range data {
		pkg.ResponseWithJson(w, http.StatusOK, data)
		return
	}
	pkg.ResponseWithJson(w, http.StatusNotFound, map[string]string{"message": "Data not found"})
}
