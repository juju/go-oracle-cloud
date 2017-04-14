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
	backupConfigurationDetailsRaw = []byte(`
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

	backupDetailsRaw = []byte(`
{
  "uri": "https://api-z999.compute.us0.oraclecloud.com/backupservice/v1/backup/Compute-acme/jack.jones@example.com/BACKUP-A",
  "runAsUser": "/Compute-acme/jack.jones@example.com",
  "name": "/Compute-acme/jack.jones@example.com/BACKUP-A",
  "backupConfigurationName": "/Compute-acme/jack.jones@example.com/backupConfigVol1",
  "volumeUri": "http://api-z999.compute.us0.oraclecloud.com/storage/volume/Compute-acme/jack.jones@example.com/vol1",
  "errorMessage": null,
  "detailedErrorMessage": null,
  "state": "SUBMITTED",
  "description": null,
  "bootable": false,
  "shared": false,
  "snapshotUri": null,
  "snapshotSize": null,
  "tagId": "0d22fec6-fc3c-4987-8021-8f0cf49b8737"
}`)[1:]
	allBackupsRaw = []byte(`
{
 "result": [
  {
    "uri": "https://api-z999.compute.us0.oraclecloud.com:443/backupservice/v1/backup/Compute-acme/jack.jones@example.com/backupConfigWeeklyVol1/930d16e4-83af-4119-9f32-691d4541e5a7",
    "runAsUser": "/Compute-acme/jack.jones@example.com",
    "name": "/Compute-acme/jack.jones@example.com/backupConfigWeeklyVol1/930d16e4-83af-4119-9f32-691d4541e5a7",
    "backupConfigurationName": "/Compute-acme/jack.jones@example.com/backupConfigWeeklyVol1",
    "volumeUri": "/storage/volume/Compute-acme/jack.jones@example.com/vol1",
    "errorMessage": "",
    "detailedErrorMessage": "",
    "state": "COMPLETED",
    "description": null,
    "bootable": false,
    "shared": false,
    "snapshotUri": "/storage/snapshot/Compute-acme/jack.jones@example.com/vol1/5fe6bc70a4c9f0fcaf5a094a01f22364f77952a225e862530c36fb59ff9aaa28",
    "snapshotSize": "1073741824b",
    "tagId": "9d2b05f1-927b-4083-92dd-b565530e373d"
  },
  {
    "uri": "https://api-z999.compute.us0.oraclecloud.com:443/backupservice/v1/backup/Compute-acme/jack.jones@example.com/vol1-BACKUP-B",
    "runAsUser": "/Compute-acme/jack.jones@example.com",
    "name": "/Compute-acme/jack.jones@example.com/vol1-BACKUP-B",
    "backupConfigurationName": "/Compute-acme/jack.jones@example.com/backupConfigWeeklyVol2",
    "volumeUri": "/storage/volume/Compute-acme/jack.jones@example.com/vol1",
    "errorMessage": "",
    "detailedErrorMessage": "",
    "state": "COMPLETED",
    "description": null,
    "bootable": false,
    "shared": false,
    "snapshotUri": "/storage/snapshot/Compute-acme/jack.jones@example.com/vol1/b2faedfd2d62b5ddbd856bf557235df49e6fcfba5ba91ccddc37893adc73757e",
    "snapshotSize": "1073741824b",
    "tagId": "22e6eaf1-3f2b-43b3-9505-63abec384e18"
  }]
  }`)[1:]

	allImageListDetailsRaw = []byte(`
  {
    "result": [
        {
            "default": 1,
            "description": "Microsoft_Windows_Server_2012_R2",
            "entries": [
                {
                    "attributes": {
                        "defaultShape": "oc4",
                        "minimumDiskSize": "27",
                        "supportedShapes": "oc1m,oc2m,oc3,oc3m,oc4,oc4m,oc5,oc5m,oc6,oc7,ocio1m,ocio2m,ocio3m,ocio4m,ocio5m,ociog1k80,ociog2k80,ociog3k80",
                        "userdata": {
                            "enable_rdp": "true"
                        },
                        "windows_kms": "kms.oraclecloud.com"
                    },
                    "machineimages": [
                        "/Compute-a432100/sgiulitti@cloudbase.com/Microsoft_Windows_Server_2012_R2"
                    ],
                    "uri": "https://compute.uscom-central-1.oraclecloud.com/imagelist/Compute-a432100/sgiulitti%40cloudbase.com/Microsoft_Windows_Server_2012_R2/entry/1",
                    "version": 1
                }
            ],
            "name": "/Compute-a432100/sgiulitti@cloudbase.com/Microsoft_Windows_Server_2012_R2",
            "uri": "https://compute.uscom-central-1.oraclecloud.com/imagelist/Compute-a432100/sgiulitti%40cloudbase.com/Microsoft_Windows_Server_2012_R2"
        },
        {
            "default": 1,
            "description": "Ubuntu.16.04-LTS.amd64.20170307",
            "entries": [
                {
                    "attributes": {
                        "defaultShape": "oc2m",
                        "minimumDiskSize": "10",
                        "supportedShapes": "oc3,oc4,oc5,oc6,oc7,oc1m,oc2m,oc3m,oc4m,oc5m",
                        "userdata": {}
                    },
                    "machineimages": [
                        "/Compute-a432100/sgiulitti@cloudbase.com/Ubuntu.16.04-LTS.amd64.20170307"
                    ],
                    "uri": "https://compute.uscom-central-1.oraclecloud.com/imagelist/Compute-a432100/sgiulitti%40cloudbase.com/Ubuntu.16.04-LTS.amd64.20170307/entry/1",
                    "version": 1
                }
            ],
            "name": "/Compute-a432100/sgiulitti@cloudbase.com/Ubuntu.16.04-LTS.amd64.20170307",
            "uri": "https://compute.uscom-central-1.oraclecloud.com/imagelist/Compute-a432100/sgiulitti%40cloudbase.com/Ubuntu.16.04-LTS.amd64.20170307"
        },
        {
            "default": 1,
            "description": "Ubuntu.14.04-LTS.amd64.20170307",
            "entries": [
                {
                    "attributes": {
                        "defaultShape": "oc2m",
                        "minimumDiskSize": "10",
                        "supportedShapes": "oc3,oc4,oc5,oc6,oc7,oc1m,oc2m,oc3m,oc4m,oc5m",
                        "userdata": {}
                    },
                    "machineimages": [
                        "/Compute-a432100/sgiulitti@cloudbase.com/Ubuntu.14.04-LTS.amd64.20170307"
                    ],
                    "uri": "https://compute.uscom-central-1.oraclecloud.com/imagelist/Compute-a432100/sgiulitti%40cloudbase.com/Ubuntu.14.04-LTS.amd64.20170307/entry/1",
                    "version": 1
                }
            ],
            "name": "/Compute-a432100/sgiulitti@cloudbase.com/Ubuntu.14.04-LTS.amd64.20170307",
            "uri": "https://compute.uscom-central-1.oraclecloud.com/imagelist/Compute-a432100/sgiulitti%40cloudbase.com/Ubuntu.14.04-LTS.amd64.20170307"
        },
        {
            "default": 1,
            "description": "Ubuntu.12.04-LTS.amd64.20170307",
            "entries": [
                {
                    "attributes": {
                        "defaultShape": "oc2m",
                        "minimumDiskSize": "10",
                        "supportedShapes": "oc3,oc4,oc5,oc6,oc7,oc1m,oc2m,oc3m,oc4m,oc5m",
                        "userdata": {}
                    },
                    "machineimages": [
                        "/Compute-a432100/sgiulitti@cloudbase.com/Ubuntu.12.04-LTS.amd64.20170307"
                    ],
                    "uri": "https://compute.uscom-central-1.oraclecloud.com/imagelist/Compute-a432100/sgiulitti%40cloudbase.com/Ubuntu.12.04-LTS.amd64.20170307/entry/1",
                    "version": 1
                }
            ],
            "name": "/Compute-a432100/sgiulitti@cloudbase.com/Ubuntu.12.04-LTS.amd64.20170307",
            "uri": "https://compute.uscom-central-1.oraclecloud.com/imagelist/Compute-a432100/sgiulitti%40cloudbase.com/Ubuntu.12.04-LTS.amd64.20170307"
        },
        {
            "default": 1,
            "description": "Microsoft_Windows_Server_2008_R2",
            "entries": [
                {
                    "attributes": {
                        "defaultShape": "oc4",
                        "minimumDiskSize": "27",
                        "supportedShapes": "oc1m,oc2m,oc3,oc3m,oc4,oc4m,oc5,oc5m,oc6,oc7,ocio1m,ocio2m,ocio3m,ocio4m,ocio5m,ociog1k80,ociog2k80,ociog3k80",
                        "userdata": {
                            "enable_rdp": "true"
                        },
                        "windows_kms": "kms.oraclecloud.com"
                    },
                    "machineimages": [
                        "/Compute-a432100/sgiulitti@cloudbase.com/Microsoft_Windows_Server_2008_R2"
                    ],
                    "uri": "https://compute.uscom-central-1.oraclecloud.com/imagelist/Compute-a432100/sgiulitti%40cloudbase.com/Microsoft_Windows_Server_2008_R2/entry/1",
                    "version": 1
                }
            ],
            "name": "/Compute-a432100/sgiulitti@cloudbase.com/Microsoft_Windows_Server_2008_R2",
            "uri": "https://compute.uscom-central-1.oraclecloud.com/imagelist/Compute-a432100/sgiulitti%40cloudbase.com/Microsoft_Windows_Server_2008_R2"
        }
    ]
}`)[1:]

	imageListDetailsRaw = []byte(`
{
            "default": 1,
            "description": "Microsoft_Windows_Server_2008_R2",
            "entries": [
                {
                    "attributes": {
                        "defaultShape": "oc4",
                        "minimumDiskSize": "27",
                        "supportedShapes": "oc1m,oc2m,oc3,oc3m,oc4,oc4m,oc5,oc5m,oc6,oc7,ocio1m,ocio2m,ocio3m,ocio4m,ocio5m,ociog1k80,ociog2k80,ociog3k80",
                        "userdata": {
                            "enable_rdp": "true"
                        },
                        "windows_kms": "kms.oraclecloud.com"
                    },
                    "machineimages": [
                        "/Compute-a432100/sgiulitti@cloudbase.com/Microsoft_Windows_Server_2008_R2"
                    ],
                    "uri": "https://compute.uscom-central-1.oraclecloud.com/imagelist/Compute-a432100/sgiulitti%40cloudbase.com/Microsoft_Windows_Server_2008_R2/entry/1",
                    "version": 1
                }
            ],
            "name": "/Compute-a432100/sgiulitti@cloudbase.com/Microsoft_Windows_Server_2008_R2",
            "uri": "https://compute.uscom-central-1.oraclecloud.com/imagelist/Compute-a432100/sgiulitti%40cloudbase.com/Microsoft_Windows_Server_2008_R2"
        }

	`)[1:]

	imageListNamesRaw = []byte(`
{
  "result": [
    "/Compute-acme/jack.jones@example.com/oel59_20G",
    "/Compute-acme/jack.jones@example.com/ol66_40GB"
  ]
}
`)[1:]

	imageListEntryRaw = []byte(`
{
  "attributes": {},
  "version": 2,
  "machineimages": ["/oracle/public/oel_6.4_2GB_v2"],
  "uri": "https://api-z999.compute.us0.oraclecloud.com/imagelist/Compute-acme/jack.jones@example.com/prodimages/entry/2"
}`)[1:]

	instanceConsoleDetailsRaw = []byte(`
{
	"timestamp": "2016-06-17T09:21:19.662570",
	"output": "k [LNKD] (IRQs *5 10 11)\r\nvgaarb: device added: ... login: ",
	"uri": "https://api-z999.compute.us0.oraclecloud.com/instanceconsole/Compute-acme/jack.jones@example.com/68a3c40c-466e-41df-a7f2-00fbfbd590e5",
	"name": "/Compute-acme/jack.jones@example.com/68a3c40c-466e-41df-a7f2-00fbfbd590e5"
}
	`)[1:]

	instanceDetailsRaw = []byte(`
{
 "list": [
  {
   "domain": "acme...",
   "placement_requirements": [
    "/system/compute/placement/default",
    "/system/compute/allow_instances"
   ],
   "ip": "10...",
   "site": "",
   "shape": "oc5",
   "imagelist": "/oracle/public/oel_6.4_60GB",
   "attributes": {
    "network": {
     "nimbula_vcable-eth0": {
      "vethernet_id": "0",
      "vethernet": "/oracle/public/default",
      "address": [
       "c6:b0:09:f4:bc:c0",
       "0.0.0.0"
      ],
      "model": "",
      "vethernet_type": "vlan",
      "id": "/Compute-acme/jack.jones@example.com/016e75e7-e911-42d1-bfe1-6a7f1b3f7908",
      "dhcp_options": []
     }
    },
    "dns": {
     "domain": "acme...",
     "hostname": "d06886.compute-acme...",
     "nimbula_vcable-eth0": "d06886.acme..."
    },
    "sshkeys": [
     "ssh-rsa AAAAB3NzaC1yc2EAAA..."
    ]
   },
   "networking": {
    "eth0": {
     "model": "",
     "dns": [
      "d06886.acme..."
     ],
     "seclists": [
      "/Compute-acme/default/default",
      "/Compute-acme/jack.jones@example.com/prod-ng"
     ],
     "vethernet": "/oracle/public/default",
     "nat": "ipreservation:/Compute-acme/jack.jones@example.com/prod-vm1"
    }
   },
   "hostname": "d06886.acme...",
   "quota_reservation": "/Compute-acme/ffc8e6d4-8f93-41f3-a062-bdbb042c3191",
   "disk_attach": "",
   "label": "Production instance 1",
   "priority": "/oracle/public/default",
   "state": "running",
   "vnc": "10...",
   "storage_attachments": [
    {
     "index": 1,
     "storage_volume_name": "/Compute-acme/jack.jones@example.com/prod-vol1",
     "name": "/Compute-acme/admin/dev1/f653a677-b566-4f92-8e93-71d47b364119/f1a67244-9abc-45d5-af69-8..."
    }
   ],
   "start_time": "2014-06-24T17:51:35Z",
   "quota": "/acme",
   "fingerprint": "19:c4:3f:2d:dc:76:b1:06:e8:88:bd:7f:a3:3b:3c:93",
   "error_reason": "",
   "sshkeys": [
    "/Compute-acme/jack.jones@example.com/acme-prod-admin"
   ],
   "tags": [
    "prod2"
   ],
   "resolvers": null,
   "account": "/Compute-acme/default",
   "name": "/Compute-acme/jack.jones@example.com/dev1/f653a677-b566-4f92-8e93-71d47b364119",
   "vcable_id": "/Compute-acme/jack.jones@example.com/016e75e7-e911-42d1-bfe1-6a7f1b3f7908",
   "uri": "http://10....",
   "reverse_dns": true,
   "entry": 1,
   "boot_order": []
  }
 ]
}`)[1:]

	allInstancesRaw = []byte(`
{
 "list": [
  {
   "domain": "acme...",
   "placement_requirements": [
    "/system/compute/placement/default",
    "/system/compute/allow_instances"
   ],
   "ip": "10...",
   "site": "",
   "shape": "oc5",
   "imagelist": "/oracle/public/oel_6.4_60GB",
   "attributes": {
    "network": {
     "nimbula_vcable-eth0": {
      "vethernet_id": "0",
      "vethernet": "/oracle/public/default",
      "address": [
       "c6:b0:09:f4:bc:c0",
       "0.0.0.0"
      ],
      "model": "",
      "vethernet_type": "vlan",
      "id": "/Compute-acme/jack.jones@example.com/016e75e7-e911-42d1-bfe1-6a7f1b3f7908",
      "dhcp_options": []
     }
    },
    "dns": {
     "domain": "acme...",
     "hostname": "d06886.compute-acme...",
     "nimbula_vcable-eth0": "d06886.acme..."
    },
    "sshkeys": [
     "ssh-rsa AAAAB3NzaC1yc2EAAA..."
    ]
   },
   "networking": {
    "eth0": {
     "model": "",
     "dns": [
      "d06886.acme..."
     ],
     "seclists": [
      "/Compute-acme/default/default",
      "/Compute-acme/jack.jones@example.com/prod-ng"
     ],
     "vethernet": "/oracle/public/default",
     "nat": "ipreservation:/Compute-acme/jack.jones@example.com/prod-vm1"
    }
   },
   "hostname": "d06886.acme...",
   "quota_reservation": "/Compute-acme/ffc8e6d4-8f93-41f3-a062-bdbb042c3191",
   "disk_attach": "",
   "label": "Production instance 1",
   "priority": "/oracle/public/default",
   "state": "running",
   "vnc": "10...",
   "storage_attachments": [
    {
     "index": 1,
     "storage_volume_name": "/Compute-acme/jack.jones@example.com/prod-vol1",
     "name": "/Compute-acme/admin/dev1/f653a677-b566-4f92-8e93-71d47b364119/f1a67244-9abc-45d5-af69-8..."
    }
   ],
   "start_time": "2014-06-24T17:51:35Z",
   "quota": "/acme",
   "fingerprint": "19:c4:3f:2d:dc:76:b1:06:e8:88:bd:7f:a3:3b:3c:93",
   "error_reason": "",
   "sshkeys": [
    "/Compute-acme/jack.jones@example.com/acme-prod-admin"
   ],
   "tags": [
    "prod2"
   ],
   "resolvers": null,
   "account": "/Compute-acme/default",
   "name": "/Compute-acme/jack.jones@example.com/dev1/f653a677-b566-4f92-8e93-71d47b364119",
   "vcable_id": "/Compute-acme/jack.jones@example.com/016e75e7-e911-42d1-bfe1-6a7f1b3f7908",
   "uri": "http://10....",
   "reverse_dns": true,
   "entry": 1,
   "boot_order": []
  }
 ]
}`)

	ipAddressAssociationDetailsRaw = []byte(`
{
 "default": 1, 
 "description": "ol 6", 
 "entries": [], 
 "uri": "https://api-z999.compute.us0.oraclecloud.com/imagelist/Compute-acme/admin/ol6",
 "name": "/Compute-acme/jack.jones@example.com/ol6"
}`)[1:]

	allIpAddressAssociationRaw = []byte(`
{
"result":[
{
  "name": "/Compute-acme/jack.jones@example.com/ipassociation1",
  "uri": "https://api-z999.compute.us0.oraclecloud.com/network/v1/ipassociation/Compute-acme/jack.jones@example.com/ipassociation1",
  "description": null,
  "tags": [],
  "vnic": "/Compute-acme/jack.jones@example.com/2e6627de-6842-49bc-9c28-21da524c297d/eth0",
  "ipAddressReservation": "/Compute-acme/jack.jones@example.com/ipreservation1"
}
]
}
`)[1:]

	ipAddressPrefixSetsRaw = []byte(`
{
 "ipAddressPrefixes": ["192.168.0.0/16"],
 "uri": "https://api-z999.compute.us0.oraclecloud.com/network/v1/ipaddressprefixset/Compute-acme/jack.jones@example.com/ipaddressprefixset1",
 "description": null,
 "tags": [],
 "name": "/Compute-acme/jack.jones@example.com/ipaddressprefixset1"
}
`)[1:]
	allIpAddressPrefixSetsRaw = []byte(`
{
  "result": [
    {
      "name": "/Compute-acme/jack.jones@example.com/ipreservation2",
      "uri": "https://api-z999.compute.us0.oraclecloud.com:443/network/v1/ipaddressprefixset/Compute-acme/jack.jones@example.com/ipreservation2",
      "description": "Sample IP address prefix set",
      "tags": [
        
      ],
      "ipAddressPrefixes": [
        "178.120.0.0/24"
      ]
    },
    {
      "name": "/Compute-acme/jack.jones@example.com/ipreservation1",
      "uri": "https://api-z999.compute.us0.oraclecloud.com:443/network/v1/ipaddressprefixset/Compute-acme/jack.jones@example.com/ipreservation1",
      "description": "updating an IP address prefix set",
      "tags": [
        
      ],
      "ipAddressPrefixes": [
        "192.168.0.0/16",
        "172.120.0.0/24"
      ]
    }
  ]
}`)[1:]
	ipAddressReservationRaw = []byte(`
	{
  "name": "/Compute-acme/jack.jones@example.com/ipreservation1",
  "uri": "https://api-z999.compute.us0.oraclecloud.com/network/v1/ipreservation/Compute-acme/jack.jones@example.com/ipreservation1",
  "description": null,
  "tags": [],
  "ipAddress": "10.252.152.90",
  "ipAddressPool": "/oracle/public/public-ippool"
}`)[1:]

	allIpAddressReservationRaw = []byte(`
{
"result":[
{
  "name": "/Compute-acme/jack.jones@example.com/ipreservation1",
  "uri": "https://api-z999.compute.us0.oraclecloud.com/network/v1/ipreservation/Compute-acme/jack.jones@example.com/ipreservation1",
  "description": null,
  "tags": [],
  "ipAddress": "10.252.152.90",
  "ipAddressPool": "/oracle/public/public-ippool"
},
{
  "name": "/Compute-acme/jack.jones@example.com/ipreservation2",
  "uri": "https://api-z999.compute.us0.oraclecloud.com/network/v1/ipreservation/Compute-acme/jack.jones@example.com/ipreservation2",
  "description": null,
  "tags": [],
  "ipAddress": "10.252.152.52",
  "ipAddressPool": "/oracle/public/cloud-ippool"
}
]
}`)[1:]
)
