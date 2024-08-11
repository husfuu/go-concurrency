package main

import (
	"fmt"
	"time"
)

func calculate(start int, end int, resultChan chan int) int {
	// proses perhitungan di dalam fungsi ini tetap sinkronus
	sum := 0
	for i := start; i <= end; i++ {
		sum += i
	}

	// kirim hasil perhitungan ke channel
	resultChan <- sum

	return sum
}

func main() {
	resultChannel := make(chan int)

	go calculate(1, 25, resultChannel)
	go calculate(26, 50, resultChannel)
	go calculate(51, 75, resultChannel)
	go calculate(76, 100, resultChannel)

	totalSum := 0
	for i := 0; i < 4; i++ {
		totalSum += <-resultChannel
	}

	time.Sleep(4 * time.Second)

	fmt.Print(totalSum)
}
