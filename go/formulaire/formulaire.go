package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Entree struct {
	reponse string
}

func main() {

	tmpl := template.Must(template.ParseFiles("formulaire.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}
		data := Entree{
			reponse: r.FormValue("reponse"),
		}
		tmpl.Execute(w, data)
		fmt.Println(data)
	})
	http.ListenAndServe(":80", nil)

}
