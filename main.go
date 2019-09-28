package main

import (
	"fmt"
	"html/template"
	"net/http"

	utils "github.com/jarzamendia/regionMaker/Utils"
)

func main() {

	utils.GenerateJSONFiles()

	templates := template.Must(template.ParseFiles("templates/index.html"))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if err := templates.ExecuteTemplate(w, "index.html", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(":8080", nil))

}
