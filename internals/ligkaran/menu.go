package lingkaran

import (
	"bufio"
	"fmt"
	"strings"
)

func Menu(scanner *bufio.Scanner) {
	for {
		fmt.Println("\n=== MENU LINGKARAN ===")
		fmt.Println("1. Hitung Luas")
		fmt.Println("2. Hitung Keliling")
		fmt.Println("0. Kembali")
		fmt.Print("Pilih: ")

		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			LuasLingkaran(scanner)
		case "2":
			KelilingLingkaran(scanner)
		case "0":
			return
		default:
			fmt.Println("Pilihan tidak valid ❌")
		}
	}
}


// func KelilingLingkaran(r float64) {
// 	keliling := fmt.Sprintf("keliling lingkaran : %.2f cm", 2*math.Pi*r)
// 	fmt.Println(keliling)
// }