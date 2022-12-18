package game

import (
	"strings"
)

type Hangman struct {
	Player          *Player
	WordFile        string
	Word            string
	WordHidden      string
	MapWord         map[string]bool
	LettersUsed     string
	MessageToPrint  string
	Attempts        int
	CurrentAttempts int
	Win             bool
	WordInit        bool
	Ft              bool
}
type Player struct {
	Name       string
	Difficulty string
	Win        int
	Lose       int
}

func (h *Hangman) CheckInputGiven(input string) {
	//Récupère l'input, si mauvais input, erreur.
	//Parcours la map des lettres découvertes, si input == key de la map, alors la valeur de clé = true
	//Si après parcours de la map pas de lettre sont trouvé, alors message d'échec
	input = ToUpper(input)
	if !h.Ft {
		h.MessageToPrint = "Entrez une lettre ou un mot"
		h.Ft = true
		h.CurrentAttempts = 11
		return
	}

	if !isAlpha(input) {
		h.MessageToPrint = "Caractère non valide"
		return
	}

	if len(input) > 1 {
		if input == h.Word {
			for k := range h.MapWord {
				h.MapWord[k] = true
			}
		} else {
			h.Attempts += 2
			h.CurrentAttempts -= 2
		}
	} else {
		if strings.Contains(h.LettersUsed, input) {
			h.MessageToPrint = "Cette lettre a déjà été utilisée"
			return
		} else {
			for k := range h.MapWord {
				if k == input {
					h.MapWord[k] = true
					h.LettersUsed += input
					h.MessageToPrint = "Bonne lettre !"
					return
				}
			}
			h.LettersUsed += input
			h.Attempts++
			h.CurrentAttempts--
			h.MessageToPrint = "Mauvaise lettre"
			return
		}
	}
}

func CheckIsFound(mapWordToFind map[string]bool) bool {
	//Vérifie si toutes les lettres du mot sont trouvées
	for key := range mapWordToFind {
		if !mapWordToFind[key] {
			return false
		}
	}
	return true
}

func CreateHidden(word string, listLetter map[string]bool) string {
	//Parcours le mot, vérifie dans la map si la lettre est trouvée. Si oui, print la lettre. Si non, print "_".
	var s string
	for _, letter := range word {
		tmp := "_"
		if listLetter[string(letter)] {
			tmp = string(letter)
		}
		s = s + tmp
	}
	return s
}

func (h *Hangman) RebootHangman() {
	h.WordInit = false
	h.Attempts = 0
	h.CurrentAttempts = 11
	h.LettersUsed = ""
	h.Ft = false
	h.Win = false
}
