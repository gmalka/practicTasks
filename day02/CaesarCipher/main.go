package main

import (
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println(caesarCipher("abcdz", 100))
}

func caesarCipher(s string, k int32) string {
    var newstr strings.Builder

    for i := 0; i < len(s); i++ {
        if s[i] >= 'a' && s[i] <= 'z' {
            newByte := s[i] + byte(k)
            if newByte > 'z' {
                newByte = ((newByte - 'a') % (26)) + 'a'
            }
            newstr.WriteByte(newByte)
        } else if s[i] >= 'A' && s[i] <= 'Z' {
            newByte := s[i] + byte(k)
            if newByte > 'Z' {
                newByte = ((newByte - 'A') % (26)) + 'A'
            }
            newstr.WriteByte(newByte)
        } else {
            newstr.WriteByte(s[i])
        }
    }

    return newstr.String()
}