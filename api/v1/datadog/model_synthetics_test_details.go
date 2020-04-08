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

// SyntheticsTestDetails TODO.
type SyntheticsTestDetails struct {
	Config *SyntheticsTestConfig `json:"config,omitempty"`
	// TODO.
	CreatedAt *string               `json:"created_at,omitempty"`
	CreatedBy *SyntheticsTestAuthor `json:"created_by,omitempty"`
	// TODO.
	Locations *[]string `json:"locations,omitempty"`
	// TODO.
	Message *string `json:"message,omitempty"`
	// TODO.
	ModifiedAt *string               `json:"modified_at,omitempty"`
	ModifiedBy *SyntheticsTestAuthor `json:"modified_by,omitempty"`
	// TODO.
	Name         *string                      `json:"name,omitempty"`
	Options      *SyntheticsTestOptions       `json:"options,omitempty"`
	OverallState *SyntheticsTestMonitorStatus `json:"overall_state,omitempty"`
	// TODO.
	PublicId *string                    `json:"public_id,omitempty"`
	Status   *SyntheticsTestPauseStatus `json:"status,omitempty"`
	// TODO.
	StepCount *int64                        `json:"stepCount,omitempty"`
	Subtype   *SyntheticsTestDetailsSubType `json:"subtype,omitempty"`
	// TODO.
	Tags *[]string                  `json:"tags,omitempty"`
	Type *SyntheticsTestDetailsType `json:"type,omitempty"`
}

// NewSyntheticsTestDetails instantiates a new SyntheticsTestDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticsTestDetails() *SyntheticsTestDetails {
	this := SyntheticsTestDetails{}
	return &this
}

