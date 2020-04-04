package main

import "fmt"

func main() {
	bestP := 0
	bestSolutions := 0
	for p := 2; p < 1000; p++ {

		solutions := 0
		for a := 2; a < (p / 3); a++ {
			if (p*(p-2*a))&(2*(p-a)) == 0 {
				solutions += 1
			}
		}

		if solutions > bestSolutions {
			bestSolutions = solutions
			bestP = p
		}
	}

	fmt.Printf("bestP = %+v\n", bestP)
	fmt.Printf("bestSolutions = %+v\n", bestSolutions)
}
