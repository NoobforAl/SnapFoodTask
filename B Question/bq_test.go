package main

import (
	"testing"
)

func TestCase1(t *testing.T) {
	arr := []int{1, 2, 3, 1, 2, 3}
	if res := foundSingle(arr); res != 0 {
		t.Error("answer fun just: 0")
	}
}

func TestCase2(t *testing.T) {
	arr := []int{1, 2, 4, 5, 2, 3, 1, 3, 4}
	if res := foundSingle(arr); res != 5 {
		t.Error("answer fun just: 5")
	}
}
