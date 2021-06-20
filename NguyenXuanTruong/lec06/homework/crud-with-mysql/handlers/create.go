// Create data
package handlers

import (
	p "CallPkgs/pkg"
	s "CallPkgs/storage"
	"encoding/json"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	// Connect với database
	db := *p.ConnectStudentDB()
	defer db.Close()

	// Kiểm tra nếu trong db chưa có bảng thì tạo bảng
	if (!db.HasTable(&s.Student{})) {
		db.CreateTable(&s.Student{})
	}

	// Sử dụng json.Decoder thay vì json.Unmarshal, bởi dữ liệu được đọc từ HTTP
	// Nếu Decode không thành công, đưa ra thông báo "Invalid body"
	var newStudent s.Student
	if err := json.NewDecoder(r.Body).Decode(&newStudent); err != nil {
		p.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}

	// Thêm vào database
	db.Create(newStudent)

	// Hiển thị kết quả trong Body
	p.ResponseWithJson(w, http.StatusCreated, newStudent)
}
