package api

import (
	"errors"
	"net/http"
	"net/url"
)

// Config represents the significat details that a client
// needs in order to interact with the oracle cloud api.
type Config struct {
	// Identify will hold the oracle cloud client identify endpoint name
	Identify string
	// Username will hold the username oracle cloud client account
	Username string
	// Password will be the password of the orcale cloud client account
	Password string
	// Endpoint will hold the base url endpoint of the oracle cloud api
	Endpoint string
}

func (c Config) validate() error {
	if c.Identify == "" {
		return errors.New("Empty identify endpoint name")
	}

	if c.Password == "" {
		return errors.New("Empty client password")
	}

	if c.Username == "" {
		return errors.New("Empty client username")
	}

	if c.Endpoint == "" {
		return errors.New("Empty endpoint url basepath")
	}

	if _, err := url.Parse(c.Endpoint); err != nil {
		return errors.New("The endpoint provided is invalid")
	}

	return nil
}

// Client holds the client credentials of the clients
// oracle cloud.
// The client needs identify name, user name and
// password in order to comunicate with the oracle
// cloud provider
type Client struct {
	identify string
	username string
	password string
	cookie   *http.Cookie
	endpoint string

	// internal http client
	http http.Client
	// internal http cookie
	// this cookie will be generated based on the client connection
}

func NewClient(cfg Config) (*Client, error) {
	var err error
	if err = cfg.validate(); err != nil {
		return nil, err
	}

	cli := &Client{
		identify: cfg.Identify,
		username: cfg.Username,
		password: cfg.Password,
		endpoint: cfg.Endpoint,
		http:     http.Client{},
	}

	return cli, nil
}

func (c Client) isAuth() bool {
	if c.cookie == nil {
		return false
	}
	return true
}
