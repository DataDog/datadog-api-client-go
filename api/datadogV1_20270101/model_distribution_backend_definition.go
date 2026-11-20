// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DistributionBackendDefinition Test-only generated schema for the dashboard versioning proof of concept.
type DistributionBackendDefinition struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	CustomLinks []WidgetCustomLink `json:"custom_links,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Description *string `json:"description,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	LegendColumns []TimeseriesWidgetLegendColumn `json:"legend_columns,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	LegendLayout *DistributionBackendDefinitionLegendLayout `json:"legend_layout,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	LegendSize *string `json:"legend_size,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Markers []WidgetMarker `json:"markers,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Notes *string `json:"notes,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Requests []DistributionBackendRequest `json:"requests"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ShowLegend *bool `json:"show_legend,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Time *WidgetTime `json:"time,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Title *string `json:"title,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TitleAlign *WidgetTextAlign `json:"title_align,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TitleSize *string `json:"title_size,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type DistributionBackendDefinitionType `json:"type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Xaxis *DistributionXAxis `json:"xaxis,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Yaxis *BackendDistributionYAxis `json:"yaxis,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewDistributionBackendDefinition instantiates a new DistributionBackendDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDistributionBackendDefinition(requests []DistributionBackendRequest, typeVar DistributionBackendDefinitionType) *DistributionBackendDefinition {
	this := DistributionBackendDefinition{}
	this.Requests = requests
	this.Type = typeVar
	return &this
}

// NewDistributionBackendDefinitionWithDefaults instantiates a new DistributionBackendDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDistributionBackendDefinitionWithDefaults() *DistributionBackendDefinition {
	this := DistributionBackendDefinition{}
	return &this
}

// GetCustomLinks returns the CustomLinks field value if set, zero value otherwise.
func (o *DistributionBackendDefinition) GetCustomLinks() []WidgetCustomLink {
	if o == nil || o.CustomLinks == nil {
		var ret []WidgetCustomLink
		return ret
	}
	return o.CustomLinks
}

// GetCustomLinksOk returns a tuple with the CustomLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionBackendDefinition) GetCustomLinksOk() (*[]WidgetCustomLink, bool) {
	if o == nil || o.CustomLinks == nil {
		return nil, false
	}
	return &o.CustomLinks, true
}

// HasCustomLinks returns a boolean if a field has been set.
func (o *DistributionBackendDefinition) HasCustomLinks() bool {
	return o != nil && o.CustomLinks != nil
}

// SetCustomLinks gets a reference to the given []WidgetCustomLink and assigns it to the CustomLinks field.
func (o *DistributionBackendDefinition) SetCustomLinks(v []WidgetCustomLink) {
	o.CustomLinks = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DistributionBackendDefinition) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionBackendDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DistributionBackendDefinition) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DistributionBackendDefinition) SetDescription(v string) {
	o.Description = &v
}

// GetLegendColumns returns the LegendColumns field value if set, zero value otherwise.
func (o *DistributionBackendDefinition) GetLegendColumns() []TimeseriesWidgetLegendColumn {
	if o == nil || o.LegendColumns == nil {
		var ret []TimeseriesWidgetLegendColumn
		return ret
	}
	return o.LegendColumns
}

// GetLegendColumnsOk returns a tuple with the LegendColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionBackendDefinition) GetLegendColumnsOk() (*[]TimeseriesWidgetLegendColumn, bool) {
	if o == nil || o.LegendColumns == nil {
		return nil, false
	}
	return &o.LegendColumns, true
}

// HasLegendColumns returns a boolean if a field has been set.
func (o *DistributionBackendDefinition) HasLegendColumns() bool {
	return o != nil && o.LegendColumns != nil
}

// SetLegendColumns gets a reference to the given []TimeseriesWidgetLegendColumn and assigns it to the LegendColumns field.
func (o *DistributionBackendDefinition) SetLegendColumns(v []TimeseriesWidgetLegendColumn) {
	o.LegendColumns = v
}

