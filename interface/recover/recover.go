package main

import (
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("err occurred: ", err)
		} else {
			panic(err)
		}
	}()
	//panic(errors.New("this is bullshit"))
	a := 5
	b := 0
	c := a / b
	fmt.Println(c)
}

func main() {
	tryRecover()
}
