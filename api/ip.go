package api

import (
	"errors"
	"fmt"

	"github.com/hoenirvili/go-oracle-cloud/response"
)

// Retrieves details of the specified IP address association.
func (c Client) IpAssociationsDetails(ip string) (resp response.IpAssociation, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	if ip == "" {
		return resp, errors.New("go-oracle-cloud: The given ip is empty")
	}

	url := fmt.Sprintf("%s/network/v1/ipassociation/Compute-%s/%s/%s",
		c.endpoint, c.identify, c.username, ip)

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

// AllIpAssociation Retrieves details of the specified IP address association.
func (c Client) AllIpAssociation() (resp response.AllIpAssociation, err error) {
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

// CreateIpAssociation creates an IP address association
// to associate an IP address reservation, a public IP address,
// with a vNIC of an instance either while creating the instance
// or when an instance is already running.
func (c Client) CreateIpAssociation(
	ipAddressReservation string,
	vnic string,
	name string,
) (resp response.IpAssociation, err error) {

	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	// construct the body for the post request
	params := response.IpAssociation{
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

// DeleteIpAssociation deletes the specified IP address association.
// Ensure that the IP address association is not being used before deleting it.
func (c Client) DeleteIpAssociation(name string) (err error) {

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

// UpdateIpAssociation updates the specified IP address association.
// You can update values for the following parameters
// of an IP address association: description, ipAddressReservation,
// vnic, and tags. If you associate an IP reservation with a vNIC while
// creating or updating the IP address association, then you can remove
// the association between the IP address reservation and vNIC by updating the IP address association.
// However, if you associate an IP reservation with an instance while creating the instance,
// then to remove the IP reservation, update the instance orchestration.
// Otherwise, whenever your instance orchestration is stopped and restarted,
// the IP reservation will again be associated with the vNIC.
func (c Client) UpdateIpAssociation(
	currentName,
	ipAddressReservation,
	vnic,
	newName string,
) (resp response.IpAssociation, err error) {

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
	params := response.IpAssociation{
		IpAddressReservation: fmt.Sprintf("/Compute-%s/%s/%s",
			c.identify, c.username, ipAddressReservation),
		Vnic: fmt.Sprintf("/Compute-%s/%s/%s",
			c.identify, c.username, vnic),
		Name: fmt.Sprintf("/Compute-%s/%s/%s",
			c.identify, c.username, newName),
	}

	url := fmt.Sprintf("%s/network/v1/ipassociation/Compute-%s/%s/%s", c.endpoint, c.identify, c.username, currentName)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		verb:   "PUT",
		treat:  defaultPostTreat,
		resp:   &resp,
		body:   params,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}
