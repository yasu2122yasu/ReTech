package main

import "fmt"

func main() {
	scores := [5][3]int{
		{43, 31, 34}, {32, 65, 75}, {43, 43, 12},
		{76, 45, 24}, {95, 142, 64},
	}

	for i := 0; i < 5; i++ {
		total := scores[i][0]+scores[i][1]+scores[i][2] > 179
		balance := scores[i][0] > 59 && scores[i][1] > 59 && scores[i][2] > 59
		talent := scores[i][0] > 89 || scores[i][1] > 89 || scores[i][2] > 89

		fmt.Printf("受験者%dさんは：\n", i+1)
		if total {
			fmt.Println("総得点で合格")
		}
		if balance {
			fmt.Println("すべての課題で合格")
		}
		if !total && !balance && talent {
			fmt.Println("一つだけ優秀")
		}

		if total && balance {
			fmt.Println("おめでとうございます。")
		} else if total || talent {
			fmt.Println("来週、挑戦できます。")
		} else {
			fmt.Println("来年、また挑戦してください。")
		}
		fmt.Println()
	}
}
