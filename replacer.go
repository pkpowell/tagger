package tagger

import "strings"

var chars = []string{
	".", " ",
	",", " ",
	"-", " ",
	"+", " ",
	"?", "",
	"!", "",
	":", " ",
	";", " ",
	"...", " ",
	// "\n", " ",
	// "\r", " ",
	// "\t", " ",
	// "'s", "",
	// "[", " ",
	// "]", " ",
	// "(", " ",
	// ")", " ",
	"'", " ",
	"‘", " ",
	"’", " ",
	// "\"", "",
}

var Replacer = strings.NewReplacer(chars...)

// func init() {
// 	strings.Fields()
// }
