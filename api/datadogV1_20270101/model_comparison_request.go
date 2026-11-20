// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ComparisonRequest Test-only generated schema for the dashboard versioning proof of concept.
type ComparisonRequest struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Directionality *ChangeDirectionality `json:"directionality,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Duration ChangeDuration `json:"duration"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Layout *ChangeLayout `json:"layout,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type *ChangeType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewComparisonRequest instantiates a new ComparisonRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewComparisonRequest(duration ChangeDuration) *ComparisonRequest {
	this := ComparisonRequest{}
	this.Duration = duration
	return &this
}

// NewComparisonRequestWithDefaults instantiates a new ComparisonRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewComparisonRequestWithDefaults() *ComparisonRequest {
	this := ComparisonRequest{}
	return &this
}

// GetDirectionality returns the Directionality field value if set, zero value otherwise.
func (o *ComparisonRequest) GetDirectionality() ChangeDirectionality {
	if o == nil || o.Directionality == nil {
		var ret ChangeDirectionality
		return ret
	}
	return *o.Directionality
}

// GetDirectionalityOk returns a tuple with the Directionality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComparisonRequest) GetDirectionalityOk() (*ChangeDirectionality, bool) {
	if o == nil || o.Directionality == nil {
		return nil, false
	}
	return o.Directionality, true
}

// HasDirectionality returns a boolean if a field has been set.
func (o *ComparisonRequest) HasDirectionality() bool {
	return o != nil && o.Directionality != nil
}

// SetDirectionality gets a reference to the given ChangeDirectionality and assigns it to the Directionality field.
func (o *ComparisonRequest) SetDirectionality(v ChangeDirectionality) {
	o.Directionality = &v
}

// GetDuration returns the Duration field value.
func (o *ComparisonRequest) GetDuration() ChangeDuration {
	if o == nil {
		var ret ChangeDuration
		return ret
	}
	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *ComparisonRequest) GetDurationOk() (*ChangeDuration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value.
func (o *ComparisonRequest) SetDuration(v ChangeDuration) {
	o.Duration = v
}

// GetLayout returns the Layout field value if set, zero value otherwise.
func (o *ComparisonRequest) GetLayout() ChangeLayout {
	if o == nil || o.Layout == nil {
		var ret ChangeLayout
		return ret
	}
	return *o.Layout
}

// GetLayoutOk returns a tuple with the Layout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComparisonRequest) GetLayoutOk() (*ChangeLayout, bool) {
	if o == nil || o.Layout == nil {
		return nil, false
	}
	return o.Layout, true
}

// HasLayout returns a boolean if a field has been set.
func (o *ComparisonRequest) HasLayout() bool {
	return o != nil && o.Layout != nil
}

// SetLayout gets a reference to the given ChangeLayout and assigns it to the Layout field.
func (o *ComparisonRequest) SetLayout(v ChangeLayout) {
	o.Layout = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ComparisonRequest) GetType() ChangeType {
	if o == nil || o.Type == nil {
		var ret ChangeType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComparisonRequest) GetTypeOk() (*ChangeType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ComparisonRequest) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given ChangeType and assigns it to the Type field.
func (o *ComparisonRequest) SetType(v ChangeType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ComparisonRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Directionality != nil {
		toSerialize["directionality"] = o.Directionality
	}
	toSerialize["duration"] = o.Duration
	if o.Layout != nil {
		toSerialize["layout"] = o.Layout
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ComparisonRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Directionality *ChangeDirectionality `json:"directionality,omitempty"`
		Duration       *ChangeDuration       `json:"duration"`
		Layout         *ChangeLayout         `json:"layout,omitempty"`
		Type           *ChangeType           `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Duration == nil {
		return fmt.Errorf("required field duration missing")
	}

	hasInvalidField := false
	if all.Directionality != nil && !all.Directionality.IsValid() {
		hasInvalidField = true
	} else {
		o.Directionality = all.Directionality
	}
	o.Duration = *all.Duration
	if all.Layout != nil && !all.Layout.IsValid() {
		hasInvalidField = true
	} else {
		o.Layout = all.Layout
	}
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
