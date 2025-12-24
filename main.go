package main

import (
	"fmt"

	"github.com/ghifar1327/koda-b5-go/internals/method"
	// minitask5 "github.com/ghifar1327/koda-b5-go/internals/minitask.5"
)

func main() {
	// minitask5.GoBufio()
	data := method.NewPerson("ghifar", "karawang", "085591710309")
	fmt.Println(data)
}
