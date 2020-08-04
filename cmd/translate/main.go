package main

import (
	"fmt"
	"os"

	flag "github.com/ogier/pflag"
)

var (
	targetLang string
	sourceText string
)

func main() {
	flag.Parse()

	// if user does not supply flags, print usage
	if flag.NFlag() == 0 {
		fmt.Printf("Usage: %s [options]\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if sourceText == "" {
		fmt.Println("Please enter text to translate")
		os.Exit(1)
	}

	fmt.Printf("Translating {%s} to: %s\n", sourceText, targetLang)

	translated := GetTranslation(sourceText, targetLang)

	fmt.Println(translated)
}

func init() {
	// TODO add language pairs

	// StringVarP params:
	// 1. the variable we want to bind this flag to,
	// 2. the double-dash flag,
	// 3. the single-dash flag,
	// 4. the default value to use if flag is not explicitly called,
	// 5. and the description of this flag
	flag.StringVarP(&targetLang, "target", "t", "ru", "Target Language")
	flag.StringVarP(&sourceText, "source", "s", "", "Source Text")
}
