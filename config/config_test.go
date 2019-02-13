package config

import "testing"

func TestInitRoutes(t *testing.T) {
	router := InitRoutes()
	if router == nil {
		t.Error("The routes could not be configured")
	}
}
