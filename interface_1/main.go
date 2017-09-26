package main

import (
	"fmt"
)

type myStr []string

func (st myStr) String() string {
	var out string
	for i, str := range st {
		out += fmt.Sprintln(i, ": ", str)

	}
	return out
}

func main() {
	s := myStr{"A", "B", "C"}
	fmt.Println(s)
}
