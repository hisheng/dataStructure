package zskiplist

import (
	"math/rand"
	"strings"
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

type Zrangespec struct {
	min, max     float64
	minex, maxex int
}
type Zlexrangespec struct {
	min, max     string
	minex, maxex int
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
	update := make([]*ZskiplistNode, ZSKIPLIST_MAXLEVEL)
	rank := make([]int, ZSKIPLIST_MAXLEVEL)
	var x *ZskiplistNode
	var i, level int

	x = zsl.Header

	for i = zsl.Level - 1; i >= 0; i-- {
		if i == zsl.Level-1 {
			rank[i] = 0
		} else {
			rank[i] = rank[i+1]
		}

		for x.Levels[i].Forward != nil && (x.Levels[i].Forward.Score < score || (x.Levels[i].Forward.Score == score && strings.Compare(x.Levels[i].Forward.Ele, ele) < 0)) {
			rank[i] += x.Levels[i].Span
			x = x.Levels[i].Forward
		}
		update[i] = x
	}

	level = ZslRandomLevel()
	if level > zsl.Level {
		for i = zsl.Level; i < level; i++ {
			rank[i] = 0
			update[i] = zsl.Header
			update[i].Levels[i].Span = zsl.Length
		}
		zsl.Level = level
	}
	x = ZslCreateNode(level, score, ele)
	for i = 0; i < level; i++ {
		x.Levels[i].Forward = update[i].Levels[i].Forward
		update[i].Levels[i].Forward = x

		x.Levels[i].Span = update[i].Levels[i].Span - (rank[0] - rank[i])
		update[i].Levels[i].Span = (rank[0] - rank[i]) + 1
	}

	for i = level; i < zsl.Level; i++ {
		update[i].Levels[i].Span++
	}

	if update[0] == zsl.Header {
		x.Backward = nil
	} else {
		x.Backward = update[0]
	}

	if x.Levels[0].Forward != nil {
		x.Levels[0].Forward.Backward = x
	} else {
		zsl.Tail = x
	}
	zsl.Length++
	return x

}

//zskiplist *zsl, zskiplistNode *x, zskiplistNode **update
func zslDeleteNode(zsl *Zskiplist, x *ZskiplistNode, update []*ZskiplistNode) {
	var i int
	for i = 0; i < zsl.Level; i++ {
		if update[i].Levels[i].Forward == x {
			update[i].Levels[i].Span += x.Levels[i].Span - 1
			update[i].Levels[i].Forward = x.Levels[i].Forward
		} else {
			update[i].Levels[i].Span -= 1
		}
	}
	if x.Levels[0].Forward != nil {
		x.Levels[i].Forward.Backward = x.Backward
	} else {
		zsl.Tail = x.Backward
	}

	for zsl.Level > 1 && zsl.Header.Levels[zsl.Level-1].Forward == nil {
		zsl.Level--
	}
	zsl.Length--
}

func zslDelete(zsl *Zskiplist, score float64, ele string, node []*ZskiplistNode) int {
	var update []*ZskiplistNode = make([]*ZskiplistNode, ZSKIPLIST_MAXLEVEL)
	var x *ZskiplistNode
	var i int
	x = zsl.Header

	for i = zsl.Level - 1; i >= 0; i-- {
		for x.Levels[i].Forward != nil && (x.Levels[i].Forward.Score < score || (x.Levels[i].Forward.Score == score && strings.Compare(x.Levels[i].Forward.Ele, ele) < 0)) {
			x = x.Levels[i].Forward
		}
		update[i] = x
	}

	x = x.Levels[0].Forward
	if x != nil && score == x.Score && strings.Compare(x.Ele, ele) == 0 {
		zslDeleteNode(zsl, x, update)
		if node == nil {
			zslFreeNode(x)
		} else {
			node = append(node, x)
			return 1
		}
	}
	return 0
}

func zslUpdateScore(zsl *Zskiplist, curscore float64, ele string, newscore float64) *ZskiplistNode {
	var update []*ZskiplistNode = make([]*ZskiplistNode, ZSKIPLIST_MAXLEVEL)
	var x *ZskiplistNode
	var i int

	x = zsl.Header
	for i = zsl.Level - 1; i >= 0; i-- {
		for x.Levels[i].Forward != nil && (x.Levels[i].Forward.Score < curscore || (x.Levels[i].Forward.Score == curscore && strings.Compare(x.Levels[0].Forward.Ele, ele) < 0)) {
			x = x.Levels[i].Forward
		}
		update[i] = x
	}

	x = x.Levels[0].Forward

	if (x.Backward == nil || x.Backward.Score < newscore) && (x.Levels[0].Forward == nil || x.Levels[0].Forward.Score > newscore) {
		x.Score = newscore
		return x
	}

	zslDeleteNode(zsl, x, update)
	var newNode *ZskiplistNode = ZslInsert(zsl, newscore, x.Ele)
	x.Ele = ""
	zslFreeNode(x)
	return newNode
}

func zslValueGteMin(value float64, spec *Zrangespec) bool {
	if spec.maxex > 0 {
		return value > spec.max
	}
	return value >= spec.max
}

func zslValueGteMax(value float64, spec *Zrangespec) bool {
	if spec.maxex > 0 {
		return value < spec.max
	}
	return value <= spec.max
}

func zslIsInRange(zsl *Zskiplist, spec *Zrangespec) bool {
	var x *ZskiplistNode
	if spec.min > spec.max || (spec.min == spec.max && (spec.minex > 0 || spec.maxex > 0)) {
		return false
	}
	x = zsl.Tail
	if x == nil || !zslValueGteMin(x.Score, spec) {
		return false
	}
	x = zsl.Header.Levels[0].Forward
	if x == nil || !zslValueGteMax(x.Score, spec) {
		return false
	}
	return true
}

func zslFirstInRange(zsl *Zskiplist, spec *Zrangespec) *ZskiplistNode {
	var x *ZskiplistNode
	var i int
	if !zslIsInRange(zsl, spec) {
		return nil
	}

	x = zsl.Header
	for i = zsl.Level - 1; i >= 0; i-- {
		for x.Levels[i].Forward != nil && !zslValueGteMin(x.Levels[i].Forward.Score, spec) {
			x = x.Levels[i].Forward
		}
	}
	x = x.Levels[0].Forward
	if !zslValueGteMax(x.Score, spec) {
		return nil
	}
	return x
}

func zslLastInRange(zsl *Zskiplist, spec *Zrangespec) *ZskiplistNode {
	var x *ZskiplistNode
	var i int
	if !zslIsInRange(zsl, spec) {
		return nil
	}
	x = zsl.Header
	for i = zsl.Level - 1; i >= 0; i-- {
		for x.Levels[i].Forward != nil && zslValueGteMax(x.Levels[i].Forward.Score, spec) {
			x = x.Levels[i].Forward
		}
	}
	return x
}

func ZslGetRank(zsl *Zskiplist, score float64, ele string) int {
	var x *ZskiplistNode
	var rank int
	var i int

	x = zsl.Header
	for i = zsl.Level - 1; i >= 0; i-- {
		for x.Levels[i].Forward != nil && (x.Levels[i].Forward.Score < score || (x.Levels[i].Forward.Score == score && (strings.Compare(x.Levels[i].Forward.Ele, ele) <= 0))) {
			rank += x.Levels[i].Span
			x = x.Levels[i].Forward
		}
		if x.Ele != "" && strings.Compare(x.Ele, ele) == 0 {
			return rank
		}
	}
	return 0
}

func ZslGetElementByRank(zsl *Zskiplist, rank int) *ZskiplistNode {
	var x *ZskiplistNode
	var traversed int = 0
	var i int

	x = zsl.Header
	for i = zsl.Level - 1; i >= 0; i-- {
		for x.Levels[0].Forward != nil && (traversed+x.Levels[i].Span) <= rank {
			traversed += x.Levels[i].Span
			x = x.Levels[i].Forward
		}
		if traversed == rank {
			return x
		}
	}
	return nil
}

//static int zslParseRange(robj *min, robj *max, zrangespec *spec) {

//func ZslDeleteRangeByScore(zsl *Zskiplist, spec Zrangespec) {
//
//}

//unsigned long zslDeleteRangeByLex(zskiplist *zsl, zlexrangespec *range, dict *dict) {

//unsigned long zslDeleteRangeByRank(zskiplist *zsl, unsigned int start, unsigned int end, dict *dict) {
