// Copyright 2017 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/hoenirvili/go-oracle-cloud/response"
)

// BackupConfigurationParams type used to feed up
// the CreateBackupConfiguration function with params.
type BackupConfigurationParams struct {

	// Description of this Backup Configuration
	Description string `json:"description,omitempty"`

	// BackupRetentionCount represents how many backups to retain
	// Minimum Value: 1
	BackupRetentionCount uint32 `json:"backupRetentionCount"`

	// Enabled when true, backups will automatically
	// be generated based on the interval.
	Enabled bool `json:"enabled"`

	// Name is the name of the backup configuration
	Name string `json:"name"`

	// VolumeUri the complete URI of the storage volume
	// that you want to backup.
	VolumeUri string `json:"volumeUri"`

	// Interval represents the interval in the backup configuration.
	// There are two kinds of Intervals. Each Interval has its own JSON format.
	// Your Interval field should look like one of the following:
	//
	// "interval":{
	//   "Hourly":{
	//     "hourlyInterval":2
	//	 }
	// }
	//
	//
	// {"DailyWeekly":
	// 	{
	//	  "daysOfWeek":["MONDAY"],
	//	  "timeOfDay":"03:15",
	// 	  "userTimeZone":"America/Los_Angeles"
	//  }
	// }
	// Days of the week is any day of the week
	// fully capitalized (MONDAY, TUESDAY, etc).
	// The user time zone is any IANA user timezone.
	//For example user time zones see List of IANA time zones.
	Interval interface{} `json:"interval"`
}

// CreateBackupConfiguration creates a new backup configuration.
// Requires authorization to create backup configurations as well
// as appropriate authorization to create snapshots from the target volume.
func (c Client) CreateBackupConfiguration(
	p BackupConfigurationParams,
) (resp response.BackupConfiguration, err error) {

	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	url := fmt.Sprintf(
		"%s/backupservice/v1/configuration/", c.endpoint)

	if p.Name == "" {
		return resp, errors.New(
			"go-oracle-cloud: Empty backup configuration name",
		)
	}

	p.Name = fmt.Sprintf("/Compute-%s/%s/%s",
		c.identify, c.username, p.Name)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		verb:   "POST",
		body:   &p,
		resp:   &resp,
		treat: func(resp *http.Response) (err error) {
			switch resp.StatusCode {
			case http.StatusCreated:
				return nil
			case http.StatusBadRequest:
				return errors.New(
					"go-oracle-cloud: Invalid backup configuration input. Volume does not exist or is not online",
				)
			case http.StatusUnauthorized:
				return errors.New("go-oracle-cloud: Unauthorized")
			case http.StatusInternalServerError:
				return errors.New(
					"go-oracle-cloud: The server encountered an error handling this request",
				)
			default:
				return fmt.Errorf(
					"go-oracle-cloud: Error api response %d %s",
					resp.StatusCode, dumpApiError(resp.Body),
				)
			}
		},
	}); err != nil {
		return resp, err
	}

	strip(&resp.Name)
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
		return errors.New(
			"go-oracle-cloud: Empty backup configuration name",
		)
	}

	url := fmt.Sprintf("%s/backupservice/v1/configuration/Compute-%s/%s/%s",
		c.endpoint, c.identify, c.username, name)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		verb:   "DELETE",
		treat: func(resp *http.Response) (err error) {
			switch resp.StatusCode {
			case http.StatusNoContent:
				return nil
			case http.StatusUnauthorized:
				return errors.New(
					"go-oracle-cloud: Cannot delete backup because the account does not have authorisation for doing this",
				)

			case http.StatusNotFound:
				return errors.New(
					"go-oracle-cloud: The URL does not refer to a valid resource",
				)

			case http.StatusConflict:
				return errors.New(
					"go-oracle-cloud: The backup configuration cannot be deleted due to associated backups or restores",
				)
			case http.StatusInternalServerError:
				return errors.New(
					"go-oracle-cloud: The server encountered an error handling this request",
				)
			default:
				return fmt.Errorf(
					"go-oracle-cloud: Error api response %d %s",
					resp.StatusCode, dumpApiError(resp.Body),
				)
			}
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
func (c Client) BackupConfigurationDetails(
	name string,
) (resp response.BackupConfiguration, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	if name == "" {
		return resp, errors.New(
			"go-oracle-cloud: Empty backup configuration name",
		)
	}

	url := fmt.Sprintf("%s/backupservice/v1/configuration/Compute-%s/%s/%s",
		c.endpoint, c.identify, c.username, name,
	)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		verb:   "GET",
		treat: func(resp *http.Response) (err error) {
			switch resp.StatusCode {
			case http.StatusOK:
				return nil
			case http.StatusUnauthorized:
				return errors.New(
					"go-oracle-cloud: Cannot delete backup because the account does not have authorisation for doing this",
				)

			case http.StatusNotFound:
				return errors.New(
					"go-oracle-cloud: The URL does not refer to a valid resource",
				)

			case http.StatusInternalServerError:
				return errors.New(
					"go-oracle-cloud: The server encountered an error handling this request",
				)
			default:
				return fmt.Errorf(
					"go-oracle-cloud: Error api response %d %s",
					resp.StatusCode, dumpApiError(resp.Body),
				)
			}
		},
		resp: &resp,
	}); err != nil {
		return resp, err
	}

	strip(&resp.Name)

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
		treat: func(resp *http.Response) (err error) {
			switch resp.StatusCode {
			case http.StatusOK:
				return nil
			case http.StatusUnauthorized:
				return errors.New(
					"go-oracle-cloud: Cannot delete backup because the account does not have authorisation for doing this",
				)
			case http.StatusInternalServerError:
				return errors.New(
					"go-oracle-cloud: The server encountered an error handling this request",
				)
			default:
				return fmt.Errorf(
					"go-oracle-cloud: Error api response %d %s",
					resp.StatusCode, dumpApiError(resp.Body),
				)
			}
		},

		resp: &resp,
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
	p BackupConfigurationParams,
	newName string,
) (resp response.BackupConfiguration, err error) {
	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	if p.Name == "" {
		return resp, errors.New(
			"go-oracle-cloud: Empty backup configuration name",
		)
	}

	if newName == "" {
		newName = p.Name
	}

	url := fmt.Sprintf("%s/backupservice/v1/configuration/Compute-%s/%s/%s",
		c.endpoint, c.identify, c.username, p.Name)

	p.Name = fmt.Sprintf("/Compute-%s/%s/%s",
		c.identify, c.username, newName)

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		verb:   "PUT",
		body:   &p,
		treat: func(resp *http.Response) (err error) {
			switch resp.StatusCode {
			case http.StatusOK:
				return nil
			case http.StatusUnauthorized:
				return errors.New(
					"go-oracle-cloud: Cannot delete backup because the account does not have authorisation for doing this",
				)
			case http.StatusInternalServerError:
				return errors.New(
					"go-oracle-cloud: The server encountered an error handling this request",
				)
			default:
				return fmt.Errorf(
					"go-oracle-cloud: Error api response %d %s",
					resp.StatusCode, dumpApiError(resp.Body),
				)
			}
		},
		resp: &resp,
	}); err != nil {
		return resp, err
	}

	strip(&resp.Name)
	return resp, nil
}
