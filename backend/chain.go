package backend

import "github.com/ethereum/go-ethereum/ethclient"

// managing to ethereum node connection
type Chain struct {
	name string
	url  string
	eth  *ethclient.Client
}

func NewChain(name string, url string) (*Chain, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}
	return &Chain{
		name: name,
		url:  url,
		eth:  client,
	}, nil
}

func (c *Chain) Name() string { return c.name }
func (c *Chain) Url() string  { return c.url }
