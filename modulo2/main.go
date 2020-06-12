package main

import (
	"html/template"
	"net/http"
	"time"
)

//Create a struct that holds information to be displayed in our HTML file
type Welcome struct {
	Name string
	Time string
}

func main() {

	welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}

	templates := template.Must(template.ParseFiles("static/index.html"))
	http.Handle("/static/", http.FileServer(http.Dir(".")))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if name := r.FormValue("name"); name != "" {
			welcome.Name = name
		}

		if err := templates.ExecuteTemplate(w, "index.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	//Start the web server, set the port to listen to 8080. Without a path it assumes localhost
	//Print any errors from starting the webserver using fmt

	http.ListenAndServe(":8080", nil)
}
