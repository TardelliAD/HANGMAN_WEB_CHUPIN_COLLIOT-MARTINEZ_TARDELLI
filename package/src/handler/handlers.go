package handler

import (
	"hangman/package/src/game"
	"html/template"
	"net/http"
)

var h game.Hangman

var p game.Player

func Fs() {
	fs := http.FileServer(http.Dir("templates/"))
	http.Handle("/templates/", http.StripPrefix("/templates/", fs))
}

func Page() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			p.Name = r.FormValue("name")
		}
		pages := []string{
			"templates/layouts/index.html",
			"templates/index.tmpl.html",
		}
		tmpl := template.Must(template.ParseFiles(pages...))
		tmpl.Execute(w, p)
	})

	http.HandleFunc("/game", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {
			//Initialisation
			if r.FormValue("button") == "EASY" {
				h.WordFile = "words"
				p.Difficulty = "Facile"
			}
			if r.FormValue("button") == "MEDIUM" {
				h.WordFile = "words2"
				p.Difficulty = "Moyen"
			}
			if r.FormValue("button") == "HARD" {
				h.WordFile = "words3"
				p.Difficulty = "Difficile"
			}
			if r.FormValue("button-reboot") == "reboot" {
				h.RebootHangman()
			}
			// Si c'est la première fois qu'on lance la partie
			// Attribut un Word à trouver
			if !h.WordInit {
				h.Word, h.MapWord = game.WordChoice(game.ReturnListWord(h.WordFile))
				h.WordInit = true
			}

			//Boucle de jeu
			h.CheckInputGiven(r.FormValue("letter"))
			h.WordHidden = game.CreateHidden(h.Word, h.MapWord)
			h.Win = game.CheckIsFound(h.MapWord)
			if h.Win {
				h.MessageToPrint = "Bien joué ! Tu as gagné"
				p.Win++
			} else if h.Attempts > 10 {
				h.MessageToPrint = "Bien essayé"
				h.WordHidden = h.Word
				p.Lose++
			}
		}
		pages := []string{
			"templates/layouts/Game.html",
			"templates/game.tmpl.html",
			"templates/pendu.tmpl.html",
			// Le scoreboard est présent dans le back mais pas à jour dans le front, une prochaine mise à jour sera nécessaire pour l'affichage
			"templates/scoreboard.tmpl.html",
		}
		tmpl := template.Must(template.ParseFiles(pages...))
		tmpl.Execute(w, h)
	})
}
