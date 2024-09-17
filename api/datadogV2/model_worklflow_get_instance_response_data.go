// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WorklflowGetInstanceResponseData The data of the instance response.
type WorklflowGetInstanceResponseData struct {
	// The attributes of the instance response data.
	Attributes *WorklflowGetInstanceResponseDataAttributes `json:"attributes,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWorklflowGetInstanceResponseData instantiates a new WorklflowGetInstanceResponseData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWorklflowGetInstanceResponseData() *WorklflowGetInstanceResponseData {
	this := WorklflowGetInstanceResponseData{}
	return &this
}

// NewWorklflowGetInstanceResponseDataWithDefaults instantiates a new WorklflowGetInstanceResponseData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWorklflowGetInstanceResponseDataWithDefaults() *WorklflowGetInstanceResponseData {
	this := WorklflowGetInstanceResponseData{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *WorklflowGetInstanceResponseData) GetAttributes() WorklflowGetInstanceResponseDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret WorklflowGetInstanceResponseDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorklflowGetInstanceResponseData) GetAttributesOk() (*WorklflowGetInstanceResponseDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *WorklflowGetInstanceResponseData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given WorklflowGetInstanceResponseDataAttributes and assigns it to the Attributes field.
func (o *WorklflowGetInstanceResponseData) SetAttributes(v WorklflowGetInstanceResponseDataAttributes) {
	o.Attributes = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o WorklflowGetInstanceResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WorklflowGetInstanceResponseData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *WorklflowGetInstanceResponseDataAttributes `json:"attributes,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
