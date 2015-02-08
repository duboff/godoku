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

func Units(unitList [][]string, squares []string) map[string][][]string {
	resp := make(map[string][][]string, len(squares))

	for _, s := range squares {
		unit := make([][]string, 3)
		i := 0

		for _, u := range unitList {
			for _, uu := range u {
				if uu == s {
					unit[i] = u
					resp[s] = unit
					i++
					break
				}
			}
		}

	}
	return resp
}

func Peers(units map[string][][]string, squares []string) map[string]map[string]bool {
	peers := make(map[string]map[string]bool, len(units))

	for _, s := range squares {
		for key, value := range units {
			peer := make(map[string]bool)

			for _, u := range value {
				

			}
	}
}

func main() {
}
