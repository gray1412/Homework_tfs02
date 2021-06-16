package main

import (
	f "CallPkg/pkgs/handleFile"
)

func main() {
	f.WriteToFile("file/output.txt", "12345678")
}
