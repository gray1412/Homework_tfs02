package handlers

import (
	p "CallPkgs/pkg"
	s "CallPkgs/storage"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Update(w http.ResponseWriter, r *http.Request) {
	// Connect với database
	db := *p.ConnectStudentDB()
	defer db.Close()

	// Truy xuất đến phần tử có Id trùng khớp với id đầu vào
	params := mux.Vars(r)
	IdUpdate, err := strconv.Atoi(params["id"])

	// Nếu không có giá trị Id nào thỏa mãn, đưa ra thông báo "Invalid todo id"
	if err != nil {
		p.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Invalid todo id"})
		return
	}

	// Kiểm tra phần tử đầu vào có thỏa mãn hay không
	// Nếu không thỏa mãn, đưa ra thông báo "Invalid body"
	var StudentUpdate s.Student
	if err := json.NewDecoder(r.Body).Decode(&StudentUpdate); err != nil {
		p.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}
	StudentUpdate.Id = uint(IdUpdate)

	// Duyệt các phần tử trong slice Students
	// Thay thế giá trị của phần tử cần thay thế bằng giá trị của StudentUpdate
	for i, stu := range s.Students {
		if stu.Id == StudentUpdate.Id {
			s.Students[i] = StudentUpdate
			p.ResponseWithJson(w, http.StatusOK, StudentUpdate)
			return
		}
	}

	p.ResponseWithJson(w, http.StatusNotFound, map[string]string{"message": "Todo not found"})
}
