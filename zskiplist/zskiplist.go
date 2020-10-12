package main

import "math/rand"

const ZSKIPLIST_MAXLEVEL = 64
const ZSKIPLIST_P = 0.25

type zskiplistNode struct {
	ele      string
	score    float64
	backward *zskiplistNode
	levels   []*zskiplistLevel
}

type zskiplistLevel struct {
	forward *zskiplistNode
	span    int
}

type zskiplist struct {
	header, tail *zskiplistNode
	length       int
	level        int
}

func zslCreate() *zskiplist {
	zsl := &zskiplist{
		level:  1,
		length: 0,
		header: zslCreateNode(ZSKIPLIST_MAXLEVEL, 0, ""),
		tail:   nil,
	}
	for j := 0; j < ZSKIPLIST_MAXLEVEL; j++ {
		zsl.header.levels[j].forward = nil
		zsl.header.levels[j].span = 0
	}
	zsl.header.backward = nil
	return zsl
}

func zslCreateNode(level int, score float64, ele string) *zskiplistNode {
	var levels = make([]*zskiplistLevel, level)
	return &zskiplistNode{
		score:  score,
		ele:    ele,
		levels: levels,
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
