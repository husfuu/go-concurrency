package main

func Calculate(start int, end int, resultChan chan int) int {
	// proses perhitungan di dalam fungsi ini tetap sinkronus
	sum := 0
	for i := start; i < end; i++ {
		sum += i
	}

	// kirim hasil perhitungan ke channel
	resultChan <- sum

	return sum
}
