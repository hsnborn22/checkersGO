package main

import "fmt"

// Define Game struct type
type Game struct {
	board       [8][8]int // board object, represented by an 8x8 matrix
	turn        int       // turn: if its equal to 0 it is white's turn(1), if its equal to 1 its black's turn (2)
	victoryFlag int       // if equal to 0, no one won yet, if equal to 1, white won (1), if equal to 2, blck won(2)
}

// auxiliary methods:
// In this section i will put some methods that are not strictly correlated to the game logic, but that I will regardless
// use throughout the program

func abs(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}

// REFERENCE:
// 1 --> white pieces
// 2 --> black pieces

// initBoard function:
// input: void
// output: [8][8]int , i.e a 2d 8x8 matrix
// Initializes a new 8x8 matrix representing a board
func initBoard() [8][8]int {
	var output [8][8]int = [8][8]int{
		{0, 2, 0, 2, 0, 2, 0, 2},
		{2, 0, 2, 0, 2, 0, 2, 0},
		{0, 2, 0, 2, 0, 2, 0, 2},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{1, 0, 1, 0, 1, 0, 1, 0},
		{0, 1, 0, 1, 0, 1, 0, 1},
		{1, 0, 1, 0, 1, 0, 1, 0},
	}
	return output
}

// initGame function:
// input: [8][8] int 2d 8x8 matrix
// output: Game object
// Initializes a new game struct
func initGame(board [8][8]int) Game {
	output := Game{
		board:       board,
		turn:        0,
		victoryFlag: 0,
	}
	return output
}

// mapPiecesForOutput function:
// input: number of type int
// output: string
// It maps a code to a character representing it in a more user friendly terminal-based representation of the board
func mapPiecesForOutput(number int) string {
	switch number {
	case 1:
		return "● "
	case 2:
		return "○ "
	default:
		return "  "
	}
}

// printBoard function:
// input: [8][8] int (a 8x8 2d matrix)
// output: void
// Function used to print out the board to the console in a nice way

func printBoard(board [8][8]int) {
	// First loop (need 2 since it's a 2d matrix)
	for i := 0; i < 8; i++ {
		// Print a \t escape character each time we start to iterate a new row to center it in the terminal
		fmt.Print("\t")
		for j := 0; j < 8; j++ {
			// Print the square with the value inside given by a string computed by the mapPiecesForOutput function (see documentation for it above)
			fmt.Printf("|%s|", mapPiecesForOutput(board[i][j]))
		}
		// Print a \n character (new line char) each time we finish to iterate a row to go to the next line for the next row.
		fmt.Print("\n")
	}
}

// Function used to calculate the vertical jump of a piece when it needs to move based on its code (i.e if it's white or black according to the program architecture)

func verticalCalculator(code int) int {
	switch code {
	case 1:
		return -1
	case 2:
		return 1
	default:
		return 0
	}
}

// Function to calculate the codes for enemy pieces given a code

func calculateEnemyCode(code int) [2]int {
	var output [2]int
	switch code {
	case 1:
		output = [2]int{2, 4}
	case 2:
		output = [2]int{1, 3}
	default:
	}
	return output
}

// Function to check if an integers is in an integer couplet (array of 2 elements)

func numIsIn(element int, twoElArray [2]int) bool {
	if twoElArray[0] == element || twoElArray[1] == element {
		return true
	} else {
		return false
	}
}

func calculatePossibleMoves(game *Game, xPos int, yPos int) [][2]int {
	var output [][2]int
	if game.board[xPos][yPos] != 0 {
		code := game.board[xPos][yPos]
		verticalAlign := verticalCalculator(code)
		enemyCode := calculateEnemyCode(code)

		if xPos+verticalAlign < 8 && xPos+verticalAlign >= 0 && yPos+1 < 8 {
			if game.board[xPos+verticalAlign][yPos+1] == 0 {
				var position [2]int = [2]int{xPos + verticalAlign, yPos + 1}
				output = append(output, position)
			} else if numIsIn(game.board[xPos+verticalAlign][yPos+1], enemyCode) {
				if xPos+(2*verticalAlign) >= 0 && xPos+(2*verticalAlign) < 8 && yPos+2 < 8 {
					if game.board[xPos+(2*verticalAlign)][yPos+2] == 0 {
						var position [2]int = [2]int{xPos + (2 * verticalAlign), yPos + 2}
						output = append(output, position)
					}
				}
			}
		}
		if xPos+verticalAlign < 8 && xPos+verticalAlign >= 0 && yPos-1 >= 0 {
			if game.board[xPos+verticalAlign][yPos-1] == 0 {
				var position [2]int = [2]int{xPos + verticalAlign, yPos - 1}
				output = append(output, position)
			} else if numIsIn(game.board[xPos+verticalAlign][yPos-1], enemyCode) {
				if xPos+(2*verticalAlign) >= 0 && xPos+(2*verticalAlign) < 8 && yPos-2 >= 0 {
					if game.board[xPos+(2*verticalAlign)][yPos-2] == 0 && xPos+(2*verticalAlign) >= 0 && xPos+(2*verticalAlign) < 8 && yPos-2 >= 0 {
						var position [2]int = [2]int{xPos + (2 * verticalAlign), yPos - 2}
						output = append(output, position)
					}
				}
			}
		}

	}
	return output
}

func isCoupleInSlice(couple [2]int, slice [][2]int) bool {
	for _,element := range slice {
		if couple == element {
			return true 
		}
	}
	return false 
}

func makeMove(game *Game, xPosInit int, yPosInit int, xPosEnd int, yPosEnd int) {
	couple := [2]int{xPosEnd,yPosEnd}
	if isCoupleInSlice(couple,calculatePossibleMoves(game,xPosInit,yPosInit)) {
		game.board[xPosEnd][yPosEnd] = game.board[xPosInit][yPosInit]
		game.board[xPosInit][yPosInit] = 0
		verticalAlign := verticalCalculator(game.board[xPosInit][yPosInit])
		if abs(xPosEnd - xPosInit) == 2 {
			var yCond int
			switch(yPosEnd) {
				case yPosInit - 2:
					yCond = yPosInit - 1
				case yPosInit + 2:
					yCond = yPosInit + 1
				default:
					yCond = 0
			}
			game.board[xPosInit + verticalAlign][yCond] = 0
		}
		switch(game.turn) {
			case 0:
				game.turn += 1
			case 1:
				game.turn -= 1
			default:
				game.turn = 0
		}
	} else {
		fmt.Println("Move not valid, try again...")
	}
}

func main() {
	board1 := initBoard()
	game1 := initGame(board1)
	for game1.victoryFlag == 0 {
		var xPosInit, yPosInit, xPosEnd, yPosEnd int
		printBoard(game1.board)
		fmt.Scanf("%1d%1d %1d%1d", &xPosInit, &yPosInit, &xPosEnd, &yPosEnd)
		makeMove(&game1, xPosInit, yPosInit,xPosEnd,yPosEnd)
	}
}
