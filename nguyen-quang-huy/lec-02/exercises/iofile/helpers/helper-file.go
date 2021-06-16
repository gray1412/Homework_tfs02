package helpers

import (
	"io/ioutil"
	"os"
)

func ReadFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return "", err
	}

	data := make([]byte, stat.Size())
	_, err = file.Read(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
func WriteFile(path string, data string) error{

	mydata := []byte(data)

	err := ioutil.WriteFile(path, mydata, 0777)
	return err
}