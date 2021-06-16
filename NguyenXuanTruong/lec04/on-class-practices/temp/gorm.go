package main
import (
	_ "github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"encoding/json"
  )
type UserInfo struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

func (u UserInfo) String() string {
	b, err := json.Marshal(u)
	if err != nil {
		return ""
	}
	return string(b)
}

func main() {
	db, err := gorm.Open("mysql", "user = root:password = 1234 dbname = Human?charset=utf8&parseTime=True&loc=Local")
  	if err != nil {
    	panic("failed to connect database") // Kiểm tra kết nối tới database
  	}
  	defer db.Close() // Để đóng cơ sở dữ liệu khi nó không được sử dụng

	// Sau khi kết nối tới database ta tạo bảng bằng lệnh CreateTable()
    db.CreateTable(&UserInfo{})

	var human = UserInfo{ID: 1, Name: "Truong"}
	db.Create(&human)

	
	u := UserInfo{
		ID:   1,
		Name: "Truong",
	}
	u2 := UserInfo{
		ID:   1,
		Name: "Truong",
	}
	fmt.Println(u, u2.String())

}