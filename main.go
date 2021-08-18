package main

import (
	"html/template"
	"log"
	"net/http"

)

func main() {
	fs := http.FileServer(http.Dir("./assets"))
	router := http.NewServeMux()
	router.Handle("/assets/", http.StripPrefix("/assets/", fs))
	router.HandleFunc("/home", home)
	router.HandleFunc("/processForm", processForm)
	server := http.Server{
		Addr:    "127.0.0.1:9093",
		Handler: router,
	}
	server.ListenAndServe()
}

func home(rw http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/index.html"))
	if err := tmpl.ExecuteTemplate(rw, "main", nil); err != nil {
		log.Fatal(err)
	}
}

func processForm(rw http.ResponseWriter, req *http.Request){
	http.Redirect(rw, req, "/home", http.StatusTemporaryRedirect)
}


