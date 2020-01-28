package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCalculate(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{99999, 100001},
	}

	for _, test := range tests {
		if output := Calculate(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}

func TestGetArticles(t *testing.T) {
	req, err := http.NewRequest("GET", "/articles", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetArticles)
	handler.ServeHTTP(rr, req)

	// Check HTTP status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check response payload
	// 1. Valid JSON
	// 2. Correct field names
	// 3. Correct types
	// 4. Correct values

}
