package main

import (
	"fmt"
)

func Call(a int) Option[string] {
	if a > 50 {
		return NewNone[string]()
	}
	return NewSome[string]("Nice!")
}

func Incoming(a Option[int]) {
	val, err := a.Get()
	if err != nil {
		val = 0
	}
	fmt.Println("val", val)
}

func main() {
	opt := Call(55)
	val, err := opt.Get()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
}
