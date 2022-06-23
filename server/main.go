package main

import "fmt"

func main() {
	prices := []int{98, 32, 32, 46, 14}
	for i := 0; i < len(prices); i++ {

		fmt.Printf("%d円¥n", prices[i])
	}
}
