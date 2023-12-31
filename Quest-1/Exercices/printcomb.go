package main

import "fmt"

func PrintComb() {
	for i := 0; i <= 7; i++ {
		for j := i + 1; j <= 8; j++ {
			for k := j + 1; k <= 9; k++ {
				fmt.Printf("%d%d%d", i, j, k)
				if i != 7 || j != 8 || k != 9 {
					fmt.Printf(", ")
				}
			}
		}
	}
	fmt.Println()
}

func main() {
	PrintComb()
}
