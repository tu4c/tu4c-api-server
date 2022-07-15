package types

type Contribution struct {
	User          string `json: "login"`
	AvartarUrl    []byte `json: "avartar_url"`
	Url           string `json: "html_url"`
	Contributions uint   `json: "contributions"`
}
