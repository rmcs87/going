package hackerrank

import (
	"sort"
	"strings"
)

func RemainderSorting(strArr []string) []string {
	sort.Slice(strArr, func(i, j int) bool {
		if len(strArr[i])%3 == len(strArr[j])%3 {
			r := strings.Compare(strArr[i], strArr[j])
			return r == -1
		}

		return len(strArr[i])%3 < len(strArr[j])%3
	})

	return strArr
}
