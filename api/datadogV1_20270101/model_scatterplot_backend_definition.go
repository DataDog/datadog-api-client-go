// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScatterplotBackendDefinition Test-only generated schema for the dashboard versioning proof of concept.
type ScatterplotBackendDefinition struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	ColorByGroups []string `json:"color_by_groups,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	CustomLinks []WidgetCustomLink `json:"custom_links,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Description *string `json:"description,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Legend *ScatterPlotLegendType `json:"legend,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Notes *string `json:"notes,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Palette *string `json:"palette,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Requests ScatterplotRequests `json:"requests"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Time *WidgetTime `json:"time,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Title *string `json:"title,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TitleAlign *WidgetTextAlign `json:"title_align,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TitleSize *string `json:"title_size,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type ScatterplotBackendDefinitionType `json:"type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Xaxis *WidgetAxis `json:"xaxis,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Yaxis *WidgetAxis `json:"yaxis,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewScatterplotBackendDefinition instantiates a new ScatterplotBackendDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScatterplotBackendDefinition(requests ScatterplotRequests, typeVar ScatterplotBackendDefinitionType) *ScatterplotBackendDefinition {
	this := ScatterplotBackendDefinition{}
	this.Requests = requests
	this.Type = typeVar
	return &this
}

// NewScatterplotBackendDefinitionWithDefaults instantiates a new ScatterplotBackendDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScatterplotBackendDefinitionWithDefaults() *ScatterplotBackendDefinition {
	this := ScatterplotBackendDefinition{}
	return &this
}

// GetColorByGroups returns the ColorByGroups field value if set, zero value otherwise.
func (o *ScatterplotBackendDefinition) GetColorByGroups() []string {
	if o == nil || o.ColorByGroups == nil {
		var ret []string
		return ret
	}
	return o.ColorByGroups
}

// GetColorByGroupsOk returns a tuple with the ColorByGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScatterplotBackendDefinition) GetColorByGroupsOk() (*[]string, bool) {
	if o == nil || o.ColorByGroups == nil {
		return nil, false
	}
	return &o.ColorByGroups, true
}

// HasColorByGroups returns a boolean if a field has been set.
func (o *ScatterplotBackendDefinition) HasColorByGroups() bool {
	return o != nil && o.ColorByGroups != nil
}

// SetColorByGroups gets a reference to the given []string and assigns it to the ColorByGroups field.
func (o *ScatterplotBackendDefinition) SetColorByGroups(v []string) {
	o.ColorByGroups = v
}

// GetCustomLinks returns the CustomLinks field value if set, zero value otherwise.
func (o *ScatterplotBackendDefinition) GetCustomLinks() []WidgetCustomLink {
	if o == nil || o.CustomLinks == nil {
		var ret []WidgetCustomLink
		return ret
	}
	return o.CustomLinks
}

// GetCustomLinksOk returns a tuple with the CustomLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScatterplotBackendDefinition) GetCustomLinksOk() (*[]WidgetCustomLink, bool) {
	if o == nil || o.CustomLinks == nil {
		return nil, false
	}
	return &o.CustomLinks, true
}

// HasCustomLinks returns a boolean if a field has been set.
func (o *ScatterplotBackendDefinition) HasCustomLinks() bool {
	return o != nil && o.CustomLinks != nil
}

// SetCustomLinks gets a reference to the given []WidgetCustomLink and assigns it to the CustomLinks field.
func (o *ScatterplotBackendDefinition) SetCustomLinks(v []WidgetCustomLink) {
	o.CustomLinks = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ScatterplotBackendDefinition) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScatterplotBackendDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ScatterplotBackendDefinition) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ScatterplotBackendDefinition) SetDescription(v string) {
	o.Description = &v
}

// GetLegend returns the Legend field value if set, zero value otherwise.
func (o *ScatterplotBackendDefinition) GetLegend() ScatterPlotLegendType {
	if o == nil || o.Legend == nil {
		var ret ScatterPlotLegendType
		return ret
	}
	return *o.Legend
}

// GetLegendOk returns a tuple with the Legend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScatterplotBackendDefinition) GetLegendOk() (*ScatterPlotLegendType, bool) {
	if o == nil || o.Legend == nil {
		return nil, false
	}
	return o.Legend, true
}

// HasLegend returns a boolean if a field has been set.
func (o *ScatterplotBackendDefinition) HasLegend() bool {
	return o != nil && o.Legend != nil
}

