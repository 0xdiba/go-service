package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestHeartbeatHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/hb", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HeartBeatHandler)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "hello\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestSquareHandler(t *testing.T) {

	testdata := []struct {
		Input  string
		Output string
		Status int
	}{
		{"foo", "Could not parse input\n", http.StatusBadRequest},
		{"0", "0", http.StatusOK},
		{"5", "25", http.StatusOK},
	}

	for _, test := range testdata {
		form := url.Values{}
		form.Add("input", test.Input)
		req, err := http.NewRequest("POST", "/square", strings.NewReader(form.Encode()))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(SquareHandler)

		handler.ServeHTTP(rr, req)
		if status := rr.Code; status != test.Status {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, test.Status)
		}

		if rr.Body.String() != test.Output {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), test.Output)
		}
	}
}
