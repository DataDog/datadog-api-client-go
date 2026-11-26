// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridDefinition Test-only generated schema for the dashboard versioning proof of concept.
type HeatgridDefinition struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Color *HeatgridColorConfig `json:"color,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	CustomLinks []WidgetCustomLink `json:"custom_links,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Description *string `json:"description,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	LabelColumn *HeatgridLabelColumnConfig `json:"label_column,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Legend *HeatgridLegendConfig `json:"legend,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Notes *string `json:"notes,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Requests []HeatgridTileDefRequest `json:"requests"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Sort HeatgridSortAndOrderBy `json:"sort"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Time *WidgetTime `json:"time,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Title *string `json:"title,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TitleAlign *WidgetTextAlign `json:"title_align,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TitleSize *string `json:"title_size,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type HeatgridDefinitionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewHeatgridDefinition instantiates a new HeatgridDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHeatgridDefinition(requests []HeatgridTileDefRequest, sort HeatgridSortAndOrderBy, typeVar HeatgridDefinitionType) *HeatgridDefinition {
	this := HeatgridDefinition{}
	this.Requests = requests
	this.Sort = sort
	this.Type = typeVar
	return &this
}

// NewHeatgridDefinitionWithDefaults instantiates a new HeatgridDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHeatgridDefinitionWithDefaults() *HeatgridDefinition {
	this := HeatgridDefinition{}
	return &this
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *HeatgridDefinition) GetColor() HeatgridColorConfig {
	if o == nil || o.Color == nil {
		var ret HeatgridColorConfig
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatgridDefinition) GetColorOk() (*HeatgridColorConfig, bool) {
	if o == nil || o.Color == nil {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *HeatgridDefinition) HasColor() bool {
	return o != nil && o.Color != nil
}

// SetColor gets a reference to the given HeatgridColorConfig and assigns it to the Color field.
func (o *HeatgridDefinition) SetColor(v HeatgridColorConfig) {
	o.Color = &v
}

// GetCustomLinks returns the CustomLinks field value if set, zero value otherwise.
func (o *HeatgridDefinition) GetCustomLinks() []WidgetCustomLink {
	if o == nil || o.CustomLinks == nil {
		var ret []WidgetCustomLink
		return ret
	}
	return o.CustomLinks
}

// GetCustomLinksOk returns a tuple with the CustomLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatgridDefinition) GetCustomLinksOk() (*[]WidgetCustomLink, bool) {
	if o == nil || o.CustomLinks == nil {
		return nil, false
	}
	return &o.CustomLinks, true
}

// HasCustomLinks returns a boolean if a field has been set.
func (o *HeatgridDefinition) HasCustomLinks() bool {
	return o != nil && o.CustomLinks != nil
}

// SetCustomLinks gets a reference to the given []WidgetCustomLink and assigns it to the CustomLinks field.
func (o *HeatgridDefinition) SetCustomLinks(v []WidgetCustomLink) {
	o.CustomLinks = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *HeatgridDefinition) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatgridDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *HeatgridDefinition) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *HeatgridDefinition) SetDescription(v string) {
	o.Description = &v
}

// GetLabelColumn returns the LabelColumn field value if set, zero value otherwise.
func (o *HeatgridDefinition) GetLabelColumn() HeatgridLabelColumnConfig {
	if o == nil || o.LabelColumn == nil {
		var ret HeatgridLabelColumnConfig
		return ret
	}
	return *o.LabelColumn
}

// GetLabelColumnOk returns a tuple with the LabelColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatgridDefinition) GetLabelColumnOk() (*HeatgridLabelColumnConfig, bool) {
	if o == nil || o.LabelColumn == nil {
		return nil, false
	}
	return o.LabelColumn, true
}

// HasLabelColumn returns a boolean if a field has been set.
func (o *HeatgridDefinition) HasLabelColumn() bool {
	return o != nil && o.LabelColumn != nil
}

// SetLabelColumn gets a reference to the given HeatgridLabelColumnConfig and assigns it to the LabelColumn field.
func (o *HeatgridDefinition) SetLabelColumn(v HeatgridLabelColumnConfig) {
	o.LabelColumn = &v
}

// GetLegend returns the Legend field value if set, zero value otherwise.
func (o *HeatgridDefinition) GetLegend() HeatgridLegendConfig {
	if o == nil || o.Legend == nil {
		var ret HeatgridLegendConfig
		return ret
	}
	return *o.Legend
}

// GetLegendOk returns a tuple with the Legend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatgridDefinition) GetLegendOk() (*HeatgridLegendConfig, bool) {
	if o == nil || o.Legend == nil {
		return nil, false
	}
	return o.Legend, true
}

// HasLegend returns a boolean if a field has been set.
func (o *HeatgridDefinition) HasLegend() bool {
	return o != nil && o.Legend != nil
}

// SetLegend gets a reference to the given HeatgridLegendConfig and assigns it to the Legend field.
func (o *HeatgridDefinition) SetLegend(v HeatgridLegendConfig) {
	o.Legend = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *HeatgridDefinition) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatgridDefinition) GetNotesOk() (*string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *HeatgridDefinition) HasNotes() bool {
	return o != nil && o.Notes != nil
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *HeatgridDefinition) SetNotes(v string) {
	o.Notes = &v
}

// GetRequests returns the Requests field value.
func (o *HeatgridDefinition) GetRequests() []HeatgridTileDefRequest {
	if o == nil {
		var ret []HeatgridTileDefRequest
		return ret
	}
	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value
// and a boolean to check if the value has been set.
func (o *HeatgridDefinition) GetRequestsOk() (*[]HeatgridTileDefRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requests, true
}

// SetRequests sets field value.
func (o *HeatgridDefinition) SetRequests(v []HeatgridTileDefRequest) {
	o.Requests = v
}

// GetSort returns the Sort field value.
func (o *HeatgridDefinition) GetSort() HeatgridSortAndOrderBy {
	if o == nil {
		var ret HeatgridSortAndOrderBy
		return ret
	}
	return o.Sort
}

// GetSortOk returns a tuple with the Sort field value
// and a boolean to check if the value has been set.
func (o *HeatgridDefinition) GetSortOk() (*HeatgridSortAndOrderBy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sort, true
}

// SetSort sets field value.
func (o *HeatgridDefinition) SetSort(v HeatgridSortAndOrderBy) {
	o.Sort = v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *HeatgridDefinition) GetTime() WidgetTime {
	if o == nil || o.Time == nil {
		var ret WidgetTime
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatgridDefinition) GetTimeOk() (*WidgetTime, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *HeatgridDefinition) HasTime() bool {
	return o != nil && o.Time != nil
}

// SetTime gets a reference to the given WidgetTime and assigns it to the Time field.
func (o *HeatgridDefinition) SetTime(v WidgetTime) {
	o.Time = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *HeatgridDefinition) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatgridDefinition) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *HeatgridDefinition) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *HeatgridDefinition) SetTitle(v string) {
	o.Title = &v
}

// GetTitleAlign returns the TitleAlign field value if set, zero value otherwise.
func (o *HeatgridDefinition) GetTitleAlign() WidgetTextAlign {
	if o == nil || o.TitleAlign == nil {
		var ret WidgetTextAlign
		return ret
	}
	return *o.TitleAlign
}

// GetTitleAlignOk returns a tuple with the TitleAlign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatgridDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool) {
	if o == nil || o.TitleAlign == nil {
		return nil, false
	}
	return o.TitleAlign, true
}

// HasTitleAlign returns a boolean if a field has been set.
func (o *HeatgridDefinition) HasTitleAlign() bool {
	return o != nil && o.TitleAlign != nil
}

// SetTitleAlign gets a reference to the given WidgetTextAlign and assigns it to the TitleAlign field.
func (o *HeatgridDefinition) SetTitleAlign(v WidgetTextAlign) {
	o.TitleAlign = &v
}

// GetTitleSize returns the TitleSize field value if set, zero value otherwise.
func (o *HeatgridDefinition) GetTitleSize() string {
	if o == nil || o.TitleSize == nil {
		var ret string
		return ret
	}
	return *o.TitleSize
}

// GetTitleSizeOk returns a tuple with the TitleSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatgridDefinition) GetTitleSizeOk() (*string, bool) {
	if o == nil || o.TitleSize == nil {
		return nil, false
	}
	return o.TitleSize, true
}

// HasTitleSize returns a boolean if a field has been set.
func (o *HeatgridDefinition) HasTitleSize() bool {
	return o != nil && o.TitleSize != nil
}

// SetTitleSize gets a reference to the given string and assigns it to the TitleSize field.
func (o *HeatgridDefinition) SetTitleSize(v string) {
	o.TitleSize = &v
}

// GetType returns the Type field value.
func (o *HeatgridDefinition) GetType() HeatgridDefinitionType {
	if o == nil {
		var ret HeatgridDefinitionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *HeatgridDefinition) GetTypeOk() (*HeatgridDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *HeatgridDefinition) SetType(v HeatgridDefinitionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o HeatgridDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Color != nil {
		toSerialize["color"] = o.Color
	}
	if o.CustomLinks != nil {
		toSerialize["custom_links"] = o.CustomLinks
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.LabelColumn != nil {
		toSerialize["label_column"] = o.LabelColumn
	}
	if o.Legend != nil {
		toSerialize["legend"] = o.Legend
	}
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
	}
	toSerialize["requests"] = o.Requests
	toSerialize["sort"] = o.Sort
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
func (o *HeatgridDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Color       *HeatgridColorConfig       `json:"color,omitempty"`
		CustomLinks []WidgetCustomLink         `json:"custom_links,omitempty"`
		Description *string                    `json:"description,omitempty"`
		LabelColumn *HeatgridLabelColumnConfig `json:"label_column,omitempty"`
		Legend      *HeatgridLegendConfig      `json:"legend,omitempty"`
		Notes       *string                    `json:"notes,omitempty"`
		Requests    *[]HeatgridTileDefRequest  `json:"requests"`
		Sort        *HeatgridSortAndOrderBy    `json:"sort"`
		Time        *WidgetTime                `json:"time,omitempty"`
		Title       *string                    `json:"title,omitempty"`
		TitleAlign  *WidgetTextAlign           `json:"title_align,omitempty"`
		TitleSize   *string                    `json:"title_size,omitempty"`
		Type        *HeatgridDefinitionType    `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Requests == nil {
		return fmt.Errorf("required field requests missing")
	}
	if all.Sort == nil {
		return fmt.Errorf("required field sort missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	o.Color = all.Color
	o.CustomLinks = all.CustomLinks
	o.Description = all.Description
	if all.LabelColumn != nil && all.LabelColumn.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LabelColumn = all.LabelColumn
	if all.Legend != nil && all.Legend.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Legend = all.Legend
	o.Notes = all.Notes
	o.Requests = *all.Requests
	if all.Sort.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Sort = *all.Sort
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
