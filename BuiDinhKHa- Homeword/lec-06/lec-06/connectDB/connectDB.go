package connectDb

import (
	"encoding/json"
	"fmt"
	"net/http"

	"NDT/memory"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@/lec06?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Connect Failed!")
		panic(err)
	}
	return db

}

func ResponseWithJson(w http.ResponseWriter, status int, object interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(object)
}

func GenerateId(data []memory.Data) uint {
	var maxId uint
	for _, data := range data {
		if data.Id > maxId {
			maxId = data.Id
		}
	}
	return maxId + 1
}
