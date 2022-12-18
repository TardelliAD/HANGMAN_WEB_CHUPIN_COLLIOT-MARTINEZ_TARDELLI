# HangManWeb

Le projet HangmanWeb consiste à réaliser un jeu du pendu dans une page web. 
Les consignes sont sur le repo suivant : https://github.com/Lyon-Ynov-Campus/YTrack/tree/master/subjects/hangman/hangman-web

## Langage

    Go
    Html
    Css

## Auteurs

    Andrea Taredelli
    Théo Colliot-Martinez
    Tao Chupin

## Règles

Le jeu consiste à retrouer un mot séléctionné au hazard dans un fichier 'dictionnaire', pour jouer vous devez entrer une lettre (minuscules ou majuscules) 
et retouver le mot caché pour cela vous aurez 10 vies, après une tentative de lettre vous perdrez 1 vie et si vous tentez plus de une lettre (pour retrouvez le 
mot directement) vous perdrez 2 vies jusqu'à arriver à 0 si cela arrive vous ajouterez 1 au compteur de défaite et retrouverez un personnage pendu, dans le cas 
ou vous trouvez le mot le compteur de victoir s'incrémente de 1 et vous pouvez rejouer. Les vies sont symbolisé par le pendu en construction.

## Instruction / Installation

Télécharger le dossier via github en zip. Décompressez le dossier et placer le où vous le souhaitez. Ouvir un terminal bash puis rendez-vous dans le dossier 
précédement télécharger puis dans le dossier "/cmd/web".

Puis lancer le programme avec la commande suivante : 'go run ./main.go'

Après cela le programme est héberger sur le port "8080" de votre machine, pour accéder au jeu cliquez sur le lien qui a apparu ("localhost:8080). 
Maintenant entrer un nom puis séléctionner une difficulté. Bonne Chance !!

Pour stopper l'hébergement rendez vous dans le bash dans le quel vous avez lancé le programme et pressez "CTRL+C".
