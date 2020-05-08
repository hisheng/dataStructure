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
	if nd == nil {
		return nil
	}
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

/*这样的方法，就会 无序了? 以后数据就乱了，所以 不是这样的改的*/
func (nd *node) Update(dt, newDt int) *node {
	//1搜索
	dtnd := nd.Search(dt)
	//2 update
	if dtnd != nil {
		dtnd.data = newDt
	}
	return dtnd
}
