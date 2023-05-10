package requests

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
)

func New(accessToken, apiEndpoint, apiVersion string) *Requester {
	return &Requester{
		Token:       accessToken,
		apiEndpoint: apiEndpoint,
		apiVersion:  apiVersion,
	}
}

func (r *Requester) MakeRequest(method string, params url.Values) ([]byte, error) {
	if params == nil {
		params = make(url.Values)
	}

	params.Add("v", r.apiVersion)
	urlRequest := fmt.Sprintf(r.apiEndpoint, method)

	if *r.Debug {
		log.Printf("[Debug] UrlRequest: %v, Method: %s, Params: %v\n", urlRequest, method, params)
	}

	req, err := http.NewRequest("POST", urlRequest, strings.NewReader(params.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", r.Token))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		if *r.Debug {
			log.Printf("[Debug] UrlRequest: %v, Method: %s, Params: %v, error request: %v\n", urlRequest, method, params, err)
		}

		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil && *r.Debug {
			log.Printf("[Debug] UrlRequest: %v, Method: %s, params: %v, error body close: %v\n", urlRequest, method, params, err)
		}
	}(resp.Body)

	decodeBody, err := io.ReadAll(resp.Body)
	if err != nil {
		if *r.Debug {
			log.Printf("[Debug] UrlRequest: %v, Method: %s, params: %v, error body decode: %v\n", urlRequest, method, params, err)
		}

		return nil, err
	}

	return decodeBody, nil
}

func (r *Requester) UploadFile(url string, file io.Reader) ([]byte, error) {
	if *r.Debug {
		log.Printf("[Debug] UrlRequest: %v, File: %s\n", url, file)
	}

	// New multipart writer.
	// https://golangbyexample.com/http-pdf-post-go/
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	fw, err := writer.CreateFormFile("photo", "photo.jpeg")
	if err != nil {
		if *r.Debug {
			log.Printf("[Debug] UrlRequest: %v, File: %s, error create writer - %v\n", url, file, err)
		}

		return nil, err
	}

	_, err = io.Copy(fw, file)
	if err != nil {
		if *r.Debug {
			log.Printf("[Debug] UrlRequest: %v, File: %s, error io copy - %v\n", url, file, err)
		}

		return nil, err
	}
	_ = writer.Close()

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		if *r.Debug {
			log.Printf("[Debug] UrlRequest: %v, File: %s, error request - %v\n", url, file, err)
		}

		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil && *r.Debug {
			log.Printf("[Debug] UrlRequest: %v, File: %s, error body close - %v\n", url, file, err)
		}
	}(resp.Body)

	decodeBody, err := io.ReadAll(resp.Body)
	if err != nil {
		if *r.Debug {
			log.Printf("[Debug] UrlRequest: %v, File: %s, error body decode - %v\n", url, file, err)
		}

		return nil, err
	}

	return decodeBody, nil
}
