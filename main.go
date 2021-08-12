package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	drawScaffold()
	fmt.Print("Choose a position from 1 to 9:")
	reader := bufio.NewReader(os.Stdin)
	number, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	number = strings.TrimSpace(number)
	posi, err := strconv.Atoi(number)
	fmt.Print(posi)
	clearScreen()
}
func drawScaffold() {
	fmt.Println(" | | ")
	fmt.Println(" | | ")
	fmt.Println(" | | ")
}
func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
