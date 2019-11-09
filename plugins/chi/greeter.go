package main

import "fmt"

type greeting string

func (g greeting) Greet() {
	fmt.Println("你好宇宙")
}

// exported as symbol named "Greeter"
var Greeter greeting
