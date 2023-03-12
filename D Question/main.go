package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func worker(c <-chan int64, s *int64, wg *sync.WaitGroup) {
	var sum int64 = 0
	defer wg.Done()

	for v := range c {
		sum += v * v
	}

	atomic.AddInt64(s, sum)
}

func sumPower2(nums []int64) int64 {
	ch := make(chan int64)
	var sum int64 = 0
	var wg sync.WaitGroup

	wg.Add(2)

	go worker(ch, &sum, &wg)
	go worker(ch, &sum, &wg)

	for _, v := range nums {
		ch <- v
	}

	close(ch)

	wg.Wait()
	return sum
}

func main() {
	nums := []int64{12, 54, 89, 21, 66, 47, 14, 285, 96}
	fmt.Println(sumPower2(nums))
}
