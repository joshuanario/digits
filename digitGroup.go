package digits

import "strings"

func DigitGroup(o int, v string, g rune, d Decimals) string {
	gChar := string([]rune{g})
	i := strings.IndexRune(v, '.')
	if i > -1 {
		end := len(v)
		trunc := int(d) + i + 1
		if end > trunc {
			end = trunc
		}
		return triplefy(-1, v[:i], gChar) + v[i:end]
	}
	return triplefy(o, v, gChar)
}
func triplefy(o int, v string, g string) string {
	if len(v) < 3 {
		return v
	}
	i := strings.IndexRune(v, '.')
	negO := 0 > o
	negI := 0 > i
	gChar := ","
	if g != "," {
		gChar = "."
	}
	if negO && negI {
		return triplefy(-1, v[:len(v)-3], gChar) + gChar + v[len(v)-3:]
	}
	if negO != negI {
		return gChar
	}
	if !negI {
		dChar := "."
		if g != "," {
			dChar = ","
		}
		return triplefy(-1, v[:i], gChar) + dChar + v[i+1:]
	}
	newO := o - 3
	modDiff := (o - len(v)) % 3
	end := len(v) - 3 - modDiff
	return triplefy(newO, v[:end], gChar) + gChar
}
