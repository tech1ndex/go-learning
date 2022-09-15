package main

import (
	"html/template"
	"net/http"
)

type TodoPageData struct {
	PageTitle string
}

func main() {
	tmpl := template.Must(template.ParseFiles("index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "ACE Quiz",
		}
		tmpl.Execute(w, data)
	})

	//Serve Static CSS
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	//Start Web Server on Port 8080
	http.ListenAndServe(":8080", nil)
}
