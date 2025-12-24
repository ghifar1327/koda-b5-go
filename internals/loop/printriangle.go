package loop

import (
	"bufio"
	"fmt"
	"strconv"
)

func PrintTriangle(scanner *bufio.Scanner) {
	fmt.Print("Masukkan tinggi segitiga: ")
	scanner.Scan()
	input := scanner.Text()

	height, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Input harus angka ❌")
		return
	}

	for y := 1; y <= height; y++ {
		result := ""
		for x := 1; x <= y; x++ {
			result += " *"
		}
		fmt.Println(result)
	}
}
