/**
@author 张海生<zhanghaisheng@qimao.com>
@dateTime   : 2021/3/11 5:40 下午
@Desc
*/
package _map

import "testing"

func TestMapArray(t *testing.T) {
	ss := []string{"aaaaaddddddddaaaaaa", "baaaaaffffffffffaaaaa", "c333333ggggggggggggg333", "d44ggggggggggg4444", "e55553333333333333333", "ffffffffff", "ggggggg", "h333", "i444444444", "j4444444444", "k444444444"}

	ms := make(map[string]string, len(ss))
	for _, s := range ss {
		if "baaaaaffffffffffaaaaa" == s {
			t.Log(1)
			ms[s] = s
		}
	}

	v, ok := ms["baaaaaffffffffffaaaaa"]
	t.Log(v, ok)

}
