package main

import "testing"

func TestCase1(t *testing.T) {
	res := callUrls([]string{
		"https://www.google.com/",
		"https://www.google.com/",
	})
	if !res.suc {
		t.Error("This request was not equal false!")
	}
}

func TestCase2(t *testing.T) {
	res := callUrls([]string{
		"https://sdfa123aaasssdfs.com/",
		"https://sdfsssda123sdfsd.com/",
		"https://www.google.com/",
		"https://stackoverflow.com/",
	})
	if res.suc {
		t.Error("This request was not equal true!")
	}
}
