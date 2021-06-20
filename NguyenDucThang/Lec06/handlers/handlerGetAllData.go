package handlers

import (
	"NDT/connectDb"
	"NDT/memory"
	"net/http"
)

func GetAllData(w http.ResponseWriter, r *http.Request) {
	db := *connectDb.ConnectDB()
	defer db.Close()
	if (!db.HasTable(&memory.Data{})) {
		db.CreateTable(&memory.Data{})
	}

	var datas []memory.Data
	db.Find(&datas)
	connectDb.ResponseWithJson(w, http.StatusOK, datas)

}
