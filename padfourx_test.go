// test the padfourx function
package sillycode

import "testing"

var padfourxtests = []struct {
	in  string
	out string
}{
	{"", "XXXX"},
	{"A", "AXXX"},
	{"AB", "ABXX"},
	{"ABC", "ABCX"},
	{"ABCD", "ABCD"},
	{"ABCDE", "ABCDE"},
}

func TestPadFourX(t *testing.T) {
	for i, tt := range padfourxtests {
		s := padfourx(tt.in)
		if s != tt.out {
			t.Errorf("%d. padfourx(%q) => %q, want %q", i, tt.in, s, tt.out)
		}
	}
}
