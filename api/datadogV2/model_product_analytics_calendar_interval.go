// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsCalendarInterval A calendar-aligned bucket definition, such as "every 1 week starting on Monday".
type ProductAnalyticsCalendarInterval struct {
	// Where each bucket starts within the calendar unit. Use an hour for `day` (for example `1am` or `14`),
	// a day name for `week` (for example `monday`), or an ordinal for `month` (for example `1st`).
	Alignment *string `json:"alignment,omitempty"`
	// Number of calendar units per bucket.
	Quantity *int64 `json:"quantity,omitempty"`
	// Timezone used to align the buckets.
	Timezone *string `json:"timezone,omitempty"`
	// Calendar unit used to bucket cohorts.
	Type ProductAnalyticsCalendarIntervalType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsCalendarInterval instantiates a new ProductAnalyticsCalendarInterval object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsCalendarInterval(typeVar ProductAnalyticsCalendarIntervalType) *ProductAnalyticsCalendarInterval {
	this := ProductAnalyticsCalendarInterval{}
	this.Type = typeVar
	return &this
}

// NewProductAnalyticsCalendarIntervalWithDefaults instantiates a new ProductAnalyticsCalendarInterval object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsCalendarIntervalWithDefaults() *ProductAnalyticsCalendarInterval {
	this := ProductAnalyticsCalendarInterval{}
	return &this
}

// GetAlignment returns the Alignment field value if set, zero value otherwise.
func (o *ProductAnalyticsCalendarInterval) GetAlignment() string {
	if o == nil || o.Alignment == nil {
		var ret string
		return ret
	}
	return *o.Alignment
}

// GetAlignmentOk returns a tuple with the Alignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsCalendarInterval) GetAlignmentOk() (*string, bool) {
	if o == nil || o.Alignment == nil {
		return nil, false
	}
	return o.Alignment, true
}

// HasAlignment returns a boolean if a field has been set.
func (o *ProductAnalyticsCalendarInterval) HasAlignment() bool {
	return o != nil && o.Alignment != nil
}

// SetAlignment gets a reference to the given string and assigns it to the Alignment field.
func (o *ProductAnalyticsCalendarInterval) SetAlignment(v string) {
	o.Alignment = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *ProductAnalyticsCalendarInterval) GetQuantity() int64 {
	if o == nil || o.Quantity == nil {
		var ret int64
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsCalendarInterval) GetQuantityOk() (*int64, bool) {
	if o == nil || o.Quantity == nil {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *ProductAnalyticsCalendarInterval) HasQuantity() bool {
	return o != nil && o.Quantity != nil
}

// SetQuantity gets a reference to the given int64 and assigns it to the Quantity field.
func (o *ProductAnalyticsCalendarInterval) SetQuantity(v int64) {
	o.Quantity = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *ProductAnalyticsCalendarInterval) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsCalendarInterval) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *ProductAnalyticsCalendarInterval) HasTimezone() bool {
	return o != nil && o.Timezone != nil
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *ProductAnalyticsCalendarInterval) SetTimezone(v string) {
	o.Timezone = &v
}

// GetType returns the Type field value.
func (o *ProductAnalyticsCalendarInterval) GetType() ProductAnalyticsCalendarIntervalType {
	if o == nil {
		var ret ProductAnalyticsCalendarIntervalType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsCalendarInterval) GetTypeOk() (*ProductAnalyticsCalendarIntervalType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ProductAnalyticsCalendarInterval) SetType(v ProductAnalyticsCalendarIntervalType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsCalendarInterval) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Alignment != nil {
		toSerialize["alignment"] = o.Alignment
	}
	if o.Quantity != nil {
		toSerialize["quantity"] = o.Quantity
	}
	if o.Timezone != nil {
		toSerialize["timezone"] = o.Timezone
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsCalendarInterval) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Alignment *string                               `json:"alignment,omitempty"`
		Quantity  *int64                                `json:"quantity,omitempty"`
		Timezone  *string                               `json:"timezone,omitempty"`
		Type      *ProductAnalyticsCalendarIntervalType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"alignment", "quantity", "timezone", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Alignment = all.Alignment
	o.Quantity = all.Quantity
	o.Timezone = all.Timezone
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
