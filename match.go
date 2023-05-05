package pinyin

func (p *Pinyin) Match(word string) []string {
	chars := []rune(word)
	if len(chars) == 1 {
		return p.Chars[chars[0]]
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
