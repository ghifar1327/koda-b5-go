package method

type person struct{
	name  string
	address string
	phone string  
}

func NewPerson(name string, address string ,phone string) *person{
	return &person{
		name : name,
		address : address,
		phone : phone,
	}
}

// func (p *person)