package getlang

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestEmptyStringFromReader(t *testing.T) {
	info := FromReader(strings.NewReader(""))
	assert.Equal(t, "und", info.LanguageCode())
}

func TestEnglishPhraseFromReader(t *testing.T) {
	info := FromReader(strings.NewReader("this is the language"))
	assert.Equal(t, "en", info.LanguageCode())
	assert.Equal(t, true, info.Confidence() > 0.75)
}

func TestEnglishPhraseUSDI(t *testing.T) {
	text := "We hold these truths to be self-evident, that all men are created equal"
	ensureClassifiedWithConfidence(
		t,
		text,
		"en",
		0.95)

	ensureClassifiedTextNamed(
		t,
		text,
		"English",
		"English")
}

func TestGermanPhraseUSDI(t *testing.T) {
	text := "Wir halten diese Wahrheiten für ausgemacht, daß alle Menschen gleich erschaffen worden"
	ensureClassifiedWithConfidence(
		t,
		text,
		"de",
		0.95)

	ensureClassifiedTextNamed(
		t,
		text,
		"German",
		"Deutsch")
}

func TestGermanMixedEnglish(t *testing.T) {
	ensureClassifiedWithConfidence(
		t,
		"Wenn wir jemand grüßen wollen, sagen wir 'How are you doing?'",
		"de",
		0.85)
}

func TestEnglishMixedGerman(t *testing.T) {
	ensureClassifiedWithConfidence(
		t,
		"If you wanted to greet someone in this language, you'd say 'wie geht es'",
		"en",
		0.35)
}

func TestEnglishMixedUkrainian(t *testing.T) {
	ensureClassifiedWithConfidence(
		t,
		"the best thing to say is своїй гідності in my opinon.",
		"en",
		0.55)
}

func TestSpanishPhraseUSDI(t *testing.T) {
	ensureClassifiedWithConfidence(
		t,
		"Sostenemos como evidentes estas verdades: que los hombres son creados iguales",
		"es",
		0.75)
}

func TestPortuguesePhraseUSDI(t *testing.T) {
	ensureClassifiedWithConfidence(
		t,
		"Consideramos estas verdades como autoevidentes, que todos os homens são criados iguais",
		"pt",
		0.95)
}

func TestPolishPhraseUDHR(t *testing.T) {
	ensureClassifiedWithConfidence(
		t,
		"Wszyscy ludzie rodzą się wolni i równi w swojej godności i prawach",
		"pl",
		0.95)
}

func TestHungarianPhraseUDHR(t *testing.T) {
	ensureClassifiedWithConfidence(
		t,
		"Minden emberi lény szabadon születik és egyenlő méltósága és joga van",
		"hu",
		0.95)
}

func TestItalianPhraseUDHR(t *testing.T) {
	ensureClassifiedWithConfidence(
		t,
		"Tutti gli esseri umani nascono liberi ed eguali in dignità e diritti",
		"it",
		0.95)
}

func TestRussianPhraseUDHR(t *testing.T) {
	ensureClassifiedWithConfidence(
		t,
		"Все люди рождаются свободными и равными в своем достоинстве и правах",
		"ru",
		0.55)
}

func TestUkrainianPhraseUDHR(t *testing.T) {
	ensureClassifiedWithConfidence(
		t,
		"Всі люди народжуються вільними і рівними у своїй гідності та правах",
		"uk",
		0.80)
}

func TestFrenchPhraseUDHR(t *testing.T) {
	ensureClassifiedWithConfidence(
		t,
		"Tous les êtres humains naissent libres et égaux",
		"fr",
		0.95)
}

func TestKoreanPhrase(t *testing.T) {
	ensureClassifiedWithConfidence(
		t,
		"원래 AB형 사람이 똑똑해",
		"ko",
		0.95)
}

func TestJapanesePhrase(t *testing.T) {
	text := "何を食べますか"
	ensureClassifiedWithConfidence(
		t,
		text,
		"ja",
		0.95)

	ensureClassifiedTextNamed(
		t,
		text,
		"Japanese",
		"日本語")
}

func TestChinesePhrase(t *testing.T) {
	text := "球的采编网络,记者遍布"
	ensureClassifiedWithConfidence(
		t,
		text,
		"zh",
		0.95)

	ensureClassifiedTextNamed(
		t,
		text,
		"Chinese",
		"中文")
}

func TestArabicPhrase(t *testing.T) {
	text := "اهتمامًا بذلك المشروع. المجموعة الوحيدة التي "
	lang := "العربية"
	ensureClassifiedWithConfidence(
		t,
		text,
		"ar",
		0.55)

	ensureClassifiedTextNamed(
		t,
		text,
		"Arabic",
		lang)
}

func TestBanglaPhrase(t *testing.T) {
	text := "এই গবেষণায় রত, তাঁদেরকে বলা হয় ভাষাবিজ্ঞানী।ভাষাবিজ্ঞানীরা নৈর্ব্যক্তিক"
	lang := "বাংলা"

	ensureClassifiedWithConfidence(
		t,
		text,
		"bn",
		0.85)

	ensureClassifiedTextNamed(
		t,
		text,
		"Bangla",
		lang)
}

func TestHindiPhrase(t *testing.T) {
	text := "ब तक लगातार चल रहा है। इसका प्रसारण प्रत्येक शनिवार और रविवार को रात 10 बजे होता है। इसका पुनः प्रसारण सोनी पल चैनल पर रात 9 बजे होता"
	lang := "हिन्दी"

	ensureClassifiedWithConfidence(
		t,
		text,
		"hi",
		0.85)

	ensureClassifiedTextNamed(
		t,
		text,
		"Hindi",
		lang)
}

func TestNonsense(t *testing.T) {
	text := "wep lvna eeii vl jkk azc nmn iuah ppl zccl c%l aa1z"
	ensureClassifiedWithConfidence(
		t,
		text,
		"und",
		0.75)

	ensureClassifiedTextNamed(
		t,
		text,
		"Unknown language",
		"")
}

func ensureClassifiedWithConfidence(t *testing.T, text string, expectedLang string, minConfidence float64) {
	info := FromString(text)

	assert.Equal(t, expectedLang, info.LanguageCode(), "Misclassified text: "+text)
	assert.Equal(t, true, info.Confidence() > minConfidence)
}

func ensureClassifiedTextNamed(t *testing.T, text string, expectedEnglishName string, expectedSelfName string) {
	info := FromString(text)

	assert.Equal(t, expectedEnglishName, info.LanguageName(), "Wrong language name: "+text)
	assert.Equal(t, expectedSelfName, info.SelfName(), "Wrong self lang name: "+text)
}
