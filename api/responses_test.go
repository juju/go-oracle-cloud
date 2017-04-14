// Copyright 2017 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package api_test

var (
	accountDetailsRaw = []byte(`
{
        "credentials": {},
        "uri": "https://api-z999.compute.us0.oraclecloud.com/account/Compute-acme/cloud_storage",
        "name": "/Compute-acme/cloud_storage",
        "description": ""
}`)[1:]

	allAccountsRaw = []byte(`
{
        "result": [{
                "credentials": {},
                "uri": "https://api-z999.compute.us0.oraclecloud.com/account/Compute-acme/cloud_storage",
                "name": "/Compute-acme/cloud_storage",
                "description": ""
        }, {
                "credentials": {},
                "uri": "https://api-z999.compute.us0.oraclecloud.com/account/Compute-acme/default",
                "name": "/Compute-acme/default",
                "description": null
        }]
}`)[1:]

	accountNamesRaw = []byte(`
{"result": ["/Compute-acme/"]}`)[1:]

	directoryAccountRaw = []byte(`
{
  "result": [
    "/Compute-acme/cloud_storage",
    "/Compute-acme/default"  ]
}`)[1:]
)
