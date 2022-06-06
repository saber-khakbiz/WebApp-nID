package code

import (
	"strconv"
)

func ValidateIranianCode(code string) bool {
	ToInt := func(str string) int {
		if icode, err := strconv.Atoi(str); err == nil {
			return icode
		}
		return -1
	}
	if icode := ToInt(code); icode >= 0 {
		if len(code) < 8 || icode == 0 {
			return false
		}

		code = ("0000" + code)[len(code)+4-10:]
		if jcode := ToInt(code[3:9]); jcode >= 0 {
			if jcode == 0 {
				return false
			}
		} else {
			return false
		}

		if ccode := ToInt(code[9:10]); ccode >= 0 {
			var scode = 0
			for i := 0; i < 9; i++ {
				temp := ToInt(code[i : i+1])
				if temp < 0 {
					return false
				}
				scode += temp * (10 - i)
			}
			scode = scode % 11
			return (scode < 2 && ccode == scode) || (scode >= 2 && ccode == (11-scode))
		} else {
			return false
		}
	} else {
		return false
	}
}
