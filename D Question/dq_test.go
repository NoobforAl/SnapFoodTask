package main

import "testing"

func TestCase1(t *testing.T) {
	nums := []int{12, 54, 89, 21, 66, 47, 14, 285, 96}

	if sumPower2(nums) != 108624 {
		t.Error("Sum not equal 108624")
	}
}
func TestCase2(t *testing.T) {
	nums := []int{1, 1, 1, 1, 1}
	if sumPower2(nums) != 5 {
		t.Error("Sum not equal 5")
	}
}
