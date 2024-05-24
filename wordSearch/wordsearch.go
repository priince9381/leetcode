package wordsearch

import "fmt"

func CheckPairedVistedSet(r int, c int, pairedSet [][]int) bool {
	for _, vistedValue := range pairedSet {
		if vistedValue[0] == r && vistedValue[1] == c {
			return true
		}
	}
	return false

}

func wordSearch(board [][]byte, word string, r int, c int, k int, rowLength int, columLenght int, pairedSet [][]int) bool {
	if k == len(word) {
		return true
	}
	if r >= rowLength || c >= columLenght || r < 0 || c < 0 || word[k] != board[r][c] || CheckPairedVistedSet(r, c, pairedSet) {
		return false
	}
	pairedSet = append(pairedSet, []int{r, c})
	res := (wordSearch(board, word, r, c+1, k+1, rowLength, columLenght, pairedSet) ||
		wordSearch(board, word, r, c-1, k+1, rowLength, columLenght, pairedSet) ||
		wordSearch(board, word, r+1, c, k+1, rowLength, columLenght, pairedSet) ||
		wordSearch(board, word, r-1, c, k+1, rowLength, columLenght, pairedSet))
	return res
}

func exist(board [][]byte, word string) bool {
	rows := len(board)
	colums := len(board[0])
	r := 0
	for r < rows {
		c := 0
		for c < colums {
			k := 0
			if wordSearch(board, word, r, c, k, rows, colums, [][]int{}) {
				return true
			}
			c++
		}
		r++
	}
	return false
}

func Testexist() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'M', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	fmt.Println(exist(board, word))
}
