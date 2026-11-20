// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NumberFormatUnitCanonical Test-only generated schema for the dashboard versioning proof of concept.
type NumberFormatUnitCanonical struct {
	// Datadog canonical unit identifier, for example byte or second.
	PerUnitName *string `json:"per_unit_name,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type NumberFormatUnitCanonicalType `json:"type"`
	// Datadog canonical unit identifier, for example byte or second.
	UnitName *string `json:"unit_name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewNumberFormatUnitCanonical instantiates a new NumberFormatUnitCanonical object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNumberFormatUnitCanonical(typeVar NumberFormatUnitCanonicalType) *NumberFormatUnitCanonical {
	this := NumberFormatUnitCanonical{}
	this.Type = typeVar
	return &this
}

// NewNumberFormatUnitCanonicalWithDefaults instantiates a new NumberFormatUnitCanonical object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNumberFormatUnitCanonicalWithDefaults() *NumberFormatUnitCanonical {
	this := NumberFormatUnitCanonical{}
	return &this
}

// GetPerUnitName returns the PerUnitName field value if set, zero value otherwise.
func (o *NumberFormatUnitCanonical) GetPerUnitName() string {
	if o == nil || o.PerUnitName == nil {
		var ret string
		return ret
	}
	return *o.PerUnitName
}

// GetPerUnitNameOk returns a tuple with the PerUnitName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberFormatUnitCanonical) GetPerUnitNameOk() (*string, bool) {
	if o == nil || o.PerUnitName == nil {
		return nil, false
	}
	return o.PerUnitName, true
}

// HasPerUnitName returns a boolean if a field has been set.
func (o *NumberFormatUnitCanonical) HasPerUnitName() bool {
	return o != nil && o.PerUnitName != nil
}

// SetPerUnitName gets a reference to the given string and assigns it to the PerUnitName field.
func (o *NumberFormatUnitCanonical) SetPerUnitName(v string) {
	o.PerUnitName = &v
}

// GetType returns the Type field value.
func (o *NumberFormatUnitCanonical) GetType() NumberFormatUnitCanonicalType {
	if o == nil {
		var ret NumberFormatUnitCanonicalType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NumberFormatUnitCanonical) GetTypeOk() (*NumberFormatUnitCanonicalType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *NumberFormatUnitCanonical) SetType(v NumberFormatUnitCanonicalType) {
	o.Type = v
}

// GetUnitName returns the UnitName field value if set, zero value otherwise.
func (o *NumberFormatUnitCanonical) GetUnitName() string {
	if o == nil || o.UnitName == nil {
		var ret string
		return ret
	}
	return *o.UnitName
}

// GetUnitNameOk returns a tuple with the UnitName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberFormatUnitCanonical) GetUnitNameOk() (*string, bool) {
	if o == nil || o.UnitName == nil {
		return nil, false
	}
	return o.UnitName, true
}

// HasUnitName returns a boolean if a field has been set.
func (o *NumberFormatUnitCanonical) HasUnitName() bool {
	return o != nil && o.UnitName != nil
}

// SetUnitName gets a reference to the given string and assigns it to the UnitName field.
func (o *NumberFormatUnitCanonical) SetUnitName(v string) {
	o.UnitName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o NumberFormatUnitCanonical) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.PerUnitName != nil {
		toSerialize["per_unit_name"] = o.PerUnitName
	}
	toSerialize["type"] = o.Type
	if o.UnitName != nil {
		toSerialize["unit_name"] = o.UnitName
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NumberFormatUnitCanonical) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		PerUnitName *string                        `json:"per_unit_name,omitempty"`
		Type        *NumberFormatUnitCanonicalType `json:"type"`
		UnitName    *string                        `json:"unit_name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	o.PerUnitName = all.PerUnitName
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.UnitName = all.UnitName

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
