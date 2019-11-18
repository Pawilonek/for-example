package main

import (
	"fmt"
)

func main() {
	strings := []string{"      |\\ |\\", "      \\ \\| |", "       \\ | |", "      .--''/", "     /o     \\", "     \\      /", "      \\    /"}

	for string := range strings {
		fmt.Println(string)
	}
}
