36. Valid Sudoku

package main

import (
    "fmt"
)

func isValidSudoku(board [][]byte) bool {
    cols := make([]map[byte]bool, 9)
    rows := make([]map[byte]bool, 9)
    squares := make([]map[byte]bool, 9)

    for i := 0; i < 9; i++ {
        cols[i] = make(map[byte]bool)
        rows[i] = make(map[byte]bool)
        squares[i] = make(map[byte]bool)
    }

    for r := 0; r < 9; r++ {
        for c := 0; c < 9; c++ {
            if board[r][c] == '.' {
                continue
            }
            if rows[r][board[r][c]] || cols[c][board[r][c]] || squares[(r/3)*3+c/3][board[r][c]] {
                return false
            }
            rows[r][board[r][c]] = true
            cols[c][board[r][c]] = true
            squares[(r/3)*3+c/3][board[r][c]] = true
        }
    }

    return true
}

func main() {
    board := [][]byte{
        []byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
        []byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
        []byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
        []byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
        []byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
        []byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
        []byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
        []byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
        []byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
    }

    fmt.Println(isValidSudoku(board)) // Output: true
}

