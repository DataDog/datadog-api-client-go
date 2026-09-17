// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsRetentionFixedTimeInterval A retention interval of fixed length, such as "7 days".
type ProductAnalyticsRetentionFixedTimeInterval struct {
	// The discriminator identifying a fixed-length retention interval.
	Type ProductAnalyticsRetentionFixedTimeIntervalType `json:"type"`
	// Time unit for a fixed-length retention interval.
	Unit ProductAnalyticsRetentionFixedTimeIntervalUnit `json:"unit"`
	// Length of the interval, expressed in `unit`.
	Value float64 `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsRetentionFixedTimeInterval instantiates a new ProductAnalyticsRetentionFixedTimeInterval object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsRetentionFixedTimeInterval(typeVar ProductAnalyticsRetentionFixedTimeIntervalType, unit ProductAnalyticsRetentionFixedTimeIntervalUnit, value float64) *ProductAnalyticsRetentionFixedTimeInterval {
	this := ProductAnalyticsRetentionFixedTimeInterval{}
	this.Type = typeVar
	this.Unit = unit
	this.Value = value
	return &this
}

// NewProductAnalyticsRetentionFixedTimeIntervalWithDefaults instantiates a new ProductAnalyticsRetentionFixedTimeInterval object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsRetentionFixedTimeIntervalWithDefaults() *ProductAnalyticsRetentionFixedTimeInterval {
	this := ProductAnalyticsRetentionFixedTimeInterval{}
	return &this
}

// GetType returns the Type field value.
func (o *ProductAnalyticsRetentionFixedTimeInterval) GetType() ProductAnalyticsRetentionFixedTimeIntervalType {
	if o == nil {
		var ret ProductAnalyticsRetentionFixedTimeIntervalType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionFixedTimeInterval) GetTypeOk() (*ProductAnalyticsRetentionFixedTimeIntervalType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ProductAnalyticsRetentionFixedTimeInterval) SetType(v ProductAnalyticsRetentionFixedTimeIntervalType) {
	o.Type = v
}

// GetUnit returns the Unit field value.
func (o *ProductAnalyticsRetentionFixedTimeInterval) GetUnit() ProductAnalyticsRetentionFixedTimeIntervalUnit {
	if o == nil {
		var ret ProductAnalyticsRetentionFixedTimeIntervalUnit
		return ret
	}
	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionFixedTimeInterval) GetUnitOk() (*ProductAnalyticsRetentionFixedTimeIntervalUnit, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value.
func (o *ProductAnalyticsRetentionFixedTimeInterval) SetUnit(v ProductAnalyticsRetentionFixedTimeIntervalUnit) {
	o.Unit = v
}

// GetValue returns the Value field value.
func (o *ProductAnalyticsRetentionFixedTimeInterval) GetValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionFixedTimeInterval) GetValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *ProductAnalyticsRetentionFixedTimeInterval) SetValue(v float64) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsRetentionFixedTimeInterval) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["type"] = o.Type
	toSerialize["unit"] = o.Unit
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsRetentionFixedTimeInterval) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type  *ProductAnalyticsRetentionFixedTimeIntervalType `json:"type"`
		Unit  *ProductAnalyticsRetentionFixedTimeIntervalUnit `json:"unit"`
		Value *float64                                        `json:"value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Unit == nil {
		return fmt.Errorf("required field unit missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"type", "unit", "value"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	if !all.Unit.IsValid() {
		hasInvalidField = true
	} else {
		o.Unit = *all.Unit
	}
	o.Value = *all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
