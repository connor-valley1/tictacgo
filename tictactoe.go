package main

import "fmt"
import "bufio"
import "os"

func updateBoard(board []string, idx int, turn int) {
	if (turn == 0) {
		board[idx] = "O"
	} else {
		board[idx] = "X"
	}
}

func printBoard(board []string) {
	for i := 0; i < 9; i +=3 {
		fmt.Println(board[i], "|", board[i+1], "|", board[i+2])
	}

}

func main() {
	board := []string{"", "", "", "", "", "", "", "", ""}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Welcome to Tic-Tac-'Go'\n")
	fmt.Printf("Tic-Tac-Toe Game created using golang\n")

	turn := 0

	for {
		printBoard(board)
		fmt.Printf("Player, choose a cell (0-8): ")
		scanner.Scan()
		input := scanner.Text()
		fmt.Printf("You choose %s\n", input)
		turn = 1 - turn
	}
}