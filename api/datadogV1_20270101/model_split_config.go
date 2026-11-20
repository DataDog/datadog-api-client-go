// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SplitConfig Test-only generated schema for the dashboard versioning proof of concept.
type SplitConfig struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	CalendarShift *CalendarShiftSplitConfig `json:"calendar_shift,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Limit float64 `json:"limit"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Sort SplitConfigSort `json:"sort"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	SplitDimensions []SplitDimension `json:"split_dimensions"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	StaticSplits [][]SplitVectorTagFilter `json:"static_splits,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewSplitConfig instantiates a new SplitConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSplitConfig(limit float64, sort SplitConfigSort, splitDimensions []SplitDimension) *SplitConfig {
	this := SplitConfig{}
	this.Limit = limit
	this.Sort = sort
	this.SplitDimensions = splitDimensions
	return &this
}

// NewSplitConfigWithDefaults instantiates a new SplitConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSplitConfigWithDefaults() *SplitConfig {
	this := SplitConfig{}
	return &this
}

// GetCalendarShift returns the CalendarShift field value if set, zero value otherwise.
func (o *SplitConfig) GetCalendarShift() CalendarShiftSplitConfig {
	if o == nil || o.CalendarShift == nil {
		var ret CalendarShiftSplitConfig
		return ret
	}
	return *o.CalendarShift
}

// GetCalendarShiftOk returns a tuple with the CalendarShift field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitConfig) GetCalendarShiftOk() (*CalendarShiftSplitConfig, bool) {
	if o == nil || o.CalendarShift == nil {
		return nil, false
	}
	return o.CalendarShift, true
}

// HasCalendarShift returns a boolean if a field has been set.
func (o *SplitConfig) HasCalendarShift() bool {
	return o != nil && o.CalendarShift != nil
}

// SetCalendarShift gets a reference to the given CalendarShiftSplitConfig and assigns it to the CalendarShift field.
func (o *SplitConfig) SetCalendarShift(v CalendarShiftSplitConfig) {
	o.CalendarShift = &v
}

// GetLimit returns the Limit field value.
func (o *SplitConfig) GetLimit() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *SplitConfig) GetLimitOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value.
func (o *SplitConfig) SetLimit(v float64) {
	o.Limit = v
}

// GetSort returns the Sort field value.
func (o *SplitConfig) GetSort() SplitConfigSort {
	if o == nil {
		var ret SplitConfigSort
		return ret
	}
	return o.Sort
}

// GetSortOk returns a tuple with the Sort field value
// and a boolean to check if the value has been set.
func (o *SplitConfig) GetSortOk() (*SplitConfigSort, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sort, true
}

// SetSort sets field value.
func (o *SplitConfig) SetSort(v SplitConfigSort) {
	o.Sort = v
}

// GetSplitDimensions returns the SplitDimensions field value.
func (o *SplitConfig) GetSplitDimensions() []SplitDimension {
	if o == nil {
		var ret []SplitDimension
		return ret
	}
	return o.SplitDimensions
}

// GetSplitDimensionsOk returns a tuple with the SplitDimensions field value
// and a boolean to check if the value has been set.
func (o *SplitConfig) GetSplitDimensionsOk() (*[]SplitDimension, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SplitDimensions, true
}

// SetSplitDimensions sets field value.
func (o *SplitConfig) SetSplitDimensions(v []SplitDimension) {
	o.SplitDimensions = v
}

// GetStaticSplits returns the StaticSplits field value if set, zero value otherwise.
func (o *SplitConfig) GetStaticSplits() [][]SplitVectorTagFilter {
	if o == nil || o.StaticSplits == nil {
		var ret [][]SplitVectorTagFilter
		return ret
	}
	return o.StaticSplits
}

// GetStaticSplitsOk returns a tuple with the StaticSplits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitConfig) GetStaticSplitsOk() (*[][]SplitVectorTagFilter, bool) {
	if o == nil || o.StaticSplits == nil {
		return nil, false
	}
	return &o.StaticSplits, true
}

// HasStaticSplits returns a boolean if a field has been set.
func (o *SplitConfig) HasStaticSplits() bool {
	return o != nil && o.StaticSplits != nil
}

// SetStaticSplits gets a reference to the given [][]SplitVectorTagFilter and assigns it to the StaticSplits field.
func (o *SplitConfig) SetStaticSplits(v [][]SplitVectorTagFilter) {
	o.StaticSplits = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SplitConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CalendarShift != nil {
		toSerialize["calendar_shift"] = o.CalendarShift
	}
	toSerialize["limit"] = o.Limit
	toSerialize["sort"] = o.Sort
	toSerialize["split_dimensions"] = o.SplitDimensions
	if o.StaticSplits != nil {
		toSerialize["static_splits"] = o.StaticSplits
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SplitConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CalendarShift   *CalendarShiftSplitConfig `json:"calendar_shift,omitempty"`
		Limit           *float64                  `json:"limit"`
		Sort            *SplitConfigSort          `json:"sort"`
		SplitDimensions *[]SplitDimension         `json:"split_dimensions"`
		StaticSplits    [][]SplitVectorTagFilter  `json:"static_splits,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Limit == nil {
		return fmt.Errorf("required field limit missing")
	}
	if all.Sort == nil {
		return fmt.Errorf("required field sort missing")
	}
	if all.SplitDimensions == nil {
		return fmt.Errorf("required field split_dimensions missing")
	}

	hasInvalidField := false
	if all.CalendarShift != nil && all.CalendarShift.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CalendarShift = all.CalendarShift
	o.Limit = *all.Limit
	if all.Sort.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Sort = *all.Sort
	o.SplitDimensions = *all.SplitDimensions
	o.StaticSplits = all.StaticSplits

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
