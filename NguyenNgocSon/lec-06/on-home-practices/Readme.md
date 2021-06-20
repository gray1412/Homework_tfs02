#Project infomation

#Cách chạy project Tfs2

-Bước 1: truy cập forder backend->main.go, tại file main.go sửa đường dẫn mySql thành của bạn
example:
```bash
var mysqlCredentials = fmt.Sprintf(
	"%s:%s@%s(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
	"root",
	"password",
	"tcp",
	"localhost",
	"3306",
	"tfs2",
)
```
-Bước 2: tạo database tfs2 trong mySql
-Bước 3: Run file main.go bằng lệnh "go run main.go"
-Bước 5: Truy cập forder frontend->html, chạy file index.html

Good luck!