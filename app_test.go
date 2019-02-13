package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestCreateOk Successful creation of a comment
func TestCreateOk(t *testing.T) {
	routes := InitRoutes()
	response := httptest.NewRecorder()

	comment := &Comment{
		Text:          "Test comments",
		User:          "test-purchase-1",
		Establishment: "test-establishment-1",
		Purchase:      "test-purchase-1",
		Score:         1,
	}
	jsonComment, _ := json.Marshal(comment)

	request, err := http.NewRequest("POST", "/comments/", bytes.NewBuffer(jsonComment))
	if err != nil {
		t.Fatal(err)
	}

	routes.ServeHTTP(response, request)
	if response.Code != http.StatusCreated {
		t.Errorf("handler returned wrong code: got %v want %v", response.Code, http.StatusCreated)
	}
}

// TestCreateValidationErrors Validation errors in the creation of a comment
func TestCreateValidationErrors(t *testing.T) {
	routes := InitRoutes()
	response := httptest.NewRecorder()

	comment := &Comment{
		Text:          "Test comments",
		User:          "test-purchase-1",
		Establishment: "test-establishment-1",
		Purchase:      "test-purchase-1",
		Score:         10,
	}
	jsonComment, _ := json.Marshal(comment)

	request, err := http.NewRequest("POST", "/comments/", bytes.NewBuffer(jsonComment))
	if err != nil {
		t.Fatal(err)
	}

	routes.ServeHTTP(response, request)

	code := response.Code
	result := response.Body.String()
	expected := "Missing or wrong request data\n"

	if code != http.StatusBadRequest {
		t.Errorf("handler returned wrong code: got %v want %v", code, http.StatusBadRequest)
	}

	if expected != result {
		t.Errorf("handler returned wrong result: got %v want %v", result, expected)
	}
}

// TestCreateDuplicated Creation of duplicate comment
func TestCreateDuplicated(t *testing.T) {
	routes := InitRoutes()
	response := httptest.NewRecorder()

	comment := &Comment{
		Text:          "Test comments",
		User:          "test-purchase-1",
		Establishment: "test-establishment-1",
		Purchase:      "test-purchase-1",
		Score:         1,
	}
	jsonComment, _ := json.Marshal(comment)

	request, err := http.NewRequest("POST", "/comments/", bytes.NewBuffer(jsonComment))
	if err != nil {
		t.Fatal(err)
	}

	routes.ServeHTTP(response, request)
	routes.ServeHTTP(response, request)

	code := response.Code
	result := response.Body.String()
	expected := "There is already a qualification for this purchase\nEOF\n"

	if code != http.StatusConflict {
		t.Errorf("handler returned wrong code: got %v want %v", code, http.StatusConflict)
	}

	if expected != result {
		t.Errorf("handler returned wrong result: got %v want %v", result, expected)
	}
}

// TestFindByPuchaseOk Get comment on a purchase
func TestFindByPuchaseOk(t *testing.T) {
	routes := InitRoutes()

	response := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/purchase/test-purchase-1/comments/", nil)
	if err != nil {
		t.Fatal(err)
	}

	routes.ServeHTTP(response, request)
	if response.Code != http.StatusOK {
		t.Errorf("handler returned wrong code: got %v want %v", response.Code, http.StatusOK)
	}
}

// TestFindByPuchaseNotFound Can not find a purchase comment
func TestFindByPuchaseNotFound(t *testing.T) {
	routes := InitRoutes()

	response := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/purchase/123456/comments/", nil)
	if err != nil {
		t.Fatal(err)
	}

	routes.ServeHTTP(response, request)

	code := response.Code
	result := response.Body.String()
	expected := "Resource not found\n"

	if code != http.StatusNotFound {
		t.Errorf("handler returned wrong code: got %v want %v", code, http.StatusBadRequest)
	}

	if expected != result {
		t.Errorf("handler returned wrong result: got %v want %v", result, expected)
	}
}

