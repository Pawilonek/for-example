package main

import "fmt"

func main() {
	queue := make(chan int)

	go func() {
		for i := 0; i < 20; i++ {
			fmt.Printf("Produced: %02d.\n", i)
			queue <- i
		}

		close(queue)
	}()

	for num := range queue {
		fmt.Printf(">: %02d.\n", num)
	}
}
