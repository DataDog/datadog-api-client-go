// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// QueryTableBackendDefinition Test-only generated schema for the dashboard versioning proof of concept.
type QueryTableBackendDefinition struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	AutoHorizontalScrolling *bool `json:"auto_horizontal_scrolling,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	CustomLinks []WidgetCustomLink `json:"custom_links,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Description *string `json:"description,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	HasSearchBar *HasSearchBar `json:"has_search_bar,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Notes *string `json:"notes,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	NumStickyColumns *float64 `json:"num_sticky_columns,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Requests []QueryTableBackendRequest `json:"requests"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Time *WidgetTime `json:"time,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Title *string `json:"title,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TitleAlign *WidgetTextAlign `json:"title_align,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TitleSize *string `json:"title_size,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type QueryTableBackendDefinitionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewQueryTableBackendDefinition instantiates a new QueryTableBackendDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewQueryTableBackendDefinition(requests []QueryTableBackendRequest, typeVar QueryTableBackendDefinitionType) *QueryTableBackendDefinition {
	this := QueryTableBackendDefinition{}
	this.Requests = requests
	this.Type = typeVar
	return &this
}

// NewQueryTableBackendDefinitionWithDefaults instantiates a new QueryTableBackendDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewQueryTableBackendDefinitionWithDefaults() *QueryTableBackendDefinition {
	this := QueryTableBackendDefinition{}
	return &this
}

// GetAutoHorizontalScrolling returns the AutoHorizontalScrolling field value if set, zero value otherwise.
func (o *QueryTableBackendDefinition) GetAutoHorizontalScrolling() bool {
	if o == nil || o.AutoHorizontalScrolling == nil {
		var ret bool
		return ret
	}
	return *o.AutoHorizontalScrolling
}

// GetAutoHorizontalScrollingOk returns a tuple with the AutoHorizontalScrolling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTableBackendDefinition) GetAutoHorizontalScrollingOk() (*bool, bool) {
	if o == nil || o.AutoHorizontalScrolling == nil {
		return nil, false
	}
	return o.AutoHorizontalScrolling, true
}

// HasAutoHorizontalScrolling returns a boolean if a field has been set.
func (o *QueryTableBackendDefinition) HasAutoHorizontalScrolling() bool {
	return o != nil && o.AutoHorizontalScrolling != nil
}

// SetAutoHorizontalScrolling gets a reference to the given bool and assigns it to the AutoHorizontalScrolling field.
func (o *QueryTableBackendDefinition) SetAutoHorizontalScrolling(v bool) {
	o.AutoHorizontalScrolling = &v
}

// GetCustomLinks returns the CustomLinks field value if set, zero value otherwise.
func (o *QueryTableBackendDefinition) GetCustomLinks() []WidgetCustomLink {
	if o == nil || o.CustomLinks == nil {
		var ret []WidgetCustomLink
		return ret
	}
	return o.CustomLinks
}

// GetCustomLinksOk returns a tuple with the CustomLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTableBackendDefinition) GetCustomLinksOk() (*[]WidgetCustomLink, bool) {
	if o == nil || o.CustomLinks == nil {
		return nil, false
	}
	return &o.CustomLinks, true
}

// HasCustomLinks returns a boolean if a field has been set.
func (o *QueryTableBackendDefinition) HasCustomLinks() bool {
	return o != nil && o.CustomLinks != nil
}

// SetCustomLinks gets a reference to the given []WidgetCustomLink and assigns it to the CustomLinks field.
func (o *QueryTableBackendDefinition) SetCustomLinks(v []WidgetCustomLink) {
	o.CustomLinks = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *QueryTableBackendDefinition) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTableBackendDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *QueryTableBackendDefinition) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *QueryTableBackendDefinition) SetDescription(v string) {
	o.Description = &v
}

// GetHasSearchBar returns the HasSearchBar field value if set, zero value otherwise.
func (o *QueryTableBackendDefinition) GetHasSearchBar() HasSearchBar {
	if o == nil || o.HasSearchBar == nil {
		var ret HasSearchBar
		return ret
	}
	return *o.HasSearchBar
}

// GetHasSearchBarOk returns a tuple with the HasSearchBar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTableBackendDefinition) GetHasSearchBarOk() (*HasSearchBar, bool) {
	if o == nil || o.HasSearchBar == nil {
		return nil, false
	}
	return o.HasSearchBar, true
}

// HasHasSearchBar returns a boolean if a field has been set.
func (o *QueryTableBackendDefinition) HasHasSearchBar() bool {
	return o != nil && o.HasSearchBar != nil
}

// SetHasSearchBar gets a reference to the given HasSearchBar and assigns it to the HasSearchBar field.
func (o *QueryTableBackendDefinition) SetHasSearchBar(v HasSearchBar) {
	o.HasSearchBar = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *QueryTableBackendDefinition) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTableBackendDefinition) GetNotesOk() (*string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *QueryTableBackendDefinition) HasNotes() bool {
	return o != nil && o.Notes != nil
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *QueryTableBackendDefinition) SetNotes(v string) {
	o.Notes = &v
}

// GetNumStickyColumns returns the NumStickyColumns field value if set, zero value otherwise.
func (o *QueryTableBackendDefinition) GetNumStickyColumns() float64 {
	if o == nil || o.NumStickyColumns == nil {
		var ret float64
		return ret
	}
	return *o.NumStickyColumns
}

// GetNumStickyColumnsOk returns a tuple with the NumStickyColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTableBackendDefinition) GetNumStickyColumnsOk() (*float64, bool) {
	if o == nil || o.NumStickyColumns == nil {
		return nil, false
	}
	return o.NumStickyColumns, true
}

