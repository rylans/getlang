package getlang

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestEmptyStringFromReader(t *testing.T) {
	info, _ := FromReader(strings.NewReader(""))
	assert.Equal(t, "und", info.LanguageCode())
}

func TestEnglishPhraseFromBigReader(t *testing.T) {
	largeText := ""
	for i := 0; i < 800; i++ {
		largeText += "this is more language as you can see "
	}
	info, _ := FromReader(strings.NewReader(largeText))
	assert.Equal(t, "en", info.LanguageCode())
	assert.Equal(t, true, info.Confidence() > 0.999)
}

func TestEnglishPhraseFromReader(t *testing.T) {
	info, _ := FromReader(strings.NewReader("this is the language"))
	assert.Equal(t, "en", info.LanguageCode())
	assert.Equal(t, true, info.Confidence() > 0.75)
}

func TestEnglishPhraseTag(t *testing.T) {
	info, _ := FromReader(strings.NewReader("this is the language"))
	tag := info.Tag()

	assert.Equal(t, "en", tag.String())
	assert.Equal(t, false, tag.IsRoot())
	assert.Equal(t, true, tag.Parent().IsRoot())
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
		"the best thing to say is своїй гідності in my opinon of this.",
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

func TestPunjabiPhrase(t *testing.T) {
	text := "ਮੇਰਾ ਨਾਮ ਭਰਤ ਹੈ."
	lang := "ਪੰਜਾਬੀ"

	ensureClassifiedWithConfidence(
		t,
		text,
		"pa",
		0.95)

	ensureClassifiedTextNamed(
		t,
		text,
		"Punjabi",
		lang)
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
		0.90)

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
		0.75)

	ensureClassifiedTextNamed(
		t,
		text,
		"Hindi",
		lang)
}

func TestGreekPhrase(t *testing.T) {
	text := "Ολοι οι άνθρωποι γεννιούνται ελεύθεροι και ίσοι στην αξιοπρέπεια και στα δικαιώματα"

	ensureClassifiedWithConfidence(
		t,
		text,
		"el",
		0.95)

	ensureClassifiedTextNamed(
		t,
		text,
		"Greek",
		"Ελληνικά")
}

func TestHebrewPhrase(t *testing.T) {
	text := "כראוי. בִּדקו את כותרת הדף"
	lang := "עברית"

	ensureClassifiedWithConfidence(
		t,
		text,
		"he",
		0.95)

	ensureClassifiedTextNamed(
		t,
		text,
		"Hebrew",
		lang)
}

func TestGujaratiPhrase(t *testing.T) {
	text := "ગુજરાતી"
	lang := "ગુજરાતી"

	ensureClassifiedWithConfidence(
		t,
		text,
		"gu",
		0.95)

	ensureClassifiedTextNamed(
		t,
		text,
		"Gujarati",
		lang)
}

func TestThaiPhrase(t *testing.T) {
	text := "ไทย ไทยไทย"
	lang := "ไทย"

	ensureClassifiedWithConfidence(
		t,
		text,
		"th",
		0.95)

	ensureClassifiedTextNamed(
		t,
		text,
		"Thai",
		lang)
}

func TestArmenianPhrase(t *testing.T) {
	text := "ըստ Գրիգորյան օրացույցի"
	lang := "հայերեն"

	ensureClassifiedWithConfidence(
		t,
		text,
		"hy",
		0.95)

	ensureClassifiedTextNamed(
		t,
		text,
		"Armenian",
		lang)
}

func TestSerbianLatinPhrase(t *testing.T) {
	text := "ljudi ne znaju jer me uglavnom vide"
	lang := "srpskohrvatski"

	ensureClassifiedWithConfidence(
		t,
		text,
		"sr",
		0.85)

	ensureClassifiedTextNamed(
		t,
		text,
		"Serbo-Croatian",
		lang)
}

func TestSerbianCyrillicPhrase(t *testing.T) {
	text := "Код животиња су ове реакције посебно важне при зарастању рана"
	lang := "српски"

	ensureClassifiedWithConfidence(
		t,
		text,
		"sr",
		0.95)

	ensureClassifiedTextNamed(
		t,
		text,
		"Serbian (Cyrillic)",
		lang)
}

func TestVietnamesePhrase(t *testing.T) {
	text := "Truyền thông Việt Nam vào dịp này đăng bài ký tên ông"
	lang := "Tiếng Việt"

	ensureClassifiedWithConfidence(
		t,
		text,
		"vi",
		0.95)

	ensureClassifiedTextNamed(
		t,
		text,
		"Vietnamese",
		lang)
}

func TestTeluguPhrase(t *testing.T) {
	text := "భారతదేశంలోని దక్షిణ"
	lang := "తెలుగు"

	ensureClassifiedWithConfidence(
		t,
		text,
		"te",
		0.95)

	ensureClassifiedTextNamed(
		t,
		text,
		"Telugu",
		lang)
}

func TestTamilPhrase(t *testing.T) {
	text := " நீளமான, கிளைக்காத"
	lang := "தமிழ்"

	ensureClassifiedWithConfidence(
		t,
		text,
		"ta",
		0.95)

	ensureClassifiedTextNamed(
		t,
		text,
		"Tamil",
		lang)
}

func TestTagalogPhrase(t *testing.T) {
	text := "ano ang nangyayari sa iyo at ang mah-ina mo ay hindi mo"
	lang := "Filipino"

	ensureClassifiedWithConfidence(
		t,
		text,
		"tl",
		0.95)

	ensureClassifiedTextNamed(
		t,
		text,
		"Filipino",
		lang)
}

func TestDutchPhrase(t *testing.T) {
	text := "Een ieder heeft, waar hij zich ook bevindt, het recht als persoon erkend te worden voor de wet"
	lang := "Nederlands"

	ensureClassifiedWithConfidence(
		t,
		text,
		"nl",
		0.95)

	ensureClassifiedTextNamed(
		t,
		text,
		"Dutch",
		lang)
}

func TestKannadaPhrase(t *testing.T) {
	text := "ನನ್ನ ಹೆಸರು ಭಾರತ್."
	lang := "ಕನ್ನಡ"

	ensureClassifiedWithConfidence(
		t,
		text,
		"kn",
		0.95)

	ensureClassifiedTextNamed(
		t,
		text,
		"Kannada",
		lang)
}

func TestMalayalamPhrase(t *testing.T) {
	text := "എന്റെ പേര് ഭാരത്."
	lang := "മലയാളം"

	ensureClassifiedWithConfidence(
		t,
		text,
		"ml",
		0.95)

	ensureClassifiedTextNamed(
		t,
		text,
		"Malayalam",
		lang)
}

func TestMarathiPhrase(t *testing.T) {
	text := "माझे नाव भरत आहे."
	lang := "मराठी"

	ensureClassifiedWithConfidence(
		t,
		text,
		"mr",
		0.95)

	ensureClassifiedTextNamed(
		t,
		text,
		"Marathi",
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
