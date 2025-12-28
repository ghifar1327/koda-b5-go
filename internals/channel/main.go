package channel

import (
	"sync"
	"time"
)

func Main() {
	messageChannel := make(chan Message)
	var wgBoard sync.WaitGroup
	var wgSender sync.WaitGroup
	wgBoard.Add(1)
	go BlackBoard(messageChannel, &wgBoard)
	messages := []Message{
		{"Ayah", "Bu.. bikin kopi"},
		{"Ibu", "Besok kita belanja"},
		{"Kakak","Aku mau motor Aerox"},
		{"Adik", "besok jalan-jalan yaa..."},
	}
	wgSender.Add(len(messages))
	for _, msg := range messages{
		MsgSender(messageChannel, msg.Sender, msg.Content, &wgSender)
	}
	wgSender.Wait()
	close(messageChannel)
	time.Sleep(time.Second)
	wgBoard.Wait()
}