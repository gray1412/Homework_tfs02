package handlers

import (
	p "CallPkgs/pkg"
	s "CallPkgs/storage"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	// Connect với database
	db := *p.ConnectStudentDB()
	defer db.Close()

	// Truy xuất đến phần tử có Id trùng khớp với id đầu vào
	params := mux.Vars(r)
	IdDelete, err := strconv.Atoi(params["id"])

	// Nếu không có giá trị Id nào thỏa mãn, đưa ra thông báo "Invalid todo id"
	if err != nil {
		p.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Invalid todo id"})
		return
	}

	// Duyệt các phần tử nằm trong Students
	// append các phần tử: tử đầu đến vị trí cần xóa, bỏ qua phần tử ở vị trí cần xóa, tiếp tục với phần còn lại
	for i, stu := range s.Students {
		if stu.Id == uint(IdDelete) {
			s.Students = append(s.Students[:i], s.Students[i+1:]...)
			p.ResponseWithJson(w, http.StatusOK, map[string]string{"message": "Todo was deleted"})
			return
		}
	}

	// Nếu phần tử cần xóa không tìm thấy, đưa ra thông báo "Todo not found"
	p.ResponseWithJson(w, http.StatusNotFound, map[string]string{"message": "Todo not found"})
}
