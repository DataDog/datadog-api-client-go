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

// SloThreshold struct for SloThreshold
type SloThreshold struct {
	// The target value for the service level indicator within the corresponding timeframe.
	Target float64 `json:"target"`
	// A string representation of the target that indicates its precision (e.g. \"99.9\"). It uses trailing zeros to show significant decimal places (e.g. \"98.00\"). Always included in service level objective responses. Ignored in create/update requests.
	TargetDisplay *string      `json:"target_display,omitempty"`
	Timeframe     SloTimeframe `json:"timeframe"`
	Warning       *float64     `json:"warning,omitempty"`
	// A string representation of the warning target (see the description of the \"target_display\" field for details). Included in service level objective responses if a warning target exists. Ignored in create/update requests.
	WarningDisplay *string `json:"warning_display,omitempty"`
}

// NewSloThreshold instantiates a new SloThreshold object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSloThreshold(target float64, timeframe SloTimeframe) *SloThreshold {
	this := SloThreshold{}
	this.Target = target
	this.Timeframe = timeframe
	return &this
}

// NewSloThresholdWithDefaults instantiates a new SloThreshold object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSloThresholdWithDefaults() *SloThreshold {
	this := SloThreshold{}
	return &this
}

// GetTarget returns the Target field value
func (o *SloThreshold) GetTarget() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Target
}

// SetTarget sets field value
func (o *SloThreshold) SetTarget(v float64) {
	o.Target = v
}

// GetTargetDisplay returns the TargetDisplay field value if set, zero value otherwise.
func (o *SloThreshold) GetTargetDisplay() string {
	if o == nil || o.TargetDisplay == nil {
		var ret string
		return ret
	}
	return *o.TargetDisplay
}

// GetTargetDisplayOk returns a tuple with the TargetDisplay field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SloThreshold) GetTargetDisplayOk() (string, bool) {
	if o == nil || o.TargetDisplay == nil {
		var ret string
		return ret, false
	}
	return *o.TargetDisplay, true
}

// HasTargetDisplay returns a boolean if a field has been set.
func (o *SloThreshold) HasTargetDisplay() bool {
	if o != nil && o.TargetDisplay != nil {
		return true
	}

	return false
}

// SetTargetDisplay gets a reference to the given string and assigns it to the TargetDisplay field.
func (o *SloThreshold) SetTargetDisplay(v string) {
	o.TargetDisplay = &v
}

// GetTimeframe returns the Timeframe field value
func (o *SloThreshold) GetTimeframe() SloTimeframe {
	if o == nil {
		var ret SloTimeframe
		return ret
	}

	return o.Timeframe
}

// SetTimeframe sets field value
func (o *SloThreshold) SetTimeframe(v SloTimeframe) {
	o.Timeframe = v
}

// GetWarning returns the Warning field value if set, zero value otherwise.
func (o *SloThreshold) GetWarning() float64 {
	if o == nil || o.Warning == nil {
		var ret float64
		return ret
	}
	return *o.Warning
}

// GetWarningOk returns a tuple with the Warning field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SloThreshold) GetWarningOk() (float64, bool) {
	if o == nil || o.Warning == nil {
		var ret float64
		return ret, false
	}
	return *o.Warning, true
}

// HasWarning returns a boolean if a field has been set.
func (o *SloThreshold) HasWarning() bool {
	if o != nil && o.Warning != nil {
		return true
	}

	return false
}

// SetWarning gets a reference to the given float64 and assigns it to the Warning field.
func (o *SloThreshold) SetWarning(v float64) {
	o.Warning = &v
}

// GetWarningDisplay returns the WarningDisplay field value if set, zero value otherwise.
func (o *SloThreshold) GetWarningDisplay() string {
	if o == nil || o.WarningDisplay == nil {
		var ret string
		return ret
	}
	return *o.WarningDisplay
}

// GetWarningDisplayOk returns a tuple with the WarningDisplay field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SloThreshold) GetWarningDisplayOk() (string, bool) {
	if o == nil || o.WarningDisplay == nil {
		var ret string
		return ret, false
	}
	return *o.WarningDisplay, true
}

// HasWarningDisplay returns a boolean if a field has been set.
func (o *SloThreshold) HasWarningDisplay() bool {
	if o != nil && o.WarningDisplay != nil {
		return true
	}

	return false
}

// SetWarningDisplay gets a reference to the given string and assigns it to the WarningDisplay field.
func (o *SloThreshold) SetWarningDisplay(v string) {
	o.WarningDisplay = &v
}

type NullableSloThreshold struct {
	Value        SloThreshold
	ExplicitNull bool
}

func (v NullableSloThreshold) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableSloThreshold) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
