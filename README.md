# getlang

[![GoDoc](https://godoc.org/github.com/rylans/getlang?status.svg)](https://godoc.org/github.com/rylans/getlang)


Detect language from text


## Getting started

Installation:
```sh
    go get -u github.com/rylans/getlang
```

example:
```go
package main

import (
	"fmt"
	"github.com/rylans/getlang"
)

func main(){
  info := getlang.FromString("Wszyscy ludzie rodzą się wolni i równi w swojej godności i prawach")
  fmt.Println(info.LanguageCode(), info.Confidence())
}
```

## Supported Languages

| Language       | ISO 639-1 | 
| -------------- | --------- |
| English        | en        |
| Spanish        | es        |
| Portuguese     | pt        |
| Italian        | it        |
| Hungarian      | hu        |
| Polish         | pl        |
| German         | de        |
| Russian        | ru        |
| Ukranian       | uk        |
| Chinese        | zh        |
| Japanese       | ja        |
| Korean         | ko        |

## Documentation
[getlang on godoc](https://godoc.org/github.com/rylans/getlang)

## License
[MIT](https://github.com/rylans/getlang/blob/master/LICENSE)

## Acknowledgements and Citations
* Thanks to [abadojack](https://github.com/abadojack) for the trigram generation logic in whatlanggo
* Cavnar, William B., and John M. Trenkle. "N-gram-based text categorization." Ann arbor mi 48113.2 (1994): 161-175.

