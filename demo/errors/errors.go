package main

import (
	"errors"
	"fmt"
)

type Stuff struct {
	values []int
}

func (s *Stuff) Get(index int) (int, error) {
	if index >= len(s.values) {
		return 0, errors.New(fmt.Sprintf("Index %v out of range", index))
	}
	return s.values[index], nil
}

func main() {
	stuff := Stuff{}
	value, err := stuff.Get(10)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}
}
