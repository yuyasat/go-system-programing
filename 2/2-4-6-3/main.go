package main

import (
	"net/http"
	"os"
)

func main() {
	request, err := http.NewRequest("GET", "http://ascii.js", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("X-TEST", "ヘッダーも追加できます ")
	request.Write(os.Stdout)

}
