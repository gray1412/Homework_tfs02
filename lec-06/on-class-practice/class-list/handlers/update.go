package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tfs-02/lec-06/on-class-practice/class-list/pkg"
	"tfs-02/lec-06/on-class-practice/class-list/storage"

	"github.com/gorilla/mux"
)

func UpdateData(w http.ResponseWriter, r *http.Request) {
	db := *pkg.ConnectDB() //ket noi DB
	defer db.Close()

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		pkg.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Invalid id"})
		return
	}

	var data storage.Data
	db.Find(&data, uint(id))
	json.NewDecoder(r.Body).Decode(&data)
	db.Save(&data)

}
