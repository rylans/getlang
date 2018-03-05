// Package getlang provides fast language detection for various languages
//
// getlang compares input text to a characteristic profile of each supported language and
// returns the language that best matches the input text
package getlang

import (
	"golang.org/x/text/language"
	"golang.org/x/text/language/display"
	"io"
	"io/ioutil"
	"math"
	"sort"
	"unicode"
)

const undeterminedRate int = 31
const undetermined string = "und"
const rescale = 0.5
const scriptCountFactor int = 3

type inUnicodeRange func(r rune) bool

// Info is the language detection result
type Info struct {
	lang        string
	probability float64
	langTag     language.Tag
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
	return display.English.Tags().Name(info.langTag)
}

// SelfName returns the name of the language in the language itself
func (info Info) SelfName() string {
	return display.Self.Name(info.langTag)
}

// FromReader detects the language from an io.Reader
//
// This function will read all bytes until an EOF is reached
func FromReader(reader io.Reader) Info {
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		panic("Error reading from reader")
	}

	return FromString(string(bytes))
}

// FromString detects the language from the given string
func FromString(text string) Info {
	langs := map[string][]string{
		"en": en,
		"es": es,
		"pt": pt,
		"hu": hu,
		"de": de,
		"it": it,
		"pl": pl,
		"ru": ru,
		"uk": uk,
		"fr": fr,
	}

	scripts := map[string]inUnicodeRange{
		"ko": isKo,
		"zh": isZh,
		"ja": isJa,
	}

	langMatches := make(map[string]int)
	langMatches[undetermined] = 2
	for k := range langs {
		langMatches[k] = 0
	}

	for k := range scripts {
		langMatches[k] = 0
	}

	trigs := sortedTrigs(text)
	for k, v := range langs {
		matchWith(k, trigs, v, langMatches)
	}

	for k, v := range scripts {
		matchScript(k, text, v, langMatches)
	}

	smx := softmax(langMatches)
	maxk := maxkey(langMatches)
	return Info{maxk, smx[maxk], language.MustParse(maxk)}
}

func softmax(mapping map[string]int) map[string]float64 {
	softmaxmap := make(map[string]float64)
	denom := 0.0
	for _, v := range mapping {
		denom += math.Exp(float64(v) * rescale)
	}
	for k := range mapping {
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

func matchScript(langname string, text string, predicate inUnicodeRange, matches map[string]int) {
	for _, rune := range text {
		if predicate(rune) {
			matches[langname] += scriptCountFactor
		}
	}
}

func matchWith(langname string, trigs []trigram, langprofile []string, matches map[string]int) {
	undeterminedCount := 0
	prof := make(map[string]int)
	for _, x := range langprofile {
		prof[x] = 1
	}

	for _, trig := range trigs {
		if _, exists := prof[trig.trigram]; exists {
			matches[langname] += trig.count
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

var isKo = func(r rune) bool {
	return unicode.Is(unicode.Hangul, r)
}

var isZh = func(r rune) bool {
	return unicode.Is(unicode.Han, r)
}

var isJa = func(r rune) bool {
	return unicode.Is(unicode.Hiragana, r) || unicode.Is(unicode.Katakana, r)
}
