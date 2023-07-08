package dog

import (
	"strings"
)

func WhenGrownUp(s string) string {
	return "When the puppy grows up it says: " + strings.ToUpper(s)
}

func Whenchild(s string) string {
	return "When the puppy is child it says: " + strings.ToLower(s)
}
