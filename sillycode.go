package sillycode

import (
	"bytes"
	"errors"
	"math"
	"regexp"
	"strings"
)

var vowels = regexp.MustCompile("(a|e|i|o|u)")

// return only even letters from string
func evenletters(word string) string {
	var buffer bytes.Buffer

	for i, e := range word {
		if math.Mod(float64(i), 2) == 0 {
			buffer.WriteString(string(e))
		}
	}
	return buffer.String()
}

// return a placeword (4 letter code from the place name)
// this is used when the place name only has one word in it
func placeword(word string) string {
	first := string(word[0])
	rest := vowels.ReplaceAllString(word[1:], "")

	word = first + rest
	wordSize := len(word)

	if wordSize < 4 {
		word = padfourx(word)
	} else if wordSize == 5 {
		first3 := word[:3]
		last := string(word[len(word)-1])
		//option := word[:4] // python had this. was part of
		// multiple return to give users options in case of some
		// kind of conflict

		word = first3 + last

	} else if wordSize > 5 {
		word = string(word[0]) + evenletters(word[1:])
	}

	return word
}

// pad place word with X's until it is 4 letters long
func padfourx(word string) string {
	wordSize := len(word)
	switch wordSize {
	case 3:
		word += "X"
	case 2:
		word += "XX"
	case 1:
		word += "XXX"
	case 0:
		word = "XXXX"
	}
	return word
}

// Turn a word from a place name into a pairword (two letter 
// abbreviation for use with another pairword)
func pairword(word string) string {
	first := string(word[0])
	rest := vowels.ReplaceAllString(word[1:], "")

	name := first + rest
	nameSize := len(name)

	if nameSize == 1 {
		name = word + "X"
	} else if nameSize == 2 {
		// two chars so we are good
	} else {
		name = first + string(rest[0])
	}

	return name
}

// Get a get the sillycode for a place name
func Placename(place string) (string, error) {
	placeSize := len(place)
	if placeSize == 0 {
		return "", errors.New("New place provided")
	}

	place = strings.ToUpper(place)

	// If the name is shorter than 4 chars, pad with X and don't strip vowels  
	if placeSize < 4 {
		place = padfourx(place)
		return place, nil
	}

	// If name is exactly 4 chars, just return it
	if placeSize == 4 {
		return place, nil
	}

	words := strings.Fields(place)
	wordsSize := len(words)

	if wordsSize == 1 {
		place = placeword(place)
	} else if wordsSize == 2 {
		name1 := pairword(words[0])
		name2 := pairword(words[1])

		place = name1 + name2
	} else {
		return "", errors.New("Don't know how to handle place names of 3+ words")
	}

	return place, nil
}
