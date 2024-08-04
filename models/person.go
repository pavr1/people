package models

import (
	"github.com/pavr1/people/config"
)

type Person struct {
	config   *config.Config
	ID       string
	Name     string
	LastName string
	Age      int
}

func NewPerson(config *config.Config) Person {
	return Person{
		config: config,
	}
}

func (p *Person) Populate(name string, lastName string, age int) {
	p.Name = name
	p.LastName = lastName
	p.Age = age
}
