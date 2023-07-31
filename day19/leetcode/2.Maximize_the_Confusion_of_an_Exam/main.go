package main

import "fmt"

func main() {
	fmt.Println(maxConsecutiveAnswers("FFFTTFTTFT", 3))
}

func maxConsecutiveAnswers(answerKey string, k int) int {
	cF, cT := 0, 0
	iF, iT := 0, 0
	kF, kT := k, k
	maxF, maxT := 0, 0

	for i := 0; i < len(answerKey); i++ {
		if answerKey[i] == 'T' {
			for kF == 0 && iF < i {
				if answerKey[iF] == 'T' {
					iF++
					cF--
					kF++
				} else {
					iF++
					cF--
				}
			}
			kF--
			cF++

			cT++
			if cF > maxF {
				maxF = cF
			}
			if cT > maxT {
				maxT = cT
			}
		} else if answerKey[i] == 'F' {
			for kT == 0 && iT < i {
				if answerKey[iT] == 'F' {
					iT++
					cT--
					kT++
				} else {
					iT++
					cT--
				}
			}

			kT--
			cT++

			cF++
			if cF > maxF {
				maxF = cF
			}
			if cT > maxT {
				maxT = cT
			}
		}
	}

	if cF > maxF {
		maxF = cF
	}
	if cT > maxT {
		maxT = cT
	}

	fmt.Println(cF, cT)
	if maxF > maxT {
		return maxF
	} else {
		return maxT
	}
}
