package main

import (
	flag "github.com/ogier/pflag"
)

var (
	sourceLang string
	targetLang string
	sourceText string
	targetText string
)

func main() {
    flag.Parse()
}

func init() {
	// StringVarP params:
	// 1. the variable we want to bind this flag to,
	// 2. the double-dash flag,
	// 3. the single-dash flag,
	// 4. the default value to use if flag is not explicitly called,
	// 5. and the description of this flag
	flag.StringVarP(&sourceLang, "source", "s", "EN", "Source Language")
	flag.StringVarP(&targetLang, "target", "t", "RU", "Target Language")
}