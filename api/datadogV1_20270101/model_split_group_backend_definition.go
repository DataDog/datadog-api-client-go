// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SplitGroupBackendDefinition Test-only generated schema for the dashboard versioning proof of concept.
type SplitGroupBackendDefinition struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Description *string `json:"description,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	HasUniformYAxes *bool `json:"has_uniform_y_axes,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Notes *string `json:"notes,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Size FORMSIZES `json:"size"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	SourceWidgetDefinition SplitGroupBackendDefinition20270101SourceWidgetDefinition `json:"source_widget_definition"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	SplitConfig SplitConfig `json:"split_config"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Time *WidgetTime `json:"time,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Title *string `json:"title,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TitleAlign *WidgetTextAlign `json:"title_align,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TitleSize *string `json:"title_size,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type SplitGroupBackendDefinitionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewSplitGroupBackendDefinition instantiates a new SplitGroupBackendDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSplitGroupBackendDefinition(size FORMSIZES, sourceWidgetDefinition SplitGroupBackendDefinition20270101SourceWidgetDefinition, splitConfig SplitConfig, typeVar SplitGroupBackendDefinitionType) *SplitGroupBackendDefinition {
	this := SplitGroupBackendDefinition{}
	this.Size = size
	this.SourceWidgetDefinition = sourceWidgetDefinition
	this.SplitConfig = splitConfig
	this.Type = typeVar
	return &this
}

// NewSplitGroupBackendDefinitionWithDefaults instantiates a new SplitGroupBackendDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSplitGroupBackendDefinitionWithDefaults() *SplitGroupBackendDefinition {
	this := SplitGroupBackendDefinition{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SplitGroupBackendDefinition) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitGroupBackendDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SplitGroupBackendDefinition) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SplitGroupBackendDefinition) SetDescription(v string) {
	o.Description = &v
}

// GetHasUniformYAxes returns the HasUniformYAxes field value if set, zero value otherwise.
func (o *SplitGroupBackendDefinition) GetHasUniformYAxes() bool {
	if o == nil || o.HasUniformYAxes == nil {
		var ret bool
		return ret
	}
	return *o.HasUniformYAxes
}

// GetHasUniformYAxesOk returns a tuple with the HasUniformYAxes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitGroupBackendDefinition) GetHasUniformYAxesOk() (*bool, bool) {
	if o == nil || o.HasUniformYAxes == nil {
		return nil, false
	}
	return o.HasUniformYAxes, true
}

// HasHasUniformYAxes returns a boolean if a field has been set.
func (o *SplitGroupBackendDefinition) HasHasUniformYAxes() bool {
	return o != nil && o.HasUniformYAxes != nil
}

// SetHasUniformYAxes gets a reference to the given bool and assigns it to the HasUniformYAxes field.
func (o *SplitGroupBackendDefinition) SetHasUniformYAxes(v bool) {
	o.HasUniformYAxes = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *SplitGroupBackendDefinition) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitGroupBackendDefinition) GetNotesOk() (*string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *SplitGroupBackendDefinition) HasNotes() bool {
	return o != nil && o.Notes != nil
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *SplitGroupBackendDefinition) SetNotes(v string) {
	o.Notes = &v
}

// GetSize returns the Size field value.
func (o *SplitGroupBackendDefinition) GetSize() FORMSIZES {
	if o == nil {
		var ret FORMSIZES
		return ret
	}
	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *SplitGroupBackendDefinition) GetSizeOk() (*FORMSIZES, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value.
func (o *SplitGroupBackendDefinition) SetSize(v FORMSIZES) {
	o.Size = v
}

// GetSourceWidgetDefinition returns the SourceWidgetDefinition field value.
func (o *SplitGroupBackendDefinition) GetSourceWidgetDefinition() SplitGroupBackendDefinition20270101SourceWidgetDefinition {
	if o == nil {
		var ret SplitGroupBackendDefinition20270101SourceWidgetDefinition
		return ret
	}
	return o.SourceWidgetDefinition
}

// GetSourceWidgetDefinitionOk returns a tuple with the SourceWidgetDefinition field value
// and a boolean to check if the value has been set.
func (o *SplitGroupBackendDefinition) GetSourceWidgetDefinitionOk() (*SplitGroupBackendDefinition20270101SourceWidgetDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceWidgetDefinition, true
}

// SetSourceWidgetDefinition sets field value.
func (o *SplitGroupBackendDefinition) SetSourceWidgetDefinition(v SplitGroupBackendDefinition20270101SourceWidgetDefinition) {
	o.SourceWidgetDefinition = v
}

// GetSplitConfig returns the SplitConfig field value.
func (o *SplitGroupBackendDefinition) GetSplitConfig() SplitConfig {
	if o == nil {
		var ret SplitConfig
		return ret
	}
	return o.SplitConfig
}

// GetSplitConfigOk returns a tuple with the SplitConfig field value
// and a boolean to check if the value has been set.
func (o *SplitGroupBackendDefinition) GetSplitConfigOk() (*SplitConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SplitConfig, true
}

// SetSplitConfig sets field value.
func (o *SplitGroupBackendDefinition) SetSplitConfig(v SplitConfig) {
	o.SplitConfig = v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *SplitGroupBackendDefinition) GetTime() WidgetTime {
	if o == nil || o.Time == nil {
		var ret WidgetTime
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitGroupBackendDefinition) GetTimeOk() (*WidgetTime, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *SplitGroupBackendDefinition) HasTime() bool {
	return o != nil && o.Time != nil
}

// SetTime gets a reference to the given WidgetTime and assigns it to the Time field.
func (o *SplitGroupBackendDefinition) SetTime(v WidgetTime) {
	o.Time = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *SplitGroupBackendDefinition) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitGroupBackendDefinition) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *SplitGroupBackendDefinition) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *SplitGroupBackendDefinition) SetTitle(v string) {
	o.Title = &v
}

