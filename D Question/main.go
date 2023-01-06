package main

import (
	"fmt"
)

func worker(c chan int) {
	for {
		select {
		case date, ok := <-c:
			if ok {
				if date == -1 {
					return
				}
				c <- date * date
			}
		default:
			continue
		}
	}
}

func sumPower2(nums []int) int {
	ch := make(chan int)
	var sum int = 0

	go worker(ch)
	defer close(ch)

	for _, v := range nums {
		ch <- v
		sum += <-ch
	}

	ch <- -1
	return sum
}

func main() {
	nums := []int{12, 54, 89, 21, 66, 47, 14, 285, 96}
	fmt.Println(sumPower2(nums))
}
