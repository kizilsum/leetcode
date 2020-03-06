/*
According to the Wikipedia's article: "The Game of Life, also known simply as Life, is a cellular automaton devised by the British mathematician John Horton Conway in 1970."
Given a board with m by n cells, each cell has an initial state live (1) or dead (0).
Each cell interacts with its eight neighbors (horizontal, vertical, diagonal)
using the following four rules (taken from the above Wikipedia article):

Any live cell with fewer than two live neighbors dies, as if caused by under-population.
Any live cell with two or three live neighbors lives on to the next generation.
Any live cell with more than three live neighbors dies, as if by over-population..
Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.

Write a function to compute the next state (after one update) of the board given its current state.
The next state is created by applying the above rules simultaneously to every cell in the current state,
where births and deaths occur simultaneously.

Example:
Input:
[
  [0,1,0],
  [0,0,1],
  [1,1,1],
  [0,0,0]
]
Output:
[
  [0,0,0],
  [1,0,1],
  [0,1,1],
  [0,1,0]
]

Follow up:
Could you solve it in-place? Remember that the board needs to be updated at the same time:
You cannot update some cells first and then use their updated values to update other cells.
In this question, we represent the board using a 2D array.
In principle, the board is infinite, which would cause problems when the active area encroaches the border of the array.
How would you address these problems?
*/

package hard

func gameOfLife(board [][]int) {
	temp := make([][]int, len(board))
	for iRow, row := range board {
		temp[iRow] = make([]int, len(row))
	}

	for iRow, row := range board {
		for iCol, col := range row {
			numberOfLiveNeighbors := getNumberOfLiveNeighbors(iRow, iCol, board)
			nextState := getNextState(col, numberOfLiveNeighbors)
			temp[iRow][iCol] = nextState
		}
	}
	for iRow, row := range board {
		for iCol := range row {
			board[iRow][iCol] = temp[iRow][iCol]
		}
	}
}

func getNumberOfLiveNeighbors(rowIndex int, colIndex int, board [][]int) int {
	rowLength := len(board)
	colLength := len(board[rowIndex])
	total := 0

	if !(rowIndex-1 < 0) {
		total += board[rowIndex-1][colIndex]
		if !(colIndex-1 < 0) {
			total += board[rowIndex-1][colIndex-1]
		}
		if !(colIndex+1 > colLength-1) {
			total += board[rowIndex-1][colIndex+1]
		}
	}
	if !(rowIndex+1 > rowLength-1) {
		total += board[rowIndex+1][colIndex]
		if !(colIndex-1 < 0) {
			total += board[rowIndex+1][colIndex-1]
		}
		if !(colIndex+1 > colLength-1) {
			total += board[rowIndex+1][colIndex+1]
		}
	}
	if !(colIndex-1 < 0) {
		total += board[rowIndex][colIndex-1]
	}
	if !(colIndex+1 > colLength-1) {
		total += board[rowIndex][colIndex+1]
	}
	return total
}

func getNextState(currentState int, numberOfLiveNeighbors int) int {
	nextState := currentState
	if currentState == 1 && (numberOfLiveNeighbors < 2 || numberOfLiveNeighbors > 3) {
		nextState = 0
	} else if currentState == 0 && numberOfLiveNeighbors == 3 {
		nextState = 1
	}
	return nextState
}
