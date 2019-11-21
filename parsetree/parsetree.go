package parsetree

import (
	"fmt"
	"io"
	"os"
)

type ParseNode struct {
	left  *ParseNode
	right *ParseNode
	data  string
}

type ParseTree struct {
	root *ParseNode
}

func (t *ParseTree) insert(data string) *ParseTree {
	left, right := &ParseNode{data: "(", left: nil, right: nil}, &ParseNode{data: ")", left: nil, right: nil}
	node := "thiswordexactly"
	if t.root == nil {
		t.root = &ParseNode{data: node, left: left, right: right}
	} else {
		t.root.insertR(data)
	}
	return t
}

func (n *ParseNode) insertL(data string) {
	if n == nil {
		return
	} else if n.left == nil {
		n.left = &ParseNode{data: data, left: nil, right: nil}
	} else {
		n.left.insertL(data)
	}
}
func (n *ParseNode) insertR(data string) {
	if n == nil {
		return
	} else if n.right == nil {
		n.right = &ParseNode{data: data, left: nil, right: nil}
	} else {
		n.right.insertR(data)
	}
}

func print(w io.Writer, node *ParseNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.data)
	print(w, node.left, ns+2, 'L')
	print(w, node.right, ns+2, 'R')
}

func Test() {
	tree := &ParseTree{}
	tree.insert("exactly thiswordexactly")
	print(os.Stdout, tree.root, 0, 'M')
}
