package main

import (
	"fmt"
	"os"

	"github.com/ifreddyrondon/go-workshop/santiago-nov2018/resources/src/13_testing/args"
	"github.com/ifreddyrondon/go-workshop/santiago-nov2018/resources/src/13_testing/divisible"
)

func main() {
	top, valid := args.GetFirstInt(os.Args)
	if !valid {
		fmt.Println("invalid arg")
	}
	fmt.Println("total", divisible.Sum(top, 3, 5))
}
