package main

import "testing"

func TestSumVer2(t *testing.T) {
	tables := []struct{
		x int
		y int 
		z int
	}{
		{1, 1, 2},
		{1, 2, 3},
		{0, 1, 1},
		{-1, 1, 0},
	}

	for _, table := range tables{
		total := Sum(table.x, table.y)
		if total != table.z {
			t.Errorf("Sum of (%d + %d) was incorrect, got: %d, want: %d.", table.x, table.y, total, table.z)
		}
	}
}
