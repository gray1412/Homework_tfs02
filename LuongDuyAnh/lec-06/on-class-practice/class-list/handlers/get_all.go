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

	pkg.ResponseWithJson(w, http.StatusOK, storage.Datas)
}
