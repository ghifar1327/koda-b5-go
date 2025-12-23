package minitask5

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	lingkaran "github.com/ghifar1327/koda-b5-go/internals/ligkaran"
	"github.com/ghifar1327/koda-b5-go/internals/loop"
	"github.com/ghifar1327/koda-b5-go/internals/profile"
	"github.com/ghifar1327/koda-b5-go/internals/sliceArray"
)

func GoBufio() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("============ MENU ==========")
		fmt.Println("1. Hitung Luas dan Lingkaran")
		fmt.Println("2. Print triangle")
		fmt.Println("3. Sisipkan angka ke array")
		fmt.Println("4. Isi biodata")
		fmt.Println("0. Exit")
		fmt.Print("Pilih menu: ")

		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			lingkaran.Menu(scanner)
		case "2":
			loop.PrintTriangle(scanner)
		case "3":
			sliceArray.InsertNumber(scanner)
		case "4":
			profile.GetProfile(scanner)
		case "0":
			fmt.Println("Keluar dari program")
			return
		default:
			fmt.Println("Pilihan tidak valid ❌")
		}
	}
}
