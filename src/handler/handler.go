package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func MakeGet(url string) (map[string]interface{}, error) {
	return Make(http.MethodGet, url, nil)
}

func MakeDelete(url string) (map[string]interface{}, error) {
	return Make(http.MethodDelete, url, nil)
}

func MakePut(url string, body io.Reader) (map[string]interface{}, error) {
	return Make(http.MethodPut, url, body)
}

func MakePatch(url string, body io.Reader) (map[string]interface{}, error) {
	return Make(http.MethodPatch, url, body)
}

func MakePost(url string, body io.Reader) (map[string]interface{}, error) {
	return Make(http.MethodPost, url, body)
}

func Make(method, url string, body io.Reader) (map[string]interface{}, error) {
	var attrs = make(map[string]interface{})
	req, err := new(method, url, body)
	if err != nil {
		return attrs, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return attrs, err
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return attrs, err
	}

	err = json.Unmarshal(raw, &attrs)
	if err != nil {
		return attrs, err
	}

	return attrs, nil
}

func new(method, url string, body io.Reader) (*http.Request, error) {
	var req *http.Request
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return req, err
	}
	req.Header.Set("content-type", "application/json")
	req.Header.Set("authentication", apiKey())
	return req, nil
}

func apiKey() string {
	api, ok := os.LookupEnv("API_KEY")
	if !ok {
		panic("set API_KEY environment variable")
	}
	return fmt.Sprintf("Bearer %s", api)
}
