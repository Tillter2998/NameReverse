package helpers

import "errors"

func ReverseString(input string) (string, error) {
	if input == "" {
		return "", errors.New("Input cannot be empty")
	}

	reversedname := ""

	for i := len(input) - 1; i >= 0; i-- {
		reversedname += string(input[i])
	}

	return reversedname, nil
}
