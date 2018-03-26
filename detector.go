// Package golangdetector measures and suggest language according to input
// text based on languages dictionary occurrences.
package golangdetector

import "github.com/lucasmenendez/gotokenizer"

const minLength = 140

// Detect function calculates the language probabilities for all languages of
// languages. To reduce noise, languages only includes English and Spanish
// wordlist.
func Detect(input string) (probs map[string]float64) {
	if len(input) < minLength {
		return
	}

	var w []string
	for _, i := range gotokenizer.Sentences(input) {
		w = append(w, gotokenizer.Words(i)...)
	}

	if len(w) == 0 {
		return
	}

	var total float64
	probs = make(map[string]float64, len(languages))
	for code, words := range languages {
		for _, c := range words {
			for _, i := range w {
				if c == i {
					probs[code]++
					total++
				}
			}
		}
	}

	for code := range probs {
		probs[code] /= total
	}

	return
}

// Suggest function calls Detect function to get the probs of all languages and
// returns the most probably language code.
func Suggest(input string) (code string) {
	var probs map[string]float64 = Detect(input)

	var max float64
	for c, p := range probs {
		if p > max {
			max, code = p, c
		}
	}

	return
}
