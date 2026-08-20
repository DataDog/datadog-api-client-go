// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsRetentionIndexTarget Selects a cohort or return period by its zero-based position in the grid.
type ProductAnalyticsRetentionIndexTarget struct {
	// The discriminator identifying a target selected by index.
	Type ProductAnalyticsRetentionIndexTargetType `json:"type"`
	// Zero-based index of the targeted cohort or return period.
	Value int64 `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsRetentionIndexTarget instantiates a new ProductAnalyticsRetentionIndexTarget object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsRetentionIndexTarget(typeVar ProductAnalyticsRetentionIndexTargetType, value int64) *ProductAnalyticsRetentionIndexTarget {
	this := ProductAnalyticsRetentionIndexTarget{}
	this.Type = typeVar
	this.Value = value
	return &this
}

// NewProductAnalyticsRetentionIndexTargetWithDefaults instantiates a new ProductAnalyticsRetentionIndexTarget object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsRetentionIndexTargetWithDefaults() *ProductAnalyticsRetentionIndexTarget {
	this := ProductAnalyticsRetentionIndexTarget{}
	return &this
}

// GetType returns the Type field value.
func (o *ProductAnalyticsRetentionIndexTarget) GetType() ProductAnalyticsRetentionIndexTargetType {
	if o == nil {
		var ret ProductAnalyticsRetentionIndexTargetType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionIndexTarget) GetTypeOk() (*ProductAnalyticsRetentionIndexTargetType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ProductAnalyticsRetentionIndexTarget) SetType(v ProductAnalyticsRetentionIndexTargetType) {
	o.Type = v
}

// GetValue returns the Value field value.
func (o *ProductAnalyticsRetentionIndexTarget) GetValue() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionIndexTarget) GetValueOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *ProductAnalyticsRetentionIndexTarget) SetValue(v int64) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsRetentionIndexTarget) MarshalJSON() ([]byte, error) {
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
func (o *ProductAnalyticsRetentionIndexTarget) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type  *ProductAnalyticsRetentionIndexTargetType `json:"type"`
		Value *int64                                    `json:"value"`
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
	o.Value = *all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
