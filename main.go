package main

import (
	"fmt"
	"math/rand"
	"os"
	"unicode"
)

var dictionary = []string{
	"Zombie",
	"Gopher",
	"United States of America",
	"Indonesia",
	"Nazism",
	"Apple",
	"Programming",
	"Python",
}

func main() {
	// Derive a word we have to guess
	targetWord := getRandomWord()
	fmt.Println(targetWord)

	// Printing game state
	guessedLetters := initializedGuessedWords(targetWord)
	hangmanState := 0
	printGameState(targetWord, guessedLetters, hangmanState)

}

func initializedGuessedWords(targetWord string) map[rune]bool {
	guessedLetters := map[rune]bool{}
	guessedLetters[unicode.ToLower(rune(targetWord[0]))] = true
	guessedLetters[unicode.ToLower(rune(targetWord[len(targetWord)-1]))] = true
	return guessedLetters
}

func getRandomWord() string {
	return dictionary[rand.Intn(len(dictionary))]
}

func printGameState(targetWord string, guessedLetters map[rune]bool, hangmanState int) {
	fmt.Println(getWordGuessingProgress(targetWord, guessedLetters))
	fmt.Println()
	fmt.Println(getHangmanDrawing(hangmanState))
	fmt.Println()
}

func getWordGuessingProgress(targetWord string, guessedLetters map[rune]bool) string {
	result := ""
	for _, word := range targetWord {
		if guessedLetters[unicode.ToLower(word)] || word == ' ' {
			result += fmt.Sprintf("%c", word)
		} else {
			result += "_"
		}
		result += " "
	}
	return result
}

func getHangmanDrawing(hangmanState int) string {
	data, err := os.ReadFile(fmt.Sprintf("states/hangman%d", hangmanState))
	if err != nil {
		panic(err)
	}
	return string(data)
}
