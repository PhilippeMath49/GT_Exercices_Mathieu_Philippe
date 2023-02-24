package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Choisissez quel est votre choix:")
	fmt.Println("1.Récupérer le contenu du .txt")
	fmt.Println("2.Ajouter du texte au .txt")
	fmt.Println("3.Supprimer le contenu")
	fmt.Println("4.Remplacer le contenu")
	fmt.Println("5.Quitter")
	//récupération de l'entrée utilisateur
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	choix, _ := strconv.Atoi(scanner.Text())
	switch choix {
	case 1:
		recuperer()
	case 2:
		ajouter()
	case 3:
		supprimer()
		fmt.Println("Votre texte est supprimé")
		retourMenu()
	case 4:
		Remplacer()
	case 5:
		os.Exit(4)

	default:
		fmt.Println("Veuillez entrer un nombre entre 1 et 5.")
	}
}

func recuperer() {
	data, err := ioutil.ReadFile("fichiers.txt") // lire le fichier text.txt
	//gestion d'erreur
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data)) //convertir les bytes en string et afficher le contenu
	main()                    //retourner au menu

}

func ajouter() {
	file, err := os.OpenFile("fichiers.txt", os.O_CREATE|os.O_APPEND, 0600)
	defer file.Close()

	if err != nil {
		panic(err)
	}
	fmt.Println("veuillez écrire ce que vous voulez ajouter.")
	//récuperer l'entrée utilisateur
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	texte := scanner.Text()
	_, err = file.WriteString(texte) // écrire dans le fichier
	if err != nil {
		panic(err)
	}
	fmt.Println("Votre texte est ajouté")
	main() //retourner au menu
}

func supprimer() {
	data, err := ioutil.ReadFile("fichiers.txt") // lire le fichier text.txt
	if err != nil {
		panic(err)
	}
	file, err := os.OpenFile("fichiers.txt", os.O_CREATE, 0600)
	defer file.Close()

	if err != nil {
		panic(err)
	}
	for i := 0; i < len(data); i++ {
		_, err = file.WriteString(" ") // écrire dans le fichier
		if err != nil {
			panic(err)
		}
	}

}

func Remplacer() {
	supprimer()
	file, err := os.OpenFile("fichiers.txt", os.O_CREATE, 0600)
	defer file.Close()

	if err != nil {
		panic(err)
	}
	fmt.Println("veuillez écrire le nouveau texte.")
	//récuperer l'entrée utilisateur
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	texte := scanner.Text()
	_, err = file.WriteString(texte) // écrire dans le fichier
	if err != nil {
		panic(err)
	}
	fmt.Println("Votre texte est changé")
	main() //retourner au menu
}

func retourMenu() {
	main()
}
