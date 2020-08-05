package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/gernest/wow"
	"github.com/gernest/wow/spin"
	flag "github.com/ogier/pflag"
)

var (
	targetLang string
	sourceText string
)

func main() {
	flag.Parse()

	availableTargetLangs := [][]string{{"ru", "Russian"},
		{"es", "Spanish"},
		{"fr", "French"},
		{"de", "German"},
		{"zh", "Chinese"},
		{"ja", "Japanese"}}

	// if user does not supply flags, print usage
	if flag.NFlag() == 0 {
		color.Green("Usage: %s [options]\n", os.Args[0])
		color.Blue("Options:")
		flag.PrintDefaults()
		color.Blue("Available target languages:")
		for i := 0; i < len(availableTargetLangs); i++ {
			fmt.Printf("%s: %s\n", availableTargetLangs[i][0], availableTargetLangs[i][1])
		}
		os.Exit(1)
	}

	if sourceText == "" {
		fmt.Println("Please enter text to translate")
		os.Exit(1)
	}

	message := fmt.Sprintf("Translating {%s} to: %s\n", sourceText, targetLang)

	color.Blue(message)

	w := wow.New(os.Stdout, spin.Get(spin.Dots), "")
	w.Start()

	translated := GetTranslation(sourceText, targetLang)

	w.PersistWith(spin.Spinner{Frames: []string{"👍"}}, " "+translated)
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
