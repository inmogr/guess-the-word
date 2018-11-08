package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Start(ori []string, shuffled []string) (score float64) {
	fmt.Println("Try to guess the correct order of the below words")
	fmt.Printf("%+v\n", ori)
	fmt.Println("Start typing a word then press enter!")

	total := 0.00
	mistakes := 0.00

	for index := 0; index < len(ori); total++ {
		// take input
		reader := bufio.NewReader(os.Stdin)
		word, _ := reader.ReadString('\n')

		//clean input => in go there is two extra characters taken with human input
		word = strings.TrimSuffix(word, "\n")
		word = strings.TrimSuffix(word, string(13))

		if word == shuffled[index] {
			index++
		} else {
			fmt.Println("Nope that was wrong, try again!")
			mistakes++
		}
	}
	score = 1.00 - (mistakes / total)
	return score
}
