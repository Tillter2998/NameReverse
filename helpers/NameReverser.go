package helpers

import "errors"

func ReverseName(name string) (string, error) {
	if name == "" {
		return "", errors.New("Name cannot be empty")
	}

	reversedname := ""

	for i := len(name) - 1; i >= 0; i-- {
		reversedname += string(name[i])
	}

	return reversedname, nil
}
