package main

import (
	"fmt"
	"database/sql"
	_"github.com/go-sql-driver/mysql" // to intitialize the sql driver
	//"client-go-master/third_party/forked/golang/template"
	"text/template"
	"net/http"
)

type Employee struct {
	ID int
	Name string
	City string
}

var tmpl = template.Must(template.ParseGlob("form*/"))  // tmpl is a pointer to the template obj.

func dbCOnn() (dB *sql.DB){
	dbDriver:="mysql"
	dbUser:="RA358832"
	dbPass:="Rajendra95@"
	dbName:="golang"

	dB,err:=sql.Open(dbDriver,dbUser+dbPass+"tcp@127.0.0.1:3307/"+dbName)
	if err!= nil{
		panic(err.Error())
	}
	fmt.Printf("Error ",err)
	return dB
}

func Index(w http.ResponseWriter,r *http.Request){
	dB:=dbCOnn()
	defer dB.Close()
	seldB,err:=dB.Query("SELECT * FROM EMPLOYEE Order by DESC")  //querry the fetched DB
	if err!=nil{
		panic(err.Error())
	}
	emp:=Employee{}
	result:=[]Employee{}
	for seldB.Next(){
		var id int
		var name,city string
		err:=seldB.Scan(&id,&name,&city)
		if err!=nil{
			panic(err.Error())
		}
		emp.ID=id
		emp.Name=name
		emp.City=city
		result=append(result,emp)
	}// continue the action till we have objects in the database
	tmpl.ExecuteTemplate(w,"Index",result)
}

func Show(w http.ResponseWriter,r *http.Request){
	dB:=dbCOnn()
	defer dB.Close()
	nID:=r.URL.Query().Get("id")
	selDb,err:=dB.Query("SELECT * FROM EMmplyee where id is ",nID)
	if err!= nil{
		panic(err.Error())
	}
	emp:=Employee{}
	for selDb.Next(){
		var id int
		var name,city string
		err:=selDb.Scan(&id,&name,&city)
		if err!=nil{
			panic(err.Error())
		}
		emp.ID=id
		emp.City=city
		emp.Name=name
	}
	tmpl.ExecuteTemplate(w,"Show",emp)
}

func New(w http.ResponseWriter,r *http.Request){
	tmpl.ExecuteTemplate(w,"New",nil)
}

func Edit(w http.ResponseWriter,r *http.Request){
	dB:=dbCOnn()
	defer dB.Close()
	nID:=r.URL.Query().Get("id")
	selDb,err:=dB.Query("SELECT * FROM EMmplyee where id is ",nID)
	if err!= nil{
		panic(err.Error())
	}
	emp:=Employee{}
	for selDb.Next(){
		var id int
		var name,city string
		err:=selDb.Scan(&id,&name,&city)
		if err!=nil{
			panic(err.Error())
		}
		emp.ID=id
		emp.City=city
		emp.Name=name
	}
	tmpl.ExecuteTemplate(w,"Edit",emp)
}

func Insert(w http.ResponseWriter,r *http.Request){
	dB:=dbCOnn()
	if r.Method=="POST"{
		name:=r.FormValue("name")
		city:=r.FormValue("city")
	}

}

func main(){
	a:=dbCOnn()
	fmt.Printf("DB has been successfully created here ")
	defer a.Close()
}

