package payments

import "fmt"

type Payment interface {
	Pay(amount int)
}

// BANK -------------
type BankPayment struct {
	history *History
}

func ViaBankPayment(h *History) *BankPayment {
	return &BankPayment{history: h}
}

func (b *BankPayment) Pay(amount int) {
	record := fmt.Sprintf("Pay via bank : %d", amount)
	b.history.Add(record)
	fmt.Println(record)
}
// ONLINE --------------- 
type OnlinePayment struct {
	history *History
}
func ViaOnlinePayment(h *History) *OnlinePayment {
	return &OnlinePayment{history: h}
}
func (o *OnlinePayment) Pay(amount int) {
	record := fmt.Sprintf("Pay via online : %d", amount)
	o.history.Add(record)
	fmt.Println(record) 
}

// FIKTIF --------------
type FiktifPayment struct {
	history *History
}
func ViaFiktifPayment(h *History) *FiktifPayment {
	return &FiktifPayment{history: h}
}
func (f *FiktifPayment) Pay(amount int) {
	record := fmt.Sprintf("Pay via fiktif : %d", amount)
	f.history.Add(record)
}
