// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridWidgetDefinition The heatgrid visualization displays values for each group over time using color.
type HeatgridWidgetDefinition struct {
	// Color configuration for continuous gradients or discrete thresholds.
	Color *HeatgridColorConfig `json:"color,omitempty"`
	// List of custom links.
	CustomLinks []WidgetCustomLink `json:"custom_links,omitempty"`
	// Description of the widget.
	Description *string `json:"description,omitempty"`
	// Configuration of the group label column.
	LabelColumn *HeatgridLabelColumn `json:"label_column,omitempty"`
	// Legend configuration for the heatgrid widget.
	Legend *HeatgridLegend `json:"legend,omitempty"`
	// Widget requests. The widget displays one formula, which can combine multiple queries.
	Requests []HeatgridWidgetRequest `json:"requests"`
	// Ordering of the heatgrid rows.
	Sort HeatgridSort `json:"sort"`
	// Time setting for the widget.
	Time *WidgetTime `json:"time,omitempty"`
	// Title of the widget.
	Title *string `json:"title,omitempty"`
	// How to align the text on the widget.
	TitleAlign *WidgetTextAlign `json:"title_align,omitempty"`
	// Size of the title.
	TitleSize *string `json:"title_size,omitempty"`
	// Type of the heatgrid widget.
	Type HeatgridWidgetDefinitionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewHeatgridWidgetDefinition instantiates a new HeatgridWidgetDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHeatgridWidgetDefinition(requests []HeatgridWidgetRequest, sort HeatgridSort, typeVar HeatgridWidgetDefinitionType) *HeatgridWidgetDefinition {
	this := HeatgridWidgetDefinition{}
	this.Requests = requests
	this.Sort = sort
	this.Type = typeVar
	return &this
}

// NewHeatgridWidgetDefinitionWithDefaults instantiates a new HeatgridWidgetDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHeatgridWidgetDefinitionWithDefaults() *HeatgridWidgetDefinition {
	this := HeatgridWidgetDefinition{}
	var typeVar HeatgridWidgetDefinitionType = HEATGRIDWIDGETDEFINITIONTYPE_HEATGRID
	this.Type = typeVar
	return &this
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *HeatgridWidgetDefinition) GetColor() HeatgridColorConfig {
	if o == nil || o.Color == nil {
		var ret HeatgridColorConfig
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatgridWidgetDefinition) GetColorOk() (*HeatgridColorConfig, bool) {
	if o == nil || o.Color == nil {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *HeatgridWidgetDefinition) HasColor() bool {
	return o != nil && o.Color != nil
}

// SetColor gets a reference to the given HeatgridColorConfig and assigns it to the Color field.
func (o *HeatgridWidgetDefinition) SetColor(v HeatgridColorConfig) {
	o.Color = &v
}

// GetCustomLinks returns the CustomLinks field value if set, zero value otherwise.
func (o *HeatgridWidgetDefinition) GetCustomLinks() []WidgetCustomLink {
	if o == nil || o.CustomLinks == nil {
		var ret []WidgetCustomLink
		return ret
	}
	return o.CustomLinks
}

// GetCustomLinksOk returns a tuple with the CustomLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatgridWidgetDefinition) GetCustomLinksOk() (*[]WidgetCustomLink, bool) {
	if o == nil || o.CustomLinks == nil {
		return nil, false
	}
	return &o.CustomLinks, true
}

// HasCustomLinks returns a boolean if a field has been set.
func (o *HeatgridWidgetDefinition) HasCustomLinks() bool {
	return o != nil && o.CustomLinks != nil
}

// SetCustomLinks gets a reference to the given []WidgetCustomLink and assigns it to the CustomLinks field.
func (o *HeatgridWidgetDefinition) SetCustomLinks(v []WidgetCustomLink) {
	o.CustomLinks = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *HeatgridWidgetDefinition) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatgridWidgetDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *HeatgridWidgetDefinition) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *HeatgridWidgetDefinition) SetDescription(v string) {
	o.Description = &v
}

// GetLabelColumn returns the LabelColumn field value if set, zero value otherwise.
func (o *HeatgridWidgetDefinition) GetLabelColumn() HeatgridLabelColumn {
	if o == nil || o.LabelColumn == nil {
		var ret HeatgridLabelColumn
		return ret
	}
	return *o.LabelColumn
}

// GetLabelColumnOk returns a tuple with the LabelColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatgridWidgetDefinition) GetLabelColumnOk() (*HeatgridLabelColumn, bool) {
	if o == nil || o.LabelColumn == nil {
		return nil, false
	}
	return o.LabelColumn, true
}

// HasLabelColumn returns a boolean if a field has been set.
func (o *HeatgridWidgetDefinition) HasLabelColumn() bool {
	return o != nil && o.LabelColumn != nil
}

// SetLabelColumn gets a reference to the given HeatgridLabelColumn and assigns it to the LabelColumn field.
func (o *HeatgridWidgetDefinition) SetLabelColumn(v HeatgridLabelColumn) {
	o.LabelColumn = &v
}

// GetLegend returns the Legend field value if set, zero value otherwise.
func (o *HeatgridWidgetDefinition) GetLegend() HeatgridLegend {
	if o == nil || o.Legend == nil {
		var ret HeatgridLegend
		return ret
	}
	return *o.Legend
}

// GetLegendOk returns a tuple with the Legend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatgridWidgetDefinition) GetLegendOk() (*HeatgridLegend, bool) {
	if o == nil || o.Legend == nil {
		return nil, false
	}
	return o.Legend, true
}

