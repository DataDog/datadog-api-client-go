// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsRetentionAggregationTarget Selects the rolled-up row that aggregates every cohort, rather than a single cohort.
type ProductAnalyticsRetentionAggregationTarget struct {
	// The discriminator identifying a target selected by aggregation.
	Type ProductAnalyticsRetentionAggregationTargetType `json:"type"`
	// The aggregation that produced the rolled-up row.
	Value string `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsRetentionAggregationTarget instantiates a new ProductAnalyticsRetentionAggregationTarget object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsRetentionAggregationTarget(typeVar ProductAnalyticsRetentionAggregationTargetType, value string) *ProductAnalyticsRetentionAggregationTarget {
	this := ProductAnalyticsRetentionAggregationTarget{}
	this.Type = typeVar
	this.Value = value
	return &this
}

// NewProductAnalyticsRetentionAggregationTargetWithDefaults instantiates a new ProductAnalyticsRetentionAggregationTarget object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsRetentionAggregationTargetWithDefaults() *ProductAnalyticsRetentionAggregationTarget {
	this := ProductAnalyticsRetentionAggregationTarget{}
	return &this
}

// GetType returns the Type field value.
func (o *ProductAnalyticsRetentionAggregationTarget) GetType() ProductAnalyticsRetentionAggregationTargetType {
	if o == nil {
		var ret ProductAnalyticsRetentionAggregationTargetType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionAggregationTarget) GetTypeOk() (*ProductAnalyticsRetentionAggregationTargetType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ProductAnalyticsRetentionAggregationTarget) SetType(v ProductAnalyticsRetentionAggregationTargetType) {
	o.Type = v
}

// GetValue returns the Value field value.
func (o *ProductAnalyticsRetentionAggregationTarget) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionAggregationTarget) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *ProductAnalyticsRetentionAggregationTarget) SetValue(v string) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsRetentionAggregationTarget) MarshalJSON() ([]byte, error) {
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
func (o *ProductAnalyticsRetentionAggregationTarget) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type  *ProductAnalyticsRetentionAggregationTargetType `json:"type"`
		Value *string                                         `json:"value"`
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
