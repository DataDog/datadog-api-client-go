/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// SyntheticsAPITestResultFull struct for SyntheticsAPITestResultFull
type SyntheticsAPITestResultFull struct {
	Check        *SyntheticsAPITestResultFullCheck `json:"check,omitempty"`
	CheckTime    *float64                          `json:"check_time,omitempty"`
	CheckVersion *int64                            `json:"check_version,omitempty"`
	ProbeDc      *string                           `json:"probe_dc,omitempty"`
	Result       *SyntheticsAPITestResultData      `json:"result,omitempty"`
	ResultId     *string                           `json:"result_id,omitempty"`
	Status       *SyntheticsTestMonitorStatus      `json:"status,omitempty"`
}

// NewSyntheticsAPITestResultFull instantiates a new SyntheticsAPITestResultFull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticsAPITestResultFull() *SyntheticsAPITestResultFull {
	this := SyntheticsAPITestResultFull{}
	return &this
}

// NewSyntheticsAPITestResultFullWithDefaults instantiates a new SyntheticsAPITestResultFull object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticsAPITestResultFullWithDefaults() *SyntheticsAPITestResultFull {
	this := SyntheticsAPITestResultFull{}
	return &this
}

// GetCheck returns the Check field value if set, zero value otherwise.
func (o *SyntheticsAPITestResultFull) GetCheck() SyntheticsAPITestResultFullCheck {
	if o == nil || o.Check == nil {
		var ret SyntheticsAPITestResultFullCheck
		return ret
	}
	return *o.Check
}

// GetCheckOk returns a tuple with the Check field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsAPITestResultFull) GetCheckOk() (*SyntheticsAPITestResultFullCheck, bool) {
	if o == nil || o.Check == nil {
		return nil, false
	}
	return o.Check, true
}

// HasCheck returns a boolean if a field has been set.
func (o *SyntheticsAPITestResultFull) HasCheck() bool {
	if o != nil && o.Check != nil {
		return true
	}

	return false
}

// SetCheck gets a reference to the given SyntheticsAPITestResultFullCheck and assigns it to the Check field.
func (o *SyntheticsAPITestResultFull) SetCheck(v SyntheticsAPITestResultFullCheck) {
	o.Check = &v
}

// GetCheckTime returns the CheckTime field value if set, zero value otherwise.
func (o *SyntheticsAPITestResultFull) GetCheckTime() float64 {
	if o == nil || o.CheckTime == nil {
		var ret float64
		return ret
	}
	return *o.CheckTime
}

// GetCheckTimeOk returns a tuple with the CheckTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsAPITestResultFull) GetCheckTimeOk() (*float64, bool) {
	if o == nil || o.CheckTime == nil {
		return nil, false
	}
	return o.CheckTime, true
}

// HasCheckTime returns a boolean if a field has been set.
func (o *SyntheticsAPITestResultFull) HasCheckTime() bool {
	if o != nil && o.CheckTime != nil {
		return true
	}

	return false
}

// SetCheckTime gets a reference to the given float64 and assigns it to the CheckTime field.
func (o *SyntheticsAPITestResultFull) SetCheckTime(v float64) {
	o.CheckTime = &v
}

// GetCheckVersion returns the CheckVersion field value if set, zero value otherwise.
func (o *SyntheticsAPITestResultFull) GetCheckVersion() int64 {
	if o == nil || o.CheckVersion == nil {
		var ret int64
		return ret
	}
	return *o.CheckVersion
}

// GetCheckVersionOk returns a tuple with the CheckVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsAPITestResultFull) GetCheckVersionOk() (*int64, bool) {
	if o == nil || o.CheckVersion == nil {
		return nil, false
	}
	return o.CheckVersion, true
}

// HasCheckVersion returns a boolean if a field has been set.
func (o *SyntheticsAPITestResultFull) HasCheckVersion() bool {
	if o != nil && o.CheckVersion != nil {
		return true
	}

	return false
}

// SetCheckVersion gets a reference to the given int64 and assigns it to the CheckVersion field.
func (o *SyntheticsAPITestResultFull) SetCheckVersion(v int64) {
	o.CheckVersion = &v
}

