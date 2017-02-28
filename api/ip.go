// Copyright 2017 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package api

import (
	"errors"
	"fmt"
	"strings"

	"github.com/hoenirvili/go-oracle-cloud/response"
)

// IPPool type describing the
// parent pool of an ip association
type IPPool string

const (
	// PublicIPPool standard ip pool for the oracle provider
	PublicIPPool IPPool = "/oracle/public/ippool"
)

// validate checks if the ippool provided is empty or not
func (i IPPool) validate() (err error) {
	if i == "" {
		return errors.New("go-oracle-cloud: Empty ip pool")
	}
	return nil
}

// prefix add the ippool: prefix for different
// ip methods that needs this in order to
// construct the correct request body
func (i *IPPool) prefix() {
	prefix := "ippool:"
	if strings.HasPrefix(string(*i), prefix) {
		return
	}

	*i = IPPool("ippool:") + *i
}

// AllIps retrieves details of all the IP networks
// that are available in the specified container.
func (c Client) AllIps() (resp response.AllIps, err error) {
	if !c.isAuth() {
		return resp, errNotAuth
	}

	url := fmt.Sprintf("%s/network/v1/ipnetwork/Compute-%s/%s/",
		c.endpoint, c.identify, c.username)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		verb:   "GET",
		resp:   &resp,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}

// IpDetails retrives details of a an IP network
// that is available in the oracle account
func (c Client) IpDetails(name string) (resp response.Ip, err error) {
	if !c.isAuth() {
		return resp, errNotAuth
	}

	if name == "" {
		return resp, errors.New("go-oracle-cloud: The given ip name is empty")
	}

	url := fmt.Sprintf("%s/network/v1/ipnetwork/Compute-%s/%s/%s",
		c.endpoint, c.identify, c.username, name)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		verb:   "GET",
		resp:   &resp,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}

// CreateIp creates an IP network. An IP network allows you to define
// an IP subnet in your account. With an IP network you can isolate
// instances by creating separate IP networks and adding instances
// to specific networks. Traffic can flow between instances within
// the same IP network, but by default each network is isolated
// from other networks and from the public Internet.
func (c Client) CreateIp(
	description string,
	ipAddressPrefix string,
	ipNetworkExchange string,
	name string,
	publicNaptEnabledFlag bool,
	tags []string,
) (resp response.Ip, err error) {

	if !c.isAuth() {
		return resp, errNotAuth
	}

	if ipAddressPrefix == "" {
		return resp, errors.New("go-oracle-cloud: Empty ipAddressPrefix")
	}

	if name == "" {
		return resp, errors.New("go-oracle-cloud: Empty ip network name")
	}

	url := fmt.Sprintf("%s/network/v1/ipnetwork/", c.endpoint)

	params := response.Ip{
		Description:       description,
		IpAddressPrefix:   ipAddressPrefix,
		IpNetworkExchange: ipNetworkExchange,
		Name:              name,
		Tags:              tags,
		PublicNaptEnabledFlag: publicNaptEnabledFlag,
	}

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		verb:   "POST",
		resp:   &resp,
		body:   params,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteIp deletes an IP network with a given name
func (c Client) DeleteIp(name string) (err error) {
	if !c.isAuth() {
		return errNotAuth
	}

	if name == "" {
		return errors.New("go-oracle-cloud: Empty ip network name")
	}

	url := fmt.Sprintf("%s/network/v1/ipnetwork/Compute-%s/%s/%s",
		c.endpoint, c.identify, c.username, name)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		verb:   "DELETE",
	}); err != nil {
		return err
	}

	return nil
}

// UpdateIp can update an IP network and change the specified IP address prefix
// for the network after you've created the network and attached instances to it.
// However, when you change an IP address prefix, it could cause the IP addresses
//currently assigned to existing instances to fall outside the specified IP network.
// If this happens, all traffic to and from those vNICs will be dropped.
// If the IP address of an instance is dynamically allocated, stopping the instance
// orchestration and restarting it will reassign a valid IP address from the IP network to the instance.
// However, if the IP address of an instance is static - that is, if the IP address
// is specified in the instance orchestration while creating the instance - then
// the IP address can't be updated by stopping the instance orchestration
// and restarting it. You would have to manually update the orchestration to assign
// a valid IP address to the vNIC attached to that IP network.
// It is therefore recommended that if you update an IP network, you
// only expand the network by specifying the same IP address prefix
// but with a shorter prefix length. For example, you can expand 192.168.1.0/24
// to 192.168.1.0/20. Don't, however, change the IP address.
// This ensures that all IP addresses that have been currently allocated
// to instances remain valid in the updated IP network.
func (c Client) UpdateIp(
	currentName string,
	newName string,
	description string,
	ipNetworkExchange string,
	ipAddressPrefix string,
	publicNaptEnabledFlag bool,
	tags []string,
) (resp response.Ip, err error) {

	if !c.isAuth() {
		return resp, errNotAuth
	}

	if currentName == "" {
		return resp, errors.New("go-oracle-cloud: Empty network ip name")
	}

	if ipAddressPrefix == "" {
		return resp, errors.New("go-oracle-cloud: Empty ipAddressPrefix ")
	}

	if newName == "" {
		newName = currentName
	}

	url := fmt.Sprintf("%s/network/v1/ipnetwork/Compute-%s/%s/%s",
		c.endpoint, c.identify, c.username, currentName)

	params := response.Ip{
		Description:       description,
		IpAddressPrefix:   ipAddressPrefix,
		IpNetworkExchange: ipNetworkExchange,
		Name:              newName,
		Tags:              tags,
		PublicNaptEnabledFlag: publicNaptEnabledFlag,
	}

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		verb:   "PUT",
		resp:   &resp,
		body:   params,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}
