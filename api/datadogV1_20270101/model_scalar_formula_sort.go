// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScalarFormulaSort Test-only generated schema for the dashboard versioning proof of concept.
type ScalarFormulaSort struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Count *float64 `json:"count,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Offset *float64 `json:"offset,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	OrderBy []ScalarFormulaOrderBy `json:"order_by,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewScalarFormulaSort instantiates a new ScalarFormulaSort object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScalarFormulaSort() *ScalarFormulaSort {
	this := ScalarFormulaSort{}
	return &this
}

// NewScalarFormulaSortWithDefaults instantiates a new ScalarFormulaSort object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScalarFormulaSortWithDefaults() *ScalarFormulaSort {
	this := ScalarFormulaSort{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ScalarFormulaSort) GetCount() float64 {
	if o == nil || o.Count == nil {
		var ret float64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScalarFormulaSort) GetCountOk() (*float64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ScalarFormulaSort) HasCount() bool {
	return o != nil && o.Count != nil
}

// SetCount gets a reference to the given float64 and assigns it to the Count field.
func (o *ScalarFormulaSort) SetCount(v float64) {
	o.Count = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *ScalarFormulaSort) GetOffset() float64 {
	if o == nil || o.Offset == nil {
		var ret float64
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScalarFormulaSort) GetOffsetOk() (*float64, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *ScalarFormulaSort) HasOffset() bool {
	return o != nil && o.Offset != nil
}

// SetOffset gets a reference to the given float64 and assigns it to the Offset field.
func (o *ScalarFormulaSort) SetOffset(v float64) {
	o.Offset = &v
}

// GetOrderBy returns the OrderBy field value if set, zero value otherwise.
func (o *ScalarFormulaSort) GetOrderBy() []ScalarFormulaOrderBy {
	if o == nil || o.OrderBy == nil {
		var ret []ScalarFormulaOrderBy
		return ret
	}
	return o.OrderBy
}

// GetOrderByOk returns a tuple with the OrderBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScalarFormulaSort) GetOrderByOk() (*[]ScalarFormulaOrderBy, bool) {
	if o == nil || o.OrderBy == nil {
		return nil, false
	}
	return &o.OrderBy, true
}

// HasOrderBy returns a boolean if a field has been set.
func (o *ScalarFormulaSort) HasOrderBy() bool {
	return o != nil && o.OrderBy != nil
}

// SetOrderBy gets a reference to the given []ScalarFormulaOrderBy and assigns it to the OrderBy field.
func (o *ScalarFormulaSort) SetOrderBy(v []ScalarFormulaOrderBy) {
	o.OrderBy = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScalarFormulaSort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
	if o.OrderBy != nil {
		toSerialize["order_by"] = o.OrderBy
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScalarFormulaSort) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Count   *float64               `json:"count,omitempty"`
		Offset  *float64               `json:"offset,omitempty"`
		OrderBy []ScalarFormulaOrderBy `json:"order_by,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.Count = all.Count
	o.Offset = all.Offset
	o.OrderBy = all.OrderBy

	return nil
}
