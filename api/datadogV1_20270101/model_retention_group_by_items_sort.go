// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RetentionGroupByItemsSort Test-only generated schema for the dashboard versioning proof of concept.
type RetentionGroupByItemsSort struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Order *QuerySortOrder `json:"order,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewRetentionGroupByItemsSort instantiates a new RetentionGroupByItemsSort object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRetentionGroupByItemsSort() *RetentionGroupByItemsSort {
	this := RetentionGroupByItemsSort{}
	return &this
}

// NewRetentionGroupByItemsSortWithDefaults instantiates a new RetentionGroupByItemsSort object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRetentionGroupByItemsSortWithDefaults() *RetentionGroupByItemsSort {
	this := RetentionGroupByItemsSort{}
	return &this
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *RetentionGroupByItemsSort) GetOrder() QuerySortOrder {
	if o == nil || o.Order == nil {
		var ret QuerySortOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionGroupByItemsSort) GetOrderOk() (*QuerySortOrder, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *RetentionGroupByItemsSort) HasOrder() bool {
	return o != nil && o.Order != nil
}

// SetOrder gets a reference to the given QuerySortOrder and assigns it to the Order field.
func (o *RetentionGroupByItemsSort) SetOrder(v QuerySortOrder) {
	o.Order = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RetentionGroupByItemsSort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RetentionGroupByItemsSort) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Order *QuerySortOrder `json:"order,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	hasInvalidField := false
	if all.Order != nil && !all.Order.IsValid() {
		hasInvalidField = true
	} else {
		o.Order = all.Order
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
