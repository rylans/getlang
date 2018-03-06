package getlang_test

import (
	"fmt"
	"github.com/rylans/getlang"
)

func ExampleInfo_Confidence() {
	short := getlang.FromString("short sentences")
	long := getlang.FromString("this sentence is a bit longer")
	fmt.Println(long.Confidence() > short.Confidence())
	// Output: true
}

func ExampleInfo_LanguageCode() {
	fmt.Println(getlang.FromString("статей на русском языке").LanguageCode())
	// Output: ru
}

func ExampleInfo_LanguageName() {
	fmt.Println(getlang.FromString("何ですか？").LanguageName())
	// Output: Japanese
}

func ExampleInfo_SelfName() {
	fmt.Println(getlang.FromString("何ですか？").SelfName())
	// Output: 日本語
}
