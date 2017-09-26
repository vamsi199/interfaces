package main

import "fmt"

type myI interface {
	add()
	subtract()
}

type myS struct {
	a, b int
}

func (m myS) subtract() {
	fmt.Println(m.a - m.b)
}

func (m myS) add() {
	fmt.Println(m.a + m.b)
}
func main() {
	var in myI
	in = myS{a: 1, b: 2}
	in.add()
	in.subtract()

}

//a,b,c will work and d cannot work
