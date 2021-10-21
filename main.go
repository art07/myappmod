package main

import (
	"fmt"

	qt "art/edu/gopy/pkg-mod/myappmod/quotes"
	str "art/edu/gopy/pkg-mod/myappmod/strings"
)

func main() {
	odds, evens := str.CountOddEven("12345")
	fmt.Println(odds, evens) // 3 2

	fmt.Println(qt.GetEmoji("turtle"))
}
