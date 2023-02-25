package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
)

func main() {
	tmpl := template.Must(template.ParseFiles("tableau.html"))

	http.HandleFunc("/", tableauWeb)
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
	tab := tableau()
	fmt.Println(tab)
}
