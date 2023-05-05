package pinyin

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/nopdan/ku"
)

type Pinyin struct {
	Chars map[rune][]string
	Words *trie
}

// 需要自己添加数据
func New() *Pinyin {
	p := new(Pinyin)
	p.Chars = make(map[rune][]string)
	p.Words = newTrie()
	return p
}

// 添加一条数据
func (p *Pinyin) AddData(word string, pinyin ...string) {
	chars := []rune(word)
	// 词组
	if len(chars) != 1 {
		p.Words.insert(word, pinyin)
		return
	}
	// 单字
	char := chars[0]
	if _, ok := p.Chars[char]; !ok {
		p.Chars[char] = pinyin
	} else {
		p.Chars[char] = append(p.Chars[char], pinyin...)
		p.Chars[char] = ku.Unique(p.Chars[char])
	}
}

// 添加数据文件
func (p *Pinyin) AddDataFile(path string) {
	rd, err := ku.Read(path)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	scan := bufio.NewScanner(rd)
	for scan.Scan() {
		line := scan.Text()
		tmp := strings.Split(line, "\t")
		if len(tmp) != 2 || line == "" {
			fmt.Printf("数据有误: %v\n", line)
			return
		}
		pinyin := strings.Split(tmp[1], " ")
		p.AddData(tmp[0], pinyin...)
	}
}