// SetLegend gets a reference to the given ScatterPlotLegendType and assigns it to the Legend field.
func (o *ScatterplotBackendDefinition) SetLegend(v ScatterPlotLegendType) {
	o.Legend = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *ScatterplotBackendDefinition) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScatterplotBackendDefinition) GetNotesOk() (*string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *ScatterplotBackendDefinition) HasNotes() bool {
	return o != nil && o.Notes != nil
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *ScatterplotBackendDefinition) SetNotes(v string) {
	o.Notes = &v
}

// GetPalette returns the Palette field value if set, zero value otherwise.
func (o *ScatterplotBackendDefinition) GetPalette() string {
	if o == nil || o.Palette == nil {
		var ret string
		return ret
	}
	return *o.Palette
}

// GetPaletteOk returns a tuple with the Palette field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScatterplotBackendDefinition) GetPaletteOk() (*string, bool) {
	if o == nil || o.Palette == nil {
		return nil, false
	}
	return o.Palette, true
}

// HasPalette returns a boolean if a field has been set.
func (o *ScatterplotBackendDefinition) HasPalette() bool {
	return o != nil && o.Palette != nil
}

// SetPalette gets a reference to the given string and assigns it to the Palette field.
func (o *ScatterplotBackendDefinition) SetPalette(v string) {
	o.Palette = &v
}

// GetRequests returns the Requests field value.
func (o *ScatterplotBackendDefinition) GetRequests() ScatterplotRequests {
	if o == nil {
		var ret ScatterplotRequests
		return ret
	}
	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value
// and a boolean to check if the value has been set.
func (o *ScatterplotBackendDefinition) GetRequestsOk() (*ScatterplotRequests, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requests, true
}

// SetRequests sets field value.
func (o *ScatterplotBackendDefinition) SetRequests(v ScatterplotRequests) {
	o.Requests = v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *ScatterplotBackendDefinition) GetTime() WidgetTime {
	if o == nil || o.Time == nil {
		var ret WidgetTime
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScatterplotBackendDefinition) GetTimeOk() (*WidgetTime, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *ScatterplotBackendDefinition) HasTime() bool {
	return o != nil && o.Time != nil
}

// SetTime gets a reference to the given WidgetTime and assigns it to the Time field.
func (o *ScatterplotBackendDefinition) SetTime(v WidgetTime) {
	o.Time = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ScatterplotBackendDefinition) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScatterplotBackendDefinition) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ScatterplotBackendDefinition) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ScatterplotBackendDefinition) SetTitle(v string) {
	o.Title = &v
}

// GetTitleAlign returns the TitleAlign field value if set, zero value otherwise.
func (o *ScatterplotBackendDefinition) GetTitleAlign() WidgetTextAlign {
	if o == nil || o.TitleAlign == nil {
		var ret WidgetTextAlign
		return ret
	}
	return *o.TitleAlign
}

// GetTitleAlignOk returns a tuple with the TitleAlign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScatterplotBackendDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool) {
	if o == nil || o.TitleAlign == nil {
		return nil, false
	}
	return o.TitleAlign, true
}

// HasTitleAlign returns a boolean if a field has been set.
func (o *ScatterplotBackendDefinition) HasTitleAlign() bool {
	return o != nil && o.TitleAlign != nil
}

// SetTitleAlign gets a reference to the given WidgetTextAlign and assigns it to the TitleAlign field.
func (o *ScatterplotBackendDefinition) SetTitleAlign(v WidgetTextAlign) {
	o.TitleAlign = &v
}

// GetTitleSize returns the TitleSize field value if set, zero value otherwise.
func (o *ScatterplotBackendDefinition) GetTitleSize() string {
	if o == nil || o.TitleSize == nil {
		var ret string
		return ret
	}
	return *o.TitleSize
}

// GetTitleSizeOk returns a tuple with the TitleSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScatterplotBackendDefinition) GetTitleSizeOk() (*string, bool) {
	if o == nil || o.TitleSize == nil {
		return nil, false
	}
	return o.TitleSize, true
}

// HasTitleSize returns a boolean if a field has been set.
func (o *ScatterplotBackendDefinition) HasTitleSize() bool {
	return o != nil && o.TitleSize != nil
}

// SetTitleSize gets a reference to the given string and assigns it to the TitleSize field.
func (o *ScatterplotBackendDefinition) SetTitleSize(v string) {
	o.TitleSize = &v
}

