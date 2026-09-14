// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeploymentRuleOptionsMonitorId A specific monitor and the groups to evaluate for it.
type DeploymentRuleOptionsMonitorId struct {
	// The exact monitor group names to evaluate. An empty array evaluates all groups.
	Groups []string `json:"groups"`
	// The monitor's decimal ID.
	Id string `json:"id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewDeploymentRuleOptionsMonitorId instantiates a new DeploymentRuleOptionsMonitorId object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeploymentRuleOptionsMonitorId(groups []string, id string) *DeploymentRuleOptionsMonitorId {
	this := DeploymentRuleOptionsMonitorId{}
	this.Groups = groups
	this.Id = id
	return &this
}

// NewDeploymentRuleOptionsMonitorIdWithDefaults instantiates a new DeploymentRuleOptionsMonitorId object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDeploymentRuleOptionsMonitorIdWithDefaults() *DeploymentRuleOptionsMonitorId {
	this := DeploymentRuleOptionsMonitorId{}
	return &this
}

// GetGroups returns the Groups field value.
func (o *DeploymentRuleOptionsMonitorId) GetGroups() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value
// and a boolean to check if the value has been set.
func (o *DeploymentRuleOptionsMonitorId) GetGroupsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Groups, true
}

// SetGroups sets field value.
func (o *DeploymentRuleOptionsMonitorId) SetGroups(v []string) {
	o.Groups = v
}

// GetId returns the Id field value.
func (o *DeploymentRuleOptionsMonitorId) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeploymentRuleOptionsMonitorId) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *DeploymentRuleOptionsMonitorId) SetId(v string) {
	o.Id = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DeploymentRuleOptionsMonitorId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["groups"] = o.Groups
	toSerialize["id"] = o.Id
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DeploymentRuleOptionsMonitorId) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Groups *[]string `json:"groups"`
		Id     *string   `json:"id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Groups == nil {
		return fmt.Errorf("required field groups missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	o.Groups = *all.Groups
	o.Id = *all.Id

	return nil
}