// NewSyntheticsTestDetailsWithDefaults instantiates a new SyntheticsTestDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticsTestDetailsWithDefaults() *SyntheticsTestDetails {
	this := SyntheticsTestDetails{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *SyntheticsTestDetails) GetConfig() SyntheticsTestConfig {
	if o == nil || o.Config == nil {
		var ret SyntheticsTestConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestDetails) GetConfigOk() (*SyntheticsTestConfig, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *SyntheticsTestDetails) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given SyntheticsTestConfig and assigns it to the Config field.
func (o *SyntheticsTestDetails) SetConfig(v SyntheticsTestConfig) {
	o.Config = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SyntheticsTestDetails) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestDetails) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SyntheticsTestDetails) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *SyntheticsTestDetails) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *SyntheticsTestDetails) GetCreatedBy() SyntheticsTestAuthor {
	if o == nil || o.CreatedBy == nil {
		var ret SyntheticsTestAuthor
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestDetails) GetCreatedByOk() (*SyntheticsTestAuthor, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *SyntheticsTestDetails) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given SyntheticsTestAuthor and assigns it to the CreatedBy field.
func (o *SyntheticsTestDetails) SetCreatedBy(v SyntheticsTestAuthor) {
	o.CreatedBy = &v
}

// GetLocations returns the Locations field value if set, zero value otherwise.
func (o *SyntheticsTestDetails) GetLocations() []string {
	if o == nil || o.Locations == nil {
		var ret []string
		return ret
	}
	return *o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestDetails) GetLocationsOk() (*[]string, bool) {
	if o == nil || o.Locations == nil {
		return nil, false
	}
	return o.Locations, true
}

// HasLocations returns a boolean if a field has been set.
func (o *SyntheticsTestDetails) HasLocations() bool {
	if o != nil && o.Locations != nil {
		return true
	}

	return false
}

// SetLocations gets a reference to the given []string and assigns it to the Locations field.
func (o *SyntheticsTestDetails) SetLocations(v []string) {
	o.Locations = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SyntheticsTestDetails) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestDetails) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SyntheticsTestDetails) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *SyntheticsTestDetails) SetMessage(v string) {
	o.Message = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *SyntheticsTestDetails) GetModifiedAt() string {
	if o == nil || o.ModifiedAt == nil {
		var ret string
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestDetails) GetModifiedAtOk() (*string, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *SyntheticsTestDetails) HasModifiedAt() bool {
	if o != nil && o.ModifiedAt != nil {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given string and assigns it to the ModifiedAt field.
func (o *SyntheticsTestDetails) SetModifiedAt(v string) {
	o.ModifiedAt = &v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *SyntheticsTestDetails) GetModifiedBy() SyntheticsTestAuthor {
	if o == nil || o.ModifiedBy == nil {
		var ret SyntheticsTestAuthor
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestDetails) GetModifiedByOk() (*SyntheticsTestAuthor, bool) {
	if o == nil || o.ModifiedBy == nil {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *SyntheticsTestDetails) HasModifiedBy() bool {
	if o != nil && o.ModifiedBy != nil {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given SyntheticsTestAuthor and assigns it to the ModifiedBy field.
func (o *SyntheticsTestDetails) SetModifiedBy(v SyntheticsTestAuthor) {
	o.ModifiedBy = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SyntheticsTestDetails) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestDetails) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SyntheticsTestDetails) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SyntheticsTestDetails) SetName(v string) {
	o.Name = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *SyntheticsTestDetails) GetOptions() SyntheticsTestOptions {
	if o == nil || o.Options == nil {
		var ret SyntheticsTestOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestDetails) GetOptionsOk() (*SyntheticsTestOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *SyntheticsTestDetails) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given SyntheticsTestOptions and assigns it to the Options field.
func (o *SyntheticsTestDetails) SetOptions(v SyntheticsTestOptions) {
	o.Options = &v
}

// GetOverallState returns the OverallState field value if set, zero value otherwise.
func (o *SyntheticsTestDetails) GetOverallState() SyntheticsTestMonitorStatus {
	if o == nil || o.OverallState == nil {
		var ret SyntheticsTestMonitorStatus
		return ret
	}
	return *o.OverallState
}

// GetOverallStateOk returns a tuple with the OverallState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestDetails) GetOverallStateOk() (*SyntheticsTestMonitorStatus, bool) {
	if o == nil || o.OverallState == nil {
		return nil, false
	}
	return o.OverallState, true
}

// HasOverallState returns a boolean if a field has been set.
func (o *SyntheticsTestDetails) HasOverallState() bool {
	if o != nil && o.OverallState != nil {
		return true
	}

	return false
}

// SetOverallState gets a reference to the given SyntheticsTestMonitorStatus and assigns it to the OverallState field.
func (o *SyntheticsTestDetails) SetOverallState(v SyntheticsTestMonitorStatus) {
	o.OverallState = &v
}

// GetPublicId returns the PublicId field value if set, zero value otherwise.
func (o *SyntheticsTestDetails) GetPublicId() string {
	if o == nil || o.PublicId == nil {
		var ret string
		return ret
	}
	return *o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestDetails) GetPublicIdOk() (*string, bool) {
	if o == nil || o.PublicId == nil {
		return nil, false
	}
	return o.PublicId, true
}

// HasPublicId returns a boolean if a field has been set.
func (o *SyntheticsTestDetails) HasPublicId() bool {
	if o != nil && o.PublicId != nil {
		return true
	}

	return false
}

// SetPublicId gets a reference to the given string and assigns it to the PublicId field.
func (o *SyntheticsTestDetails) SetPublicId(v string) {
	o.PublicId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SyntheticsTestDetails) GetStatus() SyntheticsTestPauseStatus {
	if o == nil || o.Status == nil {
		var ret SyntheticsTestPauseStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestDetails) GetStatusOk() (*SyntheticsTestPauseStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SyntheticsTestDetails) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given SyntheticsTestPauseStatus and assigns it to the Status field.
func (o *SyntheticsTestDetails) SetStatus(v SyntheticsTestPauseStatus) {
	o.Status = &v
}

// GetStepCount returns the StepCount field value if set, zero value otherwise.
func (o *SyntheticsTestDetails) GetStepCount() int64 {
	if o == nil || o.StepCount == nil {
		var ret int64
		return ret
	}
	return *o.StepCount
}

// GetStepCountOk returns a tuple with the StepCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestDetails) GetStepCountOk() (*int64, bool) {
	if o == nil || o.StepCount == nil {
		return nil, false
	}
	return o.StepCount, true
}

// HasStepCount returns a boolean if a field has been set.
func (o *SyntheticsTestDetails) HasStepCount() bool {
	if o != nil && o.StepCount != nil {
		return true
	}

	return false
}

// SetStepCount gets a reference to the given int64 and assigns it to the StepCount field.
func (o *SyntheticsTestDetails) SetStepCount(v int64) {
	o.StepCount = &v
}

// GetSubtype returns the Subtype field value if set, zero value otherwise.
func (o *SyntheticsTestDetails) GetSubtype() SyntheticsTestDetailsSubType {
	if o == nil || o.Subtype == nil {
		var ret SyntheticsTestDetailsSubType
		return ret
	}
	return *o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestDetails) GetSubtypeOk() (*SyntheticsTestDetailsSubType, bool) {
	if o == nil || o.Subtype == nil {
		return nil, false
	}
	return o.Subtype, true
}

// HasSubtype returns a boolean if a field has been set.
func (o *SyntheticsTestDetails) HasSubtype() bool {
	if o != nil && o.Subtype != nil {
		return true
	}

	return false
}

// SetSubtype gets a reference to the given SyntheticsTestDetailsSubType and assigns it to the Subtype field.
func (o *SyntheticsTestDetails) SetSubtype(v SyntheticsTestDetailsSubType) {
	o.Subtype = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SyntheticsTestDetails) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestDetails) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SyntheticsTestDetails) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *SyntheticsTestDetails) SetTags(v []string) {
	o.Tags = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SyntheticsTestDetails) GetType() SyntheticsTestDetailsType {
	if o == nil || o.Type == nil {
		var ret SyntheticsTestDetailsType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestDetails) GetTypeOk() (*SyntheticsTestDetailsType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SyntheticsTestDetails) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given SyntheticsTestDetailsType and assigns it to the Type field.
func (o *SyntheticsTestDetails) SetType(v SyntheticsTestDetailsType) {
	o.Type = &v
}

func (o SyntheticsTestDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if o.Locations != nil {
		toSerialize["locations"] = o.Locations
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.ModifiedAt != nil {
		toSerialize["modified_at"] = o.ModifiedAt
	}
	if o.ModifiedBy != nil {
		toSerialize["modified_by"] = o.ModifiedBy
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.OverallState != nil {
		toSerialize["overall_state"] = o.OverallState
	}
	if o.PublicId != nil {
		toSerialize["public_id"] = o.PublicId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.StepCount != nil {
		toSerialize["stepCount"] = o.StepCount
	}
	if o.Subtype != nil {
		toSerialize["subtype"] = o.Subtype
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableSyntheticsTestDetails struct {
	value *SyntheticsTestDetails
	isSet bool
}

func (v NullableSyntheticsTestDetails) Get() *SyntheticsTestDetails {
	return v.value
}

func (v *NullableSyntheticsTestDetails) Set(val *SyntheticsTestDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsTestDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsTestDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsTestDetails(val *SyntheticsTestDetails) *NullableSyntheticsTestDetails {
	return &NullableSyntheticsTestDetails{value: val, isSet: true}
}

func (v NullableSyntheticsTestDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsTestDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
