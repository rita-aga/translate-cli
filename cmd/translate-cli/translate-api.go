package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"unicode/utf8"
)

type Translation struct {
	Text       string `json:"text"`
	Translated string `json:"translated"`
	TargetLang string `json:"target_lang"`
	SourceLang string `json:"source_lang"`
	Response   string `json:"response"`
	Status     int    `json:"status"`
}

func GetTranslation(text string, target string) string {

	// TODO check if source and target are valid languages

	url := "https://kiaratranslate.p.rapidapi.com/get_translated/"

	payloadStr := fmt.Sprintf("{ \"input\": \"%s\", \"lang\": \"%s\"}", text, target)

	payload := strings.NewReader(payloadStr)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("x-rapidapi-host", "kiaratranslate.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "e6342aca15mshbb1f1d6218075bbp163e51jsnddd97459538f")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("accept", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error reading data: %s\n", err)
	}

	var translation Translation
	json.Unmarshal(body, &translation)

	translatedStr := translation.Translated
	decodedStr := ""

	// loop through runes in translated string and decode them
	for len(translatedStr) > 0 {
		r, size := utf8.DecodeRuneInString(translatedStr)

		translatedStr = translatedStr[size:]
		decodedStr += string(r)
	}

	return decodedStr

}
