package main

import (
	"fmt"
	"os"
	"strconv"
)

func printBits(num int) string {
	result := ""
	slice := []int{}
	for num > 0 {
		mod := num % 2

		slice = append(slice, mod)
		num /= 2
	}
	for i := 0; i < len(slice)-1; i++ {
		for j := i + 1; j < len(slice); j++ {
			slice[i], slice[j] = slice[j], slice[i]
		}
	}
	for _, v := range slice {
		result += strconv.Itoa(v)
	}
	padding := 8 - len(slice)

	for i := 0; i < padding; i++ {
		result = "0" + result
	}
	return result
}

func main() {
	if len(os.Args) != 2 {
		return
	}
	args := os.Args[1]
	num, err := strconv.Atoi(args)
	if err != nil {
		fmt.Print("00000000")
		return
	}
	fmt.Print(printBits(num))
}
