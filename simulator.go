package main

import (
	"math/rand"
)

func randBoard(board [][]int) {
	numRow := len(board)
    numCol := len(board[0])

	for i := 0; i<numRow; i++ {
		for j := 0; j<numCol; j++ {
			p := rand.Float64()
			if p < 0.15 {
				board[i][j] = 1
			} else {
				board[i][j] = 0
			}
		}
	}
}

func gameOfLife(board [][]int) {
    numRow := len(board)
    numCol := len(board[0])
	
    toModify := [][2]int{}
    directions := [][]int{{1, 0}, {0,1}, {-1, 0}, {0, -1}, {1, 1}, {-1, -1}, {-1, 1}, {1, -1}}
    for row := 0; row<numRow;row++ {
        for col := 0; col < numCol;col++ {
            liveN := 0
            for _, coords := range directions {
                nr := row + coords[0]
                nc := col + coords[1]
            
                if nr >= numRow || nc >= numCol || nr < 0 || nc < 0 {
                    continue
                }
                if board[nr][nc] == 1 {
                    liveN += 1
                }
            }
            
            switch {
                case liveN == 3 && board[row][col] == 0:
                    toModify = append(toModify, [2]int{row,col})
                case liveN < 2 && board[row][col] == 1:
                    toModify = append(toModify, [2]int{row,col})
                case liveN > 3 && board[row][col] == 1:
                    toModify = append(toModify, [2]int{row,col})
            }
            
        }
    }
    
    for _, coords := range toModify {
        r, c := coords[0], coords[1]
        if board[r][c] == 0 {
            board[r][c] = 1
        } else {
            board[r][c] = 0
        }
    }
}