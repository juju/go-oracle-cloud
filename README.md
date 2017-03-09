# go-oracle-cloud

Client library providing a simple interface with the oracle cloud IAAS rest api.

[![Build Status](https://travis-ci.org/juju/go-oracle-cloud.svg?branch=master)](https://travis-ci.org/juju/go-oracle-cloud) [![GoDoc](https://godoc.org/github.com/juju/go-oracle-cloud?status.svg)](https://godoc.org/github.com/juju/go-oracle-cloud) [![Go Report Card](https://goreportcard.com/badge/github.com/juju/go-oracle-cloud)](https://goreportcard.com/report/github.com/juju/go-oracle-cloud)

## Example client authentication

```go
package main

import (
	"fmt"
	oracle "github.com/juju/go-oracle-cloud/api"
)

func main() {
	// create the configuration for the client
	cfg := oracle.Config{
		Username: "oracle@username.com",
		Password: "oraclepassword",
		Identify: "qbitq",
		Endpoint: "https://api-z52.compute.us2.oraclecloud.com",
	}

	// create a new client based on the configuration
	cli, err := oracle.NewClient(cfg)
	if err != nil {
		fmt.Println(err)
		return
	}

	// authenticate with the client
	err = cli.Authenticate()
	if err != nil {
		fmt.Println(err)
		return
	}


	// code
}


```
