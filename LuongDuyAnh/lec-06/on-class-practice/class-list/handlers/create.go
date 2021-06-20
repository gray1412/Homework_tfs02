package handlers

import (
	"encoding/json"
	"net/http"
	"tfs-02/lec-06/on-class-practice/class-list/pkg"
	"tfs-02/lec-06/on-class-practice/class-list/storage"
)

func CreateData(w http.ResponseWriter, r *http.Request) {
	//ket noi DB
	db := *pkg.ConnectDB()
	defer db.Close()
	if (!db.HasTable(&storage.Data{})) {
		db.CreateTable(&storage.Data{})
	}

	var newData storage.Data
	if err := json.NewDecoder(r.Body).Decode(&newData); err != nil {
		pkg.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}

	//them vao DB
	db.Create(newData)
	pkg.ResponseWithJson(w, http.StatusCreated, newData)

}
