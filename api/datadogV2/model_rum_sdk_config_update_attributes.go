// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMSdkConfigUpdateAttributes Attributes of the RUM SDK configuration to update.
type RUMSdkConfigUpdateAttributes struct {
	// The RUM SDK settings to apply when updating a configuration.
	Rum RUMSdkConfigRUMUpdateAttributes `json:"rum"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRUMSdkConfigUpdateAttributes instantiates a new RUMSdkConfigUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRUMSdkConfigUpdateAttributes(rum RUMSdkConfigRUMUpdateAttributes) *RUMSdkConfigUpdateAttributes {
	this := RUMSdkConfigUpdateAttributes{}
	this.Rum = rum
	return &this
}

// NewRUMSdkConfigUpdateAttributesWithDefaults instantiates a new RUMSdkConfigUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRUMSdkConfigUpdateAttributesWithDefaults() *RUMSdkConfigUpdateAttributes {
	this := RUMSdkConfigUpdateAttributes{}
	return &this
}

// GetRum returns the Rum field value.
func (o *RUMSdkConfigUpdateAttributes) GetRum() RUMSdkConfigRUMUpdateAttributes {
	if o == nil {
		var ret RUMSdkConfigRUMUpdateAttributes
		return ret
	}
	return o.Rum
}

// GetRumOk returns a tuple with the Rum field value
// and a boolean to check if the value has been set.
func (o *RUMSdkConfigUpdateAttributes) GetRumOk() (*RUMSdkConfigRUMUpdateAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rum, true
}

// SetRum sets field value.
func (o *RUMSdkConfigUpdateAttributes) SetRum(v RUMSdkConfigRUMUpdateAttributes) {
	o.Rum = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RUMSdkConfigUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["rum"] = o.Rum

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RUMSdkConfigUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Rum *RUMSdkConfigRUMUpdateAttributes `json:"rum"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Rum == nil {
		return fmt.Errorf("required field rum missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"rum"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Rum.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Rum = *all.Rum

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
