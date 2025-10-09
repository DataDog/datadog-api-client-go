// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableRowResourceDataAttributes The definition of `TableRowResourceDataAttributes` object.
type TableRowResourceDataAttributes struct {
	// The values of the row.
	Values map[string]interface{} `json:"values,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTableRowResourceDataAttributes instantiates a new TableRowResourceDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTableRowResourceDataAttributes() *TableRowResourceDataAttributes {
	this := TableRowResourceDataAttributes{}
	return &this
}

// NewTableRowResourceDataAttributesWithDefaults instantiates a new TableRowResourceDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTableRowResourceDataAttributesWithDefaults() *TableRowResourceDataAttributes {
	this := TableRowResourceDataAttributes{}
	return &this
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *TableRowResourceDataAttributes) GetValues() map[string]interface{} {
	if o == nil || o.Values == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableRowResourceDataAttributes) GetValuesOk() (*map[string]interface{}, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return &o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *TableRowResourceDataAttributes) HasValues() bool {
	return o != nil && o.Values != nil
}

// SetValues gets a reference to the given map[string]interface{} and assigns it to the Values field.
func (o *TableRowResourceDataAttributes) SetValues(v map[string]interface{}) {
	o.Values = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TableRowResourceDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TableRowResourceDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Values map[string]interface{} `json:"values,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"values"})
	} else {
		return err
	}
	o.Values = all.Values

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
