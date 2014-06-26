package tree

import (
	"github.com/davecgh/go-spew/spew"
)

type treenode struct {
	value       interface{}
	parent      *treenode
	left, right *treenode
}

type CompareFunc func(a, b interface{}) int
type VisitFunc func(a interface{})

type Tree struct {
	root  *treenode
	size  int
	cfunc CompareFunc
}

func MakeTree(cf CompareFunc) *Tree {
	return &Tree{cfunc: cf}
}

func insertnode(parent, node *treenode, cf CompareFunc) {
	compval := cf(parent.value, node.value)

	if compval == 0 {
		parent.value = node.value
	} else if compval > 0 {
		if parent.left == nil {
			parent.left = node
			node.parent = parent
		} else {
			insertnode(parent.left, node, cf)
		}
	} else {
		if parent.right == nil {
			parent.right = node
			node.parent = parent
		} else {
			insertnode(parent.right, node, cf)
		}
	}
}

func (t *Tree) Insert(val interface{}) {
	node := &treenode{value: val}
	if t.root == nil {
		t.root = node
	} else {
		insertnode(t.root, node, t.cfunc)
	}
}

func foreach(node *treenode, vf VisitFunc) {
	if node == nil {
		return
	}

	foreach(node.left, vf)
	vf(node.value)
	foreach(node.right, vf)
}

func (t *Tree) ForEach(vf VisitFunc) {
	foreach(t.root, vf)
}

func foreachreverse(node *treenode, vf VisitFunc) {
	if node == nil {
		return
	}

	foreachreverse(node.right, vf)
	vf(node.value)
	foreachreverse(node.left, vf)
}

func (t *Tree) ForEachReverse(vf VisitFunc) {
	foreachreverse(t.root, vf)
}

func (t *Tree) findnode(node *treenode, val interface{}) (*treenode, bool) {
	for node != nil {
		compval := t.cfunc(node.value, val)
		if compval == 0 {
			return node, true
		}
		if compval > 0 {
			node = node.left
		} else {
			node = node.right
		}
	}
	return nil, false
}

func (t *Tree) Find(val interface{}) (interface{}, bool) {
	if node, ok := t.findnode(t.root, val); ok {
		return node.value, true
	}
	return nil, false
}

func (t *Tree) unlinknode(node *treenode) {

	var parentsNodePtr **treenode

	if node.parent == nil {
		parentsNodePtr = &t.root
	} else {
		if node.parent.left == node {
			parentsNodePtr = &node.parent.left
		} else {
			parentsNodePtr = &node.parent.right
		}
	}

	// leaf node
	if node.left == nil && node.right == nil {
		*parentsNodePtr = nil
	} else if node.left != nil && node.right == nil {

	} else if node.left == nil && node.right != nil {

	} else {

	}
}

func (t *Tree) Delete(val interface{}) (interface{}, bool) {
	if node, ok := t.findnode(t.root, val); ok {
		t.unlinknode(node)
		return node.value, true
	}
	return nil, false
}

func (t *Tree) Dump() {
	spew.Dump(t)
}
