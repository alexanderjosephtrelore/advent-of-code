package day5

import (
	"fmt"

	"github.com/alexanderjosephtrelore/advent-of-code/io"
)

func Solve(inputFile string) {
	input, err := io.ReadFileIntSlice(inputFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	p1 := Day5(input)
	p2 := Day5Bonus(input)
	fmt.Printf("d5p1 = %d\nd5p2 = %d\n", p1, p2)
}

func Day5(input []int) int {
	counter := 0
	for i := 0; 0 <= i && i < len(input); counter++ {
		input[i]++
		i = i + input[i] - 1
	}
	return counter
}

func Day5Bonus(input []int) int {
	counter := 0
	for i := 0; 0 <= i && i < len(input); counter++ {
		if input[i] >= 3 {
			input[i]--
			i = i + input[i] + 1
		} else {
			input[i]++
			i = i + input[i] - 1
		}
	}
	return counter
}
