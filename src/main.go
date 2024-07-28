package main

import (
	"fmt"
	"time"
)

func GetOutput(input string) string {
	return fmt.Sprintf("Value: %s", input)
}

func main() {
	fmt.Printf(GetOutput("Start\n"))
	for i := 0; i < 1000000; i++ {
		fmt.Printf("%s\n", GetOutput(fmt.Sprintf("%d", i)))
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf(GetOutput("End\n"))
}
