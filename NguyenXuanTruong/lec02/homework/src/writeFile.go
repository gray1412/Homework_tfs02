package main
 
import (
	"fmt"
	"os"
)
 
func main() {
    filename := "file/output.txt"
    data := "She got me going psycho"
 
    file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)

	if err == nil {
        fmt.Println("\nĐọc file thành công")
		_, err1 := fmt.Fprintln(file, "\n"+ data)
		if err1 == nil {
			fmt.Println("\nGhi vào file thành công")
		} else {
            fmt.Println("\nGhi vào file thất bại")
        }
		file.Close()
	} else {
		//File không tồn tại, tạo file mới
		file.Close()

		file2, err2 := os.Create(filename)
		if err2 == nil {
			fmt.Println("\nTạo file mới thành công")
			file2.Close()
			return
		} else {
            fmt.Println("\nTạo file mới thất bại")
            file2.Close()
        }

		_, err3 := file2.WriteString(data)
		if err3 == nil {
			fmt.Println("\nGhi vào file mới thành công")
			file2.Close()
			return
		} else {
			fmt.Println("\nGhi vào file mới thất bại")
			file2.Close()
		}
	}
}