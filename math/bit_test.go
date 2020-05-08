package math

import (
	"testing"
)

/*测试 bit*/

func TestBit(t *testing.T) {
	t.Log("start")
	t.Logf("%b", 0) //0
	t.Logf("%b", 1) //1
	t.Logf("%b", 2) //10
	t.Logf("%b", 3) //11

	/*
		int是带符号的，表示范围是：-2147483648到2147483648，即-2^31到2^31次方。
		uint则是不带符号的，表示范围是：2^32即0到4294967295。
	*/
	var one uint
	one = 10000
	t.Logf("%b", one) //10011100010000  13位，uint最多是32位

	//左移 i << n; 对i的所有位向左移n次
	/*
		1<<N = 2^N
		-----------
		1 左移多少位等于2的多少次方

		1024>>N = 1024/2^N
		---------------
		右移N位 相当于除以2的N次方
	*/
	t.Log("左移")
	var i int64
	t.Logf("%b", i) //0
	i = 1
	t.Logf("%b %d", i, i)       // 1 1
	t.Logf("%b %d", i<<1, i<<1) //10 2 左移动 每个移位位置代表2的幂,左移动增加
	t.Logf("%b %d", i<<2, i<<2) // 100 4

	// & AND   同一位上都是1 才1 否咋0
	t.Log("& AND")
	c, d := 1, 1
	t.Logf("c&d %b %b %b", c&d, c, d) //c&d 1 1 1
	c, d = 1, 0
	t.Logf("c&d %b %b %b", c&d, c, d) //c&d 0 1 0
	c, d = 5, 3
	t.Logf("c&d %b %b %b", c&d, c, d) //c&d 1 101 011
	c, d = 5, 4
	t.Logf("c&d %b %b %b", c&d, c, d) //c&d 100 101 100  十进制4
	c, d = 50, 43
	t.Logf("c&d %b %b %b", c&d, c, d) //c&d 100010 110010 101011 十进制34

	// | OR  同一位上有一个是1 就1 都是0 才0
	t.Log("| OR")
	c, d = 0, 0
	t.Logf("c|d %b %b %b", c|d, c, d) //c|d 0 0 0
	c, d = 1, 1
	t.Logf("c|d %b %b %b", c|d, c, d) //c|d 1 1 1
	c, d = 1, 0
	t.Logf("c|d %b %b %b", c|d, c, d) //c|d 1 1 0
	c, d = 5, 3
	t.Logf("c|d %b %b %b", c|d, c, d) //c|d 111 101 011  十进制7
	c, d = 5, 4
	t.Logf("c|d %b %b %b", c|d, c, d) //c|d 101 101 100  十进制5
	c, d = 50, 43
	t.Logf("c|d %b %b %b", c|d, c, d) //c|d 111011 110010 101011  十进制59

	// ^   XOR （异或） 假设给定 a, b 当且只有 a!=b 时  XOR(a,b) = 1 否则等于 0

	//判断写法 二元运算符号
	//a |= b  ----->  a = a | b  , a 或者 b 只要有一个为 1, 那么，a 的最终结果就为 1
	//a &= b  ----->  a = a & b  , a 和 b 二者必须都为 1, 那么，a 的最终结果才为 1
	//a ^= b  ----->  a = a ^ b  , 当且仅当 a 和 b 的值不一致时，a 的最终结果才为1，否则为0
	t.Log("判断写法 运算符号")
	a, b := 1, 0
	a |= b
	t.Logf("%b %b", a, b) // 1 0
	a, b = 1, 0
	a &= b
	t.Logf("%b %b", a, b) // 0 0
	a, b = 1, 0
	a ^= b
	t.Logf("%b %b", a, b) //  1 0

	a, b = 2, 3
	t.Logf("%b %b", a, b) //10 11
	a |= b
	t.Logf("%b %b", a, b) //11 11
	a, b = 2, 3
	a &= b
	t.Logf("%b %b", a, b) // 10 11
	a, b = 2, 3
	a ^= b
	t.Logf("%b %b", a, b) // 1 11

	//高级
	index := 256

	type BitSets []int64
	bs := make(BitSets, index/64+1)
	t.Logf("%b", bs) // [0 0 0 0 0]

	index, bit := index/64, index%64
	bs[index] |= 1 << bit
	t.Logf("%b", bs) // [0 0 0 0 1]

	t.Log("奇偶数")
	t.Log(IsOdd(101)) // true
}

//odd number ：指不能被2整除的数 ，数学表达形式为：2k+1， 奇数可以分为正奇数和负奇
func IsOdd(i int) bool {
	return (i & 1) == 1 //i是否为奇数取决于二进制的最后一位是1还是0 是1则为奇数 0则为偶数
}
