package handlers

import (
	"asslec6p2/database"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	//get id
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		responseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Invalid ID !"})
		return
	}

	//connect database
	db := *database.ConnectToDatabase()
	defer db.Close()

	//delete
	//nếu id không tồn tại trong database thì db.Delete vẫn chạy nhưng không có vấn đề gì xảy ra
	db.Delete(&database.Student{}, id)

	//response
	responseWithJson(w, http.StatusOK, map[string]string{"message": "Delete Successful !"})

	fmt.Println("Delete Successful !")
}
