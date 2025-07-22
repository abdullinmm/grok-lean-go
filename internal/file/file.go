package file

import (
	"errors"
	"fmt"
	"os"
	"unicode/utf8"
)

// ReadFile reads the contents of a file and returns them as a string.
func ReadFile(filePath string) (string, error) {

	data, err := os.ReadFile(filePath)
	if err != nil {
		var pathError *os.PathError
		if errors.As(err, &pathError) {
			if os.IsNotExist(err) {
				return "", fmt.Errorf("file not found: %s", filePath)
			} else if os.IsPermission(err) {
				return "", fmt.Errorf("permission denied: %s", filePath)
			}
			return "", fmt.Errorf("IO error with file '%s': %v", filePath, pathError.Err)
		}
		return "", err
	}

	// Checking the validity of UTF-8
	if !utf8.Valid(data) {
		return "", fmt.Errorf("invalid UTF-8 encoding in file '%s'", filePath)
	}

	return string(data), nil
}
