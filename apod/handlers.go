package apod

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

const (
	layoutISO = "2006-01-02"
	weblayout = "21/02/2006"
)

func ApodHandler(writer http.ResponseWriter, request *http.Request) {
	var date string
	currentTime := time.Now()
	switch request.Method {
	case "GET":
		{
			date = currentTime.Format(layoutISO)
		}
	case "POST":
		{
			date = request.FormValue("date")
		}
	}
	jresp := LookUpAPOD(date)
	page := template.Must(template.ParseFiles("templates/apod.html"))
	page.Execute(writer, jresp)

}

func logHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s:%s:%s:%s\n", r.RemoteAddr, r.UserAgent(), r.Method, r.URL.Path)
		handler.ServeHTTP(w, r)
	})
}

func StartApod() {

	mux := http.NewServeMux()
	rh := http.RedirectHandler("https://timesofindia.indiatimes.com", 307)
	mux.HandleFunc("/apod", ApodHandler)
	mux.Handle("/news", rh)
	log.Println("Listening in port 9090. . .")
	err := http.ListenAndServeTLS(":9090", "../certs/apod.crt", "../certs/apod.key", logHandler(mux))
	if err != nil {
		log.Fatal("ListenAndServeTLS: ", err)
	}
}