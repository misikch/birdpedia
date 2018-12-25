package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(handler)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returns wrong status code: got %v, want %v", status, http.StatusOK)
	}

	expected := `Hello World!`

	actual := recorder.Body.String()

	if expected != actual {
		t.Errorf("Handler returned unexpected body: got %v want %v", actual, expected)
	}


}
