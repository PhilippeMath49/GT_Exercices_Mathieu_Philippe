package main

import (
	"net/http"
	"text/template"
)

type Data struct {
	SwitchPage2 func()
	SwitchPage3 func()
	backToMain  func()
}

func main() {
	tmpl1 := template.Must(template.ParseFiles("page.html"))
	tmpl2 := template.Must(template.ParseFiles("page2.html"))
	tmpl3 := template.Must(template.ParseFiles("page3.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl1.Execute(w, nil)
			return
		}
	})
	http.HandleFunc("/page2", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl2.Execute(w, nil)
			return
		}
	})
	http.HandleFunc("/page3", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl3.Execute(w, nil)
			return
		}

	})
	http.ListenAndServe(":80", nil)
}

func SwitchPage2(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/page2", http.StatusSeeOther)
}

func SwitchPage3(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/page3", http.StatusSeeOther)
}

func backToMain(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
