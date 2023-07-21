package main

func main() {
	digitSum("1234", 2)
}

func digitSum(s string, k int) string {
	b := []byte(s)
	for len(b) > k {
		i := k
		t := 0
		sum := 0
		for {
			var l []byte
			if len(b) > i {
				l = b[i-k : i]
			} else {
				if len(b) <= i-k {
					break
				}
				l = b[i-k:]
			}

			for _, v := range l {
				sum += int((v - '0'))
			}
			in := make([]byte, 0, 10)
			if sum == 0 {
				in = append(in, '0')
			}
			for sum != 0 {
				p := byte(sum % 10)
				in = append(in, p+'0')
				sum /= 10
			}
			for k := len(in) - 1; k >= 0; k-- {
				b[t] = in[k]
				t++
			}

			i += k
		}
		b = b[:t]
	}

	return string(b)
}