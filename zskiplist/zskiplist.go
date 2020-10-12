package zskiplist

import (
	"math/rand"
)

const ZSKIPLIST_MAXLEVEL = 64
const ZSKIPLIST_P = 0.25

type ZskiplistNode struct {
	Ele      string
	Score    float64
	Backward *ZskiplistNode
	Levels   map[int]*ZskiplistLevel
}

type ZskiplistLevel struct {
	Forward *ZskiplistNode
	Span    int
}

type Zskiplist struct {
	Header, Tail *ZskiplistNode
	Length       int
	Level        int
}

func ZslCreate() *Zskiplist {
	zsl := &Zskiplist{
		Level:  1,
		Length: 0,
		Header: ZslCreateNode(ZSKIPLIST_MAXLEVEL, 0, ""),
		Tail:   nil,
	}
	zsl.Header.Backward = nil
	return zsl
}

func ZslCreateNode(level int, score float64, ele string) *ZskiplistNode {
	var levels map[int]*ZskiplistLevel = make(map[int]*ZskiplistLevel, 0)
	for j := 0; j < level; j++ {
		l := &ZskiplistLevel{
			Forward: nil,
			Span:    0,
		}
		levels[j] = l
	}
	return &ZskiplistNode{
		Score:  score,
		Ele:    ele,
		Levels: levels,
	}
}

func ZslRandomLevel() int {
	level := 1
	for float64(rand.Int()&0xFFFF) < (ZSKIPLIST_P * 0xFFFF) {
		level += 1
	}
	if level < ZSKIPLIST_MAXLEVEL {
		return level
	}
	return ZSKIPLIST_MAXLEVEL
}

func ZslInsert(zsl *Zskiplist, score float64, ele string) *ZskiplistNode {
	//update := make([]*zskiplistNode, ZSKIPLIST_MAXLEVEL)
	//rank := make([]int, ZSKIPLIST_MAXLEVEL)
	//var x zskiplistNode
	level := ZslRandomLevel()
	x := ZslCreateNode(level, score, ele)
	x.Backward = zsl.Tail
	if zsl.Tail != nil {
		zsl.Tail.Levels[0].Forward = x
	}
	zsl.Tail = x
	return x
}
