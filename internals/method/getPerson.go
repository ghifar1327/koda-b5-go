package method

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func GetPerson(scanner *bufio.Scanner) {
	reader := bufio.NewReader(os.Stdin)

	// Input Name
	fmt.Print("Insert Name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	// Input Address
	fmt.Print("Insert Address: ")
	address, _ := reader.ReadString('\n')
	address = strings.TrimSpace(address)

	// Input Phone
	fmt.Print("Insert Phoen Numbar: ")
	phone, _ := reader.ReadString('\n')
	phone = strings.TrimSpace(phone)

	// Buat Person pakai constructor
	person := NewPerson(name, address, phone)

	fmt.Println("\n=== DATA PERSON ===")
	fmt.Println(person.PrintMethod())

	fmt.Println(person.Greet())

	// setter
	fmt.Print("\nMasukkan nama baru: ")
	newName, _ := reader.ReadString('\n')
	newName = strings.TrimSpace(newName)
	person.SetName(newName)

	fmt.Println("\n=== NAMA BERHASIL DIUBAH ===")
	fmt.Println(person.Greet())
}