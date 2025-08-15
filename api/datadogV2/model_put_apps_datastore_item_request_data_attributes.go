// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PutAppsDatastoreItemRequestDataAttributes The definition of `PutAppsDatastoreItemRequestDataAttributes` object.
type PutAppsDatastoreItemRequestDataAttributes struct {
	// The `attributes` `value`.
	Value map[string]interface{} `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPutAppsDatastoreItemRequestDataAttributes instantiates a new PutAppsDatastoreItemRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPutAppsDatastoreItemRequestDataAttributes(value map[string]interface{}) *PutAppsDatastoreItemRequestDataAttributes {
	this := PutAppsDatastoreItemRequestDataAttributes{}
	this.Value = value
	return &this
}

// NewPutAppsDatastoreItemRequestDataAttributesWithDefaults instantiates a new PutAppsDatastoreItemRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPutAppsDatastoreItemRequestDataAttributesWithDefaults() *PutAppsDatastoreItemRequestDataAttributes {
	this := PutAppsDatastoreItemRequestDataAttributes{}
	return &this
}

// GetValue returns the Value field value.
func (o *PutAppsDatastoreItemRequestDataAttributes) GetValue() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *PutAppsDatastoreItemRequestDataAttributes) GetValueOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *PutAppsDatastoreItemRequestDataAttributes) SetValue(v map[string]interface{}) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PutAppsDatastoreItemRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PutAppsDatastoreItemRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Value *map[string]interface{} `json:"value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"value"})
	} else {
		return err
	}
	o.Value = *all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
