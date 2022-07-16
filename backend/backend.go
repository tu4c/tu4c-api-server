package backend

type Backend struct {
	contracts map[string]Contract
	git       *GitConn
}

func NewBackend() *Backend {
	return &Backend{}
}

func (b *Backend) EnrollVerify(enrollInfo types.EnrollInfo) (bool, error) {
	contributors, err := b.Contributors(enrollInfo.Url)
	if err != nil {
		return false, err
	}
	for _, contributor := range contributors {
		if contributor.User == enrollInfo.Subject {
			return true, nil
		}
	}
	// it's not error. need that result is 'false'
	return false, nil
}

func (b *Backend) EnrollInfo(enrollId uint) (types.EnrollInfo, error) {
	return b.contracts["issuer"].EnrollInfo(enrollId)
}

func (b *Backend) Contributors(url string) (types.Contributor, error) {
	return b.git.contributor(url)
}
