package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
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

func TestRouterForNotExistingRoute (t *testing.T) {
	r := makeRouter()

	mockServer := httptest.NewServer(r)

	resp, err := http.Post(mockServer.URL + "/", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("unexpected status code. Geot %v, exptected %v", resp.StatusCode, http.StatusMethodNotAllowed)
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	respString := string(b)
	expected := ""

	if respString != expected {
		t.Errorf("Response body should be empty string, got: %v", respString)
	}
}

func TestStaticFileServer(t *testing.T) {
	r := makeRouter()

	mockServer := httptest.NewServer(r)

	resp, err := http.Get(mockServer.URL + "/assets/")

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code. Expected: %v, got: %v", http.StatusOK, resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")

	expectedContentType := "text/html; charset=utf-8"

	if expectedContentType != contentType {
		t.Errorf("unexpected content-type. Expected: %v, got: %v", expectedContentType, contentType)
	}
}

func TestGetBirdsHandler(t *testing.T) {
	r := makeRouter()

	mockServer := httptest.NewServer(r)

	b := Bird{
		Species:"s1",
		Description:"s1d",
	}

	birds = append(birds, b)

	resp, err := http.Get(mockServer.URL + "/bird")

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code. Expected: %v, got: %v", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	bodyString := string(body)

	expected := `[{"species":"s1","description":"s1d"}]`

	if bodyString != expected {
		t.Errorf("Unexpected body. Expected: %v, got: %v", expected, bodyString)
	}
}

func TestCreateBirdHandler(t *testing.T) {
	birds = []Bird {
		{"sparrow", "A small harmless bird"},
	}

	form := newCreateBirdForm()

	req, err := http.NewRequest("POST", "", bytes.NewBufferString(form.Encode()))

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))

	recorder := httptest.NewRecorder()

	ht := http.HandlerFunc(createBirdHandler)

	ht.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusFound {
		t.Errorf("unexpected status code. Expected: %v, got: %v", http.StatusFound, status)
	}

	expected := Bird{"eagle", "A bird of prey"}

	actual := birds[1]

	if actual != expected {
		t.Errorf("handler should create new bird, but somithing went wrong")
	}
}

func newCreateBirdForm() *url.Values {
	form := url.Values{}
	form.Set("species", "eagle")
	form.Set("description", "A bird of prey")

	return &form
}
