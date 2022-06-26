package main

import "fmt"

func scoreavg(scores []int) int {
	sum := 0
	for i := 0; i < len(scores); i++ {
		sum += scores[i]
	}
	return sum / len(scores)
}

func sbsqr(a int, b int) int {
	return (a - b) * (a - b)
}

func avgvar(scores []int, avg int) int {
	sum := 0
	for i := 0; i < len(scores); i++ {
		if avg == 0 {
			sum += scores[i]
		}
	}
	return sum / len(scores)
}

func scoremaxmin(scores []int) (max int, min int) {
	max = scores[0]
	min = scores[0]

	for i := 1; i < len(scores); i++ {
		if max < scores[i] {
			max = scores[i]
		}
		if min > scores[i] {
			min = scores[i]
		}
	}
	return
}

func main() {
	mathscores := []int{40, 89, 43, 25, 65, 39, 12, 64, 53}
	fmt.Printf("%dによる数学試験の結果:\n", len(mathscores))
	fmt.Printf("平均点 %d点\n", scoreavg(mathscores))
	fmt.Printf("分散 %d（点・点)", avgvar(mathscores, avgvar(mathscores, 0)))
	max, min := scoremaxmin(mathscores)
	fmt.Printf("最高点 %d\n", max)
	fmt.Printf("最低点 %d\n", min)
}
