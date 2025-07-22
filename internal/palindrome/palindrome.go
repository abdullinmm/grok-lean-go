package palindrome

import "fmt"

// Palindrome checks if a string is a palindrome.
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
