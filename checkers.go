package main

import "fmt"

type Game struct {
	board       [8][8]int // board object, represented by an 8x8 matrix
	turn        int       // turn: if its equal to 0 it is white's turn(1), if its equal to 1 its black's turn (2)
	victoryFlag int       // if equal to 0, no one won yet, if equal to 1, white won (1), if equal to 2, blck won(2)
}

// REFERENCE:
// 1 --> white pieces
// 2 --> black pieces

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

func initGame(board [8][8]int) Game {
	output := Game{
		board:       board,
		turn:        0,
		victoryFlag: 0,
	}
	return output
}

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

func printBoard(board [8][8]int) {
	for i := 0; i < 8; i++ {
		fmt.Print("\t")
		for j := 0; j < 8; j++ {
			fmt.Printf("|%s|", mapPiecesForOutput(board[i][j]))
		}
		fmt.Print("\n")
	}
}

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

// trofea grill

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

func makeMove(game *Game, xPosInit int, yPosInit int, xPosEnd int, yPosEnd int) {
	fmt.Println("Haha")
}

func main() {
	board1 := initBoard()
	game1 := initGame(board1)
	for game1.victoryFlag == 0 {
		var xPosInit, yPosInit, xPosEnd, yPosEnd int
		printBoard(game1.board)
		fmt.Scanf("%1d%1d %1d%1d", &xPosInit, &yPosInit, &xPosEnd, &yPosEnd)
		fmt.Println(calculatePossibleMoves(&game1, xPosInit, yPosInit))
	}
}
