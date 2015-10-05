package sandbox

// messing with various aspects of The Go language
// Here Be Dragons

import (
	//	"errors"
	"fmt"
	"reflect"
	"testing"
)

// try/catch/finally impl.
// dont do this. its wrong. maybe

func Throw(e interface{}) {
	panic(e)
}

type Tryblock struct {
	try      func()
	catchers []interface{}
}

type UncaughtException struct {
}

func Try(f func()) *Tryblock {
	return &Tryblock{try: f}
}

func (t *Tryblock) Catch(f interface{}) *Tryblock {
	t.catchers = append(t.catchers, f)
	return t
}

func convertTo(i interface{}, to reflect.Type) (interface{}, bool) {
	from := reflect.TypeOf(i)
	if from.ConvertibleTo(to) {
		return reflect.ValueOf(i).Convert(to).Interface(), true
	}
	return nil, false
}

func callFunction(f interface{}, arg interface{}) bool {
	fval := reflect.ValueOf(f)

	if fval.Type().NumIn() != 1 {
		return false
	}

	if arg, ok := convertTo(arg, fval.Type().In(0)); ok {
		argVals := []reflect.Value{reflect.ValueOf(arg)}
		fval.Call(argVals)
		return true
	}

	return false
}

func (t *Tryblock) Finally(finally func()) {
	defer func() {
		if r := recover(); r != nil {
			called := false
			for _, f := range t.catchers {
				called = callFunction(f, r)
				if called {
					break
				}
			}

			finally()

			if !called {
				panic(r)
			}
		}
	}()

	t.try()
	finally()
}

func TestThrow(t *testing.T) {

	Try(func() {
		Throw("dddd")
	}).Catch(func(e int) {
		fmt.Println("caught:", reflect.TypeOf(e), e)
	}).Catch(func(e string) {
		fmt.Println("caught:", reflect.TypeOf(e), e)
	}).Catch(func(e error) {
		fmt.Println("caught:", reflect.TypeOf(e), e)
	}).Finally(func() {
		fmt.Println("finally")
	})
}