// HasNumStickyColumns returns a boolean if a field has been set.
func (o *QueryTableBackendDefinition) HasNumStickyColumns() bool {
	return o != nil && o.NumStickyColumns != nil
}

// SetNumStickyColumns gets a reference to the given float64 and assigns it to the NumStickyColumns field.
func (o *QueryTableBackendDefinition) SetNumStickyColumns(v float64) {
	o.NumStickyColumns = &v
}

// GetRequests returns the Requests field value.
func (o *QueryTableBackendDefinition) GetRequests() []QueryTableBackendRequest {
	if o == nil {
		var ret []QueryTableBackendRequest
		return ret
	}
	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value
// and a boolean to check if the value has been set.
func (o *QueryTableBackendDefinition) GetRequestsOk() (*[]QueryTableBackendRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requests, true
}

// SetRequests sets field value.
func (o *QueryTableBackendDefinition) SetRequests(v []QueryTableBackendRequest) {
	o.Requests = v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *QueryTableBackendDefinition) GetTime() WidgetTime {
	if o == nil || o.Time == nil {
		var ret WidgetTime
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTableBackendDefinition) GetTimeOk() (*WidgetTime, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *QueryTableBackendDefinition) HasTime() bool {
	return o != nil && o.Time != nil
}

// SetTime gets a reference to the given WidgetTime and assigns it to the Time field.
func (o *QueryTableBackendDefinition) SetTime(v WidgetTime) {
	o.Time = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *QueryTableBackendDefinition) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTableBackendDefinition) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *QueryTableBackendDefinition) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *QueryTableBackendDefinition) SetTitle(v string) {
	o.Title = &v
}

// GetTitleAlign returns the TitleAlign field value if set, zero value otherwise.
func (o *QueryTableBackendDefinition) GetTitleAlign() WidgetTextAlign {
	if o == nil || o.TitleAlign == nil {
		var ret WidgetTextAlign
		return ret
	}
	return *o.TitleAlign
}

// GetTitleAlignOk returns a tuple with the TitleAlign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTableBackendDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool) {
	if o == nil || o.TitleAlign == nil {
		return nil, false
	}
	return o.TitleAlign, true
}

// HasTitleAlign returns a boolean if a field has been set.
func (o *QueryTableBackendDefinition) HasTitleAlign() bool {
	return o != nil && o.TitleAlign != nil
}

// SetTitleAlign gets a reference to the given WidgetTextAlign and assigns it to the TitleAlign field.
func (o *QueryTableBackendDefinition) SetTitleAlign(v WidgetTextAlign) {
	o.TitleAlign = &v
}

// GetTitleSize returns the TitleSize field value if set, zero value otherwise.
func (o *QueryTableBackendDefinition) GetTitleSize() string {
	if o == nil || o.TitleSize == nil {
		var ret string
		return ret
	}
	return *o.TitleSize
}

// GetTitleSizeOk returns a tuple with the TitleSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTableBackendDefinition) GetTitleSizeOk() (*string, bool) {
	if o == nil || o.TitleSize == nil {
		return nil, false
	}
	return o.TitleSize, true
}

// HasTitleSize returns a boolean if a field has been set.
func (o *QueryTableBackendDefinition) HasTitleSize() bool {
	return o != nil && o.TitleSize != nil
}

// SetTitleSize gets a reference to the given string and assigns it to the TitleSize field.
func (o *QueryTableBackendDefinition) SetTitleSize(v string) {
	o.TitleSize = &v
}

// GetType returns the Type field value.
func (o *QueryTableBackendDefinition) GetType() QueryTableBackendDefinitionType {
	if o == nil {
		var ret QueryTableBackendDefinitionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *QueryTableBackendDefinition) GetTypeOk() (*QueryTableBackendDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *QueryTableBackendDefinition) SetType(v QueryTableBackendDefinitionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o QueryTableBackendDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AutoHorizontalScrolling != nil {
		toSerialize["auto_horizontal_scrolling"] = o.AutoHorizontalScrolling
	}
	if o.CustomLinks != nil {
		toSerialize["custom_links"] = o.CustomLinks
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.HasSearchBar != nil {
		toSerialize["has_search_bar"] = o.HasSearchBar
	}
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
	}
	if o.NumStickyColumns != nil {
		toSerialize["num_sticky_columns"] = o.NumStickyColumns
	}
	toSerialize["requests"] = o.Requests
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
func (o *QueryTableBackendDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AutoHorizontalScrolling *bool                            `json:"auto_horizontal_scrolling,omitempty"`
		CustomLinks             []WidgetCustomLink               `json:"custom_links,omitempty"`
		Description             *string                          `json:"description,omitempty"`
		HasSearchBar            *HasSearchBar                    `json:"has_search_bar,omitempty"`
		Notes                   *string                          `json:"notes,omitempty"`
		NumStickyColumns        *float64                         `json:"num_sticky_columns,omitempty"`
		Requests                *[]QueryTableBackendRequest      `json:"requests"`
		Time                    *WidgetTime                      `json:"time,omitempty"`
		Title                   *string                          `json:"title,omitempty"`
		TitleAlign              *WidgetTextAlign                 `json:"title_align,omitempty"`
		TitleSize               *string                          `json:"title_size,omitempty"`
		Type                    *QueryTableBackendDefinitionType `json:"type"`
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
	o.AutoHorizontalScrolling = all.AutoHorizontalScrolling
	o.CustomLinks = all.CustomLinks
	o.Description = all.Description
	if all.HasSearchBar != nil && !all.HasSearchBar.IsValid() {
		hasInvalidField = true
	} else {
		o.HasSearchBar = all.HasSearchBar
	}
	o.Notes = all.Notes
	o.NumStickyColumns = all.NumStickyColumns
	o.Requests = *all.Requests
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
