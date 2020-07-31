package main

import (

	"net/http"
	"html/template"
	"time"
	"log"
	//"fmt"
)

const (

	layoutISO = "2006-01-02"
	weblayout = "21/02/2006"
)


func ApodHandler(writer http.ResponseWriter, request *http.Request) {
	var date string
	currentTime := time.Now()
	switch request.Method{
	case "GET": {
		date = currentTime.Format(layoutISO)
		}
	case "POST": {
		date = request.FormValue("date")
		}
	}
	jresp := LookUpAPOD(date)
	page := template.Must(template.ParseFiles("templates/apod.html"))
	page.Execute(writer, jresp)
	

}

func main() {

	rh := http.RedirectHandler("https://www.ndtv.com", 307)
	http.HandleFunc("/apod", ApodHandler)
	log.Println("Listening in port 9090. . .")

	http.ListenAndServe(":9090",nil)
}

