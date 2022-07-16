package types

type Contributor struct {
	User          string `json: "login"`
	AvartarUrl    []byte `json: "avartar_url"`
	Url           string `json: "html_url"`
	Contributions uint   `json: "contributions"`
}

type EnrollInfo struct {
	Url     string
	Subject string
}

func (e *EnrollInfo) SetSuject(user string) {
	e.Subject = user
}
