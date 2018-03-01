package getlang

type Info struct {
  lang string
  probability float64
}

func FromString(text string) Info {
  return Info{"en", 1.0}
}
