// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SplitConfigSort Test-only generated schema for the dashboard versioning proof of concept.
type SplitConfigSort struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Compute *SplitSortCompute `json:"compute,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Order SplitConfigSortOrder `json:"order"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewSplitConfigSort instantiates a new SplitConfigSort object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSplitConfigSort(order SplitConfigSortOrder) *SplitConfigSort {
	this := SplitConfigSort{}
	this.Order = order
	return &this
}

// NewSplitConfigSortWithDefaults instantiates a new SplitConfigSort object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSplitConfigSortWithDefaults() *SplitConfigSort {
	this := SplitConfigSort{}
	return &this
}

// GetCompute returns the Compute field value if set, zero value otherwise.
func (o *SplitConfigSort) GetCompute() SplitSortCompute {
	if o == nil || o.Compute == nil {
		var ret SplitSortCompute
		return ret
	}
	return *o.Compute
}

// GetComputeOk returns a tuple with the Compute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitConfigSort) GetComputeOk() (*SplitSortCompute, bool) {
	if o == nil || o.Compute == nil {
		return nil, false
	}
	return o.Compute, true
}

// HasCompute returns a boolean if a field has been set.
func (o *SplitConfigSort) HasCompute() bool {
	return o != nil && o.Compute != nil
}

// SetCompute gets a reference to the given SplitSortCompute and assigns it to the Compute field.
func (o *SplitConfigSort) SetCompute(v SplitSortCompute) {
	o.Compute = &v
}

// GetOrder returns the Order field value.
func (o *SplitConfigSort) GetOrder() SplitConfigSortOrder {
	if o == nil {
		var ret SplitConfigSortOrder
		return ret
	}
	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *SplitConfigSort) GetOrderOk() (*SplitConfigSortOrder, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value.
func (o *SplitConfigSort) SetOrder(v SplitConfigSortOrder) {
	o.Order = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SplitConfigSort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Compute != nil {
		toSerialize["compute"] = o.Compute
	}
	toSerialize["order"] = o.Order
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SplitConfigSort) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Compute *SplitSortCompute     `json:"compute,omitempty"`
		Order   *SplitConfigSortOrder `json:"order"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Order == nil {
		return fmt.Errorf("required field order missing")
	}

	hasInvalidField := false
	if all.Compute != nil && all.Compute.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Compute = all.Compute
	if !all.Order.IsValid() {
		hasInvalidField = true
	} else {
		o.Order = *all.Order
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
