package api

import (
	"errors"
	"fmt"
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
		return errors.New("go-oracle-cloud: Empty identify endpoint name")
	}

	if c.Password == "" {
		return errors.New("go-oracle-cloud: Empty client password")
	}

	if c.Username == "" {
		return errors.New("go-oracle-cloud: Empty client username")
	}

	if c.Endpoint == "" {
		return errors.New("go-oracle-cloud: Empty endpoint url basepath")
	}

	if _, err := url.Parse(c.Endpoint); err != nil {
		return errors.New("go-oracle-cloud: The endpoint provided is invalid")
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

// RefreshCookie refreshes the authentication tokens that expires usually around 30 minutes.
// This request extends the expiry of a valid authentication
// token by 30 minutes from the time you run the command.
// It extends the expiry of the current authentication token,
// but not beyond the session expiry time, which is 3 hours.
func (c *Client) RefreshCookie() (err error) {
	url := fmt.Sprintf("%s/%s/", c.endpoint, "refresh")
	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		verb:   "GET",
		url:    url,
		treat: func(resp *http.Response) (err error) {
			if resp.StatusCode != http.StatusNoContent {
				return fmt.Errorf(
					"go-oracle-cloud: Error api response %d %s",
					resp.StatusCode, dumpApiError(resp.Body),
				)
			}

			// take the new refresh cookies
			cookies := resp.Cookies()
			if len(cookies) != 1 {
				return fmt.Errorf("go-oracle-cloud: Invalid number of session cookies: %q", cookies)
			}

			// take the cookie
			c.cookie = cookies[0]
			return nil
		},
	}); err != nil {
		return err
	}

	return nil
}
