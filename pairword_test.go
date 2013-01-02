// test the pairword function
package sillycode

import "testing"

var pairwordtests = []struct {
	in  string
	out string
}{
	{"A", "AX"},
	{"AB", "AB"},
	{"NORTH", "NO"},
}

func TestPairword(t *testing.T) {
	for i, tt := range pairwordtests {
		s := pairword(tt.in)
		if s != tt.out {
			t.Errorf("%d. pairword(%q) => %q, want %q", i, tt.in, s, tt.out)
		}
	}
}
