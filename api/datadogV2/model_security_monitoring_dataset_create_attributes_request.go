// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringDatasetCreateAttributesRequest Attributes for creating a security monitoring dataset.
type SecurityMonitoringDatasetCreateAttributesRequest struct {
	// The definition of a dataset, including its data source, name, indexes, and columns.
	Definition SecurityMonitoringDatasetDefinition `json:"definition"`
	// A description of the dataset (maximum 255 characters).
	Description string `json:"description"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringDatasetCreateAttributesRequest instantiates a new SecurityMonitoringDatasetCreateAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringDatasetCreateAttributesRequest(definition SecurityMonitoringDatasetDefinition, description string) *SecurityMonitoringDatasetCreateAttributesRequest {
	this := SecurityMonitoringDatasetCreateAttributesRequest{}
	this.Definition = definition
	this.Description = description
	return &this
}

// NewSecurityMonitoringDatasetCreateAttributesRequestWithDefaults instantiates a new SecurityMonitoringDatasetCreateAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringDatasetCreateAttributesRequestWithDefaults() *SecurityMonitoringDatasetCreateAttributesRequest {
	this := SecurityMonitoringDatasetCreateAttributesRequest{}
	return &this
}

// GetDefinition returns the Definition field value.
func (o *SecurityMonitoringDatasetCreateAttributesRequest) GetDefinition() SecurityMonitoringDatasetDefinition {
	if o == nil {
		var ret SecurityMonitoringDatasetDefinition
		return ret
	}
	return o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetCreateAttributesRequest) GetDefinitionOk() (*SecurityMonitoringDatasetDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Definition, true
}

// SetDefinition sets field value.
func (o *SecurityMonitoringDatasetCreateAttributesRequest) SetDefinition(v SecurityMonitoringDatasetDefinition) {
	o.Definition = v
}

// GetDescription returns the Description field value.
func (o *SecurityMonitoringDatasetCreateAttributesRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetCreateAttributesRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *SecurityMonitoringDatasetCreateAttributesRequest) SetDescription(v string) {
	o.Description = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringDatasetCreateAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["definition"] = o.Definition
	toSerialize["description"] = o.Description

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringDatasetCreateAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Definition  *SecurityMonitoringDatasetDefinition `json:"definition"`
		Description *string                              `json:"description"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Definition == nil {
		return fmt.Errorf("required field definition missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"definition", "description"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Definition.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Definition = *all.Definition
	o.Description = *all.Description

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
