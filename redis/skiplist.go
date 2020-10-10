package redis

import (
	"strings"
)

const ZSKIPLIST_MAXLEVEL = 64
const ZSKIPLIST_P = 0.25
type zskiplistNode struct {
	ele string
	score float64
	backward *zskiplistNode
	level []zskiplistLevel
}

type zskiplistLevel struct {
	forward *zskiplistNode
	span int64
}

type zskiplist struct {
	header,tail * zskiplistNode
	length int64
	level int
}

type zset struct {
	dict *dict
	zsl *zskiplist
}

type zrangespec struct {
	min,max  float64
	minex,maxex int   ///// (32 (34    32<x<34 minex=1  maxex=1，就是这个地方用的
}





func zslCreateNode(level int,score float64,ele string) *zskiplistNode {
	levels := make([]zskiplistLevel,level)
	return &zskiplistNode{
		ele: ele,
		score: score,
		level: levels,
		backward: nil,
	}
}

func zslCreate() *zskiplist  {
	var zsl *zskiplist
	var j int
	zsl = &zskiplist{
		level: 1,
		length: 0,
		header: zslCreateNode(64,0,""),
		tail: nil,
	}
	for j = 0;j<ZSKIPLIST_MAXLEVEL;j++ {
		zsl.header.level[j].forward = nil
		zsl.header.level[j].span = 0
	}
	zsl.header.backward = nil
	return zsl
}

func zslRandomLevel() int {
	var level int = 1
	return level
}

func zslInsert(zsl *zskiplist,score float64,ele string) *zskiplistNode {
	var update = make([]*zskiplistNode,64)
	var x *zskiplistNode
	var i int
	var level int
	var rank = make([]int64,64)

	//1 update[],rank[],level
	x = zsl.header
	for i = zsl.level-1;i>=0;i-- {
		if i == zsl.level-1 {
			rank[i] = 0
		}else {
			rank[i] = rank[i+1]
		}
		for x.level[i].forward != nil && (x.level[i].forward.score < score || (x.level[i].forward.score == score && strings.Compare(x.level[i].forward.ele,ele)<0)) {
			rank[i] += x.level[i].span
			x = x.level[i].forward
		}
		update[i] = x
	}

	level = zslRandomLevel()
	if level > zsl.level {
		for i = zsl.level;i<level;i++ {
			rank[i] =  0
			update[i] = zsl.header
			update[i].level[i].span = zsl.length
		}
	}
	zsl.level = level

	//2创建节点，更新值
	x = zslCreateNode(level,score,ele)
	for i=0;i<level;i++ {
		x.level[i].forward = update[i].level[i].forward
		update[i].level[i].forward = x

		x.level[i].span = update[i].level[i].span - (rank[0] - rank[i])
		update[i].level[i].span = (rank[0] - rank[i]) + 1
	}
	for i = level;i<zsl.level;i++ {
		update[i].level[i].span++
	}

	//3 更新backward或tail
	if update[0] ==zsl.header {
		x.backward = nil
	}else{
		x.backward = update[0]
	}

	if x.level[0].forward != nil {
		x.level[0].forward.backward = x
	}else {
		zsl.tail = x
	}

	//4 更新length
	return x
}

















