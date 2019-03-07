package main

import (
	"net/http"
	"log"
	"fmt"
)

func main(){
	http.HandleFunc("/",response_handler)
	err:=http.ListenAndServe(":8080",nil)
	if err!=nil{
		log.Fatalf("Error in fetching the request")
	}
}
func response_handler(w http.ResponseWriter,r *http.Request){
	token:=r.Header.Get("X-Access-Token")
	if token=="magic"{
		fmt.Fprintf(w,"************** Token verified************")
		fmt.Fprintf(w,"Authorization Successful")
		log.Println(" Welcome")
	}else {
		fmt.Fprintf(w,"************** Token NOT verified************")
		fmt.Fprintf(w,"You are NOT Authorized User")
	}

}
