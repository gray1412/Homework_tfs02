package main

import (
	"fmt"
	"log"
	"regexp"
	"runtime"
	"storage/handlers"
	storage "storage/storage"
	"time"
)

func main() {
	//time to run func
	defer TimeTrack(time.Now())
	// storage.Migrate()
	handlers.ReadData()

	//search by mysql
	db := storage.Connect()

	// LIKE
	db.Debug().Where("title LIKE ?", "%beautiful%").Find(&storage.Review{})
	// SELECT * FROM reviews WHERE title LIKE '%Stunning%';
}

func TimeTrack(start time.Time) {
	elapsed := time.Since(start)

	// Skip this function, and fetch the PC and file for its parent.
	pc, _, _, _ := runtime.Caller(1)

	// Retrieve a function object this functions parent.
	funcObj := runtime.FuncForPC(pc)

	// Regex to extract just the function name (and not the module path).
	runtimeFunc := regexp.MustCompile(`^.*\.(.*)$`)
	name := runtimeFunc.ReplaceAllString(funcObj.Name(), "$1")

	log.Println(fmt.Sprintf("%s took %s", name, elapsed))
}
