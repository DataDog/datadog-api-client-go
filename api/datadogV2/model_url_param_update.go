// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UrlParamUpdate The definition of `UrlParamUpdate` object.
type UrlParamUpdate struct {
	// Should the header be deleted.
	Deleted *bool `json:"deleted,omitempty"`
	// Name for tokens.
	Name string `json:"name"`
	// The `UrlParamUpdate` `value`.
	Value *string `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUrlParamUpdate instantiates a new UrlParamUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUrlParamUpdate(name string) *UrlParamUpdate {
	this := UrlParamUpdate{}
	this.Name = name
	return &this
}

// NewUrlParamUpdateWithDefaults instantiates a new UrlParamUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUrlParamUpdateWithDefaults() *UrlParamUpdate {
	this := UrlParamUpdate{}
	return &this
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *UrlParamUpdate) GetDeleted() bool {
	if o == nil || o.Deleted == nil {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UrlParamUpdate) GetDeletedOk() (*bool, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *UrlParamUpdate) HasDeleted() bool {
	return o != nil && o.Deleted != nil
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *UrlParamUpdate) SetDeleted(v bool) {
	o.Deleted = &v
}

// GetName returns the Name field value.
func (o *UrlParamUpdate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UrlParamUpdate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *UrlParamUpdate) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *UrlParamUpdate) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UrlParamUpdate) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *UrlParamUpdate) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *UrlParamUpdate) SetValue(v string) {
	o.Value = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UrlParamUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Deleted != nil {
		toSerialize["deleted"] = o.Deleted
	}
	toSerialize["name"] = o.Name
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UrlParamUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Deleted *bool   `json:"deleted,omitempty"`
		Name    *string `json:"name"`
		Value   *string `json:"value,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"deleted", "name", "value"})
	} else {
		return err
	}
	o.Deleted = all.Deleted
	o.Name = *all.Name
	o.Value = all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
