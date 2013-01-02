// test the even letters function
package sillycode

import "testing"

var eventests = []struct {
	in  string
	out string
}{
	{"a", "a"},
	{"ab", "a"},
	{"abc", "ac"},
	{"abcd", "ac"},
	{"abcde", "ace"},
}

func TestEvenLetters(t *testing.T) {
	for i, tt := range eventests {
		s := evenletters(tt.in)
		if s != tt.out {
			t.Errorf("%d. evenletters(%q) => %q, want %q", i, tt.in, s, tt.out)
		}
	}
}
