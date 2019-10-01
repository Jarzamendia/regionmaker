package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/jarzamendia/regionmaker/utils"
)

//Config Config struct
type Config struct {
	Region string
	File   string
}

func main() {

	var prod = Config{

		Region: "Produção",
		File:   "static/prod_data.json",
	}

	var hom = Config{

		Region: "Homologação",
		File:   "static/hom_data.json",
	}

	var des = Config{

		Region: "Desenvolvimento",
		File:   "static/des_data.json",
	}

	var sup = Config{

		Region: "Suporte",
		File:   "static/sup_data.json",
	}

	var cms = Config{

		Region: "CMS",
		File:   "static/cms_data.json",
	}

	var stream = Config{

		Region: "Stream",
		File:   "static/stream_data.json",
	}

	templates := template.Must(template.ParseGlob("templates/*.html"))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if err := templates.ExecuteTemplate(w, "prod.html", prod); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	http.HandleFunc("/hom", func(w http.ResponseWriter, r *http.Request) {

		if err := templates.ExecuteTemplate(w, "hom.html", hom); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	http.HandleFunc("/des", func(w http.ResponseWriter, r *http.Request) {

		if err := templates.ExecuteTemplate(w, "des.html", des); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	http.HandleFunc("/sup", func(w http.ResponseWriter, r *http.Request) {

		if err := templates.ExecuteTemplate(w, "sup.html", sup); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	http.HandleFunc("/cms", func(w http.ResponseWriter, r *http.Request) {

		if err := templates.ExecuteTemplate(w, "cms.html", cms); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	http.HandleFunc("/stream", func(w http.ResponseWriter, r *http.Request) {

		if err := templates.ExecuteTemplate(w, "stream.html", stream); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	http.HandleFunc("/generate", func(w http.ResponseWriter, r *http.Request) {

		utils.GenerateFiles()
	})

	fmt.Println("Listening")

	fmt.Println(http.ListenAndServe(":8080", nil))

}
