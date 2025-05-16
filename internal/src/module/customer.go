package module

type Customer struct {
	Email    Email
	Cpf      Cpf
	Age      Age
	IsActive bool
}

// GetEmail get the value from entity Email
func (customer *Customer) GetEmail() string {
	return customer.Email.Email
}

// GetCpf get the value from entity Cpf
func (customer *Customer) GetCpf() string {
	return customer.Cpf.Cpf
}

// GetAge get the value from entity Age
func (customer *Customer) GetAge() int {
	return customer.Age.Age
}
