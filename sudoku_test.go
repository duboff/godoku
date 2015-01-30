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
