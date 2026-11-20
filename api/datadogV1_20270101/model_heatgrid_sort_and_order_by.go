// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridSortAndOrderBy Test-only generated schema for the dashboard versioning proof of concept.
type HeatgridSortAndOrderBy struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	NestingDisplay HeatgridSortAndOrderByNestingDisplay `json:"nesting_display"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	SortBy FlatSortBy `json:"sort_by"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewHeatgridSortAndOrderBy instantiates a new HeatgridSortAndOrderBy object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHeatgridSortAndOrderBy(nestingDisplay HeatgridSortAndOrderByNestingDisplay, sortBy FlatSortBy) *HeatgridSortAndOrderBy {
	this := HeatgridSortAndOrderBy{}
	this.NestingDisplay = nestingDisplay
	this.SortBy = sortBy
	return &this
}

// NewHeatgridSortAndOrderByWithDefaults instantiates a new HeatgridSortAndOrderBy object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHeatgridSortAndOrderByWithDefaults() *HeatgridSortAndOrderBy {
	this := HeatgridSortAndOrderBy{}
	return &this
}

// GetNestingDisplay returns the NestingDisplay field value.
func (o *HeatgridSortAndOrderBy) GetNestingDisplay() HeatgridSortAndOrderByNestingDisplay {
	if o == nil {
		var ret HeatgridSortAndOrderByNestingDisplay
		return ret
	}
	return o.NestingDisplay
}

// GetNestingDisplayOk returns a tuple with the NestingDisplay field value
// and a boolean to check if the value has been set.
func (o *HeatgridSortAndOrderBy) GetNestingDisplayOk() (*HeatgridSortAndOrderByNestingDisplay, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NestingDisplay, true
}

// SetNestingDisplay sets field value.
func (o *HeatgridSortAndOrderBy) SetNestingDisplay(v HeatgridSortAndOrderByNestingDisplay) {
	o.NestingDisplay = v
}

// GetSortBy returns the SortBy field value.
func (o *HeatgridSortAndOrderBy) GetSortBy() FlatSortBy {
	if o == nil {
		var ret FlatSortBy
		return ret
	}
	return o.SortBy
}

// GetSortByOk returns a tuple with the SortBy field value
// and a boolean to check if the value has been set.
func (o *HeatgridSortAndOrderBy) GetSortByOk() (*FlatSortBy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SortBy, true
}

// SetSortBy sets field value.
func (o *HeatgridSortAndOrderBy) SetSortBy(v FlatSortBy) {
	o.SortBy = v
}

// MarshalJSON serializes the struct using spec logic.
func (o HeatgridSortAndOrderBy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["nesting_display"] = o.NestingDisplay
	toSerialize["sort_by"] = o.SortBy
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HeatgridSortAndOrderBy) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NestingDisplay *HeatgridSortAndOrderByNestingDisplay `json:"nesting_display"`
		SortBy         *FlatSortBy                           `json:"sort_by"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.NestingDisplay == nil {
		return fmt.Errorf("required field nesting_display missing")
	}
	if all.SortBy == nil {
		return fmt.Errorf("required field sort_by missing")
	}

	hasInvalidField := false
	if !all.NestingDisplay.IsValid() {
		hasInvalidField = true
	} else {
		o.NestingDisplay = *all.NestingDisplay
	}
	o.SortBy = *all.SortBy

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
