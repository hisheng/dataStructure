/**
猜拳
@author 张海生<zhanghaisheng@qimao.com>
@dateTime   : 2021/3/25 3:48 下午
@Desc
*/
package quan

import (
	"bufio"
	"os"
	"testing"
)

func TestQuan(t *testing.T) {
	// 创建一个map 指定key为string类型 val为int类型
	counts := make(map[string]int)
	// 从标准输入流中接收输入数据
	input := bufio.NewScanner(os.Stdin)

	t.Logf("Please type in something:\n")

	// 逐行扫描
	for input.Scan() {
		line := input.Text()

		// 输入bye时 结束
		if line == "bye" {
			break
		}

		// 更新key对应的val 新key对应的val是默认0值
		counts[line]++
	}

	// 遍历map统计数据
	for line, n := range counts {
		t.Logf("%d : %s\n", n, line)
	}
}
