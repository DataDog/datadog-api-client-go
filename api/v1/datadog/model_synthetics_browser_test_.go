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

// SyntheticsBrowserTest Object containing details about a Synthetic browser test.
type SyntheticsBrowserTest struct {
	Config *SyntheticsBrowserTestConfig `json:"config,omitempty"`
	// Array of locations used to run the test.
	Locations *[]string `json:"locations,omitempty"`
	// Notification message associated with the test.
	Message *string `json:"message,omitempty"`
	// The associated monitor ID.
	MonitorId *int64 `json:"monitor_id,omitempty"`
	// Name of the test.
	Name    *string                `json:"name,omitempty"`
	Options *SyntheticsTestOptions `json:"options,omitempty"`
	// The public ID of the test.
	PublicId *string                    `json:"public_id,omitempty"`
	Status   *SyntheticsTestPauseStatus `json:"status,omitempty"`
	// The steps of the test.
	Steps *[]SyntheticsStep `json:"steps,omitempty"`
	// Array of tags attached to the test.
	Tags *[]string                  `json:"tags,omitempty"`
	Type *SyntheticsBrowserTestType `json:"type,omitempty"`
}

// NewSyntheticsBrowserTest instantiates a new SyntheticsBrowserTest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticsBrowserTest() *SyntheticsBrowserTest {
	this := SyntheticsBrowserTest{}
	var type_ SyntheticsBrowserTestType = SYNTHETICSBROWSERTESTTYPE_BROWSER
	this.Type = &type_
	return &this
}

// NewSyntheticsBrowserTestWithDefaults instantiates a new SyntheticsBrowserTest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticsBrowserTestWithDefaults() *SyntheticsBrowserTest {
	this := SyntheticsBrowserTest{}
	var type_ SyntheticsBrowserTestType = SYNTHETICSBROWSERTESTTYPE_BROWSER
	this.Type = &type_
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *SyntheticsBrowserTest) GetConfig() SyntheticsBrowserTestConfig {
	if o == nil || o.Config == nil {
		var ret SyntheticsBrowserTestConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsBrowserTest) GetConfigOk() (*SyntheticsBrowserTestConfig, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *SyntheticsBrowserTest) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given SyntheticsBrowserTestConfig and assigns it to the Config field.
func (o *SyntheticsBrowserTest) SetConfig(v SyntheticsBrowserTestConfig) {
	o.Config = &v
}

// GetLocations returns the Locations field value if set, zero value otherwise.
func (o *SyntheticsBrowserTest) GetLocations() []string {
	if o == nil || o.Locations == nil {
		var ret []string
		return ret
	}
	return *o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsBrowserTest) GetLocationsOk() (*[]string, bool) {
	if o == nil || o.Locations == nil {
		return nil, false
	}
	return o.Locations, true
}

// HasLocations returns a boolean if a field has been set.
func (o *SyntheticsBrowserTest) HasLocations() bool {
	if o != nil && o.Locations != nil {
		return true
	}

	return false
}

// SetLocations gets a reference to the given []string and assigns it to the Locations field.
func (o *SyntheticsBrowserTest) SetLocations(v []string) {
	o.Locations = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SyntheticsBrowserTest) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsBrowserTest) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SyntheticsBrowserTest) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *SyntheticsBrowserTest) SetMessage(v string) {
	o.Message = &v
}

// GetMonitorId returns the MonitorId field value if set, zero value otherwise.
func (o *SyntheticsBrowserTest) GetMonitorId() int64 {
	if o == nil || o.MonitorId == nil {
		var ret int64
		return ret
	}
	return *o.MonitorId
}

// GetMonitorIdOk returns a tuple with the MonitorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsBrowserTest) GetMonitorIdOk() (*int64, bool) {
	if o == nil || o.MonitorId == nil {
		return nil, false
	}
	return o.MonitorId, true
}

// HasMonitorId returns a boolean if a field has been set.
func (o *SyntheticsBrowserTest) HasMonitorId() bool {
	if o != nil && o.MonitorId != nil {
		return true
	}

	return false
}

