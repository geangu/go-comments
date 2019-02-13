package main

import (
	"net/http"

	"./config"
)

func main() {
	_ = http.ListenAndServe(":3000", config.InitRoutes())
}
