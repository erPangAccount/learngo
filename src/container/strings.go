package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes，我爱golang!"
	fmt.Println(len(s))                                   //19
	fmt.Println("rune count:", utf8.RuneCount([]byte(s))) //rune count: 13
	fmt.Println("rune count:", utf8.RuneCountInString(s)) //rune count: 13

	_bytes := []byte(s)
	for len(_bytes) > 0 {
		ch, size := utf8.DecodeRune(_bytes)
		_bytes = _bytes[size:]
		fmt.Printf("(%c %d)", ch, size) //(Y 1)(e 1)(s 1)(， 3)(我 3)(爱 3)(g 1)(o 1)(l 1)(a 1)(n 1)(g 1)(! 1)
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c)", i, ch) //(Y 1)(e 1)(s 1)(， 3)(我 3)(爱 3)(g 1)(o 1)(l 1)(a 1)(n 1)(g 1)(! 1)
	}
}
