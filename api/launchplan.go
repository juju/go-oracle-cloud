package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/hoenirvili/go-oracle-cloud/response"
)

// InstanceParams used to feed the CreateInstance function
type InstanceParams struct {
	// Shape represents every instance in oracle cloud has a predefined shape
	// in order to create a virtual instance
	// here it can be specify the computing power shape
	Shape string
	// Imagelist is the virtual image that will be used
	// in order to init the instance
	Imagelist string
	// Name is the name of the instance
	Name string
	// Label
	Label string
	// SSHKeys that will be installed on the Instance
	SSHKeys []string
}

func (c Client) CreateInstance(params []InstanceParams) (resp response.LaunchPlan, err error) {
	if params == nil || len(params) == 0 {
		return resp, errors.New("go-oracle-cloud: Empty slice of instance parameters")
	}

	if !c.isAuth() {
		return resp, ErrNotAuth
	}

	// no need to make this type public inside this package
	type args struct {
		Shape     string   `json:"shape"`
		Imagelist string   `json:"imagelist"`
		Name      string   `json:"name"`
		Label     string   `json:"label"`
		SSHKeys   []string `json:"sshkeys"`
	}

	n := len(params)
	instanceArgs := make([]args, n, n)

	// here we are constructing the post body json
	for i := 0; i < n; i++ {
		// add the imagelist
		if params[i].Imagelist == "" {
			return resp, errors.New(
				"go-oracle-cloud: Empty image list in instance parameters",
			)
		}
		instanceArgs[i].Imagelist = fmt.Sprintf(
			"/Compute-%s/%s/%s", c.identify, c.username, params[i].Imagelist,
		)

		// add the label
		if params[i].Label == "" {
			return resp, errors.New(
				"go-oracle-cloud: Empty label in instance parameters",
			)
		}
		instanceArgs[i].Label = params[i].Label

		// make the name oracle cloud complaint
		instanceArgs[i].Name = fmt.Sprintf("/Compute-%s/%s/%s",
			c.identify, c.username, params[i].Name)

		// add the shape
		instanceArgs[i].Shape = params[i].Shape

		// add the ssh keys
		keys := len(params[i].SSHKeys)
		instanceArgs[i].SSHKeys = make([]string, keys, keys)
		for j := 0; j < keys; j++ {
			instanceArgs[i].SSHKeys[j] = fmt.Sprintf("/Compute-%s/%s/%s",
				c.identify, c.username, params[i].SSHKeys[j],
			)
		}
	}

	url := fmt.Sprintf("%s/launchplan/", c.endpoint)
	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		verb:   "POST",
		body:   &instanceArgs,
		treat: func(resp *http.Response) (err error) {
			type Instances struct {
				Instances string `json:"instances,omitempty"`
			}
			type m struct {
				Message Instances `json:"message,omitempty"`
			}
			var errOut m
			if resp.StatusCode != http.StatusCreated {
				json.NewDecoder(resp.Body).Decode(&errOut)
				return fmt.Errorf(
					"go-oracle-cloud: Error Api response %d %s",
					resp.StatusCode, errOut.Message.Instances,
				)
			}

			return nil
		},
		resp: &resp,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}
