/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"time"
)

// DashboardSummaryDashboards struct for DashboardSummaryDashboards
type DashboardSummaryDashboards struct {
	AuthorHandle *string    `json:"author_handle,omitempty"`
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	Description  *string    `json:"description,omitempty"`
	Id           *string    `json:"id,omitempty"`
	IsReadOnly   *bool      `json:"is_read_only,omitempty"`
	LayoutType   *string    `json:"layout_type,omitempty"`
	ModifiedAt   *time.Time `json:"modified_at,omitempty"`
	Title        *string    `json:"title,omitempty"`
	Url          *string    `json:"url,omitempty"`
}

// NewDashboardSummaryDashboards instantiates a new DashboardSummaryDashboards object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardSummaryDashboards() *DashboardSummaryDashboards {
	this := DashboardSummaryDashboards{}
	return &this
}

// NewDashboardSummaryDashboardsWithDefaults instantiates a new DashboardSummaryDashboards object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardSummaryDashboardsWithDefaults() *DashboardSummaryDashboards {
	this := DashboardSummaryDashboards{}
	return &this
}

// GetAuthorHandle returns the AuthorHandle field value if set, zero value otherwise.
func (o *DashboardSummaryDashboards) GetAuthorHandle() string {
	if o == nil || o.AuthorHandle == nil {
		var ret string
		return ret
	}
	return *o.AuthorHandle
}

// GetAuthorHandleOk returns a tuple with the AuthorHandle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSummaryDashboards) GetAuthorHandleOk() (*string, bool) {
	if o == nil || o.AuthorHandle == nil {
		return nil, false
	}
	return o.AuthorHandle, true
}

// HasAuthorHandle returns a boolean if a field has been set.
func (o *DashboardSummaryDashboards) HasAuthorHandle() bool {
	if o != nil && o.AuthorHandle != nil {
		return true
	}

	return false
}

// SetAuthorHandle gets a reference to the given string and assigns it to the AuthorHandle field.
func (o *DashboardSummaryDashboards) SetAuthorHandle(v string) {
	o.AuthorHandle = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DashboardSummaryDashboards) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSummaryDashboards) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DashboardSummaryDashboards) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *DashboardSummaryDashboards) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DashboardSummaryDashboards) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSummaryDashboards) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DashboardSummaryDashboards) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DashboardSummaryDashboards) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DashboardSummaryDashboards) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSummaryDashboards) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DashboardSummaryDashboards) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DashboardSummaryDashboards) SetId(v string) {
	o.Id = &v
}

// GetIsReadOnly returns the IsReadOnly field value if set, zero value otherwise.
func (o *DashboardSummaryDashboards) GetIsReadOnly() bool {
	if o == nil || o.IsReadOnly == nil {
		var ret bool
		return ret
	}
	return *o.IsReadOnly
}

// GetIsReadOnlyOk returns a tuple with the IsReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSummaryDashboards) GetIsReadOnlyOk() (*bool, bool) {
	if o == nil || o.IsReadOnly == nil {
		return nil, false
	}
	return o.IsReadOnly, true
}

// HasIsReadOnly returns a boolean if a field has been set.
func (o *DashboardSummaryDashboards) HasIsReadOnly() bool {
	if o != nil && o.IsReadOnly != nil {
		return true
	}

	return false
}

// SetIsReadOnly gets a reference to the given bool and assigns it to the IsReadOnly field.
func (o *DashboardSummaryDashboards) SetIsReadOnly(v bool) {
	o.IsReadOnly = &v
}

// GetLayoutType returns the LayoutType field value if set, zero value otherwise.
func (o *DashboardSummaryDashboards) GetLayoutType() string {
	if o == nil || o.LayoutType == nil {
		var ret string
		return ret
	}
	return *o.LayoutType
}

// GetLayoutTypeOk returns a tuple with the LayoutType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSummaryDashboards) GetLayoutTypeOk() (*string, bool) {
	if o == nil || o.LayoutType == nil {
		return nil, false
	}
	return o.LayoutType, true
}

// HasLayoutType returns a boolean if a field has been set.
func (o *DashboardSummaryDashboards) HasLayoutType() bool {
	if o != nil && o.LayoutType != nil {
		return true
	}

	return false
}

// SetLayoutType gets a reference to the given string and assigns it to the LayoutType field.
func (o *DashboardSummaryDashboards) SetLayoutType(v string) {
	o.LayoutType = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *DashboardSummaryDashboards) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSummaryDashboards) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *DashboardSummaryDashboards) HasModifiedAt() bool {
	if o != nil && o.ModifiedAt != nil {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *DashboardSummaryDashboards) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *DashboardSummaryDashboards) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSummaryDashboards) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *DashboardSummaryDashboards) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *DashboardSummaryDashboards) SetTitle(v string) {
	o.Title = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *DashboardSummaryDashboards) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSummaryDashboards) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *DashboardSummaryDashboards) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *DashboardSummaryDashboards) SetUrl(v string) {
	o.Url = &v
}

func (o DashboardSummaryDashboards) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthorHandle != nil {
		toSerialize["author_handle"] = o.AuthorHandle
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsReadOnly != nil {
		toSerialize["is_read_only"] = o.IsReadOnly
	}
	if o.LayoutType != nil {
		toSerialize["layout_type"] = o.LayoutType
	}
	if o.ModifiedAt != nil {
		toSerialize["modified_at"] = o.ModifiedAt
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableDashboardSummaryDashboards struct {
	value *DashboardSummaryDashboards
	isSet bool
}

func (v NullableDashboardSummaryDashboards) Get() *DashboardSummaryDashboards {
	return v.value
}

func (v *NullableDashboardSummaryDashboards) Set(val *DashboardSummaryDashboards) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardSummaryDashboards) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardSummaryDashboards) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardSummaryDashboards(val *DashboardSummaryDashboards) *NullableDashboardSummaryDashboards {
	return &NullableDashboardSummaryDashboards{value: val, isSet: true}
}

func (v NullableDashboardSummaryDashboards) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardSummaryDashboards) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
