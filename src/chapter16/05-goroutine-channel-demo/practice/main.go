package main

import "fmt"

func putData(numChan chan int) {
	for i := 1; i <= 2000; i++ {
		numChan <- i
	}

	close(numChan)
}

func cal(numChan chan int, resChan chan map[int]int) {
	for {
		num, ok := <-numChan

		if !ok {
			break
		}
		sum := 0
		//calculate
		for i := 1; i <= num; i++ {
			sum += i
		}

		res := make(map[int]int)
		res[num] = sum
		resChan <- res
	}
}

func main() {
	numChan := make(chan int, 2000)

	go putData(numChan)

	resChan := make(chan map[int]int, 2000)
	for i := 1; i <= 8; i++ {
		go cal(numChan, resChan)
	}

	for {
		if len(resChan) == 2000 {
			close(resChan)
			break
		}
	}

	for m := range resChan {
		fmt.Println(m)
	}
}