// HasLegend returns a boolean if a field has been set.
func (o *HeatgridWidgetDefinition) HasLegend() bool {
	return o != nil && o.Legend != nil
}

// SetLegend gets a reference to the given HeatgridLegend and assigns it to the Legend field.
func (o *HeatgridWidgetDefinition) SetLegend(v HeatgridLegend) {
	o.Legend = &v
}

// GetRequests returns the Requests field value.
func (o *HeatgridWidgetDefinition) GetRequests() []HeatgridWidgetRequest {
	if o == nil {
		var ret []HeatgridWidgetRequest
		return ret
	}
	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value
// and a boolean to check if the value has been set.
func (o *HeatgridWidgetDefinition) GetRequestsOk() (*[]HeatgridWidgetRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requests, true
}

// SetRequests sets field value.
func (o *HeatgridWidgetDefinition) SetRequests(v []HeatgridWidgetRequest) {
	o.Requests = v
}

// GetSort returns the Sort field value.
func (o *HeatgridWidgetDefinition) GetSort() HeatgridSort {
	if o == nil {
		var ret HeatgridSort
		return ret
	}
	return o.Sort
}

// GetSortOk returns a tuple with the Sort field value
// and a boolean to check if the value has been set.
func (o *HeatgridWidgetDefinition) GetSortOk() (*HeatgridSort, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sort, true
}

// SetSort sets field value.
func (o *HeatgridWidgetDefinition) SetSort(v HeatgridSort) {
	o.Sort = v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *HeatgridWidgetDefinition) GetTime() WidgetTime {
	if o == nil || o.Time == nil {
		var ret WidgetTime
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatgridWidgetDefinition) GetTimeOk() (*WidgetTime, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *HeatgridWidgetDefinition) HasTime() bool {
	return o != nil && o.Time != nil
}

// SetTime gets a reference to the given WidgetTime and assigns it to the Time field.
func (o *HeatgridWidgetDefinition) SetTime(v WidgetTime) {
	o.Time = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *HeatgridWidgetDefinition) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatgridWidgetDefinition) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *HeatgridWidgetDefinition) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *HeatgridWidgetDefinition) SetTitle(v string) {
	o.Title = &v
}

// GetTitleAlign returns the TitleAlign field value if set, zero value otherwise.
func (o *HeatgridWidgetDefinition) GetTitleAlign() WidgetTextAlign {
	if o == nil || o.TitleAlign == nil {
		var ret WidgetTextAlign
		return ret
	}
	return *o.TitleAlign
}

// GetTitleAlignOk returns a tuple with the TitleAlign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatgridWidgetDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool) {
	if o == nil || o.TitleAlign == nil {
		return nil, false
	}
	return o.TitleAlign, true
}

// HasTitleAlign returns a boolean if a field has been set.
func (o *HeatgridWidgetDefinition) HasTitleAlign() bool {
	return o != nil && o.TitleAlign != nil
}

// SetTitleAlign gets a reference to the given WidgetTextAlign and assigns it to the TitleAlign field.
func (o *HeatgridWidgetDefinition) SetTitleAlign(v WidgetTextAlign) {
	o.TitleAlign = &v
}

// GetTitleSize returns the TitleSize field value if set, zero value otherwise.
func (o *HeatgridWidgetDefinition) GetTitleSize() string {
	if o == nil || o.TitleSize == nil {
		var ret string
		return ret
	}
	return *o.TitleSize
}

// GetTitleSizeOk returns a tuple with the TitleSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatgridWidgetDefinition) GetTitleSizeOk() (*string, bool) {
	if o == nil || o.TitleSize == nil {
		return nil, false
	}
	return o.TitleSize, true
}

// HasTitleSize returns a boolean if a field has been set.
func (o *HeatgridWidgetDefinition) HasTitleSize() bool {
	return o != nil && o.TitleSize != nil
}

// SetTitleSize gets a reference to the given string and assigns it to the TitleSize field.
func (o *HeatgridWidgetDefinition) SetTitleSize(v string) {
	o.TitleSize = &v
}

// GetType returns the Type field value.
func (o *HeatgridWidgetDefinition) GetType() HeatgridWidgetDefinitionType {
	if o == nil {
		var ret HeatgridWidgetDefinitionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *HeatgridWidgetDefinition) GetTypeOk() (*HeatgridWidgetDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *HeatgridWidgetDefinition) SetType(v HeatgridWidgetDefinitionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o HeatgridWidgetDefinition) MarshalJSON() ([]byte, error) {
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
func (o *HeatgridWidgetDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Color       *HeatgridColorConfig          `json:"color,omitempty"`
		CustomLinks []WidgetCustomLink            `json:"custom_links,omitempty"`
		Description *string                       `json:"description,omitempty"`
		LabelColumn *HeatgridLabelColumn          `json:"label_column,omitempty"`
		Legend      *HeatgridLegend               `json:"legend,omitempty"`
		Requests    *[]HeatgridWidgetRequest      `json:"requests"`
		Sort        *HeatgridSort                 `json:"sort"`
		Time        *WidgetTime                   `json:"time,omitempty"`
		Title       *string                       `json:"title,omitempty"`
		TitleAlign  *WidgetTextAlign              `json:"title_align,omitempty"`
		TitleSize   *string                       `json:"title_size,omitempty"`
		Type        *HeatgridWidgetDefinitionType `json:"type"`
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
