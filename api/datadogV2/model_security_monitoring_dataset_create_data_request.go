// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringDatasetCreateDataRequest Data for creating a security monitoring dataset.
type SecurityMonitoringDatasetCreateDataRequest struct {
	// Attributes for creating a security monitoring dataset.
	Attributes SecurityMonitoringDatasetCreateAttributesRequest `json:"attributes"`
	// Type for security monitoring dataset objects.
	Type SecurityMonitoringDatasetType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringDatasetCreateDataRequest instantiates a new SecurityMonitoringDatasetCreateDataRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringDatasetCreateDataRequest(attributes SecurityMonitoringDatasetCreateAttributesRequest, typeVar SecurityMonitoringDatasetType) *SecurityMonitoringDatasetCreateDataRequest {
	this := SecurityMonitoringDatasetCreateDataRequest{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewSecurityMonitoringDatasetCreateDataRequestWithDefaults instantiates a new SecurityMonitoringDatasetCreateDataRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringDatasetCreateDataRequestWithDefaults() *SecurityMonitoringDatasetCreateDataRequest {
	this := SecurityMonitoringDatasetCreateDataRequest{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *SecurityMonitoringDatasetCreateDataRequest) GetAttributes() SecurityMonitoringDatasetCreateAttributesRequest {
	if o == nil {
		var ret SecurityMonitoringDatasetCreateAttributesRequest
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetCreateDataRequest) GetAttributesOk() (*SecurityMonitoringDatasetCreateAttributesRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *SecurityMonitoringDatasetCreateDataRequest) SetAttributes(v SecurityMonitoringDatasetCreateAttributesRequest) {
	o.Attributes = v
}

// GetType returns the Type field value.
func (o *SecurityMonitoringDatasetCreateDataRequest) GetType() SecurityMonitoringDatasetType {
	if o == nil {
		var ret SecurityMonitoringDatasetType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetCreateDataRequest) GetTypeOk() (*SecurityMonitoringDatasetType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SecurityMonitoringDatasetCreateDataRequest) SetType(v SecurityMonitoringDatasetType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringDatasetCreateDataRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringDatasetCreateDataRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *SecurityMonitoringDatasetCreateAttributesRequest `json:"attributes"`
		Type       *SecurityMonitoringDatasetType                    `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
