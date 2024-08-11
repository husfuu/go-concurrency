package main

import (
	"fmt"
	"testing"
	"time"
)

func TestParalelCalculate(t *testing.T) {
	resultChannel := make(chan int)

	go calculate(1, 10, resultChannel)

	fmt.Println("udah")

	fmt.Print(resultChannel)

	time.Sleep(1 * time.Second)
}
