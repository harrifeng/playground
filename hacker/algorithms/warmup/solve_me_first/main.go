package main

import (
	"fmt"
	"os"
)

func solveMeFirst(a uint32, b uint32) uint32 {
	return (a + b)
}

func main() {

	// prepare input output----------------->
	infile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	os.Stdin = infile

	outfile, err := os.Open("output.txt")
	_ = outfile
	if err != nil {
		panic(err)
	}
	// prepare input output----------------->
	var a, b, res uint32
	fmt.Scanf("%v\n%v", &a, &b)
	res = solveMeFirst(a, b)
	fmt.Println(res)
	os.Exit(0)
}
