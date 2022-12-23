package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

//==================================================================================================

type Translation struct {
	Translations []struct {
		DetectedSourceLanguage string `json:"detected_source_language"`
		Text                   string `json:"text"`
	}
}

// a translate func that will return japanese character

// ==================================================================================================
func TranslateToJapanese(s string) string {
	data := url.Values{
		"text":        {s},
		"target_lang": {"JA"},
	}

	req, err := http.NewRequest("POST", "https://api-free.deepl.com/v2/translate", bytes.NewBufferString(data.Encode()))
	if err != nil {
		fmt.Println(err)
		return ""
	}

	req.Header.Set("Authorization", "DeepL-Auth-Key "+DeeplToken)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()

	var t Translation
	err = json.NewDecoder(resp.Body).Decode(&t)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	fmt.Println(t.Translations[0].Text)
	return t.Translations[0].Text
}

//==================================================================================================

func TranslateToEnglish(s string) string {
	data := url.Values{
		"text":        {s},
		"target_lang": {"EN"},
	}

	req, err := http.NewRequest("POST", "https://api-free.deepl.com/v2/translate", bytes.NewBufferString(data.Encode()))
	if err != nil {
		fmt.Println(err)
		return ""
	}

	req.Header.Set("Authorization", "DeepL-Auth-Key "+DeeplToken)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()

	// decode resp.Body to struct
	var t Translation
	err = json.NewDecoder(resp.Body).Decode(&t)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	fmt.Println(t.Translations[0].Text)
	return t.Translations[0].Text

}
