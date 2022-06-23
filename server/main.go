package main

import "fmt"

func main() {
	prices := []int{98, 54, 43, 32, 12}
	items := []string{"消しゴム", "鉛筆", "生ゴミ", "煙突", "ケース"}

	for i := 0; i < len(prices); i++ {
		fmt.Printf("%s: %d円/n", items[i], prices[i])
	}

}
