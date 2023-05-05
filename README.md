<div align="center">

<img src="logo.png"  width="150" height="150"> </img>

## 汉字转拼音

[![GitHub Repo stars](https://img.shields.io/github/stars/flowerime/pinyin)](https://github.com/flowerime/pinyin/stargazers)
[![GitHub forks](https://img.shields.io/github/forks/flowerime/pinyin)](https://github.com/flowerime/pinyin/network/members)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/flowerime/pinyin)](https://github.com/flowerime/pinyin/releases)
![GitHub repo size](https://img.shields.io/github/repo-size/flowerime/pinyin)
![GitHub](https://img.shields.io/github/license/flowerime/pinyin)

主要针对拼音输入法词库

</div>

- 一个字有多个音，但是参与组词时只有一个音。
- 一个词只有一组音。

## 安装

```sh
go get -u github.com/flowerime/pinyin
```

## 使用

默认不带数据，需要使用 `AddData` 和 `AddDataFile` 添加。

```go
// 导入
import "github.com/flowerime/pinyin"

// 基础
func main() {
	p := pinyin.New()

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

```

## 数据格式

纯文本格式，词(或单字)与拼音之间用 tab 分开，拼音中间用空格分隔。

[例子](test.txt)

## 逻辑

### 添加数据

单字：追加新添加的读音  
词组：覆盖旧的读音

### 匹配

- 词库中有的直接匹配
- 若没有则按词长，最长匹配
- 匹配到单字只取第一个读音

