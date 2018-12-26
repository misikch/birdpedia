package main

import (
	"io/ioutil"
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

func TestRouter(t *testing.T) {
	r := makeRouter()

	mockServer := httptest.NewServer(r)

	resp, err := http.Get(mockServer.URL + "/")

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Unexpected status code. Expected 200, got %v", resp.StatusCode)
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	responseString := string(b)

	expected := "Hello World!"

	if responseString != expected {
		t.Errorf("unexpected response body. Got: %v, expected: %v", responseString, expected)
	}
}
