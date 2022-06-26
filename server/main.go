package main

import "fmt"

func assert(fn func(int, int)) int,
	a int, b int, result int) bool {
	if fn(a, b) == result {
		return true
		}
		return false
	}