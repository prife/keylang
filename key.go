package main

import (
	"fmt"
	"os"
	"text/scanner"
)

const (
	tok_eof = -(iota + 1)

	//keyword
	tok_def
	tok_extern

	tok_ident
	tok_number
)

var scan scanner.Scanner

func LexInit() {
	scan.Init(os.Stdin)
}

func Token() (ret rune) {
	tok := scan.Scan()
	ret = tok
	switch tok {
	case scanner.Ident:
		toktxt := scan.TokenText()
		switch toktxt {
		case "def":
			ret = tok_def
		case "extern":
			ret = tok_extern
		default:
			ret = tok_ident
		}
	case scanner.String:
		ret = tok_ident
	case scanner.RawString:
		ret = tok_ident
	case scanner.Int:
		ret = tok_number
	case scanner.EOF:
		ret = tok_eof
	}
	return ret
}

func Text() string {
	return scan.TokenText()
}

func MainLoop() {
	var tok rune
	for tok != tok_eof {
		tok = Token()
		fmt.Printf("tok:%v, text:%s\n", tok, Text())
	}
}

func main() {
	LexInit()
	MainLoop()
}
