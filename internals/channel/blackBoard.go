package channel

import (
	"fmt"
	"sync"
	"time"
)

func BlackBoard(c <-chan Message , wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("papan tulis dibuka")
	for i := range c{
		time.Sleep(time.Second)
		fmt.Printf("%s: %s\n", i.Sender, i.Content)
	}
	fmt.Println("papan tulis ditutup")
}