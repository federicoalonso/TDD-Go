package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCalculate(t *testing.T) {
	if Calculate(2) != 4 {
		t.Error("Expexted 2 + 2 to equal 4")
	}
}

func TestTableCalculate(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{4, 6},
		{999, 1001},
	}

	for _, test := range tests {
		if output := Calculate(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}

type PlusResult struct {
	x        int
	y        int
	expected int
}

var plusResult = []PlusResult{
	{1, 1, 2},
	{1, 5, 6},
	{9, 1, 10},
}

func TestPlus(t *testing.T) {
	for _, test := range plusResult {
		result := Plus(test.x, test.y)
		if result != test.expected {
			t.Fatal("Expected result not given")
		}
	}
}

func TestHttpRequest(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{ \"status\": \"good\" }")
	}

	req := httptest.NewRequest("GET", "https://tutorialedge.net", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	if 200 != resp.StatusCode {
		t.Fatal("Status Code not OK")
	}
}
