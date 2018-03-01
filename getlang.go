package getlang

import (
	"sort"
	"unicode"
)

type Info struct {
	lang        string
	probability float64
}

func FromString(text string) Info {
	langs := map[string][]string{
		"en": en,
		"es": es,
	}

	langMatches := make(map[string]int)
	for k, _ := range langs {
		langMatches[k] = 1
		// Plus one smoothing
	}

	trigs := sortedTrigs(text)
	for k, v := range langs {
		matchWith(k, trigs, v, langMatches)
	}

	return Info{maxkey(langMatches), 1.0}
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
	prof := make(map[string]int)
	for _, x := range langprofile {
		prof[x] = 1
	}

	for _, trig := range trigs {
		if _, exists := prof[trig.trigram]; exists {
			matches[langname]++
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
