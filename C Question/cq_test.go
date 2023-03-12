package main

import (
	"testing"
)

func TestCase1(t *testing.T) {
	r := NewRequester(&[]string{
		"https://www.google.com/",
		"https://www.google.com/",
	})

	if !r.Run() {
		t.Error("This request was not equal false!")
	}
}

func TestCase2(t *testing.T) {
	r := NewRequester(&[]string{
		"https://sdfa123aaasssdfs.com/",
		"https://sdfsssda123sdfsd.com/",
		"https://www.google.com/",
	})

	if r.Run() {
		t.Error("This request was equal false!")
	}

}
