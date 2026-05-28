// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OverviewItemDataAttributes Attributes of a single tile in the End User Device Monitoring overview dashboard.
type OverviewItemDataAttributes struct {
	// Human-readable name of the overview tile.
	Name *string `json:"name,omitempty"`
	// Numeric value displayed on the overview tile.
	Value *int64 `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOverviewItemDataAttributes instantiates a new OverviewItemDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOverviewItemDataAttributes() *OverviewItemDataAttributes {
	this := OverviewItemDataAttributes{}
	return &this
}

// NewOverviewItemDataAttributesWithDefaults instantiates a new OverviewItemDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOverviewItemDataAttributesWithDefaults() *OverviewItemDataAttributes {
	this := OverviewItemDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OverviewItemDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverviewItemDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OverviewItemDataAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OverviewItemDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *OverviewItemDataAttributes) GetValue() int64 {
	if o == nil || o.Value == nil {
		var ret int64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverviewItemDataAttributes) GetValueOk() (*int64, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *OverviewItemDataAttributes) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given int64 and assigns it to the Value field.
func (o *OverviewItemDataAttributes) SetValue(v int64) {
	o.Value = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OverviewItemDataAttributes) MarshalJSON() ([]byte, error) {
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
func (o *OverviewItemDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name  *string `json:"name,omitempty"`
		Value *int64  `json:"value,omitempty"`
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
