package leetcode

import "fmt"

func hexadecimal(num int) string {
	return fmt.Sprintf("%x", uint32(num))
}
