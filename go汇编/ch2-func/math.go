/**
@author 张海生<zhanghaisheng@qimao.com>
@dateTime   : 2021/3/24 4:43 下午
@Desc
*/
package main

import "fmt"

func add(a, b int) int // 汇编函数声明

func sub(a, b int) int // 汇编函数声明

func mul(a, b int) int // 汇编函数声明

func main() {
	fmt.Println(add(10, 11))
	fmt.Println(sub(99, 15))
	fmt.Println(mul(11, 12))
}
