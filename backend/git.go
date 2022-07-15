package backend

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

// need rate limit.
// avoid DOS
type GitConn struct {
}

// it's not concern about queryString. and get only one contributions.
func (g *GitConn) contribution(given string) (*types.Contribution, error) {
	_, err := url.Parse(given)
	if err != nil {
		return nil, errors.New("")
	}
	resp, err := http.Get(given)
	if err != nil {
		return nil, errors.New("")
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var contribution types.Contribution
	err = json.Unmarshal(data, &contribution)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
}