// SetMonitorId gets a reference to the given int64 and assigns it to the MonitorId field.
func (o *SyntheticsBrowserTest) SetMonitorId(v int64) {
	o.MonitorId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SyntheticsBrowserTest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsBrowserTest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SyntheticsBrowserTest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SyntheticsBrowserTest) SetName(v string) {
	o.Name = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *SyntheticsBrowserTest) GetOptions() SyntheticsTestOptions {
	if o == nil || o.Options == nil {
		var ret SyntheticsTestOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsBrowserTest) GetOptionsOk() (*SyntheticsTestOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *SyntheticsBrowserTest) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given SyntheticsTestOptions and assigns it to the Options field.
func (o *SyntheticsBrowserTest) SetOptions(v SyntheticsTestOptions) {
	o.Options = &v
}

// GetPublicId returns the PublicId field value if set, zero value otherwise.
func (o *SyntheticsBrowserTest) GetPublicId() string {
	if o == nil || o.PublicId == nil {
		var ret string
		return ret
	}
	return *o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsBrowserTest) GetPublicIdOk() (*string, bool) {
	if o == nil || o.PublicId == nil {
		return nil, false
	}
	return o.PublicId, true
}

// HasPublicId returns a boolean if a field has been set.
func (o *SyntheticsBrowserTest) HasPublicId() bool {
	if o != nil && o.PublicId != nil {
		return true
	}

	return false
}

// SetPublicId gets a reference to the given string and assigns it to the PublicId field.
func (o *SyntheticsBrowserTest) SetPublicId(v string) {
	o.PublicId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SyntheticsBrowserTest) GetStatus() SyntheticsTestPauseStatus {
	if o == nil || o.Status == nil {
		var ret SyntheticsTestPauseStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsBrowserTest) GetStatusOk() (*SyntheticsTestPauseStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SyntheticsBrowserTest) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given SyntheticsTestPauseStatus and assigns it to the Status field.
func (o *SyntheticsBrowserTest) SetStatus(v SyntheticsTestPauseStatus) {
	o.Status = &v
}

// GetSteps returns the Steps field value if set, zero value otherwise.
func (o *SyntheticsBrowserTest) GetSteps() []SyntheticsStep {
	if o == nil || o.Steps == nil {
		var ret []SyntheticsStep
		return ret
	}
	return *o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsBrowserTest) GetStepsOk() (*[]SyntheticsStep, bool) {
	if o == nil || o.Steps == nil {
		return nil, false
	}
	return o.Steps, true
}

// HasSteps returns a boolean if a field has been set.
func (o *SyntheticsBrowserTest) HasSteps() bool {
	if o != nil && o.Steps != nil {
		return true
	}

	return false
}

// SetSteps gets a reference to the given []SyntheticsStep and assigns it to the Steps field.
func (o *SyntheticsBrowserTest) SetSteps(v []SyntheticsStep) {
	o.Steps = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SyntheticsBrowserTest) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsBrowserTest) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SyntheticsBrowserTest) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *SyntheticsBrowserTest) SetTags(v []string) {
	o.Tags = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SyntheticsBrowserTest) GetType() SyntheticsBrowserTestType {
	if o == nil || o.Type == nil {
		var ret SyntheticsBrowserTestType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsBrowserTest) GetTypeOk() (*SyntheticsBrowserTestType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SyntheticsBrowserTest) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given SyntheticsBrowserTestType and assigns it to the Type field.
func (o *SyntheticsBrowserTest) SetType(v SyntheticsBrowserTestType) {
	o.Type = &v
}

func (o SyntheticsBrowserTest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.Locations != nil {
		toSerialize["locations"] = o.Locations
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.MonitorId != nil {
		toSerialize["monitor_id"] = o.MonitorId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.PublicId != nil {
		toSerialize["public_id"] = o.PublicId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Steps != nil {
		toSerialize["steps"] = o.Steps
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableSyntheticsBrowserTest struct {
	value *SyntheticsBrowserTest
	isSet bool
}

func (v NullableSyntheticsBrowserTest) Get() *SyntheticsBrowserTest {
	return v.value
}

func (v *NullableSyntheticsBrowserTest) Set(val *SyntheticsBrowserTest) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsBrowserTest) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsBrowserTest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsBrowserTest(val *SyntheticsBrowserTest) *NullableSyntheticsBrowserTest {
	return &NullableSyntheticsBrowserTest{value: val, isSet: true}
}

func (v NullableSyntheticsBrowserTest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsBrowserTest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
