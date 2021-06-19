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
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		pkg.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Invalid todo id"})
		return
	}

	var updateData storage.Data
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		pkg.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}
	updateData.Id = uint(id)

	for i, data := range storage.Datas {
		if data.Id == uint(id) {
			storage.Datas[i] = updateData
			pkg.ResponseWithJson(w, http.StatusOK, updateData)
			return
		}
	}

	pkg.ResponseWithJson(w, http.StatusNotFound, map[string]string{"message": "Todo not found"})
}
