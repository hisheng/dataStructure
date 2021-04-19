/**
@author 张海生<zhanghaisheng@qimao.com>
@dateTime   : 2021/4/19 10:04 上午
@Desc
*/
package sync

import (
	"sync"
	"testing"
)

func TestSync(t *testing.T) {
	//1 互斥锁的机制  临界区，临界区就是一个被共享的资源，这部分被保护起来。
	mutex := sync.Mutex{} //{0 0}
	t.Log(mutex)

	var mu sync.Mutex //{0 0}
	t.Log(mu)

	m := new(sync.Mutex) //&{0 0}
	t.Log(m)
}
