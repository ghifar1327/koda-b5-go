package method

import "fmt"

type Person struct {
	Name    string
	Address string
	Phone   string
}

// Constructor
func NewPerson(name, address, phone string) *Person {
	return &Person{
		Name:    name,
		Address: address,
		Phone:   phone,
	}
}

// Method Print
func (p Person) PrintMethod()string{
	return fmt.Sprintf("Name: %s\n Address: %s\n Phone: %s\n", p.Name, p.Address, p.Phone)
}

// Method Greet
func (p Person) Greet() string {
	return fmt.Sprintf("Hello, my name is %s!", p.Name)
}

// Setter Name
func (p *Person) SetName(name string) {
	p.Name = name
}