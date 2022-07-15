// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadog

import (
	"encoding/json"
	"fmt"
)

// SecurityMonitoringSignalIncidentsUpdateData Data containing the patch for changing the related incidents of a signal.
type SecurityMonitoringSignalIncidentsUpdateData struct {
	// Attributes describing the new list of related signals for a security signal.
	Attributes SecurityMonitoringSignalIncidentsUpdateAttributes `json:"attributes"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}

// NewSecurityMonitoringSignalIncidentsUpdateData instantiates a new SecurityMonitoringSignalIncidentsUpdateData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringSignalIncidentsUpdateData(attributes SecurityMonitoringSignalIncidentsUpdateAttributes) *SecurityMonitoringSignalIncidentsUpdateData {
	this := SecurityMonitoringSignalIncidentsUpdateData{}
	this.Attributes = attributes
	return &this
}

// NewSecurityMonitoringSignalIncidentsUpdateDataWithDefaults instantiates a new SecurityMonitoringSignalIncidentsUpdateData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringSignalIncidentsUpdateDataWithDefaults() *SecurityMonitoringSignalIncidentsUpdateData {
	this := SecurityMonitoringSignalIncidentsUpdateData{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *SecurityMonitoringSignalIncidentsUpdateData) GetAttributes() SecurityMonitoringSignalIncidentsUpdateAttributes {
	if o == nil {
		var ret SecurityMonitoringSignalIncidentsUpdateAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalIncidentsUpdateData) GetAttributesOk() (*SecurityMonitoringSignalIncidentsUpdateAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *SecurityMonitoringSignalIncidentsUpdateData) SetAttributes(v SecurityMonitoringSignalIncidentsUpdateAttributes) {
	o.Attributes = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringSignalIncidentsUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringSignalIncidentsUpdateData) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		Attributes *SecurityMonitoringSignalIncidentsUpdateAttributes `json:"attributes"`
	}{}
	all := struct {
		Attributes SecurityMonitoringSignalIncidentsUpdateAttributes `json:"attributes"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Attributes == nil {
		return fmt.Errorf("Required field attributes missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.Attributes = all.Attributes
	return nil
}
