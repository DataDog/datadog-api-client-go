// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ColumnSortObject Sort object
type ColumnSortObject struct {
	// Limit number of items displayed
	Count *float64 `json:"count,omitempty"`
	// Order criteria
	OrderBy []ColumnOrderObject `json:"order_by,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewColumnSortObject instantiates a new ColumnSortObject object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewColumnSortObject() *ColumnSortObject {
	this := ColumnSortObject{}
	return &this
}

// NewColumnSortObjectWithDefaults instantiates a new ColumnSortObject object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewColumnSortObjectWithDefaults() *ColumnSortObject {
	this := ColumnSortObject{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ColumnSortObject) GetCount() float64 {
	if o == nil || o.Count == nil {
		var ret float64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ColumnSortObject) GetCountOk() (*float64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ColumnSortObject) HasCount() bool {
	return o != nil && o.Count != nil
}

// SetCount gets a reference to the given float64 and assigns it to the Count field.
func (o *ColumnSortObject) SetCount(v float64) {
	o.Count = &v
}

// GetOrderBy returns the OrderBy field value if set, zero value otherwise.
func (o *ColumnSortObject) GetOrderBy() []ColumnOrderObject {
	if o == nil || o.OrderBy == nil {
		var ret []ColumnOrderObject
		return ret
	}
	return o.OrderBy
}

// GetOrderByOk returns a tuple with the OrderBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ColumnSortObject) GetOrderByOk() (*[]ColumnOrderObject, bool) {
	if o == nil || o.OrderBy == nil {
		return nil, false
	}
	return &o.OrderBy, true
}

// HasOrderBy returns a boolean if a field has been set.
func (o *ColumnSortObject) HasOrderBy() bool {
	return o != nil && o.OrderBy != nil
}

// SetOrderBy gets a reference to the given []ColumnOrderObject and assigns it to the OrderBy field.
func (o *ColumnSortObject) SetOrderBy(v []ColumnOrderObject) {
	o.OrderBy = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ColumnSortObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.OrderBy != nil {
		toSerialize["order_by"] = o.OrderBy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ColumnSortObject) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Count   *float64            `json:"count,omitempty"`
		OrderBy []ColumnOrderObject `json:"order_by,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"count", "order_by"})
	} else {
		return err
	}
	o.Count = all.Count
	o.OrderBy = all.OrderBy

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
