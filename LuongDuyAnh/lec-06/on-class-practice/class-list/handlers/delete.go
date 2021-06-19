package handlers

import (
	"net/http"
	"strconv"
	"tfs-02/lec-06/on-class-practice/class-list/pkg"
	"tfs-02/lec-06/on-class-practice/class-list/storage"

	"github.com/gorilla/mux"
)

func DeleteData(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		pkg.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Invalid todo id"})
		return
	}

	for i, data := range storage.Datas {
		if data.Id == uint(id) {
			storage.Datas = append(storage.Datas[:i], storage.Datas[i+1:]...)
			pkg.ResponseWithJson(w, http.StatusOK, map[string]string{"message": "Todo was deleted"})
			return
		}
	}

	pkg.ResponseWithJson(w, http.StatusNotFound, map[string]string{"message": "Todo not found"})
}
