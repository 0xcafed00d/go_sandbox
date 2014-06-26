package tree

import (
	"fmt"
	"testing"
)

func IntCompareFunc(a, b interface{}) int {
	if a.(int) < b.(int) {
		return -1
	} else if a.(int) == b.(int) {
		return 0
	} else {
		return 1
	}
}

func TestTree(t *testing.T) {

	tree := MakeTree(IntCompareFunc)
	tree.Insert(4)
	tree.Insert(3)
	tree.Insert(6)
	tree.Insert(9)

	if val, ok := tree.Find(6); ok {
		t.Log("tree.Find(9) returns: ", val.(int), ok)
	} else {
		t.Error("Find(9) failed")
	}

	tree.ForEach(func(a interface{}) {
		fmt.Print(a, " ")
	})

	fmt.Println()

	tree.ForEachReverse(func(a interface{}) {
		fmt.Print(a, " ")
	})

	tree.Dump()

	fmt.Println()
}
