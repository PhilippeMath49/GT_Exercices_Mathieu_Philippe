package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("combien voulez vous d'allumette? (minimum 4)")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	for n < 4 {
		fmt.Println("entrer un entier supérieur à 3")
		scanner.Scan()
		n, _ = strconv.Atoi(scanner.Text())
	}
	tour(n)

}

func allumette(nbAllumette int) {
	var n = nbAllumette
	for i := 0; i < n; i++ {
		fmt.Print("|")
	}
	fmt.Println()
}

func tour(nbAllumette int) {
	for t := 0; nbAllumette > 0; t++ {
		allumette(nbAllumette)
		fmt.Println("Veuillez choisir un nombre d'alumette à enlever  entre 1 et 3.")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		choix, _ := strconv.Atoi(scanner.Text())
		switch choix {
		case 1:
			nbAllumette--
		case 2:
			nbAllumette -= 2
		case 3:
			nbAllumette -= 3
		default:
			fmt.Println("veuillez choisir un nombre entre 1 et 3.")
		}
	}
	fmt.Println("tu as perdu")
	rejouer()
}

func rejouer() {
	fmt.Println("Voulez vous rejouer ?")
	fmt.Println("1.Oui 2.Non")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	choix, _ := strconv.Atoi(scanner.Text())
	switch choix {
	case 1:
		main()
	case 2:
		os.Exit(4)
	default:
		fmt.Println("Veuillez entrer 1 ou 2.")
	}

}
