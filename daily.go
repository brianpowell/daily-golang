package daily

import (
	"errors"
	"strings"
)

type (
	Client struct {
		Config     Config
		DomainLink string
	}

	Config struct {
		Domain string
		Token  string
	}
)

var (
	API        = "https://api.daily.co/v1"
	DomainLink = "https://{domain}.daily.co/"
)

// Create a New Client
func NewClient(config Config) (*Client, error) {
	c := &Client{}
	// Quick validation
	if err := ConfigValidate(config); err != nil {
		return nil, err
	}
	// Generate some values - only DomainLink now
	c.DomainLink = strings.Replace(DomainLink, "{domain}", config.Domain, 1)
	c.Config = config
	return c, nil
}

// Validate the Config values are set
func ConfigValidate(config Config) error {
	var err error
	if config.Domain == "" {
		err = errors.New("missing domain config value")
	}
	if config.Token == "" {
		err = errors.New("missing token config value")
	}
	return err
}
