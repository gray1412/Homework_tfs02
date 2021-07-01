package main

import (
	"fmt"
	"tfs02/database"
)

func main() {
	database.Migrate()
	fmt.Println("done")
}
