package module

type Customer struct {
	Email    Email
	Cpf      Cpf
	Age      Age
	IsActive bool
}

func (customer *Customer) GetEmail() string {
	return customer.Email.Email
}

func (customer *Customer) GetCpf() string {
	return customer.Cpf.Cpf
}

func (customer *Customer) GetAge() int {
	return customer.Age.Age
}
