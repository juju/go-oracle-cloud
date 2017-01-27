package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) Authenticate() (err error) {
	if c.isAuth() {
		return ErrAlreadyAuth
	}

	// build the json struct authentication
	auth := map[string]string{
		"user":     fmt.Sprintf("/Compute-%s/%s", c.identify, c.username),
		"password": c.password,
	}

	// build the body based on the json struct encoded provided
	body, err := json.Marshal(auth)
	if err != nil {
		return err
	}

	// build the new url based on the authentication and base endpoint
	url := fmt.Sprintf("%s/%s", c.endpoint, "authenticate/")
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/oracle-compute-v3+json")
	resp, err := c.http.Do(req)
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			err = closeErr
		}
	}()
	if err != nil {
		return err
	}

	// If the operation is successfull then we will recive 204 http status
	// if this is not the case then we should stop and return a friendly error
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf(
			"Invalid authentication request, getting status %d",
			resp.StatusCode,
		)
	}

	// the orcale api uses cookies to manage sessions
	// once a cookie is taken then we can make
	// more connections to other resources of the api
	cookies := resp.Cookies()
	if len(cookies) != 1 {
		return fmt.Errorf("Invalid number of session cookies: %q", cookies)
	}

	// add the session cookie
	c.cookie = cookies[0]

	return nil
}
