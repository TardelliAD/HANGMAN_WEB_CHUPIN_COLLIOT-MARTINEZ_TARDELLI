package packageFolder

import (
	"fmt"
)

func InputUser() string {
	fmt.Print("enter text:")
	var text string
	fmt.Scanln(&text)
	return text

}

// func ValidInput(letter byte, word string) int {
// 	newWord := make([]byte, len(word))
// 	for i := range word {
// 		if letter == word[i] {
// 			return i + 1
// 		} else {
// 			fmt.Println("Lettre non présente")
// 		}
// 	}
// 	return 0
// }
