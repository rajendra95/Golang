package main

import (
	"testing"

	"net/http"
	"net/http/httptest"
	"log"
)

func TestResponse_handler(t *testing.T){
	req,err:=http.NewRequest("GET","/",nil)
	if err!=nil{
		t.Fatalf("Error in request",err)
	}
	tt := []struct{

		Testname string
		rec *httptest.ResponseRecorder
		request *http.Request
		accessToken string
		ExpectedStatusCode int
	}{
		{Testname:"Correct Header Included", rec: httptest.NewRecorder(),request:req,accessToken:"magic",ExpectedStatusCode:http.StatusOK},
		{Testname:"Correct Header Empty", rec: httptest.NewRecorder(),request:req,accessToken:"",ExpectedStatusCode:http.StatusOK},
	}

	for _,tc:=range tt{

		req.Header.Set("X-Access-Token",tc.accessToken)
		log.Println(req.Header)
		response_handler(tc.rec,tc.request)
		res:=tc.rec.Result()
		defer res.Body.Close()
		log.Println(res.StatusCode)
		if res.StatusCode!= tc.ExpectedStatusCode{
			t.Fatalf("Expected statuscode %v Got %v ",tc.ExpectedStatusCode,res.StatusCode)
			t.Fail()
		}

	}
}

