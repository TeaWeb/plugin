package apps

// 分析字符串中的参数
func ParseArgs(s string) (args []string) {
	quotesBegin := false
	quotesEscaped := false
	var lastQuote rune
	lastArg := ""
	for index, character := range s {
		if character == '"' || character == '\'' {
			if quotesEscaped {
				lastArg += string(character)
				quotesEscaped = false
				continue
			}
			if quotesBegin {
				if lastQuote == character { // 引号结束
					quotesBegin = false
				} else {
					// 视为参数的一部分
					lastArg += string(character)
				}
			} else {
				quotesBegin = true
				lastQuote = character
			}
		} else if character == '\\' {
			if len(s) > index+1 && (s[index+1:index+2][0] == '"' || s[index+1:index+2][0] == '\'') {
				quotesEscaped = true
			} else {
				lastArg += string(character)
			}
		} else if character == ' ' || character == '\t' || character == '\n' || character == '\r' {
			if quotesBegin { // 如果在引号中，则视为参数的一部分
				lastArg += string(character)
			} else { // 如果不在引号中，则认为参数已结束
				if len(lastArg) > 0 {
					args = append(args, lastArg)
					lastArg = ""
				}
			}
		} else {
			lastArg += string(character)
		}
	}

	if len(lastArg) > 0 {
		args = append(args, lastArg)
	}

	return args
}
