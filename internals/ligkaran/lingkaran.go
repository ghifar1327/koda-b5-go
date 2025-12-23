package lingkaran

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
)

func LuasLingkaran(scanner *bufio.Scanner) {
	fmt.Print("Masukkan jari-jari: ")
	scanner.Scan()

	r, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		fmt.Println("Input harus angka ❌")
		return
	}

	luas := math.Pi * r * r
	fmt.Printf("Luas lingkaran: %.2f cm²\n", luas)
}

func KelilingLingkaran(scanner *bufio.Scanner) {
	fmt.Print("Masukkan jari-jari: ")
	scanner.Scan()

	r, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		fmt.Println("Input harus angka ❌")
		return
	}

	keliling := 2 * math.Pi * r
	fmt.Printf("Keliling lingkaran: %.2f cm\n", keliling)
}
