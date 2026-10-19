// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AgentHealthHostIssuesAttributes Active Agent Health issues for a resource.
type AgentHealthHostIssuesAttributes struct {
	// The active issues and their reconciliation metadata.
	Issues []AgentHealthHostIssue `json:"issues"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAgentHealthHostIssuesAttributes instantiates a new AgentHealthHostIssuesAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAgentHealthHostIssuesAttributes(issues []AgentHealthHostIssue) *AgentHealthHostIssuesAttributes {
	this := AgentHealthHostIssuesAttributes{}
	this.Issues = issues
	return &this
}

// NewAgentHealthHostIssuesAttributesWithDefaults instantiates a new AgentHealthHostIssuesAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAgentHealthHostIssuesAttributesWithDefaults() *AgentHealthHostIssuesAttributes {
	this := AgentHealthHostIssuesAttributes{}
	return &this
}

// GetIssues returns the Issues field value.
func (o *AgentHealthHostIssuesAttributes) GetIssues() []AgentHealthHostIssue {
	if o == nil {
		var ret []AgentHealthHostIssue
		return ret
	}
	return o.Issues
}

// GetIssuesOk returns a tuple with the Issues field value
// and a boolean to check if the value has been set.
func (o *AgentHealthHostIssuesAttributes) GetIssuesOk() (*[]AgentHealthHostIssue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Issues, true
}

// SetIssues sets field value.
func (o *AgentHealthHostIssuesAttributes) SetIssues(v []AgentHealthHostIssue) {
	o.Issues = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AgentHealthHostIssuesAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["issues"] = o.Issues

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AgentHealthHostIssuesAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Issues *[]AgentHealthHostIssue `json:"issues"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Issues == nil {
		return fmt.Errorf("required field issues missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"issues"})
	} else {
		return err
	}
	o.Issues = *all.Issues

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
