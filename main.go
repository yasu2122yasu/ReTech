package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	arg := flag.Arg(0)
	fmt.Printf("Hello,World! %s\n", arg)
}