// GetTitleAlign returns the TitleAlign field value if set, zero value otherwise.
func (o *SplitGroupBackendDefinition) GetTitleAlign() WidgetTextAlign {
	if o == nil || o.TitleAlign == nil {
		var ret WidgetTextAlign
		return ret
	}
	return *o.TitleAlign
}

// GetTitleAlignOk returns a tuple with the TitleAlign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitGroupBackendDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool) {
	if o == nil || o.TitleAlign == nil {
		return nil, false
	}
	return o.TitleAlign, true
}

// HasTitleAlign returns a boolean if a field has been set.
func (o *SplitGroupBackendDefinition) HasTitleAlign() bool {
	return o != nil && o.TitleAlign != nil
}

// SetTitleAlign gets a reference to the given WidgetTextAlign and assigns it to the TitleAlign field.
func (o *SplitGroupBackendDefinition) SetTitleAlign(v WidgetTextAlign) {
	o.TitleAlign = &v
}

// GetTitleSize returns the TitleSize field value if set, zero value otherwise.
func (o *SplitGroupBackendDefinition) GetTitleSize() string {
	if o == nil || o.TitleSize == nil {
		var ret string
		return ret
	}
	return *o.TitleSize
}

// GetTitleSizeOk returns a tuple with the TitleSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitGroupBackendDefinition) GetTitleSizeOk() (*string, bool) {
	if o == nil || o.TitleSize == nil {
		return nil, false
	}
	return o.TitleSize, true
}

// HasTitleSize returns a boolean if a field has been set.
func (o *SplitGroupBackendDefinition) HasTitleSize() bool {
	return o != nil && o.TitleSize != nil
}

// SetTitleSize gets a reference to the given string and assigns it to the TitleSize field.
func (o *SplitGroupBackendDefinition) SetTitleSize(v string) {
	o.TitleSize = &v
}

// GetType returns the Type field value.
func (o *SplitGroupBackendDefinition) GetType() SplitGroupBackendDefinitionType {
	if o == nil {
		var ret SplitGroupBackendDefinitionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SplitGroupBackendDefinition) GetTypeOk() (*SplitGroupBackendDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SplitGroupBackendDefinition) SetType(v SplitGroupBackendDefinitionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SplitGroupBackendDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.HasUniformYAxes != nil {
		toSerialize["has_uniform_y_axes"] = o.HasUniformYAxes
	}
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
	}
	toSerialize["size"] = o.Size
	toSerialize["source_widget_definition"] = o.SourceWidgetDefinition
	toSerialize["split_config"] = o.SplitConfig
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
func (o *SplitGroupBackendDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description            *string                                                    `json:"description,omitempty"`
		HasUniformYAxes        *bool                                                      `json:"has_uniform_y_axes,omitempty"`
		Notes                  *string                                                    `json:"notes,omitempty"`
		Size                   *FORMSIZES                                                 `json:"size"`
		SourceWidgetDefinition *SplitGroupBackendDefinition20270101SourceWidgetDefinition `json:"source_widget_definition"`
		SplitConfig            *SplitConfig                                               `json:"split_config"`
		Time                   *WidgetTime                                                `json:"time,omitempty"`
		Title                  *string                                                    `json:"title,omitempty"`
		TitleAlign             *WidgetTextAlign                                           `json:"title_align,omitempty"`
		TitleSize              *string                                                    `json:"title_size,omitempty"`
		Type                   *SplitGroupBackendDefinitionType                           `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Size == nil {
		return fmt.Errorf("required field size missing")
	}
	if all.SourceWidgetDefinition == nil {
		return fmt.Errorf("required field source_widget_definition missing")
	}
	if all.SplitConfig == nil {
		return fmt.Errorf("required field split_config missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	o.Description = all.Description
	o.HasUniformYAxes = all.HasUniformYAxes
	o.Notes = all.Notes
	if !all.Size.IsValid() {
		hasInvalidField = true
	} else {
		o.Size = *all.Size
	}
	o.SourceWidgetDefinition = *all.SourceWidgetDefinition
	if all.SplitConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SplitConfig = *all.SplitConfig
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
