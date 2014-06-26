package main

import (
	"fmt"
)

// http://play.golang.org/p/nM5NT_sDFE

type PeekableBuffer struct {
	buffer     []byte
	begin, end int
}

func (b *PeekableBuffer) Peek() []byte {
	return b.buffer[b.begin:b.end]
}

func (b *PeekableBuffer) Consume(amount int) {
	b.begin += amount
}

func (b *PeekableBuffer) Write(data []byte) {
	if len(b.buffer)-b.end < len(data) {
		usedlen := b.end - b.begin
		if len(b.buffer)-usedlen >= len(data) {
			//fmt.Println(">>Move")
			b.end = copy(b.buffer, b.buffer[b.begin:b.end])
			b.begin = 0
		} else {
			//fmt.Println(">>Re-alloc")
			newbuf := make([]byte, usedlen+len(data))
			b.end = copy(newbuf, b.buffer[b.begin:b.end])
			b.begin = 0
			b.buffer = newbuf
		}
	}
	b.end += copy(b.buffer[b.end:], data)
}

/*
func AdvancePastWS(buffer []byte) []byte {
	for {
		if r, sz = utf8.DecodeRune(buffer); r != utf8.RuneError {

		} else {
			return buffer
		}
	}
}

func AdvanceTo(buffer []byte, s string) ([]byte, bool) {

}

func AdvancePast(buffer []byte, s string) ([]byte, bool) {

}
*/

func dumpPB(pb *PeekableBuffer) {
	fmt.Printf("buffer: [%s] used: [%s] markers: [%d,%d]\n", string(pb.buffer), string(pb.Peek()), pb.begin, pb.end)
}

func main() {

	pb := PeekableBuffer{}

	s1 := []byte("wibble")
	pb.Write(s1)
	dumpPB(&pb)
	pb.Consume(3)
	dumpPB(&pb)
	pb.Write([]byte("1234"))
	dumpPB(&pb)
	pb.Consume(5)
	dumpPB(&pb)

	for i := 0; i < 10; i++ {
		pb.Write([]byte("X"))
		dumpPB(&pb)
	}

	pb.Consume(5)
	dumpPB(&pb)
	for i := 0; i < 10; i++ {
		pb.Write([]byte("O"))
		dumpPB(&pb)
	}

}
