package leetcode

func uniqueMorseRepresentations(words []string) int {
	output := 0
	mapOfMorse := make(map[string]bool)

	for _, word := range words {
		morse := toMorse(word)
		mapOfMorse[morse] = true
	}

	output = len(mapOfMorse)

	return output
}

func toMorse(word string) string {
	morseTable := map[string]string{
		"a": ".-",
		"b": "-...",
		"c": "-.-.",
		"d": "-..",
		"e": ".",
		"f": "..-.",
		"g": "--.",
		"h": "....",
		"i": "..",
		"j": ".---",
		"k": "-.-",
		"l": ".-..",
		"m": "--",
		"n": "-.",
		"o": "---",
		"p": ".--.",
		"q": "--.-",
		"r": ".-.",
		"s": "...",
		"t": "-",
		"u": "..-",
		"v": "...-",
		"w": ".--",
		"x": "-..-",
		"y": "-.--",
		"z": "--..",
	}

	output := ""
	for _, char := range word {
		output += morseTable[string(char)]
	}

	return output
}
