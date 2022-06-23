package main

import "fmt"

func main() {
	prices := []int{32, 32, 32, 132, 432}
	items := []string{"消しゴム", "物体", "ゴミ", "ノート", "鉛筆"}
	sum := 0

	for i := 0; i < len(prices); i++ {
		sum += prices[i]
		fmt.Printf("%s: %d円｜小計： %d円\n", items[i], prices[i], sum)
	}
	fmt.Printf("総計： %d円¥n", sum)
}
