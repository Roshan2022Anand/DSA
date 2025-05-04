package main

import (
	"dsa/easy"
	"os"
)

func main() {
	prlbm := os.Args[1]

	switch prlbm {
	case "e1":
		easy.RunTwoSum()
	case "e2":
		easy.RunValidAnagram()
	}
}
