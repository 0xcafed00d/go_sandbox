package sandbox

// messing with various aspects of The Go language
// Here Be Dragons

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestRange(t *testing.T) {

	slice := []int{8, 7, 6, 5, 4, 3, 2, 1}

	// v is a copy and the original slice is not modified
	for _, v := range slice {
		v += 1
	}

	// use index to access elements in original range to modify them
	for i := range slice {
		slice[i] += 1
	}
}

func printerror(err error) {
	fmt.Println(err)
}

func TestEmbed1(t *testing.T) {
	// embed interface into struct
	type wrapper struct {
		error
	}

	w := wrapper{}
	w.error = errors.New("test")

	// struct implements embedded interface
	printerror(w)
}

func TestPanic(t *testing.T) {
	// defer will trigger on function exit
	defer func() {
		// recover returns interface{} and if not nil is that arg passed to panic
		if r := recover(); r != nil {
			fmt.Println("Caught Panic:", reflect.TypeOf(r), r)
		}
	}()

	panic(3)
}

func SetOnError(errp *error) {
	if r := recover(); r != nil {
		if err, ok := r.(error); ok {
			*errp = err
		} else {
			panic(r)
		}
	}
}

func OnError(f func(err error)) {
	if r := recover(); r != nil {
		if err, ok := r.(error); ok {
			f(err)
		} else {
			panic(r)
		}
	}
}

func ReturnError(err error) {
	if err != nil {
		panic(err)
	}
}

func TestPanic2(t *testing.T) {

	test := func() (err error) {
		defer SetOnError(&err)

		ReturnError(nil)
		ReturnError(errors.New("this is an error1"))
		ReturnError(errors.New("this is an error2"))

		return
	}

	fmt.Println(test())
}

func TestPanic3(t *testing.T) {

	test := func() (i int, err error) {

		defer OnError(func(e error) {
			i, err = 0, e
		})

		ReturnError(nil)
		ReturnError(errors.New("this is an error1"))
		ReturnError(errors.New("this is an error2"))

		return
	}

	fmt.Println(test())
}
