// test the placeword function
package sillycode

import "testing"

var placewordtests = []struct {
	in  string
	out string
}{
	{"Denver", "Dnvr"},
	{"Ashburn", "Asbn"},
	{"Madison", "Mdsn"},
	{"Erin", "ErnX"},
	{"Oconomowoc", "Ocmc"},
	{"Reeseville", "Rsvl"},
}

func TestPlaceword(t *testing.T) {
	for i, tt := range placewordtests {
		s := placeword(tt.in)
		if s != tt.out {
			t.Errorf("%d. placeword(%q) => %q, want %q", i, tt.in, s, tt.out)
		}
	}
}
