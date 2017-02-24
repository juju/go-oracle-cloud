// Copyright 2017 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.
package api

import (
	"errors"
	"fmt"

	"github.com/hoenirvili/go-oracle-cloud/response"
)

// IpAddressAssociation retrives details of the specified IP address association.
func (c Client) IpAddressAssociationDetails(name string) (resp response.IpAddressAssociation, err error) {

	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	if name == "" {
		return resp, errors.New("go-oracle-cloud: The given ip name is empty")
	}

	url := fmt.Sprintf("%s/network/v1/ipassociation/Compute-%s/%s/%s",
		c.endpoint, c.identify, c.username, name)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		verb:   "GET",
		treat:  defaultTreat,
		resp:   &resp,
	}); err != nil {
		return resp, err
	}

	return resp, nil

}

// AllIpAddressAssociation Retrieves details of the specified IP address association.
func (c Client) AllIpAddressAssociation() (resp response.AllIpAddressAssociation, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	url := fmt.Sprintf("%s/network/v1/ipassociation/Compute-%s/%s/",
		c.endpoint, c.identify, c.username)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		verb:   "GET",
		treat:  defaultTreat,
		resp:   &resp,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}

// CreateIpAddressAssociation creates an IP address association
// to associate an IP address reservation, a public IP address,
// with a vNIC of an instance either while creating the instance
// or when an instance is already running.
func (c Client) CreateIpAddressAssociation(
	ipAddressReservation string,
	vnic string,
	name string,
) (resp response.IpAddressAssociation, err error) {

	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	if ipAddressReservation == "" {
		return resp, errors.New("go-oracle-cloud: Empty ip address reservation")
	}

	if name == "" {
		return resp, errors.New("go-oracle-cloud: Empty ip name")
	}

	if vnic == "" {
		return resp, errors.New("go-oracle-cloud: Empty vnic name")
	}

	// construct the body for the post request
	params := response.IpAddressAssociation{
		IpAddressReservation: fmt.Sprintf("/Compute-%s/%s/%s",
			c.identify, c.username, ipAddressReservation),
		Vnic: fmt.Sprintf("/Compute-%s/%s/%s",
			c.identify, c.username, vnic),
		Name: fmt.Sprintf("/Compute-%s/%s/%s",
			c.identify, c.username, name),
	}

	url := fmt.Sprintf("%s/network/v1/ipassociation/", c.endpoint)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		verb:   "POST",
		treat:  defaultPostTreat,
		resp:   &resp,
		body:   params,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteIpAddressAssociation deletes the specified IP address association.
// Ensure that the IP address association is not being used before deleting it.
func (c Client) DeleteIpAddressAssociation(name string) (err error) {

	if !c.isAuth() {
		return ErrNotAuth
	}

	if name == "" {
		return errors.New("go-oracle-cloud: The given name is empty")
	}

	url := fmt.Sprintf("%s/network/v1/ipassociation/Compute-%s/%s/%s",
		c.endpoint, c.identify, c.username, name)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		verb:   "DELETE",
		treat:  defaultDeleteTreat,
	}); err != nil {
		return err
	}

	return nil
}

// UpdateIpAddressAssociation updates the specified IP address association.
// You can update values for the following parameters
// of an IP address association: description, ipAddressReservation,
// vnic, and tags. If you associate an IP reservation with a vNIC while
// creating or updating the IP address association, then you can remove
// the association between the IP address reservation and vNIC by updating the IP address association.
// However, if you associate an IP reservation with an instance while creating the instance,
// then to remove the IP reservation, update the instance orchestration.
// Otherwise, whenever your instance orchestration is stopped and restarted,
// the IP reservation will again be associated with the vNIC.
func (c Client) UpdateIpAddressAssociation(
	currentName,
	ipAddressReservation,
	vnic,
	newName string,
) (resp response.IpAddressAssociation, err error) {

	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	if currentName == "" {
		return resp, errors.New("go-oracle-cloud: Given name is empty")
	}

	if ipAddressReservation == "" {
		return resp, errors.New("go-oracle-cloud: Given ip address reservation is empty")
	}

	if vnic == "" {
		return resp, errors.New("go-oracle-cloud: Given vnic is empty")
	}

	if newName == "" {
		newName = currentName
	}

	// construct the body for the post request
	params := response.IpAddressAssociation{
		IpAddressReservation: fmt.Sprintf("/Compute-%s/%s/%s",
			c.identify, c.username, ipAddressReservation),

		Vnic: fmt.Sprintf("/Compute-%s/%s/%s",
			c.identify, c.username, vnic),

		Name: fmt.Sprintf("/Compute-%s/%s/%s",
			c.identify, c.username, newName),
	}

	url := fmt.Sprintf("%s/network/v1/ipassociation/Compute-%s/%s/%s",
		c.endpoint, c.identify, c.username, currentName)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		verb:   "PUT",
		treat:  defaultTreat,
		resp:   &resp,
		body:   params,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}
