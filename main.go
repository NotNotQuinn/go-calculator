package main

import (
	"fmt"
	"os"

	calc "github.com/NotNotQuinn/go-calculator/calculator"
)

func main() {
	calculator := calc.NewCalculator(os.Stdin)
	ch := make(chan calc.IntAndError, 5)
	go calculator.ReadInput(ch)

	for res := range ch {
		if err := res.Err; err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			continue
		}
		fmt.Println(res.Num)
	}
}
