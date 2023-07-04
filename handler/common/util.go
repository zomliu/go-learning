package common

import "unicode"

// IsDigit 判断字符串是否是数字
func IsDigit(s string) bool {
	for _, x := range s {
		if !unicode.IsDigit(x) {
			return false
		}
	}
	return true
}
