package plfoe_mod_1

import (
	"fmt"
	"rsc.io/quote"
)

func Hello() string {
	fmt.Println("Hello, world - from sdk remote.")
	return quote.Hello()
}

//func Hello() string {
//	return "Hello, world."
//}
