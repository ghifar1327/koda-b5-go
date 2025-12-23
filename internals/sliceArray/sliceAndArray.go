package sliceArray

import (
	"bufio"
	"fmt"
	"strconv"
)

func InsertNumber(scanner *bufio.Scanner) {
	fmt.Print("Masukkan angka yang ingin di sisipkan: ")

	scanner.Scan()
	input := scanner.Text()
	num , err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Input harus angka ❌")
	}

	var number = []int{50, 75, 66, 20, 32, 90}
	number = append(number[0:3], append([]int{num}, number[3:6]...)...)
	// fmt.Print(number)
	for i := range number {
		fmt.Println(number[i])
	}
}