// GetLegendLayout returns the LegendLayout field value if set, zero value otherwise.
func (o *DistributionBackendDefinition) GetLegendLayout() DistributionBackendDefinitionLegendLayout {
	if o == nil || o.LegendLayout == nil {
		var ret DistributionBackendDefinitionLegendLayout
		return ret
	}
	return *o.LegendLayout
}

// GetLegendLayoutOk returns a tuple with the LegendLayout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionBackendDefinition) GetLegendLayoutOk() (*DistributionBackendDefinitionLegendLayout, bool) {
	if o == nil || o.LegendLayout == nil {
		return nil, false
	}
	return o.LegendLayout, true
}

// HasLegendLayout returns a boolean if a field has been set.
func (o *DistributionBackendDefinition) HasLegendLayout() bool {
	return o != nil && o.LegendLayout != nil
}

// SetLegendLayout gets a reference to the given DistributionBackendDefinitionLegendLayout and assigns it to the LegendLayout field.
func (o *DistributionBackendDefinition) SetLegendLayout(v DistributionBackendDefinitionLegendLayout) {
	o.LegendLayout = &v
}

// GetLegendSize returns the LegendSize field value if set, zero value otherwise.
func (o *DistributionBackendDefinition) GetLegendSize() string {
	if o == nil || o.LegendSize == nil {
		var ret string
		return ret
	}
	return *o.LegendSize
}

// GetLegendSizeOk returns a tuple with the LegendSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionBackendDefinition) GetLegendSizeOk() (*string, bool) {
	if o == nil || o.LegendSize == nil {
		return nil, false
	}
	return o.LegendSize, true
}

// HasLegendSize returns a boolean if a field has been set.
func (o *DistributionBackendDefinition) HasLegendSize() bool {
	return o != nil && o.LegendSize != nil
}

// SetLegendSize gets a reference to the given string and assigns it to the LegendSize field.
func (o *DistributionBackendDefinition) SetLegendSize(v string) {
	o.LegendSize = &v
}

// GetMarkers returns the Markers field value if set, zero value otherwise.
func (o *DistributionBackendDefinition) GetMarkers() []WidgetMarker {
	if o == nil || o.Markers == nil {
		var ret []WidgetMarker
		return ret
	}
	return o.Markers
}

// GetMarkersOk returns a tuple with the Markers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionBackendDefinition) GetMarkersOk() (*[]WidgetMarker, bool) {
	if o == nil || o.Markers == nil {
		return nil, false
	}
	return &o.Markers, true
}

// HasMarkers returns a boolean if a field has been set.
func (o *DistributionBackendDefinition) HasMarkers() bool {
	return o != nil && o.Markers != nil
}

// SetMarkers gets a reference to the given []WidgetMarker and assigns it to the Markers field.
func (o *DistributionBackendDefinition) SetMarkers(v []WidgetMarker) {
	o.Markers = v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *DistributionBackendDefinition) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionBackendDefinition) GetNotesOk() (*string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *DistributionBackendDefinition) HasNotes() bool {
	return o != nil && o.Notes != nil
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *DistributionBackendDefinition) SetNotes(v string) {
	o.Notes = &v
}

// GetRequests returns the Requests field value.
func (o *DistributionBackendDefinition) GetRequests() []DistributionBackendRequest {
	if o == nil {
		var ret []DistributionBackendRequest
		return ret
	}
	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value
// and a boolean to check if the value has been set.
func (o *DistributionBackendDefinition) GetRequestsOk() (*[]DistributionBackendRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requests, true
}

// SetRequests sets field value.
func (o *DistributionBackendDefinition) SetRequests(v []DistributionBackendRequest) {
	o.Requests = v
}

// GetShowLegend returns the ShowLegend field value if set, zero value otherwise.
func (o *DistributionBackendDefinition) GetShowLegend() bool {
	if o == nil || o.ShowLegend == nil {
		var ret bool
		return ret
	}
	return *o.ShowLegend
}

// GetShowLegendOk returns a tuple with the ShowLegend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionBackendDefinition) GetShowLegendOk() (*bool, bool) {
	if o == nil || o.ShowLegend == nil {
		return nil, false
	}
	return o.ShowLegend, true
}

