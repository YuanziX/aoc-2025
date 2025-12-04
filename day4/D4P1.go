package main

func D4P1(input [][]rune) int {
	n, m := len(input), len(input[0])
	ans := 0

	for i := range n {
		for j := range m {
			if input[i][j] != '@' {
				continue
			}

			ct := 0
			// up
			if i > 0 && input[i-1][j] == '@' {
				ct++
			}
			// down
			if i < n-1 && input[i+1][j] == '@' {
				ct++
			}
			// left
			if j > 0 && input[i][j-1] == '@' {
				ct++
			}
			// right
			if j < m-1 && input[i][j+1] == '@' {
				ct++
			}
			// upper left
			if i > 0 && j > 0 && input[i-1][j-1] == '@' {
				ct++
			}
			// upper right
			if i > 0 && j < m-1 && input[i-1][j+1] == '@' {
				ct++
			}
			// lower left
			if i < n-1 && j > 0 && input[i+1][j-1] == '@' {
				ct++
			}
			// lower right
			if i < n-1 && j < m-1 && input[i+1][j+1] == '@' {
				ct++
			}

			if ct < 4 {
				ans++
			}

		}
	}
	return ans
}
