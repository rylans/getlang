package getlang

import (
	"math"
	"sort"
	"unicode"
)

const undeterminedRate int = 23
const undetermined string = "und"
const rescale = 0.5

type Info struct {
	lang        string
	probability float64
}

// LanguageCode returns the ISO 639-1 code for the detected language
func (info Info) LanguageCode() string {
	return info.lang
}

// Confidence returns a measure of reliability for the language classification
func (info Info) Confidence() float64 {
	return info.probability
}

// LanguageName returns the English name of the detected language
func (info Info) LanguageName() string {
	switch info.lang {
	case "en":
		return "English"
	case "es":
		return "Spanish"
	case "pt":
		return "Portuguese"
	case "hu":
		return "Hungarian"
	case "de":
		return "German"
	case "it":
		return "Italian"
	case "pl":
		return "Polish"
	case undetermined:
		return "Undetermined language"
	}
	panic("Missing language code " + info.lang)
}

func FromString(text string) Info {
	langs := map[string][]string{
		"en": en,
		"es": es,
		"pt": pt,
		"hu": hu,
		"de": de,
		"it": it,
		"pl": pl,
	}

	langMatches := make(map[string]int)
	langMatches[undetermined] = 0
	for k, _ := range langs {
		langMatches[k] = 1
		// Plus one smoothing
	}

	trigs := sortedTrigs(text)
	for k, v := range langs {
		matchWith(k, trigs, v, langMatches)
	}

	smx := softmax(langMatches)
	maxk := maxkey(langMatches)
	return Info{maxk, smx[maxk]}
}

func softmax(mapping map[string]int) map[string]float64 {
	softmaxmap := make(map[string]float64)
	denom := 0.0
	for _, v := range mapping {
		denom += math.Exp(float64(v) * rescale)
	}
	for k, _ := range mapping {
		softmaxmap[k] = math.Exp(rescale*float64(mapping[k])) / denom
	}
	return softmaxmap
}

func maxkey(mapping map[string]int) string {
	max := 0
	key := ""
	for k, v := range mapping {
		if v > max {
			max = v
			key = k
		}
	}
	return key
}

func matchWith(langname string, trigs []trigram, langprofile []string, matches map[string]int) {
	undeterminedCount := 0
	prof := make(map[string]int)
	for _, x := range langprofile {
		prof[x] = 1
	}

	for _, trig := range trigs {
		if _, exists := prof[trig.trigram]; exists {
			matches[langname]++
		} else {
			undeterminedCount++
			if (undeterminedCount % undeterminedRate) == 0 {
				matches[undetermined]++
			}
		}
	}
}

func countedTrigrams(text string) map[string]int {
	var r1, r2, r3 rune
	trigrams := map[string]int{}
	var txt []rune

	for _, r := range text {
		txt = append(txt, unicode.ToLower(toTrigramChar(r)))
	}
	txt = append(txt, ' ')

	r1 = ' '
	r2 = txt[0]
	for i := 1; i < len(txt); i++ {
		r3 = txt[i]
		if !(r2 == ' ' && (r1 == ' ' || r3 == ' ')) {
			trigram := []rune{}
			trigram = append(trigram, r1)
			trigram = append(trigram, r2)
			trigram = append(trigram, r3)
			if trigrams[string(trigram)] == 0 {
				trigrams[string(trigram)] = 1
			} else {
				trigrams[string(trigram)]++
			}
		}
		r1, r2 = r2, r3
	}

	return trigrams
}

type trigram struct {
	trigram string
	count   int
}

func sortedTrigs(s string) []trigram {
	counterMap := countedTrigrams(s)
	trigrams := make([]trigram, len(counterMap))

	i := 0
	for tg, count := range counterMap {
		trigrams[i] = trigram{tg, count}
		i++
	}
	sort.SliceStable(trigrams, func(i, j int) bool {
		if trigrams[i].count == trigrams[j].count {
			return trigrams[i].trigram < trigrams[j].trigram
		}
		return trigrams[i].count > trigrams[j].count
	})
	return trigrams
}

func toTrigramChar(ch rune) rune {
	if unicode.IsPunct(ch) || unicode.IsSpace(ch) {
		return ' '
	}
	return ch
}
