# go-oracle-cloud

Client library providing a simple interface with the oracle cloud IAAS rest api.


## Example client authentication

```go
	package main

	import (
		"fmt"

		oracle "github.com/hoenirvili/go-oracle-cloud/api"
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
