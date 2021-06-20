package handlers

import (
	"CallPkgs/pkg"
	p "CallPkgs/pkg"
	s "CallPkgs/storage"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetSingle(w http.ResponseWriter, r *http.Request) {
	// Connect với database
	db := *p.ConnectStudentDB()
	defer db.Close()

	// Truy xuất đến phần tử có Id trùng khớp với id đầu vào
	params := mux.Vars(r)
	IdGetSingle, err := strconv.Atoi(params["id"])

	// Nếu không có giá trị Id nào thỏa mãn, đưa ra thông báo "Invalid todo id"
	if err != nil {
		p.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Invalid todo id"})
		return
	}

	// Duyệt các phần tử trong slice student, đưa phần tử cần tìm ra thông báo
	for _, stu := range s.Students {
		if stu.Id == uint(IdGetSingle) {
			p.ResponseWithJson(w, http.StatusOK, stu)
			return
		}
	}

	// Nếu phần tử cần tìm không tìm thấy, in ra thông báo "Data not found"
	pkg.ResponseWithJson(w, http.StatusNotFound, map[string]string{"message": "Data not found"})
}
