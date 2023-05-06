package pinyin

// trie 树 map 实现
type trie struct {
	ch     map[rune]*trie // children
	pinyin []string
	freq   int
}

func newTrie() *trie {
	t := new(trie)
	t.ch = make(map[rune]*trie, 1000)
	return t
}

// 插入一个词
func (t *trie) insert(word string, pinyin []string, freq int) {
	for _, v := range word {
		if t.ch == nil {
			t.ch = make(map[rune]*trie)
			t.ch[v] = new(trie)
		} else if t.ch[v] == nil {
			t.ch[v] = new(trie)
		}
		t = t.ch[v]
	}
	// 插入同一个词会覆盖前面的
	if freq >= t.freq {
		t.freq = freq
		t.pinyin = pinyin
	}
}

// 按词最长匹配
// 返回 匹配成功的（字数，pinyin）
func (t *trie) match(text []rune) (int, []string) {
	var length int
	var pinyin []string
	for p := 0; p < len(text); {
		t = t.ch[text[p]]
		p++
		if t == nil {
			break
		}
		if len(t.pinyin) != 0 {
			length = p
			pinyin = t.pinyin
		}
	}
	return length, pinyin
}
