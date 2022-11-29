package main

import (
	"fmt"
)

func GetEqualPadded(s string, length int) string {
	padding := length - len(s)
	if padding >= 0 {
		return fmt.Sprintf("'%*s'", padding, s)
	}

	return fmt.Sprintf("…%s", s[len(s)-(length-1):])
}
