package redis

/*
通过 一个链表 来实现 redis 的 zset
*/

type sortlist struct {
	key string
	head *sortlistNode
	tail *sortlistNode
	len int
}

type sortlistNode struct {
	member string
	score int
	pre *sortlistNode
	next *sortlistNode
}


// zadd kaoshi:score  hisheng  100
// zadd kaoshi:score  zhang  99
// zadd kaoshi:score  hai  98

func (s *sortlist) Zadd(member string,score int)  {
	//1 节点
	node := &sortlistNode{
		member: member,
		score: score,
	}

	//2 查找 找出 zadd 的位置
	insertNode := s.Find(score)

	//3 插入
	node.pre = insertNode
	node.next = insertNode.next
	insertNode.next = node

	//4 len
	s.len++
}

func (s *sortlist) Find(score int) *sortlistNode{
	return &sortlistNode{
		member: "",
		score: 0,
	}
}





