package profile

import (
	"bufio"
	"fmt"
	"strconv"
)

func GetProfile(scanner *bufio.Scanner) {
	// person := profile{
	// 	name:         "Muhammad Ghifar Ramdhani",
	// 	img:          "picture.jpg",
	// 	email:        "givartorreto@gamil.com",
	// 	age:          26,
	// 	phone_number: +6285591710309,
	// 	isMarrid:     false,
	// 	education: []education{
	// 		{
	// 			name:     "SDN Bayur Kidul 1",
	// 			isPassed: true,
	// 		},
	// 		{
	// 			name:     "SMPN 1 Karawang Timur",
	// 			isPassed: true,
	// 		},
	// 		{
	// 			name:     "SMAN 2 Karawang",
	// 			isPassed: true,
	// 		},
	// 	},
	// }
	// // fmt.Println(person)
	// fmt.Printf("nama : %s\nimg : %s\nalamat email : %s\numur : %d\nno telepon : %d\nstatus pernikahan : %v\npendidikan %v", person.name, person.img,person.email, person.age,person.phone_number,person.isMarrid, person.education)

	var person profile

	fmt.Print("Nama: ")
	scanner.Scan()
	person.name = scanner.Text()

	fmt.Print("Image: ")
	scanner.Scan()
	person.img = scanner.Text()

	fmt.Print("Email: ")
	scanner.Scan()
	person.email = scanner.Text()

	fmt.Print("Umur: ")
	scanner.Scan()
	age, _ := strconv.Atoi(scanner.Text())
	person.age = age

	fmt.Print("No Telepon: ")
	scanner.Scan()
	phoneNumber, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	person.phone_number = phoneNumber

	fmt.Print("Status menikah (true/false): ")
	scanner.Scan()
	isMarried, _ := strconv.ParseBool(scanner.Text())
	person.isMarrid = isMarried

	// education ============================================================================
	fmt.Print("Jumlah pendidikan: ")
	scanner.Scan()
	totalEdu, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < totalEdu; i++ {
		var edu education

		fmt.Printf("Nama sekolah ke-%d: ", i+1)
		scanner.Scan()
		edu.name = scanner.Text()

		fmt.Print("Lulus? (true/false): ")
		scanner.Scan()
		edu.isPassed, _ = strconv.ParseBool(scanner.Text())

		person.education = append(person.education, edu)
	}
	fmt.Println("\n\n\n===== BIODATA =====")
	fmt.Printf("Nama : %s\n", person.name)
	fmt.Printf("Email: %s\n", person.email)
	fmt.Printf("Email: %s\n", person.img)
	fmt.Printf("Umur : %d\n", person.age)
	fmt.Printf("Telp : %d\n", person.phone_number)
	fmt.Printf("Menikah: %v\n", person.isMarrid)
	fmt.Println("Pendidikan:")
	for _, edu := range person.education {
		fmt.Printf("- %s (Lulus: %v)\n", edu.name, edu.isPassed)
	}
	fmt.Println("\n================")

}