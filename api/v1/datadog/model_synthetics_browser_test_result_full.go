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

// SyntheticsBrowserTestResultFull Object returned describing a browser test result.
type SyntheticsBrowserTestResultFull struct {
	Check *SyntheticsBrowserTestResultFullCheck `json:"check,omitempty"`
	// When the browser test was conducted.
	CheckTime *float64 `json:"check_time,omitempty"`
	// Version of the browser test used.
	CheckVersion *int64 `json:"check_version,omitempty"`
	// Location from which the browser test was performed.
	ProbeDc *string                          `json:"probe_dc,omitempty"`
	Result  *SyntheticsBrowserTestResultData `json:"result,omitempty"`
	// ID of the browser test result.
	ResultId *string                      `json:"result_id,omitempty"`
	Status   *SyntheticsTestMonitorStatus `json:"status,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:-`
}

// NewSyntheticsBrowserTestResultFull instantiates a new SyntheticsBrowserTestResultFull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticsBrowserTestResultFull() *SyntheticsBrowserTestResultFull {
	this := SyntheticsBrowserTestResultFull{}
	return &this
}

// NewSyntheticsBrowserTestResultFullWithDefaults instantiates a new SyntheticsBrowserTestResultFull object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticsBrowserTestResultFullWithDefaults() *SyntheticsBrowserTestResultFull {
	this := SyntheticsBrowserTestResultFull{}
	return &this
}

// GetCheck returns the Check field value if set, zero value otherwise.
func (o *SyntheticsBrowserTestResultFull) GetCheck() SyntheticsBrowserTestResultFullCheck {
	if o == nil || o.Check == nil {
		var ret SyntheticsBrowserTestResultFullCheck
		return ret
	}
	return *o.Check
}

// GetCheckOk returns a tuple with the Check field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsBrowserTestResultFull) GetCheckOk() (*SyntheticsBrowserTestResultFullCheck, bool) {
	if o == nil || o.Check == nil {
		return nil, false
	}
	return o.Check, true
}

// HasCheck returns a boolean if a field has been set.
func (o *SyntheticsBrowserTestResultFull) HasCheck() bool {
	if o != nil && o.Check != nil {
		return true
	}

	return false
}

// SetCheck gets a reference to the given SyntheticsBrowserTestResultFullCheck and assigns it to the Check field.
func (o *SyntheticsBrowserTestResultFull) SetCheck(v SyntheticsBrowserTestResultFullCheck) {
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

// GetCheckTimeOk returns a tuple with the CheckTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsBrowserTestResultFull) GetCheckTimeOk() (*float64, bool) {
	if o == nil || o.CheckTime == nil {
		return nil, false
	}
	return o.CheckTime, true
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

// GetCheckVersionOk returns a tuple with the CheckVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsBrowserTestResultFull) GetCheckVersionOk() (*int64, bool) {
	if o == nil || o.CheckVersion == nil {
		return nil, false
	}
	return o.CheckVersion, true
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

// GetProbeDcOk returns a tuple with the ProbeDc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsBrowserTestResultFull) GetProbeDcOk() (*string, bool) {
	if o == nil || o.ProbeDc == nil {
		return nil, false
	}
	return o.ProbeDc, true
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

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsBrowserTestResultFull) GetResultOk() (*SyntheticsBrowserTestResultData, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
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

// GetResultIdOk returns a tuple with the ResultId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsBrowserTestResultFull) GetResultIdOk() (*string, bool) {
	if o == nil || o.ResultId == nil {
		return nil, false
	}
	return o.ResultId, true
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

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsBrowserTestResultFull) GetStatusOk() (*SyntheticsTestMonitorStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
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

func (o SyntheticsBrowserTestResultFull) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
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

func (o *SyntheticsBrowserTestResultFull) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Check        *SyntheticsBrowserTestResultFullCheck `json:"check,omitempty"`
		CheckTime    *float64                              `json:"check_time,omitempty"`
		CheckVersion *int64                                `json:"check_version,omitempty"`
		ProbeDc      *string                               `json:"probe_dc,omitempty"`
		Result       *SyntheticsBrowserTestResultData      `json:"result,omitempty"`
		ResultId     *string                               `json:"result_id,omitempty"`
		Status       *SyntheticsTestMonitorStatus          `json:"status,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if v := all.Status; v != nil && !v.IsValid() {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.Check = all.Check
	o.CheckTime = all.CheckTime
	o.CheckVersion = all.CheckVersion
	o.ProbeDc = all.ProbeDc
	o.Result = all.Result
	o.ResultId = all.ResultId
	o.Status = all.Status
	return nil
}

type NullableSyntheticsBrowserTestResultFull struct {
	value *SyntheticsBrowserTestResultFull
	isSet bool
}

func (v NullableSyntheticsBrowserTestResultFull) Get() *SyntheticsBrowserTestResultFull {
	return v.value
}

func (v *NullableSyntheticsBrowserTestResultFull) Set(val *SyntheticsBrowserTestResultFull) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsBrowserTestResultFull) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsBrowserTestResultFull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsBrowserTestResultFull(val *SyntheticsBrowserTestResultFull) *NullableSyntheticsBrowserTestResultFull {
	return &NullableSyntheticsBrowserTestResultFull{value: val, isSet: true}
}

func (v NullableSyntheticsBrowserTestResultFull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsBrowserTestResultFull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
