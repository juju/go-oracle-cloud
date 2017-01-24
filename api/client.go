package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

// defines the region where the oracle cloud
// infrastracture is geo locatted.
type Region string

const (
	US1 = "us1"
	US3 = "us3"
)

// base url for oracle cloud api
var base = "https://api-z999.compute.%s.oraclecloud.com"

// Config represents the significat details that a client
// needs in order to interact with the oracle cloud api.
type Config struct {
	// Identify will hold the oracle cloud client identify endpoint name
	Identify string
	// Username will hold the username oracle cloud client account
	Username string
	// Password will be the password of the orcale cloud client account
	Password string
	// Region will be the region on what oracle infrastructure we wish to use
	Region Region
}

func (c Config) validate() (bool, error) {
	if c.Identify == "" {
		return false, errors.New("Empty identify endpoint name")
	}

	if c.Password == "" {
		return false, errors.New("Empty client password")
	}

	if c.Username == "" {
		return false, errors.New("Empty client username")
	}

	if c.Region == "" {
		return false, errors.New("Empty cloud region")
	}

	return true, nil
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
	r        *http.Request
}

func NewClient(cfg Config) (*Client, error) {
	ok, err := cfg.validate()
	if !ok {
		return nil, err
	}

	cli := &Client{
		identify: cfg.Identify,
		username: cfg.Username,
		password: cfg.Password,
		endpoint: fmt.Sprintf("https://api-z999.compute.%s.oraclecloud.com", cfg.Region),
	}

	if err = cli.newRequest(); err != nil {
		return nil, err
	}

	return cli, nil
}

func (c *Client) newRequest() error {
	var err error
	c.r, err = http.NewRequest("", c.endpoint, nil)
	if err != nil {
		return err
	}

	c.r.Header.Del("Agent") // remove the golang agent crap

	// add oracle comaptible headers
	c.r.Header.Add("Content-Type", "application/oracle-compute-v3+json")
	c.r.Header.Add("user", c.username)
	c.r.Header.Add("password", c.password)

	fmt.Println(c.r)
	os.Exit(1)
	if c.r.Cookie == nil {
		c.newCookie()
	}

	c.r.AddCookie(c.cookie)

	return nil
}

func (c *Client) newCookie() error {
	sufix := "/authenticate/"
	url := fmt.Sprintf(c.endpoint, sufix)

	credentials := map[string]string{"username": c.username, "password": c.password}
	body, err := json.Marshal(credentials)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Cannot generate cookie, status code %d", resp.StatusCode)
	}

	cookies := resp.Cookies()
	fmt.Println(cookies)
	//TODO
	if resp.Body.Close() != nil {
		return errors.New("Cannot close response body cookie")
	}

	return nil
}
