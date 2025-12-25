package payments

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func Cashier(scanner *bufio.Scanner) {
	history := NewHistory()
	bank := ViaBankPayment(history)
	online := ViaOnlinePayment(history)
	fiktif := ViaFiktifPayment(history)

	for {
		fmt.Println("\nPilih Metode Perbayaran")
		fmt.Println("\n1. Payment via Bank")
		fmt.Println("2. Payment via Online")
		fmt.Println("3. Payment via Fiktif")
		fmt.Println("0. Exit")
		fmt.Print("\nPilih: ")

		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		if choice == "0" {
			if len(history.GetHistory())==0{
				fmt.Println("\nhistory masih kosong")
				continue
			}
			fmt.Println("\nHistori pembayaran:")
			for i, h := range history.GetHistory() {
				fmt.Printf("%d. %v\n", i+1, h)
			}
			return
		}

		var payment Payment

		switch choice {
		case "1":
			payment = bank
		case "2":
			payment = online
		case "3":
			payment = fiktif
		default:
			fmt.Println("input tidak sesuai")
			continue
		}

		fmt.Print("Masukkan nominal: ")
		scanner.Scan()
		amount, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if err != nil || amount <= 0 {
			fmt.Println("input harus berupa angka dan lebih dari 0")
			continue
		}

		payment.Pay(amount)
	}
}
