// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GeomapBackendDefinition Test-only generated schema for the dashboard versioning proof of concept.
type GeomapBackendDefinition struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	CustomLinks []WidgetCustomLink `json:"custom_links,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Description *string `json:"description,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Notes *string `json:"notes,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Requests []GeomapBackendRequest `json:"requests"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Style GeomapStyle `json:"style"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Time *WidgetTime `json:"time,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Title *string `json:"title,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TitleAlign *WidgetTextAlign `json:"title_align,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TitleSize *string `json:"title_size,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type GeomapBackendDefinitionType `json:"type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	View GeomapView `json:"view"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewGeomapBackendDefinition instantiates a new GeomapBackendDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGeomapBackendDefinition(requests []GeomapBackendRequest, style GeomapStyle, typeVar GeomapBackendDefinitionType, view GeomapView) *GeomapBackendDefinition {
	this := GeomapBackendDefinition{}
	this.Requests = requests
	this.Style = style
	this.Type = typeVar
	this.View = view
	return &this
}

// NewGeomapBackendDefinitionWithDefaults instantiates a new GeomapBackendDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGeomapBackendDefinitionWithDefaults() *GeomapBackendDefinition {
	this := GeomapBackendDefinition{}
	return &this
}

// GetCustomLinks returns the CustomLinks field value if set, zero value otherwise.
func (o *GeomapBackendDefinition) GetCustomLinks() []WidgetCustomLink {
	if o == nil || o.CustomLinks == nil {
		var ret []WidgetCustomLink
		return ret
	}
	return o.CustomLinks
}

// GetCustomLinksOk returns a tuple with the CustomLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeomapBackendDefinition) GetCustomLinksOk() (*[]WidgetCustomLink, bool) {
	if o == nil || o.CustomLinks == nil {
		return nil, false
	}
	return &o.CustomLinks, true
}

// HasCustomLinks returns a boolean if a field has been set.
func (o *GeomapBackendDefinition) HasCustomLinks() bool {
	return o != nil && o.CustomLinks != nil
}

// SetCustomLinks gets a reference to the given []WidgetCustomLink and assigns it to the CustomLinks field.
func (o *GeomapBackendDefinition) SetCustomLinks(v []WidgetCustomLink) {
	o.CustomLinks = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GeomapBackendDefinition) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeomapBackendDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GeomapBackendDefinition) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GeomapBackendDefinition) SetDescription(v string) {
	o.Description = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *GeomapBackendDefinition) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeomapBackendDefinition) GetNotesOk() (*string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *GeomapBackendDefinition) HasNotes() bool {
	return o != nil && o.Notes != nil
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *GeomapBackendDefinition) SetNotes(v string) {
	o.Notes = &v
}

// GetRequests returns the Requests field value.
func (o *GeomapBackendDefinition) GetRequests() []GeomapBackendRequest {
	if o == nil {
		var ret []GeomapBackendRequest
		return ret
	}
	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value
// and a boolean to check if the value has been set.
func (o *GeomapBackendDefinition) GetRequestsOk() (*[]GeomapBackendRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requests, true
}

// SetRequests sets field value.
func (o *GeomapBackendDefinition) SetRequests(v []GeomapBackendRequest) {
	o.Requests = v
}

// GetStyle returns the Style field value.
func (o *GeomapBackendDefinition) GetStyle() GeomapStyle {
	if o == nil {
		var ret GeomapStyle
		return ret
	}
	return o.Style
}

// GetStyleOk returns a tuple with the Style field value
// and a boolean to check if the value has been set.
func (o *GeomapBackendDefinition) GetStyleOk() (*GeomapStyle, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Style, true
}

// SetStyle sets field value.
func (o *GeomapBackendDefinition) SetStyle(v GeomapStyle) {
	o.Style = v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GeomapBackendDefinition) GetTime() WidgetTime {
	if o == nil || o.Time == nil {
		var ret WidgetTime
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeomapBackendDefinition) GetTimeOk() (*WidgetTime, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GeomapBackendDefinition) HasTime() bool {
	return o != nil && o.Time != nil
}

// SetTime gets a reference to the given WidgetTime and assigns it to the Time field.
func (o *GeomapBackendDefinition) SetTime(v WidgetTime) {
	o.Time = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *GeomapBackendDefinition) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeomapBackendDefinition) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *GeomapBackendDefinition) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *GeomapBackendDefinition) SetTitle(v string) {
	o.Title = &v
}

