// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeploymentGateRuleFailureMonitor Failed monitor reference.
type DeploymentGateRuleFailureMonitor struct {
	// Monitor ID.
	Id string `json:"id"`
	// Monitor state that caused the failure.
	State string `json:"state"`
	// URL for inspecting the monitor during the evaluation window.
	Url string `json:"url"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDeploymentGateRuleFailureMonitor instantiates a new DeploymentGateRuleFailureMonitor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeploymentGateRuleFailureMonitor(id string, state string, url string) *DeploymentGateRuleFailureMonitor {
	this := DeploymentGateRuleFailureMonitor{}
	this.Id = id
	this.State = state
	this.Url = url
	return &this
}

// NewDeploymentGateRuleFailureMonitorWithDefaults instantiates a new DeploymentGateRuleFailureMonitor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDeploymentGateRuleFailureMonitorWithDefaults() *DeploymentGateRuleFailureMonitor {
	this := DeploymentGateRuleFailureMonitor{}
	return &this
}

// GetId returns the Id field value.
func (o *DeploymentGateRuleFailureMonitor) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeploymentGateRuleFailureMonitor) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *DeploymentGateRuleFailureMonitor) SetId(v string) {
	o.Id = v
}

// GetState returns the State field value.
func (o *DeploymentGateRuleFailureMonitor) GetState() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *DeploymentGateRuleFailureMonitor) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value.
func (o *DeploymentGateRuleFailureMonitor) SetState(v string) {
	o.State = v
}

// GetUrl returns the Url field value.
func (o *DeploymentGateRuleFailureMonitor) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *DeploymentGateRuleFailureMonitor) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value.
func (o *DeploymentGateRuleFailureMonitor) SetUrl(v string) {
	o.Url = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DeploymentGateRuleFailureMonitor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["state"] = o.State
	toSerialize["url"] = o.Url

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DeploymentGateRuleFailureMonitor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id    *string `json:"id"`
		State *string `json:"state"`
		Url   *string `json:"url"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.State == nil {
		return fmt.Errorf("required field state missing")
	}
	if all.Url == nil {
		return fmt.Errorf("required field url missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "state", "url"})
	} else {
		return err
	}
	o.Id = *all.Id
	o.State = *all.State
	o.Url = *all.Url

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
