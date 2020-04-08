package data

type node struct {
	data  int
	left  *node
	right *node
}

func NewNode(data int) *node {
	return &node{data: data}
}

func (nd *node) Insert(newNode *node) {
	if newNode.data == nd.data {
		return
	}
	if newNode.data > nd.data {
		if nd.right == nil {
			nd.right = newNode
		} else {
			nd.right.Insert(newNode)
		}
	} else {
		if nd.left == nil {
			nd.left = newNode
		} else {
			nd.left.Insert(newNode)
		}
	}
}

func (nd *node) Search(dt int) *node {
	if dt == nd.data {
		return nd
	}

	if dt > nd.data {
		return nd.right.Search(dt)
	}

	if dt < nd.data {
		return nd.left.Search(dt)
	}

	return nil
}
