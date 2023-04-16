package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.StatusCode)
}
