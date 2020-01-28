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

// SyntheticsBrowserTestResultFull struct for SyntheticsBrowserTestResultFull
type SyntheticsBrowserTestResultFull struct {
	Check        *SyntheticsApiTestResultFullCheck `json:"check,omitempty"`
	CheckTime    *float64                          `json:"check_time,omitempty"`
	CheckVersion *int64                            `json:"check_version,omitempty"`
	ProbeDc      *string                           `json:"probe_dc,omitempty"`
	Result       *SyntheticsBrowserTestResultData  `json:"result,omitempty"`
	ResultId     *string                           `json:"result_id,omitempty"`
	Status       *SyntheticsTestMonitorStatus      `json:"status,omitempty"`
}

// GetCheck returns the Check field value if set, zero value otherwise.
func (o *SyntheticsBrowserTestResultFull) GetCheck() SyntheticsApiTestResultFullCheck {
	if o == nil || o.Check == nil {
		var ret SyntheticsApiTestResultFullCheck
		return ret
	}
	return *o.Check
}

// GetCheckOk returns a tuple with the Check field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsBrowserTestResultFull) GetCheckOk() (SyntheticsApiTestResultFullCheck, bool) {
	if o == nil || o.Check == nil {
		var ret SyntheticsApiTestResultFullCheck
		return ret, false
	}
	return *o.Check, true
}

// HasCheck returns a boolean if a field has been set.
func (o *SyntheticsBrowserTestResultFull) HasCheck() bool {
	if o != nil && o.Check != nil {
		return true
	}

	return false
}

// SetCheck gets a reference to the given SyntheticsApiTestResultFullCheck and assigns it to the Check field.
func (o *SyntheticsBrowserTestResultFull) SetCheck(v SyntheticsApiTestResultFullCheck) {
	o.Check = &v
}

// GetCheckTime returns the CheckTime field value if set, zero value otherwise.
func (o *SyntheticsBrowserTestResultFull) GetCheckTime() float64 {
	if o == nil || o.CheckTime == nil {
		var ret float64
		return ret
	}
	return *o.CheckTime
}

// GetCheckTimeOk returns a tuple with the CheckTime field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsBrowserTestResultFull) GetCheckTimeOk() (float64, bool) {
	if o == nil || o.CheckTime == nil {
		var ret float64
		return ret, false
	}
	return *o.CheckTime, true
}

// HasCheckTime returns a boolean if a field has been set.
func (o *SyntheticsBrowserTestResultFull) HasCheckTime() bool {
	if o != nil && o.CheckTime != nil {
		return true
	}

	return false
}

// SetCheckTime gets a reference to the given float64 and assigns it to the CheckTime field.
func (o *SyntheticsBrowserTestResultFull) SetCheckTime(v float64) {
	o.CheckTime = &v
}

// GetCheckVersion returns the CheckVersion field value if set, zero value otherwise.
func (o *SyntheticsBrowserTestResultFull) GetCheckVersion() int64 {
	if o == nil || o.CheckVersion == nil {
		var ret int64
		return ret
	}
	return *o.CheckVersion
}

// GetCheckVersionOk returns a tuple with the CheckVersion field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsBrowserTestResultFull) GetCheckVersionOk() (int64, bool) {
	if o == nil || o.CheckVersion == nil {
		var ret int64
		return ret, false
	}
	return *o.CheckVersion, true
}

// HasCheckVersion returns a boolean if a field has been set.
func (o *SyntheticsBrowserTestResultFull) HasCheckVersion() bool {
	if o != nil && o.CheckVersion != nil {
		return true
	}

	return false
}

// SetCheckVersion gets a reference to the given int64 and assigns it to the CheckVersion field.
func (o *SyntheticsBrowserTestResultFull) SetCheckVersion(v int64) {
	o.CheckVersion = &v
}

// GetProbeDc returns the ProbeDc field value if set, zero value otherwise.
func (o *SyntheticsBrowserTestResultFull) GetProbeDc() string {
	if o == nil || o.ProbeDc == nil {
		var ret string
		return ret
	}
	return *o.ProbeDc
}

// GetProbeDcOk returns a tuple with the ProbeDc field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsBrowserTestResultFull) GetProbeDcOk() (string, bool) {
	if o == nil || o.ProbeDc == nil {
		var ret string
		return ret, false
	}
	return *o.ProbeDc, true
}

// HasProbeDc returns a boolean if a field has been set.
func (o *SyntheticsBrowserTestResultFull) HasProbeDc() bool {
	if o != nil && o.ProbeDc != nil {
		return true
	}

	return false
}

// SetProbeDc gets a reference to the given string and assigns it to the ProbeDc field.
func (o *SyntheticsBrowserTestResultFull) SetProbeDc(v string) {
	o.ProbeDc = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *SyntheticsBrowserTestResultFull) GetResult() SyntheticsBrowserTestResultData {
	if o == nil || o.Result == nil {
		var ret SyntheticsBrowserTestResultData
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsBrowserTestResultFull) GetResultOk() (SyntheticsBrowserTestResultData, bool) {
	if o == nil || o.Result == nil {
		var ret SyntheticsBrowserTestResultData
		return ret, false
	}
	return *o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *SyntheticsBrowserTestResultFull) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given SyntheticsBrowserTestResultData and assigns it to the Result field.
func (o *SyntheticsBrowserTestResultFull) SetResult(v SyntheticsBrowserTestResultData) {
	o.Result = &v
}

// GetResultId returns the ResultId field value if set, zero value otherwise.
func (o *SyntheticsBrowserTestResultFull) GetResultId() string {
	if o == nil || o.ResultId == nil {
		var ret string
		return ret
	}
	return *o.ResultId
}

// GetResultIdOk returns a tuple with the ResultId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsBrowserTestResultFull) GetResultIdOk() (string, bool) {
	if o == nil || o.ResultId == nil {
		var ret string
		return ret, false
	}
	return *o.ResultId, true
}

// HasResultId returns a boolean if a field has been set.
func (o *SyntheticsBrowserTestResultFull) HasResultId() bool {
	if o != nil && o.ResultId != nil {
		return true
	}

	return false
}

// SetResultId gets a reference to the given string and assigns it to the ResultId field.
func (o *SyntheticsBrowserTestResultFull) SetResultId(v string) {
	o.ResultId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SyntheticsBrowserTestResultFull) GetStatus() SyntheticsTestMonitorStatus {
	if o == nil || o.Status == nil {
		var ret SyntheticsTestMonitorStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsBrowserTestResultFull) GetStatusOk() (SyntheticsTestMonitorStatus, bool) {
	if o == nil || o.Status == nil {
		var ret SyntheticsTestMonitorStatus
		return ret, false
	}
	return *o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SyntheticsBrowserTestResultFull) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given SyntheticsTestMonitorStatus and assigns it to the Status field.
func (o *SyntheticsBrowserTestResultFull) SetStatus(v SyntheticsTestMonitorStatus) {
	o.Status = &v
}

type NullableSyntheticsBrowserTestResultFull struct {
	Value        SyntheticsBrowserTestResultFull
	ExplicitNull bool
}

func (v NullableSyntheticsBrowserTestResultFull) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableSyntheticsBrowserTestResultFull) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
