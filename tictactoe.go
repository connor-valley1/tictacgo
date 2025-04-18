package main

import "fmt"
import "bufio"
import "os"
import "strconv"
import "math/rand"
import "time"

func updateBoard(board []string, idx int, turn int) {
	if (turn == 0) {
		board[idx] = "O"
	} else {
		board[idx] = "X"
	}
}

func printBoard(board []string) {
	for i := 0; i < 9; i +=3 {
		fmt.Println("", board[i], "|", board[i+1], "|", board[i+2], "")
		if (i < 6) {
			fmt.Println("---|---|---")
		}
	}
}

// need check win function
func checkGameOver(board []string) bool {
	for i := 0; i < 9; i += 3 {
		if (board[i] == board[i+1] && board[i] == board[i+2]) {
			if (board[i] == "X") {
				fmt.Println("Computer Wins! Try again next time!")
				return true
			}
			if (board[i] == "O") {
				fmt.Println("You Win! Well Played!")
				return true
			}
		}
	}
	for i := 0; i < 3; i += 1 {
		if (board[i] == board[i+3] && board[i+6] == board[i]) {
			if (board[i] == "X") {
				fmt.Println("Computer Wins! Try again next time!")
				return true
			}
			if (board[i] == "O") {
				fmt.Println("You Win! Well Played!")
				return true
			}
		}
	}
	if ((board[0] == board[4] && board[0] == board[8]) || (board[2] == board[4] && board[2] == board[6])) {
		if (board[4] == "X") {
				fmt.Println("Computer Wins! Try again next time!")
				return true
			}
			if (board[4] == "O") {
				fmt.Println("You Win! Well Played!")
				return true
			}
	}

	for i := 0; i < 9; i += 1 {
		if (board[i] == " ") {
			return false
		}
	}
	fmt.Println("Cats Game, better luck next time!")
	return true
}

// need computer player (random plays)
func computerPlay(board []string) {
	var indexes []int
	for i := 0; i < 9; i += 1 {
		if (board[i] == " ") {
			indexes = append(indexes, i)
		}
	}
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(len(indexes) - 1)
	updateBoard(board, indexes[num], 1)
}

func main() {
	
	// Initialize a string array for holding the board
	board := []string{" ", " ", " ", " ", " ", " ", " ", " ", " "}


	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Welcome to Tic-Tac-'Go'\n")
	fmt.Printf("Tic-Tac-Toe Game created using golang\n")

	turn := 0

	for {
		printBoard(board)
		over := checkGameOver(board)
		if (over) {
			break
		}
		if (turn == 1) {
			fmt.Println("Computer Plays:")
			computerPlay(board)
			turn = 1 - turn
			continue
		}
		fmt.Printf("Player, choose a cell (0-8): ")
		scanner.Scan()
		input, err := strconv.Atoi(scanner.Text())
		if (err != nil) {
			fmt.Println("Bad Input")
			continue
		}
		if (board[input] != " ") {
			fmt.Println("Pick different space!")
			continue
		}
		updateBoard(board, input, turn)
		turn = 1 - turn
	}
}