package main

import (
	"html/template"
	"math/rand"
	"net/http"
)

type Tab struct {
	Tableau []int
}

func main() {

	http.HandleFunc("/", tableauWeb)
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	http.ListenAndServe(":80", nil)
}

func tableau() []int {
	//on genere le nombre de dimmension du tableau, avec minimum 3 et maximum 17
	n := rand.Intn(15) + 3
	tab := make([]int, n)
	//on remplis chaque élément du tableau avec un int aléatoire
	for i := 0; i < n; i++ {
		tab[i] = rand.Int()
	}
	return tab
}

func tableauWeb(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("tableau.html"))
	data := Tab{
		Tableau: tableau(),
	}
	tmpl.Execute(w, data)

}
