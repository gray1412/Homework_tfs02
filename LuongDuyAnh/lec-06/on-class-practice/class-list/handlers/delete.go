package handlers

import (
	"net/http"
	"strconv"
	"tfs-02/lec-06/on-class-practice/class-list/pkg"
	"tfs-02/lec-06/on-class-practice/class-list/storage"

	"github.com/gorilla/mux"
)

func DeleteData(w http.ResponseWriter, r *http.Request) {
	db := *pkg.ConnectDB()
	defer db.Close()

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		pkg.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Invalid todo id"})
		return
	}

	idDel := uint(id)
	var datas []storage.Data
	db.Delete(&datas, idDel) //xoa tu bang data where id=id
}
