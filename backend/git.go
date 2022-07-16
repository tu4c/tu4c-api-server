package backend

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// in contract: https://api.github.com/repos/[Owner]/[Repo]/contributors?
const (
	DefaultRankSet = 500
	QueryPrefix    = "per_page="
	QueryPostfix   = "&page=1"
)

// need rate limit.
// avoid DOS
type GitConn struct {
}

func (g *GitConn) QueryString() string {
	return fmt.Sprintf(QueryPrefix + string(DefaultRankSet) + QueryPostfix)
}

// it's not concern about queryString. and get only one contributions.
func (g *GitConn) contributiors(given string) (*[]types.Contributor, error) {
	_, err := url.Parse(given)
	if err != nil {
		return nil, errors.New("")
	}
	// get list
	resp, err := http.Get(given + g.QueryString())
	if err != nil {
		return nil, errors.New("")
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var contributors *[]types.Contributor
	err = json.Unmarshal(data, contributors)
	if err != nil {
		return nil, err
	}
	return contributors, nil
}
