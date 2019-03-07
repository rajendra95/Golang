package main

import (
	"net/http"
	"log"
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
		log.Println(" Welcome")
	}
	if token!= "magic" {
		log.Fatalf("You are NOT Authorized User")
	}

}
