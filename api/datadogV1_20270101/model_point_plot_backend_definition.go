// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PointPlotBackendDefinition Test-only generated schema for the dashboard versioning proof of concept.
type PointPlotBackendDefinition struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Description *string `json:"description,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Legend *ScatterPlotLegendType `json:"legend,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Markers []Marker `json:"markers,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Notes *string `json:"notes,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Requests []PointPlotDataProjectionRequest `json:"requests"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Time *WidgetTime `json:"time,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Title *string `json:"title,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TitleAlign *WidgetTextAlign `json:"title_align,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TitleSize *string `json:"title_size,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type PointPlotBackendDefinitionType `json:"type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Yaxis *Axis `json:"yaxis,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewPointPlotBackendDefinition instantiates a new PointPlotBackendDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPointPlotBackendDefinition(requests []PointPlotDataProjectionRequest, typeVar PointPlotBackendDefinitionType) *PointPlotBackendDefinition {
	this := PointPlotBackendDefinition{}
	this.Requests = requests
	this.Type = typeVar
	return &this
}

// NewPointPlotBackendDefinitionWithDefaults instantiates a new PointPlotBackendDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPointPlotBackendDefinitionWithDefaults() *PointPlotBackendDefinition {
	this := PointPlotBackendDefinition{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PointPlotBackendDefinition) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PointPlotBackendDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PointPlotBackendDefinition) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PointPlotBackendDefinition) SetDescription(v string) {
	o.Description = &v
}

// GetLegend returns the Legend field value if set, zero value otherwise.
func (o *PointPlotBackendDefinition) GetLegend() ScatterPlotLegendType {
	if o == nil || o.Legend == nil {
		var ret ScatterPlotLegendType
		return ret
	}
	return *o.Legend
}

// GetLegendOk returns a tuple with the Legend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PointPlotBackendDefinition) GetLegendOk() (*ScatterPlotLegendType, bool) {
	if o == nil || o.Legend == nil {
		return nil, false
	}
	return o.Legend, true
}

// HasLegend returns a boolean if a field has been set.
func (o *PointPlotBackendDefinition) HasLegend() bool {
	return o != nil && o.Legend != nil
}

// SetLegend gets a reference to the given ScatterPlotLegendType and assigns it to the Legend field.
func (o *PointPlotBackendDefinition) SetLegend(v ScatterPlotLegendType) {
	o.Legend = &v
}

// GetMarkers returns the Markers field value if set, zero value otherwise.
func (o *PointPlotBackendDefinition) GetMarkers() []Marker {
	if o == nil || o.Markers == nil {
		var ret []Marker
		return ret
	}
	return o.Markers
}

// GetMarkersOk returns a tuple with the Markers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PointPlotBackendDefinition) GetMarkersOk() (*[]Marker, bool) {
	if o == nil || o.Markers == nil {
		return nil, false
	}
	return &o.Markers, true
}

// HasMarkers returns a boolean if a field has been set.
func (o *PointPlotBackendDefinition) HasMarkers() bool {
	return o != nil && o.Markers != nil
}

// SetMarkers gets a reference to the given []Marker and assigns it to the Markers field.
func (o *PointPlotBackendDefinition) SetMarkers(v []Marker) {
	o.Markers = v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *PointPlotBackendDefinition) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PointPlotBackendDefinition) GetNotesOk() (*string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *PointPlotBackendDefinition) HasNotes() bool {
	return o != nil && o.Notes != nil
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *PointPlotBackendDefinition) SetNotes(v string) {
	o.Notes = &v
}

// GetRequests returns the Requests field value.
func (o *PointPlotBackendDefinition) GetRequests() []PointPlotDataProjectionRequest {
	if o == nil {
		var ret []PointPlotDataProjectionRequest
		return ret
	}
	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value
// and a boolean to check if the value has been set.
func (o *PointPlotBackendDefinition) GetRequestsOk() (*[]PointPlotDataProjectionRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requests, true
}

// SetRequests sets field value.
func (o *PointPlotBackendDefinition) SetRequests(v []PointPlotDataProjectionRequest) {
	o.Requests = v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *PointPlotBackendDefinition) GetTime() WidgetTime {
	if o == nil || o.Time == nil {
		var ret WidgetTime
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PointPlotBackendDefinition) GetTimeOk() (*WidgetTime, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *PointPlotBackendDefinition) HasTime() bool {
	return o != nil && o.Time != nil
}

// SetTime gets a reference to the given WidgetTime and assigns it to the Time field.
func (o *PointPlotBackendDefinition) SetTime(v WidgetTime) {
	o.Time = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *PointPlotBackendDefinition) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PointPlotBackendDefinition) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *PointPlotBackendDefinition) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *PointPlotBackendDefinition) SetTitle(v string) {
	o.Title = &v
}

