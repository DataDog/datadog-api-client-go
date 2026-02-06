// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IntegrationAssignmentDataAttributesRequestAssignment
type IntegrationAssignmentDataAttributesRequestAssignment struct {
	// Map of Jira issue URLs to lists of finding IDs.
	Jira map[string][]string `json:"jira"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIntegrationAssignmentDataAttributesRequestAssignment instantiates a new IntegrationAssignmentDataAttributesRequestAssignment object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIntegrationAssignmentDataAttributesRequestAssignment(jira map[string][]string) *IntegrationAssignmentDataAttributesRequestAssignment {
	this := IntegrationAssignmentDataAttributesRequestAssignment{}
	this.Jira = jira
	return &this
}

// NewIntegrationAssignmentDataAttributesRequestAssignmentWithDefaults instantiates a new IntegrationAssignmentDataAttributesRequestAssignment object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIntegrationAssignmentDataAttributesRequestAssignmentWithDefaults() *IntegrationAssignmentDataAttributesRequestAssignment {
	this := IntegrationAssignmentDataAttributesRequestAssignment{}
	return &this
}

// GetJira returns the Jira field value.
func (o *IntegrationAssignmentDataAttributesRequestAssignment) GetJira() map[string][]string {
	if o == nil {
		var ret map[string][]string
		return ret
	}
	return o.Jira
}

// GetJiraOk returns a tuple with the Jira field value
// and a boolean to check if the value has been set.
func (o *IntegrationAssignmentDataAttributesRequestAssignment) GetJiraOk() (*map[string][]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Jira, true
}

// SetJira sets field value.
func (o *IntegrationAssignmentDataAttributesRequestAssignment) SetJira(v map[string][]string) {
	o.Jira = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IntegrationAssignmentDataAttributesRequestAssignment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["jira"] = o.Jira

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IntegrationAssignmentDataAttributesRequestAssignment) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Jira *map[string][]string `json:"jira"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Jira == nil {
		return fmt.Errorf("required field jira missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"jira"})
	} else {
		return err
	}
	o.Jira = *all.Jira

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
