package problem500To599

func A541reverseStr(s string, k int) string {
	res := ""
	tmp := ""
	b := true
	for i, ch := range s {
		if b {
			tmp = string(ch) + tmp
		} else {
			res += string(ch)
		}
		if i%k == k-1 {
			if b {
				res += tmp
				tmp = ""
			}
			b = !b
		}
	}
	if b {
		res += tmp
	}
	return res
}
