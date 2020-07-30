package main

import "fmt"

// Maximal square
// O (n * m)

func maximalSquare(matrix [][]byte) int {
	fmt.Println(matrix)
	if len(matrix) == 0 {
		return 0
	}

	columns := len(matrix[0])
	rows := len(matrix)

	// run time requires make, bc len is unknown
	dp := make([][]int, rows+1)
	for i := range dp {
		dp[i] = make([]int, columns+1)
	}

	maxlen := 0

	for i := 1; i <= rows; i++ {
		for j := 1; j <= columns; j++ {
			if matrix[i-1][j-1] == '1' {
				// notice this is the min of min
				dp[i][j] = min(min(dp[i][j-1], dp[i-1][j]), dp[i-1][j-1]) + 1
				maxlen = max(maxlen, dp[i][j])
			}
		}
	}

	fmt.Println(dp)

	return maxlen * maxlen
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
