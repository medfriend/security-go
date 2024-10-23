package util

import "strconv"

func StringToUint(s string) (uint, error) {

	val, err := strconv.ParseUint(s, 10, 64) // Base 10 y 64 bits
	if err != nil {
		return 0, err
	}

	return uint(val), nil
}
