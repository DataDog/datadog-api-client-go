// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateVariantAttributes Attributes for updating a variant.
type UpdateVariantAttributes struct {
	// The display name of the variant.
	Name *string `json:"name,omitempty"`
	// The value of the variant as a string.
	Value *string `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpdateVariantAttributes instantiates a new UpdateVariantAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpdateVariantAttributes() *UpdateVariantAttributes {
	this := UpdateVariantAttributes{}
	return &this
}

// NewUpdateVariantAttributesWithDefaults instantiates a new UpdateVariantAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpdateVariantAttributesWithDefaults() *UpdateVariantAttributes {
	this := UpdateVariantAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateVariantAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVariantAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateVariantAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateVariantAttributes) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *UpdateVariantAttributes) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVariantAttributes) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *UpdateVariantAttributes) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *UpdateVariantAttributes) SetValue(v string) {
	o.Value = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpdateVariantAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UpdateVariantAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name  *string `json:"name,omitempty"`
		Value *string `json:"value,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "value"})
	} else {
		return err
	}
	o.Name = all.Name
	o.Value = all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
