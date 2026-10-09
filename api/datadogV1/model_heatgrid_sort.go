// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridSort Ordering of the heatgrid rows.
type HeatgridSort struct {
	// Display groups as flat rows.
	NestingDisplay HeatgridNestingDisplay `json:"nesting_display"`
	// Sort rows by aggregated value or group label.
	SortBy HeatgridSortBy `json:"sort_by"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewHeatgridSort instantiates a new HeatgridSort object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHeatgridSort(nestingDisplay HeatgridNestingDisplay, sortBy HeatgridSortBy) *HeatgridSort {
	this := HeatgridSort{}
	this.NestingDisplay = nestingDisplay
	this.SortBy = sortBy
	return &this
}

// NewHeatgridSortWithDefaults instantiates a new HeatgridSort object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHeatgridSortWithDefaults() *HeatgridSort {
	this := HeatgridSort{}
	return &this
}

// GetNestingDisplay returns the NestingDisplay field value.
func (o *HeatgridSort) GetNestingDisplay() HeatgridNestingDisplay {
	if o == nil {
		var ret HeatgridNestingDisplay
		return ret
	}
	return o.NestingDisplay
}

// GetNestingDisplayOk returns a tuple with the NestingDisplay field value
// and a boolean to check if the value has been set.
func (o *HeatgridSort) GetNestingDisplayOk() (*HeatgridNestingDisplay, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NestingDisplay, true
}

// SetNestingDisplay sets field value.
func (o *HeatgridSort) SetNestingDisplay(v HeatgridNestingDisplay) {
	o.NestingDisplay = v
}

// GetSortBy returns the SortBy field value.
func (o *HeatgridSort) GetSortBy() HeatgridSortBy {
	if o == nil {
		var ret HeatgridSortBy
		return ret
	}
	return o.SortBy
}

// GetSortByOk returns a tuple with the SortBy field value
// and a boolean to check if the value has been set.
func (o *HeatgridSort) GetSortByOk() (*HeatgridSortBy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SortBy, true
}

// SetSortBy sets field value.
func (o *HeatgridSort) SetSortBy(v HeatgridSortBy) {
	o.SortBy = v
}

// MarshalJSON serializes the struct using spec logic.
func (o HeatgridSort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["nesting_display"] = o.NestingDisplay
	toSerialize["sort_by"] = o.SortBy
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HeatgridSort) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NestingDisplay *HeatgridNestingDisplay `json:"nesting_display"`
		SortBy         *HeatgridSortBy         `json:"sort_by"`
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
