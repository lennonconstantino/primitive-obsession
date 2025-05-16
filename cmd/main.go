package main

import (
	"fmt"
	"primitive-obsession/internal/src/module"
)

func main() {
	fmt.Println("Entry code here.")

	// Client Code
	var email module.Email
	email.Email = "lennon@email.com"
	var cpf module.Cpf
	cpf.Cpf = "123.456.789-09"
	var age module.Age
	age.Age = 36

	if erro := email.Validation(); erro != nil {
		fmt.Println("Error domain. ->", erro)
		return
	}

	if erro := cpf.ValidationCpf(); erro != nil {
		fmt.Println("Error domain. ->", erro)
		return
	}

	if erro := age.Validate(); erro != nil {
		fmt.Println("Error domain. ->", erro)
		return
	}

	var customer module.Customer
	customer.Email = email
	customer.Cpf = cpf
	customer.Age = age

	fmt.Println("Customer info: ")
	fmt.Println("-- Email: ", customer.GetEmail())
	fmt.Println("-- Cpf: ", customer.GetCpf())
	fmt.Println("-- Age: ", customer.GetAge())
}
