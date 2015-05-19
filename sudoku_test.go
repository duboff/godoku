package sudoku

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCross(t *testing.T) {
	testCross := Cross("abc", "12")
	expectedCross := []string{"a1", "a2", "b1", "b2", "c1", "c2"}

	assert.Equal(t, 6, len(testCross))
	assert.Equal(t, expectedCross, testCross)
}

func TestUnitList(t *testing.T) {
	unitList := UnitList("ABCDEFGHI", "123456789")
	firstRow := []string{"A1", "B1", "C1", "D1", "E1", "F1", "G1", "H1", "I1"}
	lastBox := []string{"G7", "G8", "G9", "H7", "H8", "H9", "I7", "I8", "I9"}

	assert.Equal(t, 27, len(unitList))
	assert.Equal(t, firstRow, unitList[0])
	assert.Equal(t, lastBox, unitList[26])
}

func TestUnits(t *testing.T) {
	squares := Cross("ABCDEFGHI", "123456789")
	unitList := UnitList("ABCDEFGHI", "123456789")
	units := Units(unitList, squares)
	expectedUnit := []string{"A2", "B2", "C2", "D2", "E2", "F2", "G2", "H2", "I2"}

	for _, s := range squares {
		assert.Equal(t, 3, len(units[s]))
	}
	assert.Equal(t, expectedUnit, units["C2"][0])
}

func TestPeers(t *testing.T) {
	squares := Cross("ABCDEFGHI", "123456789")
	unitList := UnitList("ABCDEFGHI", "123456789")
	units := Units(unitList, squares)
	peers := Peers(units, squares)

	for _, s := range squares {
		assert.Equal(t, 20, len(peers[s]))
	}
	assert.Equal(t, "A2", peers["C2"][0])
	assert.Equal(t, "B2", peers["C2"][1])
	assert.Equal(t, "D2", peers["C2"][2])
}

func TestParseGrid(t *testing.T) {
	squares := Cross("ABC", "123")
	grid := "1...2...3"
	expectedGrid := ParseGrid(grid, squares)

	assert.Equal(t, "1", expectedGrid["A1"])
	assert.Equal(t, "3", expectedGrid["C3"])
	assert.Equal(t, "2", expectedGrid["B2"])
}

func TestEliminate(t *testing.T) {
	values := make(map[string]string)
	values["C1"] = "..3..7..9"

	assert.Equal(t, values, Eliminate(values, "C1", "5"))

	newValues := make(map[string]string)
	for k, v := range values {
		newValues[k] = v
	}
	newValues["C1"] = "..3....9"

	assert.Equal(t, newValues, Eliminate(values, "C1", "7"))

	values["C2"] = "7"
	assert.Nil(t, Eliminate(values, "C2", "7"))
}
