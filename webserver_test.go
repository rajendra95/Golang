package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"io/ioutil"
	"strconv"
	"bytes"
)

func TestIndexHandler( t *testing.T){

	tt := []struct{
		name string
		value string
		double int
		status int
	}{
		{name: "double of two",value : "3" ,status:http.StatusOK},
	}
	for _, tc := range tt {
		req, err := http.NewRequest("GET", "localhost:8080/double?v="+ tc.value, nil)
		if err != nil {
			t.Fatalf("we could not create the request %v", err)
		}
		rec := httptest.NewRecorder()
		index_handler(rec, req)
		res := rec.Result()
		defer res.Body.Close()
		if res.StatusCode != tc.status{
			t.Fatalf("expected response code %v, got %v", tc.status,res.StatusCode)
		}
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("Could not read the response")
		}
		d, err := strconv.Atoi(string(bytes.TrimSpace(b)))
		if err != nil {
			t.Fatalf("expected integer, got %s", d)
		}
		if d != 4 {
			t.Fatalf("expected 4, got %v", d)
		}
	}



}

