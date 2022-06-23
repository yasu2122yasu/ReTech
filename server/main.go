package main

import "fmt"

func main() {
	scores := [][3]int{
		{30, 23, 45}, {34, 45, 43}, {32, 53, 41},
		{32, 32, 52}, {64, 14, 54}}

	for i := 0; i < len(scores); i++ {
		sum := 0
		for k := 0; k < 3; k++ {
			sum += scores[i][k]
		}
		fmt.Printf("受験者%d: 平均値%d点\n", i+1, sum/3)
	}
}
