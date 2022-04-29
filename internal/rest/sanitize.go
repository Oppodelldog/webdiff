package rest

import "strings"

var badCharacters = []string{
	"..",
	".",
	"<!--",
	"-->",
	"<",
	">",
	"'",
	"\"",
	"&",
	"$",
	"#",
	"{", "}", "[", "]", "=",
	";", "?", "%20", "%22",
	"%3c",   // <
	"%253c", // <
	"%3e",   // >
	"",      // > -- fill in with % 0 e - without spaces in between
	"%28",   // (
	"%29",   // )
	"%2528", // (
	"%26",   // &
	"%24",   // $
	"%3f",   // ?
	"%3b",   // ;
	"%3d",   // =
	"./",
	"/",
}

func removeBadCharacters(input string, dictionary []string) string {
	for _, badChar := range dictionary {
		input = strings.ReplaceAll(input, badChar, "")
	}

	return input
}

func sanitizeFilename(name string) string {
	var badDictionary = badCharacters

	if name == "" {
		return name
	}

	name = strings.TrimSpace(name)
	name = strings.Replace(name, " ", "", -1)
	name = removeBadCharacters(name, badDictionary)
	name = strings.Replace(name, "\\", "", -1)

	return name
}
