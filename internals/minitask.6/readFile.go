package minitask6

import (
	// "bufio"
	"fmt"
	"os"
)

func FileReader(){
file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}else{
		result := fmt.Sprintf("file %v", file)
		println(result)
	}
	// defer file.Close()

	// fmt.Println("File berhasil dibuka:", file.Name())
}