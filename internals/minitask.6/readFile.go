package minitask6

import (
	"fmt"
	"os"
)

func FileReader() {
	func ()  {
		if r := recover(); r != nil{
			fmt.Printf("recoveri panic %v", r)
		}
	}()
	file, err := os.Open("file.txt")
	if err != nil {
		panic("Error:")
	}
	defer file.Close()
	
	data, err := os.ReadFile("file.txt")
	if err != nil {
		panic("Error membaca file")
	}

	fmt.Println(string(data))
}