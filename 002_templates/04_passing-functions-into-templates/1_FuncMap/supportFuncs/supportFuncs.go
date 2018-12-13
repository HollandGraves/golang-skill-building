package supportFuncs

import "strings"

/*


==========================================================================
LIST OF FUNCTIONS I'VE MADE
==========================================================================

func firstLetters(s string) string


*/

// FirstLetters :
// 1. split the string into single words within a slice of string
// 2. declare var to hold byte type of each first letter of each word
// 3. loop through slice of string, append each first letter byte
//	  to the var slice of byte just declared
// 4. turn slice of bytes back into string
// 5. return string
func FirstLetters(s string, sep string) string {
	// 1.
	fl := strings.Split(s, sep)

	// 2.
	var firstLetterSlice []byte

	// 3.
	for _, firLet := range fl {
		firstLetterSlice = append(firstLetterSlice, firLet[0])
	}

	// 4,
	fls := string(firstLetterSlice)

	// 5.
	return fls
}
