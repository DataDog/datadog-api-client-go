// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TreemapBackendDefinition Test-only generated schema for the dashboard versioning proof of concept.
type TreemapBackendDefinition struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	ColorBy *TreemapBackendDefinitionColorBy `json:"color_by,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	CustomLinks []WidgetCustomLink `json:"custom_links,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Description *string `json:"description,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	GroupBy *TreemapBackendDefinitionGroupBy `json:"group_by,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Notes *string `json:"notes,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Requests []TreemapBackendRequest `json:"requests"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	SizeBy *TreemapBackendDefinitionSizeBy `json:"size_by,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Time *WidgetTime `json:"time,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Title *string `json:"title,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TitleAlign *WidgetTextAlign `json:"title_align,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TitleSize *string `json:"title_size,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type TreemapBackendDefinitionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTreemapBackendDefinition instantiates a new TreemapBackendDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTreemapBackendDefinition(requests []TreemapBackendRequest, typeVar TreemapBackendDefinitionType) *TreemapBackendDefinition {
	this := TreemapBackendDefinition{}
	this.Requests = requests
	this.Type = typeVar
	return &this
}

// NewTreemapBackendDefinitionWithDefaults instantiates a new TreemapBackendDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTreemapBackendDefinitionWithDefaults() *TreemapBackendDefinition {
	this := TreemapBackendDefinition{}
	return &this
}

// GetColorBy returns the ColorBy field value if set, zero value otherwise.
func (o *TreemapBackendDefinition) GetColorBy() TreemapBackendDefinitionColorBy {
	if o == nil || o.ColorBy == nil {
		var ret TreemapBackendDefinitionColorBy
		return ret
	}
	return *o.ColorBy
}

// GetColorByOk returns a tuple with the ColorBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TreemapBackendDefinition) GetColorByOk() (*TreemapBackendDefinitionColorBy, bool) {
	if o == nil || o.ColorBy == nil {
		return nil, false
	}
	return o.ColorBy, true
}

// HasColorBy returns a boolean if a field has been set.
func (o *TreemapBackendDefinition) HasColorBy() bool {
	return o != nil && o.ColorBy != nil
}

// SetColorBy gets a reference to the given TreemapBackendDefinitionColorBy and assigns it to the ColorBy field.
func (o *TreemapBackendDefinition) SetColorBy(v TreemapBackendDefinitionColorBy) {
	o.ColorBy = &v
}

// GetCustomLinks returns the CustomLinks field value if set, zero value otherwise.
func (o *TreemapBackendDefinition) GetCustomLinks() []WidgetCustomLink {
	if o == nil || o.CustomLinks == nil {
		var ret []WidgetCustomLink
		return ret
	}
	return o.CustomLinks
}

// GetCustomLinksOk returns a tuple with the CustomLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TreemapBackendDefinition) GetCustomLinksOk() (*[]WidgetCustomLink, bool) {
	if o == nil || o.CustomLinks == nil {
		return nil, false
	}
	return &o.CustomLinks, true
}

// HasCustomLinks returns a boolean if a field has been set.
func (o *TreemapBackendDefinition) HasCustomLinks() bool {
	return o != nil && o.CustomLinks != nil
}

// SetCustomLinks gets a reference to the given []WidgetCustomLink and assigns it to the CustomLinks field.
func (o *TreemapBackendDefinition) SetCustomLinks(v []WidgetCustomLink) {
	o.CustomLinks = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TreemapBackendDefinition) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TreemapBackendDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TreemapBackendDefinition) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TreemapBackendDefinition) SetDescription(v string) {
	o.Description = &v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *TreemapBackendDefinition) GetGroupBy() TreemapBackendDefinitionGroupBy {
	if o == nil || o.GroupBy == nil {
		var ret TreemapBackendDefinitionGroupBy
		return ret
	}
	return *o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TreemapBackendDefinition) GetGroupByOk() (*TreemapBackendDefinitionGroupBy, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *TreemapBackendDefinition) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given TreemapBackendDefinitionGroupBy and assigns it to the GroupBy field.
func (o *TreemapBackendDefinition) SetGroupBy(v TreemapBackendDefinitionGroupBy) {
	o.GroupBy = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *TreemapBackendDefinition) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TreemapBackendDefinition) GetNotesOk() (*string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *TreemapBackendDefinition) HasNotes() bool {
	return o != nil && o.Notes != nil
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *TreemapBackendDefinition) SetNotes(v string) {
	o.Notes = &v
}

// GetRequests returns the Requests field value.
func (o *TreemapBackendDefinition) GetRequests() []TreemapBackendRequest {
	if o == nil {
		var ret []TreemapBackendRequest
		return ret
	}
	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value
// and a boolean to check if the value has been set.
func (o *TreemapBackendDefinition) GetRequestsOk() (*[]TreemapBackendRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requests, true
}

// SetRequests sets field value.
func (o *TreemapBackendDefinition) SetRequests(v []TreemapBackendRequest) {
	o.Requests = v
}

// GetSizeBy returns the SizeBy field value if set, zero value otherwise.
func (o *TreemapBackendDefinition) GetSizeBy() TreemapBackendDefinitionSizeBy {
	if o == nil || o.SizeBy == nil {
		var ret TreemapBackendDefinitionSizeBy
		return ret
	}
	return *o.SizeBy
}

// GetSizeByOk returns a tuple with the SizeBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TreemapBackendDefinition) GetSizeByOk() (*TreemapBackendDefinitionSizeBy, bool) {
	if o == nil || o.SizeBy == nil {
		return nil, false
	}
	return o.SizeBy, true
}

// HasSizeBy returns a boolean if a field has been set.
func (o *TreemapBackendDefinition) HasSizeBy() bool {
	return o != nil && o.SizeBy != nil
}

// SetSizeBy gets a reference to the given TreemapBackendDefinitionSizeBy and assigns it to the SizeBy field.
func (o *TreemapBackendDefinition) SetSizeBy(v TreemapBackendDefinitionSizeBy) {
	o.SizeBy = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *TreemapBackendDefinition) GetTime() WidgetTime {
	if o == nil || o.Time == nil {
		var ret WidgetTime
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TreemapBackendDefinition) GetTimeOk() (*WidgetTime, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *TreemapBackendDefinition) HasTime() bool {
	return o != nil && o.Time != nil
}

// SetTime gets a reference to the given WidgetTime and assigns it to the Time field.
func (o *TreemapBackendDefinition) SetTime(v WidgetTime) {
	o.Time = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *TreemapBackendDefinition) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TreemapBackendDefinition) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *TreemapBackendDefinition) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *TreemapBackendDefinition) SetTitle(v string) {
	o.Title = &v
}

// GetTitleAlign returns the TitleAlign field value if set, zero value otherwise.
func (o *TreemapBackendDefinition) GetTitleAlign() WidgetTextAlign {
	if o == nil || o.TitleAlign == nil {
		var ret WidgetTextAlign
		return ret
	}
	return *o.TitleAlign
}

// GetTitleAlignOk returns a tuple with the TitleAlign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TreemapBackendDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool) {
	if o == nil || o.TitleAlign == nil {
		return nil, false
	}
	return o.TitleAlign, true
}

// HasTitleAlign returns a boolean if a field has been set.
func (o *TreemapBackendDefinition) HasTitleAlign() bool {
	return o != nil && o.TitleAlign != nil
}

// SetTitleAlign gets a reference to the given WidgetTextAlign and assigns it to the TitleAlign field.
func (o *TreemapBackendDefinition) SetTitleAlign(v WidgetTextAlign) {
	o.TitleAlign = &v
}

// GetTitleSize returns the TitleSize field value if set, zero value otherwise.
func (o *TreemapBackendDefinition) GetTitleSize() string {
	if o == nil || o.TitleSize == nil {
		var ret string
		return ret
	}
	return *o.TitleSize
}

// GetTitleSizeOk returns a tuple with the TitleSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TreemapBackendDefinition) GetTitleSizeOk() (*string, bool) {
	if o == nil || o.TitleSize == nil {
		return nil, false
	}
	return o.TitleSize, true
}

// HasTitleSize returns a boolean if a field has been set.
func (o *TreemapBackendDefinition) HasTitleSize() bool {
	return o != nil && o.TitleSize != nil
}

// SetTitleSize gets a reference to the given string and assigns it to the TitleSize field.
func (o *TreemapBackendDefinition) SetTitleSize(v string) {
	o.TitleSize = &v
}

// GetType returns the Type field value.
func (o *TreemapBackendDefinition) GetType() TreemapBackendDefinitionType {
	if o == nil {
		var ret TreemapBackendDefinitionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TreemapBackendDefinition) GetTypeOk() (*TreemapBackendDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *TreemapBackendDefinition) SetType(v TreemapBackendDefinitionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TreemapBackendDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ColorBy != nil {
		toSerialize["color_by"] = o.ColorBy
	}
	if o.CustomLinks != nil {
		toSerialize["custom_links"] = o.CustomLinks
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
	}
	toSerialize["requests"] = o.Requests
	if o.SizeBy != nil {
		toSerialize["size_by"] = o.SizeBy
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.TitleAlign != nil {
		toSerialize["title_align"] = o.TitleAlign
	}
	if o.TitleSize != nil {
		toSerialize["title_size"] = o.TitleSize
	}
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TreemapBackendDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ColorBy     *TreemapBackendDefinitionColorBy `json:"color_by,omitempty"`
		CustomLinks []WidgetCustomLink               `json:"custom_links,omitempty"`
		Description *string                          `json:"description,omitempty"`
		GroupBy     *TreemapBackendDefinitionGroupBy `json:"group_by,omitempty"`
		Notes       *string                          `json:"notes,omitempty"`
		Requests    *[]TreemapBackendRequest         `json:"requests"`
		SizeBy      *TreemapBackendDefinitionSizeBy  `json:"size_by,omitempty"`
		Time        *WidgetTime                      `json:"time,omitempty"`
		Title       *string                          `json:"title,omitempty"`
		TitleAlign  *WidgetTextAlign                 `json:"title_align,omitempty"`
		TitleSize   *string                          `json:"title_size,omitempty"`
		Type        *TreemapBackendDefinitionType    `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Requests == nil {
		return fmt.Errorf("required field requests missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	if all.ColorBy != nil && !all.ColorBy.IsValid() {
		hasInvalidField = true
	} else {
		o.ColorBy = all.ColorBy
	}
	o.CustomLinks = all.CustomLinks
	o.Description = all.Description
	if all.GroupBy != nil && !all.GroupBy.IsValid() {
		hasInvalidField = true
	} else {
		o.GroupBy = all.GroupBy
	}
	o.Notes = all.Notes
	o.Requests = *all.Requests
	if all.SizeBy != nil && !all.SizeBy.IsValid() {
		hasInvalidField = true
	} else {
		o.SizeBy = all.SizeBy
	}
	o.Time = all.Time
	o.Title = all.Title
	if all.TitleAlign != nil && !all.TitleAlign.IsValid() {
		hasInvalidField = true
	} else {
		o.TitleAlign = all.TitleAlign
	}
	o.TitleSize = all.TitleSize
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
