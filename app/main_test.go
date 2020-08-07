package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestHeadersHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/headers", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	req.Header.Add("User-Agent", "go")
	handler := http.HandlerFunc(HeadersHandler)
	handler.ServeHTTP(rec, req)

	// Non 200 Check
	if status := rec.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Response Check
	expected := `User-Agent: go`
	if strings.TrimSuffix(rec.Body.String(), "\n") != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestNotFoundHandler(t *testing.T) {
	args := &ArgsHandler{filePath: "/path/doesnt/exist"}

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(args.Handler)
	handler.ServeHTTP(rec, req)

	// Test 404
	if status := rec.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}
func TestFoundHandler(t *testing.T) {

	file, _ := os.Create("/tmp/exists")
	file.WriteString("go")
	file.Close()

	args := &ArgsHandler{filePath: "/tmp/exists"}

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(args.Handler)
	handler.ServeHTTP(rec, req)

	os.Remove("/tmp/exists")

	// Test 404
	if status := rec.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}

	// Response Check
	expected := `go`
	if strings.TrimSuffix(rec.Body.String(), "\n") != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}
