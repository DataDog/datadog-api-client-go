// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsRetentionPeriod A return period definition, such as "1 week".
type ProductAnalyticsRetentionPeriod struct {
	// Time unit of the period, such as `day`, `week`, `month`, or `year`.
	Unit *string `json:"unit,omitempty"`
	// Length of the period, expressed in `unit`.
	Value *int64 `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsRetentionPeriod instantiates a new ProductAnalyticsRetentionPeriod object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsRetentionPeriod() *ProductAnalyticsRetentionPeriod {
	this := ProductAnalyticsRetentionPeriod{}
	return &this
}

// NewProductAnalyticsRetentionPeriodWithDefaults instantiates a new ProductAnalyticsRetentionPeriod object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsRetentionPeriodWithDefaults() *ProductAnalyticsRetentionPeriod {
	this := ProductAnalyticsRetentionPeriod{}
	return &this
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *ProductAnalyticsRetentionPeriod) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionPeriod) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *ProductAnalyticsRetentionPeriod) HasUnit() bool {
	return o != nil && o.Unit != nil
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *ProductAnalyticsRetentionPeriod) SetUnit(v string) {
	o.Unit = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ProductAnalyticsRetentionPeriod) GetValue() int64 {
	if o == nil || o.Value == nil {
		var ret int64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionPeriod) GetValueOk() (*int64, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ProductAnalyticsRetentionPeriod) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given int64 and assigns it to the Value field.
func (o *ProductAnalyticsRetentionPeriod) SetValue(v int64) {
	o.Value = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsRetentionPeriod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
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
func (o *ProductAnalyticsRetentionPeriod) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Unit  *string `json:"unit,omitempty"`
		Value *int64  `json:"value,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"unit", "value"})
	} else {
		return err
	}
	o.Unit = all.Unit
	o.Value = all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
