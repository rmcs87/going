package hackerrank

import (
	"sort"
	"strings"
)

func customSorting(strArr []string) []string {

	sort.Slice(strArr, func(i, j int) bool {
		if len(strArr[i])%2 != 0 && len(strArr[j])%2 == 0 {
			return true
		} else if len(strArr[i])%2 == 0 && len(strArr[j])%2 != 0 {
			return false
		} else if len(strArr[i])%2 != 0 && len(strArr[j])%2 != 0 {
			if len(strArr[i]) > len(strArr[j]) {
				return false
			} else if len(strArr[i]) < len(strArr[j]) {
				return true
			} else {
				r := strings.Compare(strArr[i], strArr[j])
				return r == -1
			}
		} else if len(strArr[i])%2 == 0 && len(strArr[j])%2 == 0 {
			if len(strArr[i]) > len(strArr[j]) {
				return true
			} else if len(strArr[i]) < len(strArr[j]) {
				return false
			} else {
				r := strings.Compare(strArr[i], strArr[j])
				return r == -1
			}
		} else {
			r := strings.Compare(strArr[i], strArr[j])
			return r == -1
		}
	})

	return strArr
}
