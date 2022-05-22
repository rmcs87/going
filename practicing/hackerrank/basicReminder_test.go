package hackerrank

import "testing"

func TestRemainderSorting(t *testing.T) {
	type test struct {
		name     string
		expected []string
		data     []string
	}
	tests := []test{
		{"basic", []string{"Oregon", "Wisconsin", "Utah", "Colorado"}, []string{"Colorado", "Utah", "Wisconsin", "Oregon"}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := RemainderSorting(test.data)
			if !Equal(got, test.expected) {
				t.Errorf("got %v but want %v", got, test.expected)
			}
		})
	}

}

func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
