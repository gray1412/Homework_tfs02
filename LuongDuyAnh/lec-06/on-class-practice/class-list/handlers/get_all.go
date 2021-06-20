package handlers

import (
	"net/http"
	"tfs-02/lec-06/on-class-practice/class-list/pkg"
	"tfs-02/lec-06/on-class-practice/class-list/storage"
)

func GetAllData(w http.ResponseWriter, r *http.Request) {
	db := *pkg.ConnectDB()
	defer db.Close()
	if (!db.HasTable(&storage.Data{})) {
		db.CreateTable(&storage.Data{})
	}

	var datas []storage.Data
	db.Find(&datas) //lay tat ca du lieu cua bang data
	pkg.ResponseWithJson(w, http.StatusOK, datas)

}
