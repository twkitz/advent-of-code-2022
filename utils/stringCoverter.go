package utils

import (
	"fmt"
	"strings"
)

func SliceToString(slice []int) string {
	return strings.Trim(strings.Replace(fmt.Sprint(slice), " ", ", ", -1), "[]")
}
