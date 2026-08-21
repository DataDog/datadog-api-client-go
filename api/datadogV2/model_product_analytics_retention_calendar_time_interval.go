// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsRetentionCalendarTimeInterval A retention interval aligned to calendar boundaries.
type ProductAnalyticsRetentionCalendarTimeInterval struct {
	// The discriminator identifying a calendar-aligned retention interval.
	Type ProductAnalyticsRetentionCalendarTimeIntervalType `json:"type"`
	// A calendar-aligned bucket definition, such as "every 1 week starting on Monday".
	Value ProductAnalyticsCalendarInterval `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsRetentionCalendarTimeInterval instantiates a new ProductAnalyticsRetentionCalendarTimeInterval object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsRetentionCalendarTimeInterval(typeVar ProductAnalyticsRetentionCalendarTimeIntervalType, value ProductAnalyticsCalendarInterval) *ProductAnalyticsRetentionCalendarTimeInterval {
	this := ProductAnalyticsRetentionCalendarTimeInterval{}
	this.Type = typeVar
	this.Value = value
	return &this
}

// NewProductAnalyticsRetentionCalendarTimeIntervalWithDefaults instantiates a new ProductAnalyticsRetentionCalendarTimeInterval object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsRetentionCalendarTimeIntervalWithDefaults() *ProductAnalyticsRetentionCalendarTimeInterval {
	this := ProductAnalyticsRetentionCalendarTimeInterval{}
	return &this
}

// GetType returns the Type field value.
func (o *ProductAnalyticsRetentionCalendarTimeInterval) GetType() ProductAnalyticsRetentionCalendarTimeIntervalType {
	if o == nil {
		var ret ProductAnalyticsRetentionCalendarTimeIntervalType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionCalendarTimeInterval) GetTypeOk() (*ProductAnalyticsRetentionCalendarTimeIntervalType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ProductAnalyticsRetentionCalendarTimeInterval) SetType(v ProductAnalyticsRetentionCalendarTimeIntervalType) {
	o.Type = v
}

// GetValue returns the Value field value.
func (o *ProductAnalyticsRetentionCalendarTimeInterval) GetValue() ProductAnalyticsCalendarInterval {
	if o == nil {
		var ret ProductAnalyticsCalendarInterval
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionCalendarTimeInterval) GetValueOk() (*ProductAnalyticsCalendarInterval, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *ProductAnalyticsRetentionCalendarTimeInterval) SetValue(v ProductAnalyticsCalendarInterval) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsRetentionCalendarTimeInterval) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["type"] = o.Type
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsRetentionCalendarTimeInterval) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type  *ProductAnalyticsRetentionCalendarTimeIntervalType `json:"type"`
		Value *ProductAnalyticsCalendarInterval                  `json:"value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"type", "value"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	if all.Value.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
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
