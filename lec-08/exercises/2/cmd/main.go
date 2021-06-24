package main

import (
	handlers "tfs-02/lec-08/exercises/2/handlers"
	storage "tfs-02/lec-08/exercises/2/storage"
)

func main() {
	storage.Migrate()
	handlers.ReadData()
}
