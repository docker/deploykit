package compute

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"

	"github.com/Sirupsen/logrus"
	"github.com/google/go-querystring/query"
)

// InstanceClient is a client for the Instance functions of the Compute API.
type InstanceClient struct {
	client        *APIClient
	compartmendID string
}

// Instance contains the instance reference from:
// https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Instance/
type Instance struct {
	// The Availability Domain the instance is running in
	AvailabilityDomain string `json:"availabilityDomain"`
	// The OCID of the compartment that contains the instance
	CompartmentID string `json:"compartmentId"`
	// A user-friendly name
	DisplayName string `json:"displayName"`
	// Additional metadata key/value pairs that you provide
	ExtendedMetadata *json.RawMessage
	// The OCID of the instance
	ID string `json:"id"`
	// The image used to boot the instance
	ImageID string `json:"imageId"`
	// iPXE script to continue the boot process.
	IpxeScript string `json:"ipxeScript"`
	// The current state of the instance.
	// PROVISIONING | RUNNING | STARTING |
	// STOPPING | STOPPED | CREATING_IMAGE | TERMINATING | TERMINATED
	LifeCycleState string `json:"lifecycleState"`
	// Custom metadata that you provide
	Metadata struct {
		PublicKey string `json:"ssh_authorized_keys"`
		UserData  string `json:"user_data"`
	} `json:"metadata"`
	// The region that contains the Availability Domain the instance is running in
	Region string `json:"region"`
	// The shape of the instance
	Shape string `json:"shape"`
	// The date and time the instance was created (RFC3339)
	TimeCreated string `json:"timeCreated"`
}

// InstancesParameters are optional parameters when listing instances
type InstancesParameters struct {
	AvailabilityDomain string `url:"availabilityDomain,omitempty"` //The name of the Availability Domain.
	DisplayName        string `url:"displayName,omitempty"`        //A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	Limit              int    `url:"limit,omitempty"`              //The maximum number of items to return in a paginated "List" call.
	Page               string `url:"page,omitempty"`               //The value of the opc-next-page response header from the previous "List" call
	Filter             *InstanceFilter
}

// NewInstanceClient provides a client interface for instance API calls
func (c *APIClient) NewInstanceClient(compartment string) *InstanceClient {
	return &InstanceClient{
		client:        c,
		compartmendID: compartment,
	}
}

// GetInstance returns a struct of an instance request given an instance ID
func (ic *InstanceClient) GetInstance(instanceID string) Instance {
	instance := Instance{}
	resp, err := ic.client.Get("/instances/" + instanceID)
	if err != nil {
		logrus.Error(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	logrus.Info("Body: ", string(body))
	if err != nil {
		logrus.Fatalf("Could not read JSON response: %s", err)
	}
	if err = json.Unmarshal(body, &instance); err != nil {
		logrus.Fatalf("Unmarshal impossible: %s", err)
	}
	return instance
}

// ListInstances returns a slice struct of all instance
func (ic *InstanceClient) ListInstances(options *InstancesParameters) []Instance {
	instances := []Instance{}
	queryString := url.QueryEscape(ic.compartmendID)
	if options != nil {
		v, _ := query.Values(*options)
		queryString = queryString + "&" + v.Encode()
	}
	resp, err := ic.client.Get(fmt.Sprintf("/instances?compartmentId=%s", queryString))
	if err != nil {
		logrus.Error(err)
	}
	logrus.Info("StatusCode: ", resp.StatusCode)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Fatalf("Could not read JSON response: %s", err)
	}
	// logrus.Info("Body: ", string(body))

	if err = json.Unmarshal(body, &instances); err != nil {
		logrus.Fatalf("Unmarshal impossible: %s", err)
	}
	if options.Filter != nil {
		instances = filterInstances(instances, *options.Filter)
	}
	return instances
}
