package pinyin

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
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
func (p *Pinyin) AddOne(word string, pinyin ...string) {
	p.addOne(word, pinyin, 1)
}

func (p *Pinyin) addOne(word string, pinyin []string, freq int) {
	chars := []rune(word)
	// 词组
	if len(chars) != 1 {
		p.Words.insert(word, pinyin, freq)
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
func (p *Pinyin) AddFile(path string) {
	rd, err := ku.Read(path)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	p.AddReader(rd)
}

func (p *Pinyin) AddData(data []byte) {
	brd := bytes.NewReader(data)
	rd := ku.NewReader(brd)
	p.AddReader(rd)
}

// 输入 utf-8 编码格式的字节流
func (p *Pinyin) AddReader(rd io.Reader) {
	scan := bufio.NewScanner(rd)
	for scan.Scan() {
		line := scan.Text()
		tmp := strings.Split(line, "\t")
		if len(tmp) < 2 || line == "" {
			fmt.Printf("数据有误: %v\n", line)
			return
		}
		freq := 1
		if len(tmp) >= 3 {
			num, err := strconv.Atoi(tmp[2])
			if err == nil {
				freq = num
			}
		}
		pinyin := strings.Split(tmp[1], " ")
		p.addOne(tmp[0], pinyin, freq)
	}
}
