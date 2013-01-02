package sillycode

import (
	"errors"
	"regexp"
	"strings"
  "math"
  "bytes"
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
	place = strings.ToUpper(place)

	// If the name is shorter than 4 chars, pad with X and don't strip vowels  
	if placeSize < 4 {
		switch placeSize {
		case 3:
			place += "X"
		case 2:
			place += "XX"
		case 1:
			place += "XXX"
		case 0:
			return "", errors.New("No place provided")
		}
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
	} else  {
		return "", errors.New("Don't know how to handle place names of 3+ words")
	}

	return place, nil
}
