package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
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
	// fmt.Println(targetWord)

	// Printing game state
	guessedLetters := initializedGuessedWords(targetWord)
	hangmanState := 0
	for !isGameOver(targetWord, guessedLetters, hangmanState) {
		printGameState(targetWord, guessedLetters, hangmanState)
		input := readInput()
		if len(input) != 1 {
			fmt.Println("Invalid input. Please use letters only....")
			continue
		}

		letter := rune(input[0])
		if isCorrectGuess(targetWord, letter) {
			guessedLetters[unicode.ToLower(letter)] = true
		} else {
			hangmanState++
		}
	}
	fmt.Println("Game Over....")
	if isWordGuessed(targetWord, guessedLetters) {
		fmt.Println("You Win!")
	} else {
		fmt.Println("You lose!")
	}
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

func isGameOver(targetWord string, guessedLetters map[rune]bool, hangmanState int) bool {
	printGameState(targetWord, guessedLetters, hangmanState)
	return isWordGuessed(targetWord, guessedLetters) || isHangmanComplete(hangmanState)
}

func isWordGuessed(targetWord string, guessedLetters map[rune]bool) bool {
	for _, word := range targetWord {
		if !guessedLetters[unicode.ToLower(word)] {
			return false
		}
	}
	return true
}

func isHangmanComplete(hangmanState int) bool {
	return hangmanState >= 9
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

func readInput() string {
	var input string
	fmt.Print("> ")
	fmt.Scanf("%s", &input)
	return strings.TrimSpace(input)
}

func isCorrectGuess(targetWord string, letter rune) bool {
	return strings.ContainsRune(targetWord, letter)
}
