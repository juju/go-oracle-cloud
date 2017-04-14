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

	aclDetailsRaw = []byte(`
{
  "description": "Sample ACL 1",
  "enabledFlag": true,
  "name": "/Compute-acme/jack.jones@example.com/acl1",
  "tags": null,
  "uri": "https://api-z999.compute.us0.oraclecloud.com/network/v1/acl/Compute-acme/jack.jones@example.com/acl1"
}`)[1:]

	allAclsRaw = []byte(`
	{
  "result": [
    {
      "name": "/Compute-acme/jack.jones@example.com/acl1",
      "uri": "https://api-z999.compute.us0.oraclecloud.com:443/network/v1/acl/Compute-acme/jack.jones@example.com/acl1",
      "description": "Updating sample ACL 1",
      "tags": [
        "test"
      ],
      "enabledFlag": false
    },
    {
      "name": "/Compute-acme/default",
      "uri": "https://api-z999.compute.us0.oraclecloud.com:443/network/v1/acl/Compute-acme/default",
      "description": null,
      "tags": [
        
      ],
      "enabledFlag": true
    }
  ]
}`)[1:]
	backupDetailsRaw = []byte(`
{
  "uri": "https://api-z999.compute.us0.oraclecloud.com/backupservice/v1/configuration/Compute-acme/jack.jones@example.com/backupConfigVol1",
  "runAsUser": "/Compute-acme/jack.jones@example.com",
  "name": "/Compute-acme/jack.jones@example.com/backupConfigVol1",
  "enabled": false,
  "backupRetentionCount": 2,
  "nextScheduledRun": "2016-08-19T05:10:44.859Z",
  "interval": {
    "Hourly": {
      "hourlyInterval": 1
    }
  },
  "volumeUri": "http://api-z999.compute.us0.oraclecloud.com/storage/volume/Compute-acme/jack.jones@example.com/vol1",
  "description": null,
  "tagId": "27f57e2d-f0f6-430d-9c04-40a58d632513"
}`)[1:]

	allBackupsRaw = []byte(`
	{
	"result": [
  {
    "uri": "https://api-z999.compute.us0.oraclecloud.com:443/backupservice/v1/configuration/Compute-acme/jack.jones@example.com/backupConfigVol1",
    "runAsUser": "/Compute-acme/jack.jones@example.com",
    "name": "/Compute-acme/jack.jones@example.com/backupConfigVol1",
    "enabled": true,
    "backupRetentionCount": 2,
    "nextScheduledRun": "2016-10-25T21:25:12.898Z",
    "interval": {
      "Hourly": {
        "hourlyInterval": 12
      }
    },
    "volumeUri": "/storage/volume/Compute-acme/jack.jones@example.com/vol1",
    "description": null,
    "tagId": "63ed3bec-5da2-42d2-9f6a-6440a5c91567"
  },
  {
    "uri": "https://api-z999.compute.us0.oraclecloud.com:443/backupservice/v1/configuration/Compute-acme/jack.jones@example.com/backupConfigWeeklyVol2",
    "runAsUser": "/Compute-acme/jack.jones@example.com",
    "name": "/Compute-acme/jack.jones@example.com/backupConfigWeeklyVol2",
    "enabled": false,
    "backupRetentionCount": 2,
    "nextScheduledRun": "2016-09-28T04:00:00Z",
    "interval": {
      "DailyWeekly": {
        "daysOfWeek": [
          "WEDNESDAY",
          "FRIDAY"
        ],
        "timeOfDay": "04:00",
        "userTimeZone": "GMT"
      }
    },
    "volumeUri": "/storage/volume/Compute-acme/jack.jones@example.com/vol1",
    "description": null,
    "tagId": "4b47cfac-ab33-44fa-a8ae-2dc3f333a44b"
  }
]
	}
`)[1:]
)
