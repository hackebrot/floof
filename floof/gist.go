package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

const url = "https://api.github.com/gists"

//gistReponse used for unmarshalling the GitHub API response body
type gistReponse struct {
	GistURL string `json:"html_url"`
}

//GistFile representing a local file
type GistFile struct {
	Name    string `json:"-"`
	Content string `json:"content"`
}

//Gist struct matching the API object for POST /gists
type Gist struct {
	Files       map[string]GistFile `json:"files"`
	Description string              `json:"description"`
	Public      bool                `json:"public"`
}

//Post request to create a gist via the GitHub API
//https://developer.github.com/v3/gists/#create-a-gist
func (g Gist) Post(username string, token string) (gistURL string, err error) {
	j, err := json.Marshal(g)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	req.SetBasicAuth(username, token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		s := []string{"POST", url, resp.Status}
		return "", errors.New(strings.Join(s, " "))
	}
	body, _ := ioutil.ReadAll(resp.Body)

	var r gistReponse
	err = json.Unmarshal(body, &r)
	if err != nil {
		return "", err
	}
	return r.GistURL, nil
}
