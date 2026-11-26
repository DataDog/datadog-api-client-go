// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetTime20270101OneOf Test-only generated schema for the dashboard versioning proof of concept.
type WidgetTime20270101OneOf struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	From float64 `json:"from"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	HideIncompleteCostData *bool `json:"hide_incomplete_cost_data,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	To float64 `json:"to"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type WidgetNewFixedSpanType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewWidgetTime20270101OneOf instantiates a new WidgetTime20270101OneOf object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWidgetTime20270101OneOf(from float64, to float64, typeVar WidgetNewFixedSpanType) *WidgetTime20270101OneOf {
	this := WidgetTime20270101OneOf{}
	this.From = from
	this.To = to
	this.Type = typeVar
	return &this
}

// NewWidgetTime20270101OneOfWithDefaults instantiates a new WidgetTime20270101OneOf object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWidgetTime20270101OneOfWithDefaults() *WidgetTime20270101OneOf {
	this := WidgetTime20270101OneOf{}
	return &this
}

// GetFrom returns the From field value.
func (o *WidgetTime20270101OneOf) GetFrom() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *WidgetTime20270101OneOf) GetFromOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value.
func (o *WidgetTime20270101OneOf) SetFrom(v float64) {
	o.From = v
}

// GetHideIncompleteCostData returns the HideIncompleteCostData field value if set, zero value otherwise.
func (o *WidgetTime20270101OneOf) GetHideIncompleteCostData() bool {
	if o == nil || o.HideIncompleteCostData == nil {
		var ret bool
		return ret
	}
	return *o.HideIncompleteCostData
}

// GetHideIncompleteCostDataOk returns a tuple with the HideIncompleteCostData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetTime20270101OneOf) GetHideIncompleteCostDataOk() (*bool, bool) {
	if o == nil || o.HideIncompleteCostData == nil {
		return nil, false
	}
	return o.HideIncompleteCostData, true
}

// HasHideIncompleteCostData returns a boolean if a field has been set.
func (o *WidgetTime20270101OneOf) HasHideIncompleteCostData() bool {
	return o != nil && o.HideIncompleteCostData != nil
}

// SetHideIncompleteCostData gets a reference to the given bool and assigns it to the HideIncompleteCostData field.
func (o *WidgetTime20270101OneOf) SetHideIncompleteCostData(v bool) {
	o.HideIncompleteCostData = &v
}

// GetTo returns the To field value.
func (o *WidgetTime20270101OneOf) GetTo() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *WidgetTime20270101OneOf) GetToOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value.
func (o *WidgetTime20270101OneOf) SetTo(v float64) {
	o.To = v
}

// GetType returns the Type field value.
func (o *WidgetTime20270101OneOf) GetType() WidgetNewFixedSpanType {
	if o == nil {
		var ret WidgetNewFixedSpanType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WidgetTime20270101OneOf) GetTypeOk() (*WidgetNewFixedSpanType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *WidgetTime20270101OneOf) SetType(v WidgetNewFixedSpanType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o WidgetTime20270101OneOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["from"] = o.From
	if o.HideIncompleteCostData != nil {
		toSerialize["hide_incomplete_cost_data"] = o.HideIncompleteCostData
	}
	toSerialize["to"] = o.To
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WidgetTime20270101OneOf) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		From                   *float64                `json:"from"`
		HideIncompleteCostData *bool                   `json:"hide_incomplete_cost_data,omitempty"`
		To                     *float64                `json:"to"`
		Type                   *WidgetNewFixedSpanType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.From == nil {
		return fmt.Errorf("required field from missing")
	}
	if all.To == nil {
		return fmt.Errorf("required field to missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	o.From = *all.From
	o.HideIncompleteCostData = all.HideIncompleteCostData
	o.To = *all.To
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
