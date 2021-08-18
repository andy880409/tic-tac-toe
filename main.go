package main

import (
	"fmt"
	"log"
)

type game struct {
	board  [9]string
	player string
}

func main() {
	var newGame game
	printBoard(newGame.board)
	posi, err := inputPosition()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(posi)

}
func inputPosition() (int, error) {
	var posi int
	fmt.Print("Enter a position from 1 to 9:")
	_, err := fmt.Scan(&posi)
	if err != nil {
		return posi, nil
	}
	return posi, nil
}
func printBoard(board [9]string) {
	for i, v := range board {
		if v == "" {
			fmt.Print(" ")
		} else {
			fmt.Print(v)
		}
		if (i+1)%3 == 0 { //2,5,8 position need a newline
			fmt.Print("\n")
		} else {
			fmt.Print("|")
		}
	}
}

// func clearScreen() {
// 	cmd := exec.Command("clear")
// 	cmd.Stdout = os.Stdout
// 	cmd.Run()
// }
