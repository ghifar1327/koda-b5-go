package channel

import (
	"sync"
	"time"
	// "time"
)

type Message struct {
	Sender  string
	Content string
}

func MsgSender(c chan<- Message, sender string, content string, wg *sync.WaitGroup){
	defer wg.Done()
	time.Sleep(time.Microsecond * 1000)
	c <- Message{
		Sender: sender,
		Content: content,
	}
}