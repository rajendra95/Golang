package main

import (
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
	"bytes"
	"fmt"
	"time"
)

// Client sends its location to the server and server is responding back with the weatherData.

type weatherData struct {
	LocationName string   `json: locationName`
	Weather      string   `json: weather`
	Temperature  int      `json: temperature`
	Celsius      bool     `json: celsius`
	TempForecast []int    `json: temp_forecast`
	Wind         windData `json: wind`
}

type windData struct {
	Direction string `json: direction`
	Speed     int    `json: speed`
}

type loc struct {
	lat float32 `json:lat`
	lon float32 `json:lon`
}

// 1st of all we need funchandler to handler the req from user

func weatherHandler(w http.ResponseWriter,r *http.Request){
	// need to create a req
	location:= loc{}
	jsn,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		log.Fatalf("Error while reading request body")
	}
	err =json.Unmarshal(jsn,&location)
	if err!=nil{
		log.Fatalf("error while decoding the incoming request")
	}
	log.Printf("the incoming request",location)

	// creation of respsonse

	weather := weatherData{
		LocationName: "Zzyzx",
		Weather:      "cloudy",
		Temperature:  31,
		Celsius:      true,
		TempForecast: []int{30, 32, 29},
		Wind: windData{
			Direction: "S",
			Speed:     20,
		},
	}
	weatherJson , err:= json.Marshal(weather)
	if err!=nil{
		log.Printf("Error occured ", err)
	}
	w.Header().Set("Content-Type","application/json") // setting up the header
	w.Write(weatherJson) // write this response into the json body

}

func server(){
	http.HandleFunc("/",weatherHandler)
	http.ListenAndServe(":8080",nil)
}

func client(){
	locJson,err:=json.Marshal(loc{lat:35.124,lon:12.456})
	if err!=nil{
		log.Fatalf("Error while marshaling")
	}
	req,err :=http.NewRequest("POST","http://localhost:8080",bytes.NewBuffer(locJson))
	req.Header.Set("Content-Type","application/json")
	client:=http.Client{
		Timeout:time.Second*2,
	}
	resp,err :=client.Do(req)
	if err!=nil{
		log.Fatal(err)
	}
	body,err:=ioutil.ReadAll(resp.Body)
	fmt.Println("response body is ",string(body))
	resp.Body.Close()

}
func main(){
	go server()
	client()
}
