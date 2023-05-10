package quotes

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func Get() (Quote, error) {
	params := make(url.Values)
	params.Add("method", "getQuote")
	params.Add("format", "json")
	params.Add("lang", "ru")

	req, err := http.NewRequest("POST", "http://api.forismatic.com/api/1.0/", strings.NewReader(params.Encode()))
	if err != nil {
		return Quote{}, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Quote{}, err
	}
	defer resp.Body.Close()

	decodeBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return Quote{}, err
	}

	var quote Quote
	if err := json.Unmarshal(decodeBody, &quote); err != nil {
		return Quote{}, err
	}

	return quote, nil
}

type Quote struct {
	QuoteText   string `json:"quoteText"`
	QuoteAuthor string `json:"quoteAuthor"`
	SenderName  string `json:"senderName"`
	SenderLink  string `json:"senderLink"`
	QuoteLink   string `json:"quoteLink"`
}
