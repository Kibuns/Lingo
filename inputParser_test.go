package main

import (
	"testing"

	"github.com/Kibuns/Lingo/Logic"
	"github.com/Kibuns/Lingo/Models"
)

func TestCorrectWord(t *testing.T) {
	var testInput Models.ParsedInput = Logic.ParseUserInput("testword", "testword")

	if !testInput.GuessedWord {
		t.Errorf("Expected GuessedWord to be true, but it was false")
	}

	for _, l := range testInput.Letters{
		if !l.IsCorrect {
			t.Errorf("Expected all letters to be correct, but %s is not", l.Char)
		}
	}
}

func TestIncorrectWord(t *testing.T) {
	const testGuess = "wort"
	const testSecret = "word"
	var testInput Models.ParsedInput = Logic.ParseUserInput(testGuess, testSecret)

	if testInput.GuessedWord {
		t.Errorf("Expected GuessedWord to be false, but it was true")
	}

	testParser(testGuess, testSecret, "GGG_", t)
}

/*
When the secret words contains a single instance of a letter,
there should be a maximum of 1 yellow or green (present / correct) letter
in the parsedInput of that guess. the amount of yellow letters of the same letter should indicate
the amount of instances of that letter in the secret word
*/
func TestDoubleYellow(t *testing.T) {
	testParser("penne", "enima", "_YY__", t)
}

func TestSingleYellow(t *testing.T) {
	testParser("tett", "test", "GG_G", t)
}


//below are generic tests based on my wordle inputs and their output
func Test1(t *testing.T) {
	testParser("scalp", "ranch", "_YY__", t)
}

func Test2(t *testing.T) {
	testParser("canty", "ranch", "YGG__", t)
}

func Test3(t *testing.T) {
	testParser("adieu", "alone", "G__Y_", t)
}

func Test4(t *testing.T) {
	testParser("apple", "alone", "G__YG", t)
}

func Test5(t *testing.T) {
	testParser("local", "iliac", "Y_YG_", t)
}

func Test6(t *testing.T) {
	testParser("scalp", "iliac", "_YYY_", t)
}

func Test7(t *testing.T) {
	testParser("milky", "iliac", "_YY__", t)
}

//7 letters
func Test8(t *testing.T) {
	testParser("errands", "anarchy", "_Y_YY__", t)
}
func Test9(t *testing.T) {
	testParser("chantry", "anarchy", "YYGY_YG", t)
}

func Test10(t *testing.T) {
	testParser("autopsy", "clothes", "__YY_Y_", t)
}
/*
This method tests the inputParser by inputting "guess", with "secret" as the secret word.
The outcome string represents a character for the correctness of each letter after being parsed.
It should be compiled using the following syntax:
"_" => Incorrect,
"G" => Correct position,
"Y" => Present but in incorrect position,
example: GY_Y_

*/
func testParser(guess string, secret string, expectedOutcome string, t *testing.T){
	var parsedInput Models.ParsedInput = Logic.ParseUserInput(guess, secret)
	testResult := createTestResult(parsedInput.Letters)
	if testResult != expectedOutcome {
		t.Errorf("Expected %s, but got %s", expectedOutcome, testResult)
	}
}


func createTestResult(letters []Models.Letter) string {
	var testResult string = ""
	for _, l := range letters {
		if l.IsCorrect {testResult += "G"}
		if !l.IsCorrect && l.IsPresent {testResult += "Y"}
		if !l.IsCorrect && !l.IsPresent {testResult += "_"}
	}
	return testResult
}
