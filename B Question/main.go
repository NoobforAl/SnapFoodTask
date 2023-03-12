package main

import "fmt"

func foundSingle(arr *[]int) int {
	keys := make(map[int]int)

	for i, v := range *arr {
		if _, ok := keys[v]; ok {
			delete(keys, v)
		} else {
			keys[v] = i
		}
	}

	for v := range keys {
		return v
	}

	return -1
}

func main() {
	arr := []int{2, 2, 5, 6, 5}
	ans := foundSingle(&arr)

	if ans != -1 {
		fmt.Printf("Value : %d", ans)
	} else {
		fmt.Println("NotFound!")
	}
}
