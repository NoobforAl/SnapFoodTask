package main

import "fmt"

type find struct {
	index, value int
}

func foundSingle(arr *[]int) *find {
	keys := make(map[int]int)
	read := make(map[int]int)

	for i, v := range *arr {
		if _, ok := keys[v]; ok {
			delete(keys, v)
		} else {
			if _, ok := read[v]; !ok {
				keys[v] = i
			} else {
				read[v] = 0
			}
		}
	}

	for v, i := range keys {
		return &find{index: i, value: v}
	}

	return nil
}

func main() {
	arr := []int{2, 2, 5, 6, 5}
	res := foundSingle(&arr)

	if res != nil {
		fmt.Printf("Index : %d\nValue : %d", res.index, res.value)
	} else {
		fmt.Println("NotFound!")
	}
}
