package palindrome

import "fmt"

func Palindrome(s string) (err error) {
	if len(s) == 0 {
		return fmt.Errorf("Empty string")
	}

	str := []rune(s)
	for i := range str {
		if str[i] != str[len(str)-i-1] {
			break
		}
	}
	return nil
}
