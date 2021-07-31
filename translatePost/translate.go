package translatePost

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"gopkg.in/ini.v1"
)

type RequestBody struct {
	Text   string `json:"text"`
	Source string `json:"source"`
	Target string `json:"target"`
}

type ResBody struct {
	TranslatedText string `json:"translatedText"`
}

func Translate_post(text, source, target string) string {
	// you can set "ja" or "en" in source and target
	requestBody := &RequestBody{
		Text:   text,
		Source: source,
		Target: target,
	}
	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		panic("Error")
	}

	cfg, err := ini.Load("translatePost/config.ini")
	if err != nil {
		panic("Error")
	}
	endpoint := cfg.Section("api").Key("url").String()

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonString))
	if err != nil {
		panic("Error")
	}

	req.Header.Set("Content-Type", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		panic("Error")
	}

	defer resp.Body.Close()

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("Error")
	}

	// fmt.Printf(string(byteArray))

	var resbody ResBody
	err = json.NewDecoder(bytes.NewReader(byteArray)).Decode(&resbody)
	if err != nil {
		panic("Error")
	}

	// fmt.Println(resbody.TranslatedText)
	return resbody.TranslatedText
}
