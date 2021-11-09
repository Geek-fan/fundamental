package engineering

type person struct {
	name string
	age int
	address string
}

type NewPersonOption func(p *person)

func NewPerson(name string, opts ...NewPersonOption) *person {
	p := &person{name: name}

	for _, opt := range opts {
		opt(p)
	}
	return p
}

func SetAge(age int) NewPersonOption {
	return func(p *person) {
		p.age = age
	}
}

func SetAddress(address string) NewPersonOption {
	return func(p *person) {
		p.address = address
	}
}