package stringService

type stringService struct {
}

func NewStringService() stringService {
	return stringService{}
}

func (s stringService) ServeString(str string) string {
	if len(str) == 0 {
		return ""
	}

	m := make(map[byte]struct{}, len(str)/10)
	maxLen, maxL, l, r := 0, 0, 0, 1
	curLen := 0

	for i := 0; i < len(str); i++ {
		if _, ok := m[str[i]]; !ok {
			m[str[i]] = struct{}{}
			curLen++
		} else {
			if curLen > maxLen {
				maxLen = curLen
				maxL = l
				r = i
			}
			for str[l] != str[i] {
				delete(m, str[l])
				curLen--
				l++
			}
			l++
		}
	}
	if curLen > maxLen {
		maxL = l
		r = len(str)
	}

	return str[maxL:r]
}