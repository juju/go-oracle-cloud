// Copyright 2017 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

var endpoints = map[string]string{
	"account":               "%s/account",
	"acl":                   "%s/network/v1/acl",
	"authenticate":          "%s/authenticate/",
	"backupconfiguration":   "%s/backupservice/v1/configuration",
	"backups":               "%s/backupservice/v1/backup",
	"imagelistentries":      "%s/imagelist/%s/entry",
	"imagelists":            "%s/imagelist",
	"instanceconsole":       "%s/instanceconsole",
	"instance":              "%s/instance",
	"ipaddressassociation":  "%s/network/v1/ipassociation",
	"ipaddressprefixsets":   "%s/network/v1/ipaddressprefixset",
	"ipaddressreservation":  "%s/network/v1/ipreservation",
	"ipassociation":         "%s/ip/association",
	"ipnetworkexchanges":    "%s/network/v1/ipnetworkexchange",
	"ipnetworks":            "%s/network/v1/ipnetwork",
	"ipreservation":         "%s/ip/reservation",
	"launchplan":            "%s/launchplan",
	"machineimages":         "%s/machineimage",
	"orchestrations":        "%s/orchestration",
	"osscontainers":         "%s/integrations/osscontainer",
	"rebootinstancerequest": "%s/rebootinstancerequest",
	"refreshtoken":          "%s/refresh",
	"restores":              "%s/backupservice/v1/restore",
	"routes":                "%s/network/v1/route",
	"secapplications":       "%s/secapplication",
	"secassociations":       "%s/secassociation",
	"seciplists":            "%s/seciplist",
	"seclists":              "%s/seclist",
	"secrules":              "%s/secrule",
	"securityprotocols":     "%s/network/v1/secprotocol",
	"securityrules":         "%s/network/v1/secrule",
	"shapes":                "%s/shape",
	"snapshots":             "%s/snapshot",
	"sshkeys":               "%s/sshkey",
	"storageattachments":    "%s/storage/attachment",
	"storageproperties":     "%s/property/storage",
	"storagesnapshots":      "%s/storage/snapshot",
	"storagevolumes":        "%s/storage/volume",
	"virtualnics":           "%s/network/v1/vnic",
	"virtualnicsets":        "%s/network/v1/vnicset",
	"vpnendpoints":          "%s/vpnendpoint",
}

// treatStatus will be used as a callback to custom check the response
// if the client decides the response contains some error codes
// it will make the Request return that error
type treatStatus func(resp *http.Response) error

func debugTreat(resp *http.Response) (err error) {
	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(raw))
	os.Exit(1)
	return nil
}

func defaultTreat(resp *http.Response) (err error) {
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf(
			"go-oracle-cloud: Error api response %d %s",
			resp.StatusCode, dumpApiError(resp.Body),
		)
	}
	return nil
}

func defaultPostTreat(resp *http.Response) (err error) {
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf(
			"go-oracle-cloud: Error api response %d %s",
			resp.StatusCode, dumpApiError(resp.Body),
		)
	}
	return nil
}

func defaultDeleteTreat(resp *http.Response) (err error) {
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf(
			"go-oracle-cloud: Error api response %d %s",
			resp.StatusCode, dumpApiError(resp.Body),
		)
	}
	return nil
}

// paramsRequest used to fill up the params for the request function
type paramsRequest struct {
	// directory is the type of directory request
	directory bool
	// use this client to do the request
	client *http.Client
	// include in the request the cookie
	cookie *http.Cookie
	// url this will not be checked so, make sure it is valid
	url string
	// verb contains an http verb POST, GET, PUT, PATH
	verb string
	// body will contain the http body request if any
	// for Get request this should be leaved nil
	body interface{}
	// treat will be used as a callback to
	// check the response and save extract things
	treat treatStatus
	// resp will contains a json object where the
	// request function will decode
	resp interface{}
}

// request function is a wrapper around building the request,
// treating exceptional errors and executing the client http connection
func request(cfg paramsRequest) (err error) {
	var buf io.Reader

	// if we have a body we assume that the body
	// should be json encoded and ready to
	// be appendend into the request
	if cfg.body != nil {
		raw, err := json.Marshal(cfg.body)
		if err != nil {
			return err
		}
		buf = bytes.NewBuffer(raw)
	}

	req, err := http.NewRequest(cfg.verb, cfg.url, buf)
	if err != nil {
		return err
	}

	// add the session cookie if there is one
	if cfg.cookie != nil {
		req.AddCookie(cfg.cookie)
	}

	// let the endpoint api know that the request
	// was made from the go oracle cloud client
	req.Header.Add("User-Agent", "go-oracle-cloud v1.0")
	// the oracle api supports listing so in order to use it
	// we should include the +directory header
	if cfg.directory {
		req.Header.Add("Accept", "application/oracle-compute-v3+directory+json")
	} else {
		req.Header.Add("Accept", "application/oracle-compute-v3+json")
	}
	// every request should let know that we accept encoded gzip responses
	req.Header.Add("Accept-Encoding", "gzip;q=1.0, identity; q=0.5")

	switch cfg.verb {
	case "POST", "PUT":
		req.Header.Add("Content-Encoding", "deflate")
		req.Header.Add("Content-Type", "application/oracle-compute-v3+json")
	case "DELETE":
		req.Header.Add("Content-Type", "application/oracle-compute-v3+json")
	case "GET":
	}

	resp, err := cfg.client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		if errClose := resp.Body.Close(); errClose != nil {
			// overwrite the previous error if any
			err = errClose
		}
	}()

	// if we have a special treat function
	// let the caller treat the response
	if cfg.treat != nil {
		if err = cfg.treat(resp); err != nil {
			return err
		}
	}

	// if we the caller tells us that the http request
	// returns an response and wants to decode the response
	// that is json format.
	if cfg.resp != nil {
		if err = json.NewDecoder(resp.Body).Decode(cfg.resp); err != nil {
			return err
		}
	}

	return nil
}
