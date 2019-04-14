package rbtree

// node is red/black
// root is black
// leaf is black
// red node's child is black
// all the simple way from node to leaf contains same number of black nodes

// a N node rbtree's height is less than 2log(N)+1

// https://blog.csdn.net/v_JULY_v/article/details/6109153
// https://juejin.im/entry/58371f13a22b9d006882902d#%E4%BA%8C%E5%8F%89%E6%9F%A5%E6%89%BE%E6%A0%91%E7%9A%84%E6%8F%92%E5%85%A5

type Node struct {
	IsRed  bool
	Parent *Node
	Left   *Node
	Right  *Node
	Val    int
}

type RBTree struct {
	Root *Node
}

func (t *RBTree) leftRotate(x *Node) {
	y := x.Right
	x.Right = y.Left
	if y.Left != nil {
		y.Left.Parent = x
	}
	y.Parent = x.Parent
	if x.Parent == nil {
		t.Root = y
	} else if x.Parent.Left == x {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}
	y.Left = x
	x.Parent = y
}

func (t *RBTree) rightRotate(x *Node) {
	y := x.Left
	x.Left = y.Right
	if y.Right != nil {
		y.Right.Parent = x
	}
	y.Parent = x.Parent
	if x.Parent == nil {
		t.Root = y
	} else if x.Parent.Left == x {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}
	y.Right = x
	x.Parent = y
}

func (t *RBTree) Insert(x *Node) {
	var cursorParent *Node
	cursor := t.Root

	for cursor != nil {
		cursorParent = cursor.Parent
		if cursor.Val > x.Val {
			cursor = cursor.Left
		} else if cursor.Val < x.Val {
			cursor = cursor.Right
		} else {
			return
		}
	}

	x.Parent = cursorParent
	if cursorParent == nil {
		t.Root = x
	} else if x.Val < cursorParent.Val {
		cursorParent.Left = x
	} else {
		cursorParent.Right = x
	}

	t.fixAfterInsertion(x)
}

func (t *RBTree) fixAfterInsertion(x *Node) {
	x.IsRed = true
	for x != nil && x != t.Root && x.Parent.IsRed {
		// 插入发生在树左侧
		if x.Parent == x.Parent.Parent.Left { // x的父亲是x爷爷的左孩子
			y := x.Parent.Parent.Right
			if y.IsRed { //变色即可
				x.Parent.IsRed = false
				y.IsRed = false
				x.Parent.Parent.IsRed = true
				x = x.Parent.Parent
			} else { //如果叔叔节点 y 不是红色，就需要右旋，让父亲节点变成根节点，爷爷节点去右子树去，然后把父亲节点变成黑色、爷爷节点变成红色
				//特殊情况：x 是父亲节点的右孩子，需要对父亲节点进行左旋，把 x 移动到左子树
				if x == x.Parent.Right {
					x = x.Parent
					t.leftRotate(x)
				}
				x.Parent.IsRed = false
				x.Parent.Parent.IsRed = true
				t.rightRotate(x.Parent.Parent)
			}
		} else { // 插入发生在树右侧
			y := x.Parent.Parent.Left
			if y.IsRed {
				x.Parent.IsRed = false
				y.IsRed = false
				x.Parent.Parent.IsRed = true
				x = x.Parent.Parent
			} else {
				if x == x.Parent.Left {
					x = x.Parent
					t.rightRotate(x)
				}
				x.Parent.IsRed = false
				x.Parent.Parent.IsRed = true
				t.leftRotate(x.Parent.Parent)
			}
		}
	}
	t.Root.IsRed = false
}
