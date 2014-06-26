package main

import (
	"fmt"
)

type TestStruct struct {
	s    []byte
	a, b int
}

func (p *TestStruct) String() string {
	return string(p.s)
}

func dump(p *TestStruct) {
	fmt.Printf("%#v", p)
}

func main() {

	ts := TestStruct{}

	s := []byte("hello")
	ts.s = s

	//dump(&ts)

	for n := 0; n < 5; n++ {
		fmt.Println(ts)
	}
}
