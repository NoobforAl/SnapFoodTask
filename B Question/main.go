package main

import (
	"log"
)

func foundSingle(arr []int) int {
	var exception int
	for _, v := range arr {
		exception ^= v
	}
	return exception
}

func main() {
	arr := []int{2, 2, 5, 6, 5}
	ans := foundSingle(arr)

	if ans != 6 {
		log.Fatalf("Wrong Answer!\nAnswer Equal 6 not equal: %d", ans)
	}
	log.Println("True Answer!")
}
