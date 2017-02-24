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

// AllIp retrieves details of all the IP networks
// that are available in the specified container.
func (c Client) AllIp() (resp response.AllIp, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	url := fmt.Sprintf("%s/network/v1/ipnetwork/Compute-%s/%s/",
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

func (c Client) IpDetails(name string) (resp response.Ip, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
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
		treat:  defaultTreat,
		resp:   &resp,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}

func (c Client) CreateIp(
	description string,
	ipAddressPrefix string,
	ipNetworkExchange string,
	name string,
	publicNaptEnabledFlag bool,
	tags []string,
) (resp response.Ip, err error) {

	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	if ipAddressPrefix == "" {
		return resp, errors.New("go-oracle-cloud: Empty ipAddressPrefix")
	}

	if name == "" {
		return resp, errors.New("go-oracle-cloud: Empty name")
	}

	url := fmt.Sprintf("%s/network/v1/ipnetwork/", c.endpoint)

	params := response.Ip{
		Description:     description,
		IpAddressPrefix: ipAddressPrefix,
		IpNetworkExchange: fmt.Sprintf("/Compute-%s/%s/%s",
			c.identify, c.username, ipNetworkExchange),
		Name: fmt.Sprintf("/Compute-%s/%s/%s",
			c.identify, c.username, name),
		Tags: tags,
		PublicNaptEnabledFlag: publicNaptEnabledFlag,
	}

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

func (c Client) DeleteIp(name string) (err error) {
	if !c.isAuth() {
		return ErrNotAuth
	}

	if name == "" {
		return errors.New("go-oracle-cloud: The given ip name is empty")
	}

	url := fmt.Sprintf("%s/network/v1/ipnetwork/Compute-%s/%s/%s",
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
	description string,
	ipNetworkExchange string,
	ipAddressPrefix string,
	newName string,
	publicNaptEnabledFlag bool,
	tags []string,
) (resp response.Ip, err error) {

	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	if currentName == "" {
		return resp, errors.New("go-oracle-cloud: Given name is empty")
	}

	if ipAddressPrefix == "" {
		return resp, errors.New("go-oracle-cloud: Given ip address prefix is empty")
	}

	if newName == "" {
		newName = currentName
	}

	url := fmt.Sprintf("%s/network/v1/ipnetwork/Compute-%s/%s/%s",
		c.endpoint, c.identify, c.username, currentName)

	params := response.Ip{
		Description:     description,
		IpAddressPrefix: ipAddressPrefix,
		IpNetworkExchange: fmt.Sprintf("/Compute-%s/%s/%s",
			c.identify, c.username, ipNetworkExchange),

		Name: fmt.Sprintf("/Compute-%s/%s/%s",
			c.identify, c.username, newName),
		Tags: tags,
		PublicNaptEnabledFlag: publicNaptEnabledFlag,
	}

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

// AllIpReservations Retrieves details of the IP reservations that are available
func (c Client) AllIpReservation() (resp response.AllIpReservation, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	url := fmt.Sprintf("%s/ip/reservation/Compute-%s/%s/",
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

// IpReservationDetails retrieves details of an IP reservation.
// You can use this request to verify whether the
// CreateIpReservation or PutIpReservatio were completed successfully.
func (c Client) IpReservationDetails(name string) (resp response.IpReservation, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	if name == "" {
		return resp, errors.New("go-oracle-cloud: Empty name provided")
	}

	url := fmt.Sprintf("%s/ip/reservation/Compute-%s/%s/%s",
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

// CreateIpReservation creates an IP reservation.
// After creating an IP reservation, you can associate it with
// an instance by using the CrateIpAddressAssociation method
func (c Client) CreateIpReservation(
	currentName string,
	newName string,
	parentpool bool,
	permanent bool,
	tags []string,
) (resp response.IpReservation, err error) {

	if newName == "" {
		newName = currentName
	}

	params := struct {
		Permanent  bool     `json:"permanent"`
		Tags       []string `json:"tags,omitempty"`
		Name       string   `json:"name"`
		Parentpool bool     `json:"parentpool"`
	}{
		Permanent: permanent,
		Tags:      tags,
		Name: fmt.Sprintf("/Compute-%s/%s/%s",
			c.identify, c.username, newName),
		Parentpool: parentpool,
	}

	url := fmt.Sprintf("%s/ip/reservation/", c.endpoint)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		body:   &params,
		verb:   "POST",
		treat:  defaultPostTreat,
		resp:   &resp,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteIpReservation deletes the ip reservation of a instance.
// When you no longer need an IP reservation, you can delete it.
// Ensure that no instance is using the IP reservation that you want to delete.
func (c Client) DeleteIpReservation(name string) (err error) {
	if !c.isAuth() {
		return ErrNotAuth
	}

	if name == "" {
		return errors.New("go-oracle-cloud: Empty name provided")
	}

	url := fmt.Sprintf("%s/ip/reservation/Compute-%s/%s/%s",
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

// UpdateIpreservation changes the permanent field of an IP reservation
// from false to true or vice versa.
// You can use this command when, for example, you want to delete an
// instance but retain its autoallocated public IP address as a permanent IP
// reservation for use later with another instance. In such a case, before
// deleting the instance, change the permanent field of the IP reservation
// from false to true.
// Note that if you change the permanent field of an IP reservation tofalse,
// and if the reservation is not associated with an instance, then
// the reservation will be deleted.
// You can also update the tags that are used to identify the IP reservation.
func (c Client) UpdateIpReservation(
	currentName string,
	newName string,
	parentpool bool,
	permanent bool,
	tags []string,
) (resp response.IpReservation, err error) {

	if currentName == "" {
		return resp, errors.New("go-oracle-cloud: Empty name provided")
	}

	if newName == "" {
		newName = currentName
	}

	params := struct {
		Permanent  bool     `json:"permanent"`
		Tags       []string `json:"tags,omitempty"`
		Name       string   `json:"name"`
		Parentpool bool     `json:"parentpool"`
	}{
		Permanent: permanent,
		Tags:      tags,
		Name: fmt.Sprintf("/Compute-%s/%s/%s",
			c.identify, c.username, newName),
		Parentpool: parentpool,
	}

	url := fmt.Sprintf("%s/ip/reservation/Compute-%s/%s/%s",
		c.endpoint, c.identify, c.username, newName)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		body:   &params,
		url:    url,
		verb:   "PUT",
		treat:  defaultTreat,
		resp:   &resp,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}

// IpAssociation retrieves the names of objects and subcontainers
// that you can access in the specified container.
func (c Client) AllIpAssociation() (resp response.AllIpAssociation, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	url := fmt.Sprintf("%s/ip/association/Compute-%s/%s/",
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

// IpAssociationDetails retrieves details of the IP associations that are
// available in the specified container
func (c Client) IpAssociationDetails(name string) (resp response.IpAssociation, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	if name == "" {
		return resp, errors.New("go-oracle-cloud: Empty ip association name")
	}

	url := fmt.Sprintf("%s/ip/association/Compute-%s/%s/%s",
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

// Creates an association between an IP address
// and the vcable ID of an instance.
func (c Client) CreateIpAssociation(
	parentpool string,
	vcable string,
) (resp response.IpAssociation, err error) {

	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	params := struct {
		Parentpool string `json:"parentpool"`
		Vcable     string `json:"vcable"`
	}{
		Parentpool: parentpool,
		Vcable: fmt.Sprintf("Compute-%s/%s/%s",
			c.identify, c.username, vcable),
	}

	url := fmt.Sprintf("%s/ip/ipassociation/", c.endpoint)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		body:   &params,
		url:    url,
		verb:   "POST",
		treat:  defaultPostTreat,
		resp:   &resp,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}

// Deletes the specified IP association
func (c Client) DeleteIpAssociation(name string) (err error) {
	if !c.isAuth() {
		return ErrNotAuth
	}

	if name == "" {
		return errors.New("go-oracle-cloud: Empty ip association name provided")
	}

	url := fmt.Sprintf("%s/ip/ipassociation/%s/%s/%s",
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
