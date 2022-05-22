package hackerrank

import "time"

var last = int64(1)
var previous = int64(2)
var counter = 0

func ModuloFibonacciSequence(requestChan chan bool, resultChan chan int) {
	for {
		<-requestChan
		time.Sleep(10 * time.Millisecond)

		if counter == 0 {
			resultChan <- int(last)
			counter++
		} else if counter == 1 {
			resultChan <- int(previous)
			counter++
		} else {
			temp := previous
			previous = fibbo(last, previous) % 1e9
			last = temp
			resultChan <- int(previous)
		}
	}
}

func fibbo(i1, i2 int64) int64 {
	return i1 + i2
}
