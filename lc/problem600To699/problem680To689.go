package problem600To699

func validPalindrome(s string) bool {
	i, j := 0, len(s)-1
	// 兼容双字节的遍历方式, 但题目是英文字母, 所以两种遍历方式都可以
	// arr := []rune(s)
	for i < j {
		if s[i] == s[j] {
			i++
			j--
		} else {
			break
		}
	}
	if i >= j {
		return true
	}
	if palindrome(i+1, j, s) {
		return true
	}
	return palindrome(i, j-1, s)
}

func palindrome(i, j int, arr string) bool {
	for i < j {
		if arr[i] == arr[j] {
			i++
			j--
		} else {
			break
		}
	}
	if i >= j {
		return true
	}
	return false
}
