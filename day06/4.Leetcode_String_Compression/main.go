package main

import "fmt"

func main() {
	b := make([]byte, 999)
	for i := 0; i < 999; i++ {
		b[i] = 'a'
	}

	compress(b)

	fmt.Println(string(b))
}

func compress(chars []byte) int {
	l, r := 0, 1

	count := 1
	for r = 1; r < len(chars); r++ {
		if chars[r] == chars[l] {
			count++
		} else if count > 1 {
			buf := []byte{}
			for count != 0 {
				buf = append(buf, byte(count%10)+'0')
				count /= 10
			}
			for j := len(buf) - 1; j >= 0; j-- {
				l++
				chars[l] = buf[j]
			}
			l++
			chars[l] = chars[r]
			count = 1
		} else {
			l++
			chars[l] = chars[r]
		}
	}
	if count > 1 {
		buf := []byte{}
		for count != 0 {
			buf = append(buf, byte(count%10)+'0')
			count /= 10
		}
		for j := len(buf) - 1; j >= 0; j-- {
			l++
			chars[l] = buf[j]
		}
		l++
	} else {
		l++
	}

	chars = chars[:l]

	return len(chars)
}
