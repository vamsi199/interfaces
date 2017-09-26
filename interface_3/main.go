package main

import "fmt"

type value interface {
}

func print(v value) {
	fmt.Println(v)

}

func main() {
	print(1)
	print("hello")

}
