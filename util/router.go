package util

import (
	"strconv"
)

func UintID(requestParam string) (uint, error) {
	id, err := strconv.Atoi(requestParam)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}
