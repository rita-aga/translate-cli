package main

import (
	"fmt"

	flag "github.com/ogier/pflag"
)

var (
	targetLang string
	sourceText string
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
	flag.StringVarP(&targetLang, "target", "t", "ru", "Target Language")

	// if user does not supply flags, print usage
	// we can clean this up later by putting this into its own function
	//  if flag.NFlag() == 0 {
	// 	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	// 	fmt.Println("Options:")
	// 	flag.PrintDefaults()
	// 	os.Exit(1)
	//  }

	// users = strings.Split(user, ",")
	// fmt.Printf("Translating to: %s\n", targetLang)

	translated := GetTranslation("it's a beautiful day", "ru")

	fmt.Println(translated)
}