// GetTitleAlign returns the TitleAlign field value if set, zero value otherwise.
func (o *PointPlotBackendDefinition) GetTitleAlign() WidgetTextAlign {
	if o == nil || o.TitleAlign == nil {
		var ret WidgetTextAlign
		return ret
	}
	return *o.TitleAlign
}

// GetTitleAlignOk returns a tuple with the TitleAlign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PointPlotBackendDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool) {
	if o == nil || o.TitleAlign == nil {
		return nil, false
	}
	return o.TitleAlign, true
}

// HasTitleAlign returns a boolean if a field has been set.
func (o *PointPlotBackendDefinition) HasTitleAlign() bool {
	return o != nil && o.TitleAlign != nil
}

// SetTitleAlign gets a reference to the given WidgetTextAlign and assigns it to the TitleAlign field.
func (o *PointPlotBackendDefinition) SetTitleAlign(v WidgetTextAlign) {
	o.TitleAlign = &v
}

// GetTitleSize returns the TitleSize field value if set, zero value otherwise.
func (o *PointPlotBackendDefinition) GetTitleSize() string {
	if o == nil || o.TitleSize == nil {
		var ret string
		return ret
	}
	return *o.TitleSize
}

// GetTitleSizeOk returns a tuple with the TitleSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PointPlotBackendDefinition) GetTitleSizeOk() (*string, bool) {
	if o == nil || o.TitleSize == nil {
		return nil, false
	}
	return o.TitleSize, true
}

// HasTitleSize returns a boolean if a field has been set.
func (o *PointPlotBackendDefinition) HasTitleSize() bool {
	return o != nil && o.TitleSize != nil
}

// SetTitleSize gets a reference to the given string and assigns it to the TitleSize field.
func (o *PointPlotBackendDefinition) SetTitleSize(v string) {
	o.TitleSize = &v
}

// GetType returns the Type field value.
func (o *PointPlotBackendDefinition) GetType() PointPlotBackendDefinitionType {
	if o == nil {
		var ret PointPlotBackendDefinitionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PointPlotBackendDefinition) GetTypeOk() (*PointPlotBackendDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *PointPlotBackendDefinition) SetType(v PointPlotBackendDefinitionType) {
	o.Type = v
}

// GetYaxis returns the Yaxis field value if set, zero value otherwise.
func (o *PointPlotBackendDefinition) GetYaxis() Axis {
	if o == nil || o.Yaxis == nil {
		var ret Axis
		return ret
	}
	return *o.Yaxis
}

// GetYaxisOk returns a tuple with the Yaxis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PointPlotBackendDefinition) GetYaxisOk() (*Axis, bool) {
	if o == nil || o.Yaxis == nil {
		return nil, false
	}
	return o.Yaxis, true
}

// HasYaxis returns a boolean if a field has been set.
func (o *PointPlotBackendDefinition) HasYaxis() bool {
	return o != nil && o.Yaxis != nil
}

// SetYaxis gets a reference to the given Axis and assigns it to the Yaxis field.
func (o *PointPlotBackendDefinition) SetYaxis(v Axis) {
	o.Yaxis = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PointPlotBackendDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Legend != nil {
		toSerialize["legend"] = o.Legend
	}
	if o.Markers != nil {
		toSerialize["markers"] = o.Markers
	}
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
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
	if o.Yaxis != nil {
		toSerialize["yaxis"] = o.Yaxis
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PointPlotBackendDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string                           `json:"description,omitempty"`
		Legend      *ScatterPlotLegendType            `json:"legend,omitempty"`
		Markers     []Marker                          `json:"markers,omitempty"`
		Notes       *string                           `json:"notes,omitempty"`
		Requests    *[]PointPlotDataProjectionRequest `json:"requests"`
		Time        *WidgetTime                       `json:"time,omitempty"`
		Title       *string                           `json:"title,omitempty"`
		TitleAlign  *WidgetTextAlign                  `json:"title_align,omitempty"`
		TitleSize   *string                           `json:"title_size,omitempty"`
		Type        *PointPlotBackendDefinitionType   `json:"type"`
		Yaxis       *Axis                             `json:"yaxis,omitempty"`
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
	o.Description = all.Description
	o.Legend = all.Legend
	o.Markers = all.Markers
	o.Notes = all.Notes
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
	if all.Yaxis != nil && all.Yaxis.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Yaxis = all.Yaxis

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
