package main

import "fmt"
import "time"

func main() {
	for {
		fmt.Println("work")

		time.Sleep(time.Second)
	}
}
