package main

import (
	"fmt"
	"net/http"
	"log"
	"strconv"
)

func main(){

	http.HandleFunc("/",index_handler)
	err:= http.ListenAndServe(":8080",nil)
	if err!=nil{
		log.Fatalf("Error in starting the server ")
	}
}

func index_handler(w http.ResponseWriter,r *http.Request){

	text:=r.FormValue("v")//it should be Fprints otherwise it wont work.
	if text ==""{
		http.Error(w,"missing literal",http.StatusBadRequest)
		return
	}
	v,err:=strconv.Atoi(text)
	if err!=nil{
		http.Error(w,"not a number ",http.StatusBadRequest)
		return
	}
	//fmt.Fprintf(w," <h1> Welcome <h1> \n")
	fmt.Fprintln(w,v*2)
}
