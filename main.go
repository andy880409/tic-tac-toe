package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const initApp = 0

var board = [15]string{" ", "|", " ", "|", " \n", " ", "|", " ", "|", " \n", " ", "|", " ", "|", " \n"}

func main() {
	drawScaffold(initApp)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Choose a position from 1 to 9:")
		number, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		number = strings.TrimSpace(number)
		posi, err := strconv.Atoi(number)
		drawScaffold(posi)
	}
	// clearScreen()
}
func drawScaffold(posi int) {
	switch posi {
	case 1:
		board[0] = "X"

	}
	for _, v := range board {
		fmt.Print(v)
	}

}

// func clearScreen() {
// 	cmd := exec.Command("clear")
// 	cmd.Stdout = os.Stdout
// 	cmd.Run()
// }
