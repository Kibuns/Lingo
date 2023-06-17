package Logic

import (
	"fmt"
	"strings"

	"github.com/Kibuns/Lingo/Models"
)

func ParseUserInput(input string, secretWord string) Models.ParsedInput {
	var parsedInput Models.ParsedInput;
	parsedInput.Letters = []Models.Letter{}

	secretWordMap := make(map[rune]int)
	for _, ch := range secretWord {
		secretWordMap[ch]++
	}

	for i, ch := range input {
		var l Models.Letter;
		l.Char = string(ch)

		fmt.Println(secretWordMap)

		if secretWordMap[ch] > 0 {
			if strings.ContainsRune(secretWord, ch) {
				//secret word contains the character! but is it in the right spot?
				if secretWord[i] == input[i] {
					l.IsCorrect = true
					l.IsPresent = true
					secretWordMap[ch]--
					fmt.Println(string(ch) + " is correct")
				} else {
					l.IsCorrect = false
					l.IsPresent = true
					secretWordMap[ch]--
					fmt.Println(string(ch) + " is present")
				}
			} else{
				l.IsCorrect = false
				l.IsPresent = false
				fmt.Println(string(ch) + " is incorrect")
			}
		} else{
			l.IsCorrect = false
			l.IsPresent = false
			fmt.Println(string(ch) + " is incorrect")
		}
		addLetterToLetters(&l, &parsedInput)
	}

	parsedInput.GuessedWord = true;
	for _, l := range parsedInput.Letters {
		if(!l.IsCorrect){
			parsedInput.GuessedWord = false;
		}
	}

	return parsedInput
}

func addLetterToLetters(letter *Models.Letter, parsedInput *Models.ParsedInput) {
	parsedInput.Letters = append(parsedInput.Letters, *letter)
}