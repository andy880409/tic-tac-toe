package main

import (
	"errors"
	"fmt"
)

type Game struct {
	board  [9]string
	player string
	turn   int
}

func main() {
	var game Game
	gameOver := false
	game.player = "X"
	var winner string
	for !gameOver {
		printBoard(game.board)
		posi := inputPosition()
		err := game.playGame(posi)
		if err != nil {
			fmt.Println(err)
			continue //skip next step
		}
		gameOver, winner = game.checkGameWinner()
	}
	printBoard(game.board)
	if winner == "" {
		fmt.Println("It`s a tie.")
	} else {
		fmt.Printf("Winner is %s.\n", winner)
	}
}

//switch player
func (game *Game) switchPlayer() {
	if game.player == "X" {
		game.player = "O"
	} else {
		game.player = "X"
	}
}

//insert posi
func (game *Game) playGame(posi int) error {
	if game.board[posi-1] != "" {
		return errors.New("Try another move.")
	}
	game.board[posi-1] = game.player
	game.turn++
	game.switchPlayer()
	return nil
}

//return whether game is over and who is winner
func (game Game) checkGameWinner() (bool, string) {
	board := game.board
	//horizontal
	for i := 0; i < 9; {
		if board[i] != "" && board[i] == board[i+1] && board[i] == board[i+2] {
			return true, board[i]
		} else {
			i += 3
		}
	}
	//vertical
	for i := 0; i < 3; i++ {
		if board[i] != "" && board[i] == board[i+3] && board[i] == board[i+6] {
			return true, board[i]
		}
	}
	//diagonal 0,4,8 and 2,4,6
	if board[0] != "" && board[0] == board[4] && board[0] == board[8] {
		return true, board[0]
	}
	if board[2] != "" && board[2] == board[4] && board[2] == board[6] {
		return true, board[0]
	}
	if game.turn == 9 {
		return true, ""
	}

	return false, ""
}

//return player input number
func inputPosition() int {
	var posi int
	fmt.Print("Enter a position from 1 to 9:")
	fmt.Scan(&posi)

	return posi
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