// GetType returns the Type field value.
func (o *ScatterplotBackendDefinition) GetType() ScatterplotBackendDefinitionType {
	if o == nil {
		var ret ScatterplotBackendDefinitionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ScatterplotBackendDefinition) GetTypeOk() (*ScatterplotBackendDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ScatterplotBackendDefinition) SetType(v ScatterplotBackendDefinitionType) {
	o.Type = v
}

// GetXaxis returns the Xaxis field value if set, zero value otherwise.
func (o *ScatterplotBackendDefinition) GetXaxis() WidgetAxis {
	if o == nil || o.Xaxis == nil {
		var ret WidgetAxis
		return ret
	}
	return *o.Xaxis
}

// GetXaxisOk returns a tuple with the Xaxis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScatterplotBackendDefinition) GetXaxisOk() (*WidgetAxis, bool) {
	if o == nil || o.Xaxis == nil {
		return nil, false
	}
	return o.Xaxis, true
}

// HasXaxis returns a boolean if a field has been set.
func (o *ScatterplotBackendDefinition) HasXaxis() bool {
	return o != nil && o.Xaxis != nil
}

// SetXaxis gets a reference to the given WidgetAxis and assigns it to the Xaxis field.
func (o *ScatterplotBackendDefinition) SetXaxis(v WidgetAxis) {
	o.Xaxis = &v
}

// GetYaxis returns the Yaxis field value if set, zero value otherwise.
func (o *ScatterplotBackendDefinition) GetYaxis() WidgetAxis {
	if o == nil || o.Yaxis == nil {
		var ret WidgetAxis
		return ret
	}
	return *o.Yaxis
}

// GetYaxisOk returns a tuple with the Yaxis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScatterplotBackendDefinition) GetYaxisOk() (*WidgetAxis, bool) {
	if o == nil || o.Yaxis == nil {
		return nil, false
	}
	return o.Yaxis, true
}

// HasYaxis returns a boolean if a field has been set.
func (o *ScatterplotBackendDefinition) HasYaxis() bool {
	return o != nil && o.Yaxis != nil
}

// SetYaxis gets a reference to the given WidgetAxis and assigns it to the Yaxis field.
func (o *ScatterplotBackendDefinition) SetYaxis(v WidgetAxis) {
	o.Yaxis = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScatterplotBackendDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ColorByGroups != nil {
		toSerialize["color_by_groups"] = o.ColorByGroups
	}
	if o.CustomLinks != nil {
		toSerialize["custom_links"] = o.CustomLinks
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Legend != nil {
		toSerialize["legend"] = o.Legend
	}
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
	}
	if o.Palette != nil {
		toSerialize["palette"] = o.Palette
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
	if o.Xaxis != nil {
		toSerialize["xaxis"] = o.Xaxis
	}
	if o.Yaxis != nil {
		toSerialize["yaxis"] = o.Yaxis
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScatterplotBackendDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ColorByGroups []string                          `json:"color_by_groups,omitempty"`
		CustomLinks   []WidgetCustomLink                `json:"custom_links,omitempty"`
		Description   *string                           `json:"description,omitempty"`
		Legend        *ScatterPlotLegendType            `json:"legend,omitempty"`
		Notes         *string                           `json:"notes,omitempty"`
		Palette       *string                           `json:"palette,omitempty"`
		Requests      *ScatterplotRequests              `json:"requests"`
		Time          *WidgetTime                       `json:"time,omitempty"`
		Title         *string                           `json:"title,omitempty"`
		TitleAlign    *WidgetTextAlign                  `json:"title_align,omitempty"`
		TitleSize     *string                           `json:"title_size,omitempty"`
		Type          *ScatterplotBackendDefinitionType `json:"type"`
		Xaxis         *WidgetAxis                       `json:"xaxis,omitempty"`
		Yaxis         *WidgetAxis                       `json:"yaxis,omitempty"`
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
	o.ColorByGroups = all.ColorByGroups
	o.CustomLinks = all.CustomLinks
	o.Description = all.Description
	o.Legend = all.Legend
	o.Notes = all.Notes
	o.Palette = all.Palette
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
	if all.Xaxis != nil && all.Xaxis.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Xaxis = all.Xaxis
	if all.Yaxis != nil && all.Yaxis.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Yaxis = all.Yaxis

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
