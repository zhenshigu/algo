package main

//平衡二叉树
type node struct {
	data   int
	left   *node
	right  *node
	height int
}

func (n *node) setLeft(left *node) {
	n.left = left
}

func (n *node) setRight(right *node) {
	n.right = right
}
func (n *node) setHeight(height int) {
	n.height = height
}
func (n *node) setData(data int) {
	n.data = data
}
func (n *node) getHeight() int {
	if n == nil {
		return 0
	}
	return n.height
}
func (n *node) leftRotation() *node {
	leftNode := n.left
	n.left = leftNode.right
	leftNode.right = n
	n.setHeight(getMaxHeight(n.left, n.right) + 1)
	leftNode.setHeight(getMaxHeight(leftNode.left, leftNode.right) + 1)
	return leftNode
}
func (n *node) rightRotation() *node {
	rightNode := n.right
	n.right = rightNode.left
	rightNode.left = n
	n.setHeight(getMaxHeight(n.left, n.right) + 1)
	rightNode.setHeight(getMaxHeight(rightNode.left, rightNode.right) + 1)
	return rightNode

}
func (n *node) leftRightRotation() *node {
	rightNode := n.right
	n.right = rightNode.leftRotation()
	return n.rightRotation()
}
func (n *node) rightLeftRotation() *node {
	leftNode := n.left
	n.left = leftNode.rightRotation()
	return n.leftRotation()
}
func (n *node) new(data int) *node {
	return &node{
		data:   data,
		height: 1,
		left:   nil,
		right:  nil,
	}
}

func (n *node) insertNode(data int) *node {
	if n == nil {
		return n.new(data)
	}
	if data < n.data {
		n.left = n.left.insertNode(data)
	} else {
		n.right = n.right.insertNode(data)
	}
	return n.rotationNode()
}

func (n *node) removeNode(data int) *node {
	if n == nil {
		panic("the node is nil")
	}
	if n.data == data {
		if n.left != nil && n.right != nil {
			leftMostNode := n.right.leftMost()
			n.data = leftMostNode.data
			n.right = n.right.removeNode(leftMostNode.data)
		} else if n.left != nil {
			n = n.left
		} else {
			n = n.right
		}
	} else if data < n.data {
		if n.left == nil {
			panic("no such node")
		}
		n.left = n.left.removeNode(data)
	} else if data > n.data {
		if n.right == nil {
			panic("no such node")
		}
		n.right = n.right.removeNode(data)
	}
	return n.rotationNode()

}
func (n *node) rotationNode() *node {
	if n == nil {
		return n
	}
	if n.left.getHeight()-n.right.getHeight() == 2 {
		if n.left.left.getHeight() > n.left.right.getHeight() {
			return n.leftRotation()
		} else {
			return n.rightLeftRotation()
		}
	} else if n.right.getHeight()-n.left.getHeight() == 2 {
		if n.right.right.getHeight() > n.right.left.getHeight() {
			return n.rightRotation()
		} else {
			return n.leftRightRotation()
		}
	}
	n.height = getMaxHeight(n.left, n.right) + 1
	return n
}
func getMaxHeight(n1 *node, n2 *node) int {
	if n1.getHeight() > n2.getHeight() {
		return n1.getHeight()
	} else {
		return n2.getHeight()
	}
}

func (n *node) leftMost() *node {
	for n != nil && n.left != nil {
		n = n.left
	}
	return n
}
