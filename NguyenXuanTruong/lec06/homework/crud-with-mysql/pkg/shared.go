// Connect with database
package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"

	s "CallPkgs/storage"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Kết nối với database
func ConnectStudentDB() *gorm.DB {
	dsn := "root:1234@/tfsstudent"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Kết nối thất bại")
	}
	return db
}

// Chuyển dữ liệu từ dạng object sang json
func ResponseWithJson(w http.ResponseWriter, status int, object interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(object) // parse to json
}

// Với mỗi Student trong slice student, kiểm tra Id của Student đó với maxId có trong db
// Nếu Id của Student đó lớn hơn max Id, có nghĩa đó là Student mới được thêm vào
// Max Id điều chỉnh bằng với Id Student mới
func GenerateId(student []s.Student) uint {
	var maxId uint
	for _, s := range student {
		if s.Id > maxId {
			maxId = s.Id
		}
	}
	return maxId + 1
}
