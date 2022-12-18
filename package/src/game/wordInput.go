package game

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"
)

func ReturnListWord(wordsfile string) []string {
	// wordsfile est le fichier que l'on va vouloir ouvrir grâce au flag
	file, err := os.Open("package/doc/" + wordsfile + ".txt")

	//Ouvre le fichier words.txt et crée un tableau avec tous les mots du fichier

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var listWord []string

	for scanner.Scan() {
		listWord = append(listWord, scanner.Text())
	}

	file.Close()

	return listWord
}

func WordChoice(listWord []string) (string, map[string]bool) {
	//Prends la liste des mots et renvois un mot aléatoire et une map des lettres du mot choisis

	mapLetter := make(map[string]bool)

	//Choisis un mot au hasard
	rand.Seed(time.Now().UnixNano())
	randomWord := listWord[rand.Intn(len(listWord))]
	randomWord = ToUpper(randomWord)

	for _, letter := range randomWord {
		mapLetter[string(letter)] = false
	}

	discoveredLetter := len(randomWord)/2 - 1

	for key := range mapLetter {
		mapLetter[key] = true
		discoveredLetter--
		if discoveredLetter <= 0 {
			break
		}
	}

	return randomWord, mapLetter
}

func ToUpper(s string) string {
	word := ""
	for _, letter := range s {
		if letter >= 'a' && letter <= 'z' {
			letter -= 32
			word += string(letter)
		} else {
			word += string(letter)
		}
	}
	return word
}

func isAlpha(s string) bool {
	if s == "" {
		return true
	}
	for _, letter := range s {
		if !(letter >= 'a' && letter <= 'z' || letter >= 'A' && letter <= 'Z') {
			return false
		}
	}
	return true
}
