package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
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
	cookie   http.Cookie
	endpoint string

	// internal http client
	http http.Client
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
	}
	endpoint := fmt.Sprintf("%s/Compute-%s/", cfg.Endpoint, cfg.Identify)
	if _, err = url.Parse(endpoint); err != nil {
		return nil, err
	}
	cli.endpoint = endpoint

	cli.http = http.Client{}

	return cli, nil
}

func (c *Client) Authenticate() (err error) {
	authenticate := map[string]string{
		"user":     fmt.Sprintf("/Compute-%s/%s", c.identify, c.username),
		"password": c.password,
	}

	body, err := json.Marshal(authenticate)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s%s", c.endpoint, "authenticate")
	fmt.Println(url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/oracle-compute-v3+json")

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}

	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			err = closeErr
		}
	}()

	// the api in this case should not return any body at all.
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("Cannot generate cookie, status code %d", resp.StatusCode)
	}

	cookies := resp.Cookies()
	fmt.Println(cookies)
	os.Exit(1)
	return nil

}