// HasShowLegend returns a boolean if a field has been set.
func (o *DistributionBackendDefinition) HasShowLegend() bool {
	return o != nil && o.ShowLegend != nil
}

// SetShowLegend gets a reference to the given bool and assigns it to the ShowLegend field.
func (o *DistributionBackendDefinition) SetShowLegend(v bool) {
	o.ShowLegend = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *DistributionBackendDefinition) GetTime() WidgetTime {
	if o == nil || o.Time == nil {
		var ret WidgetTime
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionBackendDefinition) GetTimeOk() (*WidgetTime, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *DistributionBackendDefinition) HasTime() bool {
	return o != nil && o.Time != nil
}

// SetTime gets a reference to the given WidgetTime and assigns it to the Time field.
func (o *DistributionBackendDefinition) SetTime(v WidgetTime) {
	o.Time = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *DistributionBackendDefinition) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionBackendDefinition) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *DistributionBackendDefinition) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *DistributionBackendDefinition) SetTitle(v string) {
	o.Title = &v
}

// GetTitleAlign returns the TitleAlign field value if set, zero value otherwise.
func (o *DistributionBackendDefinition) GetTitleAlign() WidgetTextAlign {
	if o == nil || o.TitleAlign == nil {
		var ret WidgetTextAlign
		return ret
	}
	return *o.TitleAlign
}

// GetTitleAlignOk returns a tuple with the TitleAlign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionBackendDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool) {
	if o == nil || o.TitleAlign == nil {
		return nil, false
	}
	return o.TitleAlign, true
}

// HasTitleAlign returns a boolean if a field has been set.
func (o *DistributionBackendDefinition) HasTitleAlign() bool {
	return o != nil && o.TitleAlign != nil
}

// SetTitleAlign gets a reference to the given WidgetTextAlign and assigns it to the TitleAlign field.
func (o *DistributionBackendDefinition) SetTitleAlign(v WidgetTextAlign) {
	o.TitleAlign = &v
}

// GetTitleSize returns the TitleSize field value if set, zero value otherwise.
func (o *DistributionBackendDefinition) GetTitleSize() string {
	if o == nil || o.TitleSize == nil {
		var ret string
		return ret
	}
	return *o.TitleSize
}

// GetTitleSizeOk returns a tuple with the TitleSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionBackendDefinition) GetTitleSizeOk() (*string, bool) {
	if o == nil || o.TitleSize == nil {
		return nil, false
	}
	return o.TitleSize, true
}

// HasTitleSize returns a boolean if a field has been set.
func (o *DistributionBackendDefinition) HasTitleSize() bool {
	return o != nil && o.TitleSize != nil
}

// SetTitleSize gets a reference to the given string and assigns it to the TitleSize field.
func (o *DistributionBackendDefinition) SetTitleSize(v string) {
	o.TitleSize = &v
}

// GetType returns the Type field value.
func (o *DistributionBackendDefinition) GetType() DistributionBackendDefinitionType {
	if o == nil {
		var ret DistributionBackendDefinitionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DistributionBackendDefinition) GetTypeOk() (*DistributionBackendDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *DistributionBackendDefinition) SetType(v DistributionBackendDefinitionType) {
	o.Type = v
}

// GetXaxis returns the Xaxis field value if set, zero value otherwise.
func (o *DistributionBackendDefinition) GetXaxis() DistributionXAxis {
	if o == nil || o.Xaxis == nil {
		var ret DistributionXAxis
		return ret
	}
	return *o.Xaxis
}

// GetXaxisOk returns a tuple with the Xaxis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionBackendDefinition) GetXaxisOk() (*DistributionXAxis, bool) {
	if o == nil || o.Xaxis == nil {
		return nil, false
	}
	return o.Xaxis, true
}

