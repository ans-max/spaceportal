package apod

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	layoutISO = "2006-01-02"
	weblayout = "21/02/2006"
	scert     = "/app/certs/apod.crt"
	skey      = "/app/certs/apod.key"
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
	page := template.Must(template.ParseFiles("apod/templates/apod.html"))
	page.Execute(writer, jresp)

}

func logHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s:%s:%s:%s\n", r.RemoteAddr, r.UserAgent(), r.Method, r.URL.Path)
		handler.ServeHTTP(w, r)
	})
}

func shutdown(writer http.ResponseWriter, req *http.Request) {
	os.Exit(0)
}

func StartApod(port string) {

	mux := http.NewServeMux()
	mux.HandleFunc("/apod", ApodHandler)
	mux.HandleFunc("/shutdown", shutdown)
	log.Printf("Listening in port %s. . .", port)
	err := http.ListenAndServeTLS(port, scert, skey, logHandler(mux))
	if err != nil {
		log.Fatal("ListenAndServeTLS: ", err)
	}
}
