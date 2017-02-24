// Copyright 2017 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/hoenirvili/go-oracle-cloud/response"
)

type backupConfiguration struct {
	VolumeUri            string   `json:"volumeUri"`
	Name                 string   `json:"name"`
	Enabled              bool     `json:"enabled"`
	BackupRetentionCount uint64   `json:"backupRetentionCount"`
	Interval             interval `json:"interval"`
}

type interval struct {
	Hourly hourly `json:"Hourly"`
}

// TODO(sgiulitti) add suport for this format also
// {"DailyWeekly":{"daysOfWeek":["MONDAY"],"timeOfDay":"03:15","userTimeZone":"America/Los_Angeles"}}
type hourly struct {
	HourlyInterval uint64 `json:"hourlyInterval"`
}

// CrreateBackupConfiguration creates a new backup configuration.
// Requires authorization to create backup configurations as well
// as appropriate authorization to create snapshots from the target volume.
func (c Client) CreateBackupConfiguration(
	volumeName, backUpName string,
	enabled bool,
	backupRetentionCount, hourlyInterval uint64,
) (resp response.BackupConfiguration, err error) {

	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	url := fmt.Sprintf("%s/backupservice/v1/configuration/", c.endpoint)

	if backUpName == "" {
		return resp, errors.New("go-oracle-cloud: Cannot create backup configuration with no name")
	}

	backupConfiguration := backupConfiguration{
		VolumeUri: fmt.Sprintf("%s/storage/volume/Compute-%s/%s/%s",
			c.endpoint, c.identify, c.username, volumeName),

		Name:                 backUpName,
		Enabled:              enabled,
		BackupRetentionCount: backupRetentionCount,
		Interval: interval{
			Hourly: hourly{
				HourlyInterval: hourlyInterval,
			},
		},
	}

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		verb:   "POST",
		body:   &backupConfiguration,
		treat:  defaultPostTreat,
		resp:   &resp,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteBackupConfiguration deletes a backup configuration.
// In order to delete the configuration all backups and restores
// related to the configuration must already be deleted.
// If disabling a backup configuration is desired, consider setting enabled to false.
func (c Client) DeleteBackupConfiguration(name string) (err error) {
	if !c.isAuth() {
		return ErrNotAuth
	}

	if name == "" {
		return errors.New("go-oracle-cloud: Cannot delete a backup configuration with empty name")
	}

	url := fmt.Sprintf("%s/backupservice/v1/configuration/Compute-%s/%s/%s",
		c.endpoint, c.identify, c.username, name)
	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		verb:   "DELETE",
		treat: func(resp *http.Response) (err error) {
			if resp.StatusCode != http.StatusNoContent {
				switch resp.StatusCode {
				case http.StatusUnauthorized:
					return fmt.Errorf(
						"go-oracle-cloud: Cannot delete backup because the account does not have authorisation for doing this",
					)

				case http.StatusNotFound:
					return fmt.Errorf(
						"go-oracle-cloud: The URL does not refer to a valid resource",
					)

				case http.StatusConflict:
					return fmt.Errorf(
						"go-oracle-cloud: The backup configuration cannot be deleted due to associated backups or restores.",
					)
				case http.StatusInternalServerError:
					return fmt.Errorf(
						"go-oracle-cloud: The server encountered an error handling this request.",
					)
				default:
					return fmt.Errorf(
						"go-oracle-cloud: Error api response %d %s",
						resp.StatusCode, dumpApiError(resp.Body),
					)
				}
			}
			return nil
		},
	}); err != nil {
		return err
	}

	return nil
}

// BackupConfigurationDetails retrieves details of the specified
// backup configuration. You can use this request to verify whether
// the CreateBackupConfiguration and UpdateBackupConfiguration
// requests were completed successfully.
func (c Client) BackupConfigurationDetails(name string) (resp response.BackupConfiguration, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	if name == "" {
		return resp, errors.New("go-oracle-cloud: Cannot return details of a empty backup configuration name")
	}

	url := fmt.Sprintf("%s/backupservice/v1/configuration/Compute-%s/%s/%s",
		c.endpoint, c.identify, c.username, name,
	)

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

// AllBackupConfiguration retrieves details for all backup
// configuration objects the current user has permission to access
func (c Client) AllBackupConfiguration() (resp []response.BackupConfiguration, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	url := fmt.Sprintf("%s/backupservice/v1/configuration/", c.endpoint)

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

// UpdateBackupConfiguration
// Modify an existing backup configuration.
// All fields, including unmodifiable fields, must be provided
// for this operation. The following fields are unmodifiable:
// volumeName, runAsUser, name.
func (c Client) UpdateBackupConfiguration(
	backUpName, volumeName, description string,
	enabled bool,
	hourlyInterval, backupRetentionCount uint64,
) (resp response.BackupConfiguration, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	url := fmt.Sprintf("%s/backupservice/v1/configuration/Compute-%s/%s/%s",
		c.endpoint, c.identify, c.username, backUpName)

	backupConfiguration := struct {
		RunAsUser            string   `json:"runAsUser"`
		VolumeUri            string   `json:"volumeUri"`
		Name                 string   `json:"name"`
		Enabled              bool     `json:"enabled"`
		BackupRetentionCount uint64   `json:"backupRetentionCount"`
		Interval             interval `json:"interval"`
		Description          string   `json:"description,omitempty"`
	}{
		RunAsUser: fmt.Sprintf("/Compute-%s/%s", c.identify, c.username),
		VolumeUri: fmt.Sprintf("%s/storage/volume/Compute-%s/%s/%s",
			c.endpoint, c.identify, c.username, volumeName),
		Name: fmt.Sprintf("/Compute-%s/%s/%s",
			c.identify, c.username, backUpName),

		Enabled:              enabled,
		BackupRetentionCount: backupRetentionCount,
		Interval: interval{
			Hourly: hourly{
				HourlyInterval: hourlyInterval,
			},
		},
		Description: description,
	}

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		verb:   "PUT",
		body:   &backupConfiguration,
		treat:  defaultPostTreat,
		resp:   &resp,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}
