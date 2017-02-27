// Copyright 2017 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package api

import (
	"errors"
	"fmt"

	"github.com/hoenirvili/go-oracle-cloud/response"
)

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
	parentpool string,
	permanent bool,
	tags []string,
) (resp response.IpReservation, err error) {

	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	if currentName == "" {
		return resp, errors.New(
			"go-oracle-cloud: Empty curret ip reservation name",
		)
	}

	if parentpool == "" {
		return resp, errors.New(
			"go-oracle-cloud: Empty pool of public IP addresses",
		)
	}

	if newName == "" {
		newName = currentName
	}

	params := struct {
		Permanent  bool     `json:"permanent"`
		Tags       []string `json:"tags,omitempty"`
		Name       string   `json:"name"`
		Parentpool string   `json:"parentpool"`
	}{
		Permanent:  permanent,
		Tags:       tags,
		Name:       newName,
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
	parentpool string,
	permanent bool,
	tags []string,
) (resp response.IpReservation, err error) {

	if currentName == "" {
		return resp, errors.New(
			"go-oracle-cloud: Empty ip reservation name provided",
		)
	}

	if newName == "" {
		newName = currentName
	}

	params := struct {
		Permanent  bool     `json:"permanent"`
		Tags       []string `json:"tags,omitempty"`
		Name       string   `json:"name"`
		Parentpool string   `json:"parentpool"`
	}{
		Permanent:  permanent,
		Tags:       tags,
		Name:       newName,
		Parentpool: parentpool,
	}

	url := fmt.Sprintf("%s/ip/reservation/Compute-%s/%s/%s",
		c.endpoint, c.identify, c.username, currentName)

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
