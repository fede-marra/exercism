package airportrobot

import "fmt"

// Write your code here.
type Greeter interface {
	LanguageName() string
	Greet(string) string
}

type German struct {
	
}

func (g German) Greet(name string) string {
	return fmt.Sprintf("Hallo %s!", name)
}

func (g German) LanguageName() string {
	return "German"
}

type Italian struct {
	
}

func (i Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao %s!", name)
}

func (i Italian) LanguageName() string {
	return "Italian"
}


type Portuguese struct {
	
}

func (p Portuguese) Greet(name string) string {
	return fmt.Sprintf("Olá %s!", name)
}

func (p Portuguese) LanguageName() string {
	return "Portuguese"
}



func SayHello(name string, g Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", g.LanguageName(), g.Greet(name))
}


// => "I can speak German: Hallo Dietrich!"


// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