// HasXaxis returns a boolean if a field has been set.
func (o *DistributionBackendDefinition) HasXaxis() bool {
	return o != nil && o.Xaxis != nil
}

// SetXaxis gets a reference to the given DistributionXAxis and assigns it to the Xaxis field.
func (o *DistributionBackendDefinition) SetXaxis(v DistributionXAxis) {
	o.Xaxis = &v
}

// GetYaxis returns the Yaxis field value if set, zero value otherwise.
func (o *DistributionBackendDefinition) GetYaxis() BackendDistributionYAxis {
	if o == nil || o.Yaxis == nil {
		var ret BackendDistributionYAxis
		return ret
	}
	return *o.Yaxis
}

// GetYaxisOk returns a tuple with the Yaxis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionBackendDefinition) GetYaxisOk() (*BackendDistributionYAxis, bool) {
	if o == nil || o.Yaxis == nil {
		return nil, false
	}
	return o.Yaxis, true
}

// HasYaxis returns a boolean if a field has been set.
func (o *DistributionBackendDefinition) HasYaxis() bool {
	return o != nil && o.Yaxis != nil
}

// SetYaxis gets a reference to the given BackendDistributionYAxis and assigns it to the Yaxis field.
func (o *DistributionBackendDefinition) SetYaxis(v BackendDistributionYAxis) {
	o.Yaxis = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DistributionBackendDefinition) MarshalJSON() ([]byte, error) {
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
	if o.LegendColumns != nil {
		toSerialize["legend_columns"] = o.LegendColumns
	}
	if o.LegendLayout != nil {
		toSerialize["legend_layout"] = o.LegendLayout
	}
	if o.LegendSize != nil {
		toSerialize["legend_size"] = o.LegendSize
	}
	if o.Markers != nil {
		toSerialize["markers"] = o.Markers
	}
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
	}
	toSerialize["requests"] = o.Requests
	if o.ShowLegend != nil {
		toSerialize["show_legend"] = o.ShowLegend
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
	if o.Xaxis != nil {
		toSerialize["xaxis"] = o.Xaxis
	}
	if o.Yaxis != nil {
		toSerialize["yaxis"] = o.Yaxis
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DistributionBackendDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CustomLinks   []WidgetCustomLink                         `json:"custom_links,omitempty"`
		Description   *string                                    `json:"description,omitempty"`
		LegendColumns []TimeseriesWidgetLegendColumn             `json:"legend_columns,omitempty"`
		LegendLayout  *DistributionBackendDefinitionLegendLayout `json:"legend_layout,omitempty"`
		LegendSize    *string                                    `json:"legend_size,omitempty"`
		Markers       []WidgetMarker                             `json:"markers,omitempty"`
		Notes         *string                                    `json:"notes,omitempty"`
		Requests      *[]DistributionBackendRequest              `json:"requests"`
		ShowLegend    *bool                                      `json:"show_legend,omitempty"`
		Time          *WidgetTime                                `json:"time,omitempty"`
		Title         *string                                    `json:"title,omitempty"`
		TitleAlign    *WidgetTextAlign                           `json:"title_align,omitempty"`
		TitleSize     *string                                    `json:"title_size,omitempty"`
		Type          *DistributionBackendDefinitionType         `json:"type"`
		Xaxis         *DistributionXAxis                         `json:"xaxis,omitempty"`
		Yaxis         *BackendDistributionYAxis                  `json:"yaxis,omitempty"`
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
	o.CustomLinks = all.CustomLinks
	o.Description = all.Description
	o.LegendColumns = all.LegendColumns
	if all.LegendLayout != nil && !all.LegendLayout.IsValid() {
		hasInvalidField = true
	} else {
		o.LegendLayout = all.LegendLayout
	}
	o.LegendSize = all.LegendSize
	o.Markers = all.Markers
	o.Notes = all.Notes
	o.Requests = *all.Requests
	o.ShowLegend = all.ShowLegend
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