// GetTitleAlign returns the TitleAlign field value if set, zero value otherwise.
func (o *GeomapBackendDefinition) GetTitleAlign() WidgetTextAlign {
	if o == nil || o.TitleAlign == nil {
		var ret WidgetTextAlign
		return ret
	}
	return *o.TitleAlign
}

// GetTitleAlignOk returns a tuple with the TitleAlign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeomapBackendDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool) {
	if o == nil || o.TitleAlign == nil {
		return nil, false
	}
	return o.TitleAlign, true
}

// HasTitleAlign returns a boolean if a field has been set.
func (o *GeomapBackendDefinition) HasTitleAlign() bool {
	return o != nil && o.TitleAlign != nil
}

// SetTitleAlign gets a reference to the given WidgetTextAlign and assigns it to the TitleAlign field.
func (o *GeomapBackendDefinition) SetTitleAlign(v WidgetTextAlign) {
	o.TitleAlign = &v
}

// GetTitleSize returns the TitleSize field value if set, zero value otherwise.
func (o *GeomapBackendDefinition) GetTitleSize() string {
	if o == nil || o.TitleSize == nil {
		var ret string
		return ret
	}
	return *o.TitleSize
}

// GetTitleSizeOk returns a tuple with the TitleSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeomapBackendDefinition) GetTitleSizeOk() (*string, bool) {
	if o == nil || o.TitleSize == nil {
		return nil, false
	}
	return o.TitleSize, true
}

// HasTitleSize returns a boolean if a field has been set.
func (o *GeomapBackendDefinition) HasTitleSize() bool {
	return o != nil && o.TitleSize != nil
}

// SetTitleSize gets a reference to the given string and assigns it to the TitleSize field.
func (o *GeomapBackendDefinition) SetTitleSize(v string) {
	o.TitleSize = &v
}

// GetType returns the Type field value.
func (o *GeomapBackendDefinition) GetType() GeomapBackendDefinitionType {
	if o == nil {
		var ret GeomapBackendDefinitionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GeomapBackendDefinition) GetTypeOk() (*GeomapBackendDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *GeomapBackendDefinition) SetType(v GeomapBackendDefinitionType) {
	o.Type = v
}

// GetView returns the View field value.
func (o *GeomapBackendDefinition) GetView() GeomapView {
	if o == nil {
		var ret GeomapView
		return ret
	}
	return o.View
}

// GetViewOk returns a tuple with the View field value
// and a boolean to check if the value has been set.
func (o *GeomapBackendDefinition) GetViewOk() (*GeomapView, bool) {
	if o == nil {
		return nil, false
	}
	return &o.View, true
}

// SetView sets field value.
func (o *GeomapBackendDefinition) SetView(v GeomapView) {
	o.View = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GeomapBackendDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CustomLinks != nil {
		toSerialize["custom_links"] = o.CustomLinks
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
	}
	toSerialize["requests"] = o.Requests
	toSerialize["style"] = o.Style
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
	toSerialize["view"] = o.View
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GeomapBackendDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CustomLinks []WidgetCustomLink           `json:"custom_links,omitempty"`
		Description *string                      `json:"description,omitempty"`
		Notes       *string                      `json:"notes,omitempty"`
		Requests    *[]GeomapBackendRequest      `json:"requests"`
		Style       *GeomapStyle                 `json:"style"`
		Time        *WidgetTime                  `json:"time,omitempty"`
		Title       *string                      `json:"title,omitempty"`
		TitleAlign  *WidgetTextAlign             `json:"title_align,omitempty"`
		TitleSize   *string                      `json:"title_size,omitempty"`
		Type        *GeomapBackendDefinitionType `json:"type"`
		View        *GeomapView                  `json:"view"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Requests == nil {
		return fmt.Errorf("required field requests missing")
	}
	if all.Style == nil {
		return fmt.Errorf("required field style missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.View == nil {
		return fmt.Errorf("required field view missing")
	}

	hasInvalidField := false
	o.CustomLinks = all.CustomLinks
	o.Description = all.Description
	o.Notes = all.Notes
	o.Requests = *all.Requests
	if all.Style.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Style = *all.Style
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
	o.View = *all.View

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
