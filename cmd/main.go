package main

import (
	"dsa/easy"
	"dsa/medium"
	"os"
)

func main() {
	prlbm := os.Args[1]

	switch prlbm {
	case "e1":
		easy.RunTwoSum()
	case "e2":
		easy.RunValidAnagram()
	case "e3":
		easy.RunContainDuplicate()
	case "e4":
		easy.Run1stUniqChar()
	case "m1":
		medium.RunGrpAnagram()
	case "m2":
		medium.RunLenOstr()
	}
}
