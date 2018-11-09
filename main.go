package main

import (
	"bufio"
	"fmt"
	"guess-the-word/game"
	"guess-the-word/generator"
	"guess-the-word/random"
	"guess-the-word/user"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome!")
	fmt.Println("Please select login or sign up to continue:")
	fmt.Println("1- login")
	fmt.Println("2- sign up")
	fmt.Println("or anything else to exit")
	fmt.Print("Enter text: ")

	user.Setup()
	reader := bufio.NewReader(os.Stdin)

	// take input
	selection, _ := reader.ReadString('\n')

	//clean input => in go there is two extra characters taken with human input
	selection = strings.TrimSuffix(selection, "\n")
	selection = strings.TrimSuffix(selection, string(13))

	uid := "Failed"

	switch selection {
	case "1":
		uid = user.Login()
		break
	case "2":
		uid = user.SignUp()
		break
	default:
		fmt.Println("Thank you for using our system")
	}

	if uid == "Failed" {
		fmt.Println("We failed to obtain your user details, please try again later")
		fmt.Println("Thank you for using our system")
		return
	}

	start(uid)
}

func start(uid string) {
	// display message
	fmt.Println("Choose type of words to guess from:")
	fmt.Println("1- Colors")
	fmt.Println("2- Directions")
	fmt.Println("3- Names")
	fmt.Print("Enter text: ")

	// take input
	reader := bufio.NewReader(os.Stdin)
	guessGame, _ := reader.ReadString('\n')

	//clean input => in go there is two extra characters taken with human input
	guessGame = strings.TrimSuffix(guessGame, "\n")
	guessGame = strings.TrimSuffix(guessGame, string(13))

	words := generator.Generate(guessGame)
	if len(words) == 0 {
		fmt.Println("We failed to find the game you have requested, please try again later")
		fmt.Println("Thank you for using our system")
		return
	}

	shuffled := random.Shuffle(words)
	score := game.Start(words, shuffled)

	fmt.Print("\n\n")
	fmt.Println("Great you guessed it all!")
	fmt.Printf("Your score is %%%.2f\n", score)
	user.AddScore(uid, score)
	total := user.GetTotalScore(uid)
	fmt.Printf("Your total score is %%%.2f\n", total)
}
