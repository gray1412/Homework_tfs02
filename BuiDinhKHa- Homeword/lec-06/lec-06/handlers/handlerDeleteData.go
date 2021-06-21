package handlers

import (
	"net/http"
	"strconv"

	"NDT/connectDb"
	"NDT/memory"

	"github.com/gorilla/mux"
)

func DeleteData(w http.ResponseWriter, r *http.Request) {
	db := *connectDb.ConnectDB()
	defer db.Close()

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		connectDb.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Invalid todo id"})
		return
	}

	idDel := uint(id)
	var datas []memory.Data
	db.Delete(&datas, idDel) //xoa tu bang data where id=id
}
