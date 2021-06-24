package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func ConnectDB() *gorm.DB {
	mysqlCredentials := fmt.Sprintf(
		"%s:%s@%s(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"root",
		"123456",
		"tcp",
		"localhost",
		"3306",
		"tfs-02",
	)
	db, err := gorm.Open("mysql", mysqlCredentials)
	if err != nil {
		fmt.Println(err)
		panic("fall to connect DB")
	}
	// defer db.Close()
	return db
}

//parse to json
func ResponseWithJson(w http.ResponseWriter, status int, object interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(object)
}
