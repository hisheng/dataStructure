/**
@author 张海生<zhanghaisheng@qimao.com>
@dateTime   : 2021/3/25 7:06 下午
@Desc
*/
package unit_test

import "testing"

func BenchmarkName(b *testing.B) {
	b.Log(b.N)
	b.ResetTimer()
	for i := 0; i < 10; i++ {
		b.Log(i)
	}
	b.StopTimer()
}
