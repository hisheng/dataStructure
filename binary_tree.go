package main

import (
	"fmt"
	"time"
)
//https://www.jianshu.com/p/c46676e3f37a
//二叉树结构
type tree struct {
	root *treeNode
	count int
}

type treeNode struct {
	date int
	left *treeNode
	right *treeNode
}

//追加元素 (广度优先，即按层级遍历后添加)
//二叉查找树的插入过程有点类似查找操作。
//新插入的数据一般都是在叶子节点上，
//所以我们只需要从根节点开始，
//依次比较要插入的数据和节点的大小关系。
//如果要插入的数据比节点的数据大，并且节点的右子树为空，就将新数据直接插到右子节点的位置；
//如果不为空，就再递归遍历右子树，查找插入位置。
//同理，如果要插入的数据比节点数值小，并且节点的左子树为空，就将新数据插入到左子节点的位置；如果不为空，就再递归遍历左子树，查找插入位置。
func (btree *tree) Insert(dt int)  {
	node := &treeNode{date:dt}
	 //判断是否有root值
	if btree.root == nil{
		btree.root = node
		btree.count ++
		return
	}

	//1在有root的树
	if btree.root.date < node.date {
		btree.root.right = node
	}else {
		btree.root.left = node
	}

	queue := []*treeNode{btree.root}

	for len(queue) != 0 {
		cur_node := queue[0]
		queue = queue[1:]

		if cur_node.left == nil {
			cur_node.left = node
			return
		} else {
			queue = append(queue, cur_node.left)
		}
		if cur_node.right == nil {
			cur_node.right = node
			return
		} else {
			queue = append(queue, cur_node.right)
		}
	}

}
////广度遍历
//func (this *Tree) BreadthTravel() {
//
//    if this.RootNode == nil {
//        return
//    }
//    queue := []*TreeNode{}
//    queue = append(queue, this.RootNode)
//
//    for len(queue) != 0 {
//        //fmt.Printf("len(queue):%d\n", len(queue))
//        cur_node := queue[0]
//        queue = queue[1:]
//
//        fmt.Printf("%v  ", cur_node.Data)
//
//        if cur_node.LeftChild != nil {
//            queue = append(queue, cur_node.LeftChild)
//        }
//        if cur_node.RightChild != nil {
//            queue = append(queue, cur_node.RightChild)
//        }
//    }
//
//}
func (btree *tree) breadthTravel()  {
	if btree.root == nil {
		return
	}
	queue := []*treeNode{}
	queue = append(queue, btree.root)

	for len(queue) != 0 {
		cur_node := queue[0]
		queue = queue[1:]

		fmt.Printf("%v  ", cur_node.date)
		time.Sleep(time.Second)

		if cur_node.left != nil {
			queue = append(queue, cur_node.left)
		}
		if cur_node.right != nil {
			queue = append(queue, cur_node.right)
		}
	}
}
