package data

import "fmt"

//https://www.cnblogs.com/coderJiebao/p/Algorithmofnotes20.html
//https://studygolang.com/articles/3174
//二叉树结构
type tree struct {
	root  *node
	count int
}

func NewTree() *tree {
	return &tree{root: nil}
}

//追加元素 (广度优先，即按层级遍历后添加)
//二叉查找树的插入过程有点类似查找操作。
//新插入的数据一般都是在叶子节点上，
//所以我们只需要从根节点开始，
//依次比较要插入的数据和节点的大小关系。
//如果要插入的数据比节点的数据大，并且节点的右子树为空，就将新数据直接插到右子节点的位置；
//如果不为空，就再递归遍历右子树，查找插入位置。
//同理，如果要插入的数据比节点数值小，并且节点的左子树为空，就将新数据插入到左子节点的位置；如果不为空，就再递归遍历左子树，查找插入位置。
func (btree *tree) Insert(dt int) {
	//判断是否有root值
	if btree.root == nil {
		btree.root = NewNode(dt)
		btree.count++ //目前的写法，好像没法统计这个count++,因为后面是子节点的添加，没有带进这个 tree
		return
	}
	btree.root.Insert(NewNode(dt))
}

func (btree *tree) Search(dt int) *node {
	return btree.root.Search(dt)
}

//广度遍历 pre-order / in-order / post-order / breadth-first

//前(中后)序遍历，是按照 根节点 在哪命名的
/*前序遍历：根结点 ---> 左子树 ---> 右子树

中序遍历：左子树---> 根结点 ---> 右子树

后序遍历：左子树 ---> 右子树 ---> 根结点
*/

//前序遍历：根节点->左子树->右子树
func PreOrder(nd *node) {
	if nd == nil {
		return
	}
	fmt.Println(nd.data)
	PreOrder(nd.left)
	PreOrder(nd.right)
}

//中序遍历：左子树---> 根结点 ---> 右子树
func InOrder(nd *node) {
	if nd == nil {
		return
	}
	PreOrder(nd.left)
	fmt.Println(nd.data)
	PreOrder(nd.right)
}

//后序遍历：左子树 ---> 右子树 ---> 根结点
func PostOrder(nd *node) {
	if nd == nil {
		return
	}
	PreOrder(nd.left)
	PreOrder(nd.right)
	fmt.Println(nd.data)
}
