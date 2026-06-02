// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumHardcodedRetentionFilterUpdateAttributes The attributes of a hardcoded retention filter that can be updated.
// Only fields whose matching flag in `cross_product_sampling_editability` is `true` can be modified.
type RumHardcodedRetentionFilterUpdateAttributes struct {
	// Partial update for cross-product retention of a hardcoded retention filter.
	// Only fields whose matching flag in `cross_product_sampling_editability` is `true` can be updated.
	CrossProductSampling *RumHardcodedCrossProductSamplingUpdate `json:"cross_product_sampling,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumHardcodedRetentionFilterUpdateAttributes instantiates a new RumHardcodedRetentionFilterUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumHardcodedRetentionFilterUpdateAttributes() *RumHardcodedRetentionFilterUpdateAttributes {
	this := RumHardcodedRetentionFilterUpdateAttributes{}
	return &this
}

// NewRumHardcodedRetentionFilterUpdateAttributesWithDefaults instantiates a new RumHardcodedRetentionFilterUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumHardcodedRetentionFilterUpdateAttributesWithDefaults() *RumHardcodedRetentionFilterUpdateAttributes {
	this := RumHardcodedRetentionFilterUpdateAttributes{}
	return &this
}

// GetCrossProductSampling returns the CrossProductSampling field value if set, zero value otherwise.
func (o *RumHardcodedRetentionFilterUpdateAttributes) GetCrossProductSampling() RumHardcodedCrossProductSamplingUpdate {
	if o == nil || o.CrossProductSampling == nil {
		var ret RumHardcodedCrossProductSamplingUpdate
		return ret
	}
	return *o.CrossProductSampling
}

// GetCrossProductSamplingOk returns a tuple with the CrossProductSampling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumHardcodedRetentionFilterUpdateAttributes) GetCrossProductSamplingOk() (*RumHardcodedCrossProductSamplingUpdate, bool) {
	if o == nil || o.CrossProductSampling == nil {
		return nil, false
	}
	return o.CrossProductSampling, true
}

// HasCrossProductSampling returns a boolean if a field has been set.
func (o *RumHardcodedRetentionFilterUpdateAttributes) HasCrossProductSampling() bool {
	return o != nil && o.CrossProductSampling != nil
}

// SetCrossProductSampling gets a reference to the given RumHardcodedCrossProductSamplingUpdate and assigns it to the CrossProductSampling field.
func (o *RumHardcodedRetentionFilterUpdateAttributes) SetCrossProductSampling(v RumHardcodedCrossProductSamplingUpdate) {
	o.CrossProductSampling = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumHardcodedRetentionFilterUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CrossProductSampling != nil {
		toSerialize["cross_product_sampling"] = o.CrossProductSampling
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumHardcodedRetentionFilterUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CrossProductSampling *RumHardcodedCrossProductSamplingUpdate `json:"cross_product_sampling,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cross_product_sampling"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.CrossProductSampling != nil && all.CrossProductSampling.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CrossProductSampling = all.CrossProductSampling

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
