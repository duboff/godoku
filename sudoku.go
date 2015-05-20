package sudoku

import (
	"strings"
)

var cols string = "123456789"
var rows string = "ABCDEFGHI"
var squares []string = Cross(rows, cols)
var unitlist [][]string = UnitList(rows, cols)
var units map[string][][]string = Units(unitlist, squares)

var peers map[string][]string = Peers(units, squares)

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

func AppendIfMissing(slice []string, i string) []string {
	for _, ele := range slice {
    if ele == i {
      return slice
    }
	}
	return append(slice, i)
}

func Peers(units map[string][][]string, squares []string) map[string][]string {
	resp := make(map[string][]string, len(squares))

	for _, s := range squares {
		unit := units[s]
		peerList := make([]string, 0)

		for _, u := range unit {
			for _, q := range u {
				if q != s {
					peerList = AppendIfMissing(peerList, q)
				}
			}
		}
		resp[s] = peerList
	}
	return resp
}

func ParseGrid(grid string, squares []string) map[string]string{
	resp := make(map[string]string, len(squares))
	for i, s := range squares {
		resp[s] = string(grid[i])
	}
	return resp
}

func Eliminate(values map[string]string, key string, candidate string) map[string]string {
	if !strings.Contains(values[key], candidate) {
		return values
	}
	values[key] = strings.Replace(values[key], candidate, "", 1)

	if len(values[key]) == 0 {
		return nil
	} else if len(values[key]) == 1 {
		d2 := values[key]
		for _, v2 := range peers[key] {
			if Eliminate(values, v2, d2) == nil {
				return nil
			}
		}
	}
	return values	
}

func main() {
}
