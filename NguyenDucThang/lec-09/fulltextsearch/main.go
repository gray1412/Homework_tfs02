package main

import (
	datautil "fultextsearch/data"
	storage "fultextsearch/storage"
)

func main() {
	storage.CreateReview()
	datautil.ReadData()
}
