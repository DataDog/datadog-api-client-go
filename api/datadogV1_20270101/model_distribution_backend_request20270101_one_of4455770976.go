// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DistributionBackendRequest20270101OneOf4455770976 Test-only generated schema for the dashboard versioning proof of concept.
type DistributionBackendRequest20270101OneOf4455770976 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Style *RequestStyle `json:"style,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	SyntheticsBatches LogQueryDefinition `json:"synthetics_batches"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewDistributionBackendRequest20270101OneOf4455770976 instantiates a new DistributionBackendRequest20270101OneOf4455770976 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDistributionBackendRequest20270101OneOf4455770976(syntheticsBatches LogQueryDefinition) *DistributionBackendRequest20270101OneOf4455770976 {
	this := DistributionBackendRequest20270101OneOf4455770976{}
	this.SyntheticsBatches = syntheticsBatches
	return &this
}

// NewDistributionBackendRequest20270101OneOf4455770976WithDefaults instantiates a new DistributionBackendRequest20270101OneOf4455770976 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDistributionBackendRequest20270101OneOf4455770976WithDefaults() *DistributionBackendRequest20270101OneOf4455770976 {
	this := DistributionBackendRequest20270101OneOf4455770976{}
	return &this
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *DistributionBackendRequest20270101OneOf4455770976) GetStyle() RequestStyle {
	if o == nil || o.Style == nil {
		var ret RequestStyle
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionBackendRequest20270101OneOf4455770976) GetStyleOk() (*RequestStyle, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *DistributionBackendRequest20270101OneOf4455770976) HasStyle() bool {
	return o != nil && o.Style != nil
}

// SetStyle gets a reference to the given RequestStyle and assigns it to the Style field.
func (o *DistributionBackendRequest20270101OneOf4455770976) SetStyle(v RequestStyle) {
	o.Style = &v
}

// GetSyntheticsBatches returns the SyntheticsBatches field value.
func (o *DistributionBackendRequest20270101OneOf4455770976) GetSyntheticsBatches() LogQueryDefinition {
	if o == nil {
		var ret LogQueryDefinition
		return ret
	}
	return o.SyntheticsBatches
}

// GetSyntheticsBatchesOk returns a tuple with the SyntheticsBatches field value
// and a boolean to check if the value has been set.
func (o *DistributionBackendRequest20270101OneOf4455770976) GetSyntheticsBatchesOk() (*LogQueryDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SyntheticsBatches, true
}

// SetSyntheticsBatches sets field value.
func (o *DistributionBackendRequest20270101OneOf4455770976) SetSyntheticsBatches(v LogQueryDefinition) {
	o.SyntheticsBatches = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DistributionBackendRequest20270101OneOf4455770976) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Style != nil {
		toSerialize["style"] = o.Style
	}
	toSerialize["synthetics_batches"] = o.SyntheticsBatches
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DistributionBackendRequest20270101OneOf4455770976) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Style             *RequestStyle       `json:"style,omitempty"`
		SyntheticsBatches *LogQueryDefinition `json:"synthetics_batches"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.SyntheticsBatches == nil {
		return fmt.Errorf("required field synthetics_batches missing")
	}

	hasInvalidField := false
	if all.Style != nil && all.Style.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Style = all.Style
	if all.SyntheticsBatches.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SyntheticsBatches = *all.SyntheticsBatches

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
