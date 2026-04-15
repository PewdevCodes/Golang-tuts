package main 

import (
	"fmt"
	"math/rand"
)

/*  
              INTERFACES 
An interface is a type that defines a set of method signatures. A type implements an interface by implementing its methods. Interfaces allow us to write flexible and reusable code by enabling polymorphism. In Go, interfaces are satisfied implicitly, meaning that a type implements an interface simply by having the required methods, without needing to declare it explicitly.

Contract between the interface and the struct: An interface defines a contract that a struct must fulfill by implementing the methods specified in the interface. This allows us to write code that can work with any type that satisfies the interface, promoting code reusability and flexibility.

*/

type Player interface {
	KickBall()
}

type FootBallPlayer struct {
	stamina int
	power  int
}

func (f FootBallPlayer) KickBall() {

	shot := f.stamina + f.power
	fmt.Println("The football player kicks the ball with a power of:", shot)

}

func main() { 
	team := make([]Player, 11)
	for i := 0 ; i < len(team) ; i ++ {
		team[i] = FootBallPlayer{
			stamina: rand.Intn(100),
			power:   rand.Intn(100),
		}
	}
	for i := 0 ; i < 11 ; i ++ {
		team[i].KickBall()
	} 
}

// anything can be a speaker as long as it has the speak method
type Speaker interface { 
	Speak() string
}

type Dog struct {
	Breed string
	Name string
}

func ( d Dog) Speak() string { 
	return fmt.Sprintf("Hi, I'm %s, a %s.", d.Name, d.Breed)
}

func makeSound(s Speaker) {
	fmt.Println(s.Speak())
}