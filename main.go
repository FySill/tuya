package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func SetCookie(w http.ResponseWriter) {
	cookie := http.Cookie{
		Name:  "name",
		Value: "tu",
	}
	http.SetCookie(w, &cookie)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	webpage, err := ioutil.ReadFile("views/home.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("home.html file error %v", err), 500)
	}
	fmt.Fprintf(w, string(webpage))
}

func CssHandler(w http.ResponseWriter, r *http.Request) {
	SetCookie(w)
	w.Header().Set("Content-type", "stylesheet")
	err := r.ParseForm()
	if err != nil {
		http.Error(w, fmt.Sprintf("parse url error %v", err), 500)
	}
	fmt.Fprintf(w, ".......\n")
	fmt.Fprintf(w, "request.Method  %v\n", r.Method)
	fmt.Fprintf(w, "request.Method  %v\n", r.RequestURI)
	fmt.Fprintf(w, "request.Cookies  %v\n", r.Cookies())
}

func main() {
	port := 8090
	portstring := strconv.Itoa(port)

	mux := http.NewServeMux()
	mux.Handle("/views/home", http.HandlerFunc(HomeHandler))
  mux.Handle("/", http.HandlerFunc(HomeHandler))
	//mux.Handle("/static/css/", http.HandlerFunc(CssHandler))
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	log.Print("Listening on port: " + portstring + "...")
	err := http.ListenAndServe(":"+portstring, mux)
	if err != nil {
		log.Print("ListenAndServe error: ", err)
	}
}
