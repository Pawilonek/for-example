package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	const pass string = "7694f4a66316e53c8cdd9d9954bd611d"

	var input string
	var inputHash string

	for {
		fmt.Print("Enter super secret password: ")
		fmt.Scanln(&input)

		hasher := md5.New()
		hasher.Write([]byte(input))
		inputHash = hex.EncodeToString(hasher.Sum(nil))

		if inputHash == pass {
			fmt.Print("¯\\_(ツ)_/¯")

			break
		}
	}
}
