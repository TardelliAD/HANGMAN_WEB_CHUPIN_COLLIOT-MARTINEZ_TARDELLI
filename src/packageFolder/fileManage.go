package packageFolder

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func ReadFile() []string {
	file, err := os.Open("./packageFolder/words.txt") // Ouverture du fichier text

	if err != nil { // si erreur d'ouverture
		log.Fatal(err) // renvoie l'erreur
	}

	scanner := bufio.NewScanner(file) // Nouvelle variable scanner

	var text []string // Tableau de string pour stocker notre text

	for scanner.Scan() { // Scan les lignes une par une
		text = append(text, scanner.Text()) // Ajout au tableau text le contenue ligne par ligne de notre fichier
	}
	file.Close() // Fermeture du fichier

	return text
}

func WordChoice(text []string) string {
	rand.Seed(time.Now().UnixNano()) // Nombre aléatoire
	randomWord := rand.Intn(len(text))

	return text[randomWord]
}

func PrintWord() string {
	word := WordChoice(ReadFile()) //Tire un mot aléatoire
	fmt.Println(word)
	secretWordTab := make([]byte, len(word)) //Crée un tableau vide pour le mot à trous
	var secretWord string                    //String du secretWord

	numLetter := len(word)/2 - 1 //Initialise la limite de mot à découvrir

	for i := numLetter; i > 0; i-- {
		rand.Seed(time.Now().UnixNano())                 // Nombre aléatoire
		time.Sleep(5 * time.Nanosecond)                  //Sleep pour pas tomber sur le même nombre trop vite
		randomLetter := rand.Intn(len(word))             //Nombre aléatoire entre 0 et la len(word)
		secretWordTab[randomLetter] = word[randomLetter] //Remplace la lettre vide de secretWordTab à l'index du nombre aléatoire, par la lettre de word à l'index du nombre aléatoire
	}

	// For qui permet de remplacer les lettre vide par des underscores et de transformer le tableau de byte en string
	for _, letter := range secretWordTab {
		if letter == 0 {
			secretWord += "_"
		} else {
			secretWord += string(letter)
		}
	}

	return secretWord
}
