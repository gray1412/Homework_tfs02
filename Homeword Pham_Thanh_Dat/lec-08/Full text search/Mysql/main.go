package main

import (
	datautil "amazon/data"
	storage "amazon/storage"
)

func main() {
	storage.Migrate()
	datautil.ReadData()
}
