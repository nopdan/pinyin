package pinyin

func (p *Pinyin) Match(word string) []string {
	chars := []rune(word)
	// 单字只返回一个读音
	if len(chars) == 1 {
		if v := p.MatchChar(chars[0]); v != nil {
			return []string{v[0]}
		}
		return []string{word}
	}

	ret := make([]string, 0, len(chars))
	for i := 0; i < len(chars); {
		len, pinyin := p.Words.match(chars[i:])
		if len == 0 {
			if py, ok := p.Chars[chars[i]]; ok {
				ret = append(ret, py[0])
			}
			i++
			continue
		}
		ret = append(ret, pinyin...)
		i += len
	}
	return ret
}

func (p *Pinyin) MatchChar(c rune) []string {
	if v, ok := p.Chars[c]; ok {
		return v
	}
	return nil
}
