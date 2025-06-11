package main

import "fmt"

// IQuackBehavior is the interface which acts like bridge between context obj and implementation obj
type IQuackBehavior interface {
	Quack()
}

// Duck is the context obj, this calls the implementations
type Duck struct {
	QuackHandler IQuackBehavior // call the implementations which compose the interface
}

// compositions
type MuteQuack struct{}

func (mq MuteQuack) Quack() {
	fmt.Println("Mute quack")
}

type NoisyQuack struct{}

func (nq NoisyQuack) Quack() {
	fmt.Println("Noisy quack")
}

// call the Quack method (decouple of the compositions)
func (duck Duck) PerformQuack() {
	duck.QuackHandler.Quack()
}

func main() {
	// case when I need noisy duck
	noisyHandler := NoisyQuack{}

	noisyDuck := Duck{
		QuackHandler: noisyHandler,
	}

	fmt.Print("Calling noisy duck method: ")
	noisyDuck.QuackHandler.Quack()

	// case when I need mute duck
	muteHandler := MuteQuack{}
	muteDuck := Duck{
		QuackHandler: muteHandler,
	}

	fmt.Print("Calling muted duck method: ")
	muteDuck.QuackHandler.Quack()

}
