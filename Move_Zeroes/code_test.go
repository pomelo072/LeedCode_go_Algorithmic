package Move_Zeroes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	MoveZeroes(nums)
	if assert.Equal(t, nums, []int{1, 3, 12, 0, 0}) {
		t.Log(nums)
	} else {
		t.Log(nums)
		t.Error("Error")
	}
}
