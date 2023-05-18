package main

import (
	"net/http"

	"example.com/test/goworkspace/src/000-personaltest/httpServerBasic2/api"
)

func main() {
	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)
}