// GetProbeDc returns the ProbeDc field value if set, zero value otherwise.
func (o *SyntheticsAPITestResultFull) GetProbeDc() string {
	if o == nil || o.ProbeDc == nil {
		var ret string
		return ret
	}
	return *o.ProbeDc
}

// GetProbeDcOk returns a tuple with the ProbeDc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsAPITestResultFull) GetProbeDcOk() (*string, bool) {
	if o == nil || o.ProbeDc == nil {
		return nil, false
	}
	return o.ProbeDc, true
}

// HasProbeDc returns a boolean if a field has been set.
func (o *SyntheticsAPITestResultFull) HasProbeDc() bool {
	if o != nil && o.ProbeDc != nil {
		return true
	}

	return false
}

// SetProbeDc gets a reference to the given string and assigns it to the ProbeDc field.
func (o *SyntheticsAPITestResultFull) SetProbeDc(v string) {
	o.ProbeDc = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *SyntheticsAPITestResultFull) GetResult() SyntheticsAPITestResultData {
	if o == nil || o.Result == nil {
		var ret SyntheticsAPITestResultData
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsAPITestResultFull) GetResultOk() (*SyntheticsAPITestResultData, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *SyntheticsAPITestResultFull) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given SyntheticsAPITestResultData and assigns it to the Result field.
func (o *SyntheticsAPITestResultFull) SetResult(v SyntheticsAPITestResultData) {
	o.Result = &v
}

// GetResultId returns the ResultId field value if set, zero value otherwise.
func (o *SyntheticsAPITestResultFull) GetResultId() string {
	if o == nil || o.ResultId == nil {
		var ret string
		return ret
	}
	return *o.ResultId
}

// GetResultIdOk returns a tuple with the ResultId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsAPITestResultFull) GetResultIdOk() (*string, bool) {
	if o == nil || o.ResultId == nil {
		return nil, false
	}
	return o.ResultId, true
}

// HasResultId returns a boolean if a field has been set.
func (o *SyntheticsAPITestResultFull) HasResultId() bool {
	if o != nil && o.ResultId != nil {
		return true
	}

	return false
}

// SetResultId gets a reference to the given string and assigns it to the ResultId field.
func (o *SyntheticsAPITestResultFull) SetResultId(v string) {
	o.ResultId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SyntheticsAPITestResultFull) GetStatus() SyntheticsTestMonitorStatus {
	if o == nil || o.Status == nil {
		var ret SyntheticsTestMonitorStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsAPITestResultFull) GetStatusOk() (*SyntheticsTestMonitorStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SyntheticsAPITestResultFull) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given SyntheticsTestMonitorStatus and assigns it to the Status field.
func (o *SyntheticsAPITestResultFull) SetStatus(v SyntheticsTestMonitorStatus) {
	o.Status = &v
}

func (o SyntheticsAPITestResultFull) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Check != nil {
		toSerialize["check"] = o.Check
	}
	if o.CheckTime != nil {
		toSerialize["check_time"] = o.CheckTime
	}
	if o.CheckVersion != nil {
		toSerialize["check_version"] = o.CheckVersion
	}
	if o.ProbeDc != nil {
		toSerialize["probe_dc"] = o.ProbeDc
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	if o.ResultId != nil {
		toSerialize["result_id"] = o.ResultId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableSyntheticsAPITestResultFull struct {
	value *SyntheticsAPITestResultFull
	isSet bool
}

func (v NullableSyntheticsAPITestResultFull) Get() *SyntheticsAPITestResultFull {
	return v.value
}

func (v *NullableSyntheticsAPITestResultFull) Set(val *SyntheticsAPITestResultFull) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsAPITestResultFull) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsAPITestResultFull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsAPITestResultFull(val *SyntheticsAPITestResultFull) *NullableSyntheticsAPITestResultFull {
	return &NullableSyntheticsAPITestResultFull{value: val, isSet: true}
}

func (v NullableSyntheticsAPITestResultFull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsAPITestResultFull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
