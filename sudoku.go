package sudoku

import ()

func Cross(A, B string) []string {
	resp := make([]string, len(A)*len(B))

	ind := 0
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			resp[ind] = string(A[i]) + string(B[j])
			ind++
		}
	}
	return resp
}

func UnitList(rows, cols string) [][]string {
	rowThrees := []string{"ABC", "DEF", "GHI"}
	colThrees := []string{"123", "456", "789"}

	resp := make([][]string, 27)

	ind := 0

	for _, c := range cols {
		resp[ind] = Cross(rows, string(c))
		ind++
	}

	for _, r := range rows {
		resp[ind] = Cross(string(r), cols)
		ind++
	}

	for _, r := range rowThrees {
		for _, c := range colThrees {
			resp[ind] = Cross(r, c)
			ind++
		}
	}
	return resp
}

func main() {
}
