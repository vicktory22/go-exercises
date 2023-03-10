package airportrobot

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

type Italian struct{}

func (i Italian) Greet(name string) string {
	return fmt.Sprintf("I can speak %s: Ciao %s!", i.LanguageName(), name)
}

func (i Italian) LanguageName() string {
	return "Italian"
}

type Portuguese struct{}

func (p Portuguese) Greet(name string) string {
	return fmt.Sprintf("I can speak %s: Olá %s!", p.LanguageName(), name)
}

func (p Portuguese) LanguageName() string {
	return "Portuguese"
}

func SayHello(name string, greeter Greeter) string {
	return greeter.Greet(name)
}
