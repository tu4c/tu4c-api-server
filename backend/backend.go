package backend

type Backend struct {
	contracts map[string]Contract
	git       *GitConn
}

func NewBackend() *Backend {
	return &Backend{}
}

func (b *Backend) Contribution(url string) (types.Contribution, error) {
	return b.git.contribution(url)
}
