package main

// https://github.com/julycoding/The-Art-Of-Programming-By-July/blob/master/ebook/zh/03.01.md

type node struct {
	isRed bool
	val   int

	parent *node
	left   *node
	right  *node
}

type rbtree struct {
	root *node
}

func (t *rbtree) leftRotate(n *node) {
	right := n.right
	n.right = right.left
	right.left = n
	right.parent = n.parent

	if n.parent == nil {
		t.root = right
	} else if n.parent.left == n {
		n.parent.left = right
	} else {
		n.parent.right = right
	}
	n.parent = right
}

func (t *rbtree) rightRotate(n *node) {
	left := n.left
	n.left = left.right
	left.right = n
	left.parent = n.parent

	if n.parent == nil {
		t.root = left
	} else if n.parent.left == n {
		n.parent.left = left
	} else {
		n.parent.right = left
	}
	n.parent = left
}

func (t *rbtree) insert(n *node) {
	if t.root == nil {
		t.root = n
		return
	}
	cursor := t.root
	for cursor != nil {
		if cursor.val > n.val {
			cursor = cursor.left
		} else {
			cursor = cursor.right
		}
	}
	n.parent = cursor.parent
	if n.val > n.parent.val {
		n.parent.right = n
	} else {
		n.parent.left = n
	}
	n.red = true
	//rebalance rbtree
	t.insertFixUp(n)
}

func (t *rbtree) insertFixUp(n *node) {

}
