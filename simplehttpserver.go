package main

import ("fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/",index_handler)
	http.HandleFunc("/about",about_handler)
	http.ListenAndServe(":8080",nil)
}

func about_handler(w http.ResponseWriter, r * http.Request){
	fmt.Fprintf(w,"Its about War!!!!")
}

func index_handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Helllo... you are here, Welcome!!!!")
}
