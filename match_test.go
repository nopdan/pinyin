package pinyin

import (
	"fmt"
	"testing"
)

func TestMatch(t *testing.T) {
	p := New()

	// // 添加数据文件
	p.AddDataFile("test.txt")
	fmt.Println(p.Match("会计师"))
	// [kuai ji shi]

	// ！！ 顺序最长匹配，人参 / 加
	fmt.Println(p.Match("一个人参加了会议"))
	// [yi ge ren shen jia le hui yi]

	// 添加一条数据
	p.AddData("一个人", "yi", "ge", "ren")
	fmt.Println(p.Match("一个人参加了会议"))
	// [yi ge ren can jia le hui yi]

	fmt.Println(p.Match("了"))
	// [le liao]
	fmt.Println(p.Match("α"))
	// []
}
