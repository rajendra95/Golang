package main

import (
	"testing"
	"net/http"
	"log"
	"net/http/httptest"

)

func TestResponse_handler(t *testing.T){

	req,err:=http.NewRequest("GET","/",nil)
	req.Header.Set("X-Access-Token","magic")
	if err!=nil{
		log.Fatalf("Error in creating the Request %v",err)
	}
	rec:= httptest.NewRecorder()

	tc:=[]struct{
		ExpectedStatusCode int
		ExpectedResponseBody []byte
	}{
		{ExpectedStatusCode:http.StatusOK,ExpectedResponseBody:[]byte("You have some magic in you\n")},
		//{AccessToken: "", ExpectedStatusCode:http.StatusForbidden, ExpectedResponseBody:[]byte(("You don't have enough magic in you\n")),request:req},
	}
	for _,c:= range tc{
		response_handler(rec,req)
		if c.ExpectedStatusCode!=req.Response.StatusCode{
			t.Errorf("the status code does not match! Expected %v Got %v",c.ExpectedStatusCode,req.Response.StatusCode)
			t.Fail()
		}
		/*
		if !bytes.Equal(c.ExpectedResponseBody,c.request.Body){
			t.Errorf("Body did not match Expected %v Got %v",c.ExpectedResponseBody,c.request.Body)
			t.Fail()
		}*/
	}


}