// TestFindByEstablishmentOk Get comments from an establishment
func TestFindByEstablishmentOk(t *testing.T) {
	routes := InitRoutes()

	response := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/establishment/test-establishment-1/comments/?start=2019-01-01T00:00:00-05:00&end=2030-01-01T00:00:00-05:00", nil)
	if err != nil {
		t.Fatal(err)
	}

	routes.ServeHTTP(response, request)

	code := response.Code

	if code != http.StatusOK {
		t.Errorf("handler returned wrong code: got %v want %v", code, http.StatusOK)
	}

}

// TestFindByEstablishmentNotFound Comments of an establishment not found
func TestFindByEstablishmentNotFound(t *testing.T) {
	routes := InitRoutes()

	response := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/establishment/test-establishment-1/comments/?start=2019-01-01T00:00:00-05:00&end=2019-01-01T00:00:00-05:00", nil)
	if err != nil {
		t.Fatal(err)
	}

	routes.ServeHTTP(response, request)

	code := response.Code
	result := response.Body.String()
	expected := "No resources found\n"

	if code != http.StatusNotFound {
		t.Errorf("handler returned wrong code: got %v want %v", code, http.StatusNotFound)
	}

	if result != expected {
		t.Errorf("handler returned wrong result: got %v want %v", result, expected)
	}
}

// TestFindByEstablishmentWrongStart Start date not valid
func TestFindByEstablishmentWrongStart(t *testing.T) {
	routes := InitRoutes()

	response := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/establishment/test-establishment-1/comments/?start=2019-01-01&end=2020-01-01", nil)
	if err != nil {
		t.Fatal(err)
	}

	routes.ServeHTTP(response, request)

	code := response.Code
	result := response.Body.String()
	expected := "The 'start' parameter is not in the expected date format RFC3339\n"

	if code != http.StatusBadRequest {
		t.Errorf("handler returned wrong code: got %v want %v", code, http.StatusBadRequest)
	}

	if result != expected {
		t.Errorf("handler returned wrong result: got %v want %v", result, expected)
	}
}

// TestFindByEstablishmentWrongEnd End date not valid
func TestFindByEstablishmentWrongEnd(t *testing.T) {
	routes := InitRoutes()

	response := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/establishment/test-establishment-1/comments/?start=2019-01-01T00:00:00-05:00&end=2020-01-01", nil)
	if err != nil {
		t.Fatal(err)
	}

	routes.ServeHTTP(response, request)

	code := response.Code
	result := response.Body.String()
	expected := "The 'end' parameter is not in the expected date format RFC3339\n"

	if code != http.StatusBadRequest {
		t.Errorf("handler returned wrong code: got %v want %v", code, http.StatusBadRequest)
	}

	if result != expected {
		t.Errorf("handler returned wrong result: got %v want %v", result, expected)
	}
}

// TestDeleteOk Successful creation of a comment
func TestDeleteNotFound(t *testing.T) {
	routes := InitRoutes()
	response := httptest.NewRecorder()

	request, err := http.NewRequest("DELETE", "/purchase/test-purchase/comments/", nil)
	if err != nil {
		t.Fatal(err)
	}

	routes.ServeHTTP(response, request)
	if response.Code != http.StatusNotFound {
		t.Errorf("handler returned wrong code: got %v want %v", response.Code, http.StatusNotFound)
	}
}

// TestDeleteOk Successful creation of a comment
func TestDeleteOk(t *testing.T) {
	routes := InitRoutes()
	response := httptest.NewRecorder()

	request, err := http.NewRequest("DELETE", "/purchase/test-purchase-1/comments/", nil)
	if err != nil {
		t.Fatal(err)
	}

	routes.ServeHTTP(response, request)
	if response.Code != http.StatusOK {
		t.Errorf("handler returned wrong code: got %v want %v", response.Code, http.StatusOK)
	}
}
