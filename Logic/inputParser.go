package Logic

import (
	"strings"

	"github.com/Kibuns/Lingo/Models"
)

func ParseUserInput(input string, secretWord string) Models.ParsedInput {
	var parsedInput Models.ParsedInput;
	parsedInput.Letters = []Models.Letter{}

	//keep track of a map in order to know if the algorithm has already assigned a Y of G to a letter in the guessed word
	secretWordMap := make(map[rune]int)
	for _, ch := range secretWord {
		secretWordMap[ch]++
	}

	//check for G's
	for i, ch := range input {
		var l Models.Letter;
		l.Char = string(ch)

		if secretWord[i] == input[i] {
			l.IsCorrect = true
			secretWordMap[ch]--
		} else {
			l.IsCorrect = false
		}

		addLetterToLetters(&l, &parsedInput)
	}

	//check for Y's
	for i := range parsedInput.Letters {
		l := &parsedInput.Letters[i]
		ch := rune(l.Char[0])
	
		if strings.ContainsRune(secretWord, ch) && secretWordMap[ch] > 0 { //if the secret contains the character, but the wordmap is 0, it means theres already that same letter somewhere else in the word assigned as either Y or G
			l.IsPresent = true
			secretWordMap[ch]--
		} 
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