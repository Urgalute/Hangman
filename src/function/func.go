package Hangman

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var Word string       //mot a deviner
var TabUnder []string //tableau d'underscore
var Win bool          //verif si on a win (si y'a plus d'underscore)
var Guessed []string  //liste des lettres/mots déjà rentrer
var Graph int         //compteur de point pour le graph

// Start Menu
func Menu() {
	ClearTerminal()
	var input string
	fmt.Println("Welcome to Hangman!")
	fmt.Println("Choisissez une option:")
	fmt.Println("1. Nouvelle game")
	fmt.Println("2. Quitter")
	fmt.Print("Votre choix: ")
	fmt.Scan(&input)
	switch input {
	case "1":
		NewGame()
	case "2":
		fmt.Println("Exiting...")
		return
	default:
		fmt.Println("Pas compris")
		Menu()
	}
}

// NewGame Menu
func NewGame() {
	Word = ""
	TabUnder = []string{}
	Win = false
	Guessed = []string{}
	Graph = 0
	ClearTerminal()
	var input string
	fmt.Println("Choisissez une difficulté...")
	fmt.Println("1. Facile")
	fmt.Println("2. Difficile")
	fmt.Print("Votre choix: ")
	fmt.Scan(&input)
	switch input {
	case "1":
		Underscore("Après-midi")
		Display()
	case "3":
		//hangmanGame("hard")
	default:
		fmt.Println("Pas compris")
		NewGame()
	}
}

// ClearTerminal
func ClearTerminal() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func Underscore(word string) {
	Word = word
	for _, i := range word {
		if i == '-' {
			TabUnder = append(TabUnder, "-")
		} else {
			TabUnder = append(TabUnder, "_")
		}
	}
}

func Display() {
	fmt.Print("Voici le mot à deviner : ")
	for _, i := range TabUnder {
		fmt.Print(i, " ")
	}
	fmt.Println("")
	Guess()
}

func Guess() {
	var input string
	fmt.Print("Entrez une lettre ou un mot: ")
	fmt.Scan(&input)
	if len(input) == 1 {
		GuessLetter(input)
	} else {
		GuessWord(input)
	}
}

func GuessLetter(letter string) {
	if !IsInGuessed(letter) {
		Guessed = append(Guessed, letter)
		for _, v := range Word {
			if string(v) == letter {
				ChangeTableau(letter)
				Display()
			} else {
				fmt.Println("Mauvaise réponse !")
				Graph += 1
				CheckLoose()
			}
		}
	} else {
		fmt.Println("Vous avez déjà essayé cette lettre!")
		Display()
	}
}

func GuessWord(word string) {
	if !IsInGuessed(word) {
		Guessed = append(Guessed, word)
		if word == Word {
			Win = true
			CheckWin()
		} else {
			fmt.Println("Mauvaise réponse!")
			Graph += 2
			CheckLoose()
		}
	} else {
		fmt.Println("Vous avez déjà essayé ce mot!")
		Display()
	}
}

func ChangeTableau(guess string) {
	for i, v := range Word {
		if string(v) == guess {
			TabUnder[i] = guess
		}
	}
}

func CheckWin() {
	ClearTerminal()
	if Win {
		fmt.Println("Vous avez gagné!")
		fmt.Println("Le mot était :", Word)
		fmt.Println("Refaire une partie ?")
		fmt.Println("1: Oui")
		fmt.Println("2: Non")
		var input string
		fmt.Print("Votre choix: ")
		fmt.Scan(&input)
		switch input {
		case "1":
			NewGame()
		case "2":
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Pas compris")
			CheckWin()
		}
	}
}

func CheckLoose() {
	ClearTerminal()
	if Graph >= 7 {
		fmt.Println("Vous avez perdu!")
		fmt.Println("Le mot était :", Word)
		fmt.Println("Refaire une partie?")
		fmt.Println("1: Oui")
		fmt.Println("2: Non")
		var input string
		fmt.Print("Votre choix: ")
		fmt.Scan(&input)
		switch input {
		case "1":
			NewGame()
		case "2":
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Pas compris")
			CheckLoose()
		}
	} else {
		Display()
	}
}

func IsInGuessed(w string) bool {
	for _, v := range Guessed {
		if v == w {
			return true
		}
	}
	return false
}
