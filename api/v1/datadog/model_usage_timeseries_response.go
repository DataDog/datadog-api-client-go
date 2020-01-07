/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"bytes"
	"encoding/json"
)

// UsageTimeseriesResponse struct for UsageTimeseriesResponse
type UsageTimeseriesResponse struct {
	Usage *[]UsageTimeseriesHour `json:"usage,omitempty"`
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *UsageTimeseriesResponse) GetUsage() []UsageTimeseriesHour {
	if o == nil || o.Usage == nil {
		var ret []UsageTimeseriesHour
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UsageTimeseriesResponse) GetUsageOk() ([]UsageTimeseriesHour, bool) {
	if o == nil || o.Usage == nil {
		var ret []UsageTimeseriesHour
		return ret, false
	}
	return *o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *UsageTimeseriesResponse) HasUsage() bool {
	if o != nil && o.Usage != nil {
		return true
	}

	return false
}

// SetUsage gets a reference to the given []UsageTimeseriesHour and assigns it to the Usage field.
func (o *UsageTimeseriesResponse) SetUsage(v []UsageTimeseriesHour) {
	o.Usage = &v
}

type NullableUsageTimeseriesResponse struct {
	Value        UsageTimeseriesResponse
	ExplicitNull bool
}

func (v NullableUsageTimeseriesResponse) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableUsageTimeseriesResponse) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
