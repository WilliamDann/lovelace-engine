package main

import (
	"strconv"
	"testing"
)

func TestCoordsToSquareNum0_0(t *testing.T) {
	value := coords_to_square_num(0, 0)
	if value != 0 {
		t.Errorf("incorrect conversion: " + strconv.FormatInt(int64(value), 10))
	}
}

func TestCoordsToSquareNum1_1(t *testing.T) {
	value := coords_to_square_num(1, 1)
	if value != 9 {
		t.Errorf("incorrect conversion: " + strconv.FormatInt(int64(value), 10))
	}
}

func TestCoordsToSquareNum7_7(t *testing.T) {
	value := coords_to_square_num(7, 7)
	if value != 63 {
		t.Errorf("incorrect conversion: " + strconv.FormatInt(int64(value), 10))
	}
}
