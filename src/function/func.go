package Hangman

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"slices"
	"strings"
)

var Word string                //mot a deviner
var TabUnder []string          //tableau d'underscore
var win bool                   //verif si on a win (si y'a plus d'underscore)
var LetterGuessedList []string //liste des lettres déjà rentrer
var WordGuessedList []string   //liste des mots déjà rentrer
var Graph int                  //compteur de point pour le graph

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
		ClearTerminal()
		fmt.Println("Exiting...")
		return
	default:
		fmt.Println("Pas compris")
		Menu()
	}
}

// NewGame Menu
func NewGame() {
	ClearTerminal()
	Word = ""
	TabUnder = []string{}
	win = false
	LetterGuessedList = []string{}
	WordGuessedList = []string{}
	Graph = 0
	var input string
	fmt.Println("Choisissez une difficulté...")
	fmt.Println("1. Facile")
	fmt.Println("2. Difficile")
	fmt.Print("Votre choix: ")
	fmt.Scan(&input)
	switch input {
	case "1":
		ClearTerminal()
		ShowTextFromFile("src/wordList/wordList_Easy.txt")
	case "2":
		ClearTerminal()
		ShowTextFromFile("src/wordList/wordList_Hard.txt")
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

// Tableau d'underscore correspondant au mot à deviner
// Params : word = mot choisis dans la liste de mot
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

// Affichage de base
func Display() {
	DisplayHangman()
	fmt.Print("Voici le mot à deviner : ")
	for _, i := range TabUnder {
		fmt.Print(i, " ")
	}
	fmt.Println("")
	fmt.Println("Lettres déjà tenté : ", LetterGuessedList)
	fmt.Println("")
	fmt.Println("Mots déjà tenté : ", WordGuessedList)
	fmt.Println("")
	Guess()
}

// Menu pour taper votre lettre/mot
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

// Check si le guess est dans le mot
func GuessLetter(letter string) {
	if !IsLetterInGuessed(letter) {
		LetterGuessedList = append(LetterGuessedList, strings.ToLower(letter))
		if IsInWord(letter) {
			ChangeTableau(letter)
			if !CheckWin() {
				ClearTerminal()
				Display()
			}
		} else {
			ClearTerminal()
			fmt.Println("Mauvaise réponse !")
			fmt.Println("")
			fmt.Println("")
			Graph += 1
			CheckLoose()
		}
	} else {
		ClearTerminal()
		fmt.Println("Vous avez déjà essayé cette lettre!")
		fmt.Println("")
		fmt.Println("")
		Display()
	}
}

// Check si la lettre est dans le mot
// Params: letter = lettre en input
func IsInWord(letter string) bool {
	for _, v := range Word {
		if strings.EqualFold(letter, string(v)) {
			return true
		}
	}
	return false
}

// Check si le guess est le mot
// Params: word = mot en input
func GuessWord(word string) {
	if !IsWordInGuessed(word) {
		WordGuessedList = append(WordGuessedList, word)
		if strings.EqualFold(word, Word) {
			win = true
			Win()
		} else {
			ClearTerminal()
			fmt.Println("Mauvaise réponse !")
			fmt.Println("")
			fmt.Println("")
			Graph += 2
			CheckLoose()
		}
	} else {
		ClearTerminal()
		fmt.Println("Vous avez déjà essayé ce mot!")
		fmt.Println("")
		fmt.Println("")
		Display()
	}
}

// Modifie le tableau d'underscore en rajoutant la lettre trouvées
// Params: guess = lettre en input
func ChangeTableau(guess string) {
	for i, v := range Word {
		if string(v+32) == guess || string(v-32) == guess || string(v) == guess {
			TabUnder[i] = string(v)
		}
	}
}

// Vérifie si la game est gagné
func CheckWin() bool {
	if slices.Contains(TabUnder, "_") {
		return false
	} else {
		win = true
		Win()
		return true
	}
}

// Affichage di win et possibilité de relancer une partie
func Win() {
	if win {
		ClearTerminal()
		fmt.Println("Vous avez gagné!")
		fmt.Println("Le mot était :", Word)
		fmt.Println(("-------------"))
		fmt.Println("")
		fmt.Println("Refaire une partie ?")
		fmt.Println("")
		fmt.Println("1: Oui")
		fmt.Println("2: Non")
		fmt.Println("")
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

// Vérifie si la game est perdue
func CheckLoose() {
	if Graph >= 10 {
		ClearTerminal()
		DisplayHangman()
		fmt.Println("Vous avez perdu!")
		fmt.Println("")
		fmt.Println("Le mot était :", Word)
		fmt.Println("")
		fmt.Println("")
		fmt.Println("Refaire une partie?")
		fmt.Println("")
		fmt.Println("1: Oui")
		fmt.Println("2: Non")
		fmt.Println("")
		var input string
		fmt.Print("Votre choix: ")
		fmt.Scan(&input)
		switch input {
		case "1":
			NewGame()
		case "2":
			ClearTerminal()
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

// Vérifie si le mot est déjà dans la liste des mots déjà rentrés
// Params: w = mot en input
func IsWordInGuessed(w string) bool {
	for _, v := range WordGuessedList {
		if strings.EqualFold(v, w) {
			return true
		}
	}
	return false
}

// Vérifie si la lettre est déjà dans la liste des lettres déjà rentrés
// Params: l = lettre en input
func IsLetterInGuessed(l string) bool {
	for _, v := range LetterGuessedList {
		if strings.EqualFold(v, l) {
			return true
		}
	}
	return false
}

// Affichage du pendu
func DisplayHangman() {
	if Graph == 0 {
		Graph = 1
	}
	if Graph >= 11 {
		Graph = 10
	}
	file, err := os.Open("src/GraphHangman/hangman.txt")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	defer file.Close()
	var lines []string
	for i := 0; i <= 9; i++ {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
	}
	set := Graph * 10
	for i := 10 * (Graph - 1); i < set; i++ {
		fmt.Println(lines[i])
	}
}

// Définie le nombre de lettres bonus en fonction de la taille du mot
func NbrRandom() {
	bonusLetters := len(Word) / 5
	RandomLetters(bonusLetters)
}

// Définis aléatoirement les lettres bonus, en évitant les doublons et les caractère spéciaux données d'office
// Params: nbr = nombre de lettres
func RandomLetters(nbr int) {
	for i := 0; i < nbr; i++ {
		id := rand.Intn(len(Word))
		if IsLetterInGuessed(string(Word[id])) {
			i--
			continue
		}
		if string(Word[id]) == "-" {
			i--
			continue
		}
		ChangeTableau(string(Word[id]))
		LetterGuessedList = append(LetterGuessedList, string(Word[id]))
	}

}

// Ouvre un fichier .txt et choisis aléatoirement un mot à l'intérieur
// Params: path = chemin du fichier .txt
func ShowTextFromFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Une erreur s'est produite lors de la recherche du fichier")
		fmt.Println("")
		fmt.Println("")
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	lines := make([]string, 0)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
	randomIndex := rand.Intn(len(lines))
	Word = lines[randomIndex]
	Underscore(Word)
	NbrRandom()
	Display()
}
