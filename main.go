package main

import (
	"net/http"
)

func main() {
	_ = http.ListenAndServe(":3000", InitRoutes())
}
