package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func main() {
	tmpl1 := template.Must(template.ParseFiles("page.html"))
	tmpl2 := template.Must(template.ParseFiles("page2.html"))
	tmpl3 := template.Must(template.ParseFiles("page3.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			errorHandler(w, r, http.StatusNotFound)
			return

		}
		if r.Method != http.MethodPost {
			tmpl1.Execute(w, nil)
			return
		}
	})
	http.HandleFunc("/page2", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/page2" {
			tmpl2.Execute(w, nil)
			return
		}
		if r.Method != http.MethodPost {
			tmpl2.Execute(w, nil)
			return
		}
	})
	http.HandleFunc("/page3", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/page3" {
			tmpl2.Execute(w, nil)
			return
		}
		if r.Method != http.MethodPost {
			tmpl3.Execute(w, nil)
			return
		}
	})

	http.ListenAndServe(":80", nil)
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "ERROR 404")
	}
}
