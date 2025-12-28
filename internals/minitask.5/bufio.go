package minitask5

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ghifar1327/koda-b5-go/internals/channel"
	"github.com/ghifar1327/koda-b5-go/internals/goroutine"
	lingkaran "github.com/ghifar1327/koda-b5-go/internals/ligkaran"
	"github.com/ghifar1327/koda-b5-go/internals/loop"
	"github.com/ghifar1327/koda-b5-go/internals/method"
	minitask6 "github.com/ghifar1327/koda-b5-go/internals/minitask.6"
	payments "github.com/ghifar1327/koda-b5-go/internals/payment"
	"github.com/ghifar1327/koda-b5-go/internals/profile"
	"github.com/ghifar1327/koda-b5-go/internals/sliceArray"
)

func GoBufio() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n============ MENU ==========")
		fmt.Println("1. Hitung Luas dan Lingkaran")
		fmt.Println("2. Print triangle")
		fmt.Println("3. Sisipkan angka ke array")
		fmt.Println("4. Isi biodata")
		fmt.Println("5. Buka file")
		fmt.Println("6. Data person 'method'")
		fmt.Println("7. Cashier")
		fmt.Println("8. Rutinitas pagi")
		fmt.Println("9. Black Board 'channel'")
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
		case "5":
			minitask6.FileReader()
		case "6":
			method.GetPerson(scanner)
		case "7":
			payments.Cashier(scanner)
		case "8":
			goroutine.Rutinitas()
		case "9":
			channel.Main()
		case "0":
			fmt.Println("Keluar dari program")
			return
		default:
			fmt.Println("Pilihan tidak valid ❌")
		}
	}
}
