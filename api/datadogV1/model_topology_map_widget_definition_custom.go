// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TopologyMapWidgetDefinitionCustom Custom topology map widget. Renders nodes and edges from inline `data` and from scalar `requests`. There is one node for each node ID found in `data.nodes`, `data.edges`, and `requests`. When `data` and several requests set the same field on the same node (such as `title` or `values[0]`), the last non-null value wins.
type TopologyMapWidgetDefinitionCustom struct {
	// List of custom links.
	CustomLinks []WidgetCustomLink `json:"custom_links,omitempty"`
	// Static nodes and edges, stored in columns. Within a table, every column must have the same number of values and a unique `id`.
	Data *TopologyMapInlineData `json:"data,omitempty"`
	// The description of the widget.
	Description *string `json:"description,omitempty"`
	// Algorithm that positions nodes and edges. The `hierarchical` layout places nodes in ranks so edges flow in one direction.
	Layout TopologyMapLayout `json:"layout"`
	// List of scalar requests that supply node data. Can be empty.
	Requests []TopologyRequestCustom `json:"requests"`
	// Time setting for the widget.
	Time *WidgetTime `json:"time,omitempty"`
	// Title of your widget.
	Title *string `json:"title,omitempty"`
	// How to align the text on the widget.
	TitleAlign *WidgetTextAlign `json:"title_align,omitempty"`
	// Size of the title.
	TitleSize *string `json:"title_size,omitempty"`
	// Type of the topology map widget.
	Type TopologyMapWidgetDefinitionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTopologyMapWidgetDefinitionCustom instantiates a new TopologyMapWidgetDefinitionCustom object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTopologyMapWidgetDefinitionCustom(layout TopologyMapLayout, requests []TopologyRequestCustom, typeVar TopologyMapWidgetDefinitionType) *TopologyMapWidgetDefinitionCustom {
	this := TopologyMapWidgetDefinitionCustom{}
	this.Layout = layout
	this.Requests = requests
	this.Type = typeVar
	return &this
}

// NewTopologyMapWidgetDefinitionCustomWithDefaults instantiates a new TopologyMapWidgetDefinitionCustom object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTopologyMapWidgetDefinitionCustomWithDefaults() *TopologyMapWidgetDefinitionCustom {
	this := TopologyMapWidgetDefinitionCustom{}
	var typeVar TopologyMapWidgetDefinitionType = TOPOLOGYMAPWIDGETDEFINITIONTYPE_TOPOLOGY_MAP
	this.Type = typeVar
	return &this
}

// GetCustomLinks returns the CustomLinks field value if set, zero value otherwise.
func (o *TopologyMapWidgetDefinitionCustom) GetCustomLinks() []WidgetCustomLink {
	if o == nil || o.CustomLinks == nil {
		var ret []WidgetCustomLink
		return ret
	}
	return o.CustomLinks
}

// GetCustomLinksOk returns a tuple with the CustomLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopologyMapWidgetDefinitionCustom) GetCustomLinksOk() (*[]WidgetCustomLink, bool) {
	if o == nil || o.CustomLinks == nil {
		return nil, false
	}
	return &o.CustomLinks, true
}

// HasCustomLinks returns a boolean if a field has been set.
func (o *TopologyMapWidgetDefinitionCustom) HasCustomLinks() bool {
	return o != nil && o.CustomLinks != nil
}

// SetCustomLinks gets a reference to the given []WidgetCustomLink and assigns it to the CustomLinks field.
func (o *TopologyMapWidgetDefinitionCustom) SetCustomLinks(v []WidgetCustomLink) {
	o.CustomLinks = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *TopologyMapWidgetDefinitionCustom) GetData() TopologyMapInlineData {
	if o == nil || o.Data == nil {
		var ret TopologyMapInlineData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopologyMapWidgetDefinitionCustom) GetDataOk() (*TopologyMapInlineData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *TopologyMapWidgetDefinitionCustom) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given TopologyMapInlineData and assigns it to the Data field.
func (o *TopologyMapWidgetDefinitionCustom) SetData(v TopologyMapInlineData) {
	o.Data = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TopologyMapWidgetDefinitionCustom) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopologyMapWidgetDefinitionCustom) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TopologyMapWidgetDefinitionCustom) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TopologyMapWidgetDefinitionCustom) SetDescription(v string) {
	o.Description = &v
}

// GetLayout returns the Layout field value.
func (o *TopologyMapWidgetDefinitionCustom) GetLayout() TopologyMapLayout {
	if o == nil {
		var ret TopologyMapLayout
		return ret
	}
	return o.Layout
}

// GetLayoutOk returns a tuple with the Layout field value
// and a boolean to check if the value has been set.
func (o *TopologyMapWidgetDefinitionCustom) GetLayoutOk() (*TopologyMapLayout, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Layout, true
}

// SetLayout sets field value.
func (o *TopologyMapWidgetDefinitionCustom) SetLayout(v TopologyMapLayout) {
	o.Layout = v
}

// GetRequests returns the Requests field value.
func (o *TopologyMapWidgetDefinitionCustom) GetRequests() []TopologyRequestCustom {
	if o == nil {
		var ret []TopologyRequestCustom
		return ret
	}
	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value
// and a boolean to check if the value has been set.
func (o *TopologyMapWidgetDefinitionCustom) GetRequestsOk() (*[]TopologyRequestCustom, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requests, true
}

// SetRequests sets field value.
func (o *TopologyMapWidgetDefinitionCustom) SetRequests(v []TopologyRequestCustom) {
	o.Requests = v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *TopologyMapWidgetDefinitionCustom) GetTime() WidgetTime {
	if o == nil || o.Time == nil {
		var ret WidgetTime
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopologyMapWidgetDefinitionCustom) GetTimeOk() (*WidgetTime, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *TopologyMapWidgetDefinitionCustom) HasTime() bool {
	return o != nil && o.Time != nil
}

// SetTime gets a reference to the given WidgetTime and assigns it to the Time field.
func (o *TopologyMapWidgetDefinitionCustom) SetTime(v WidgetTime) {
	o.Time = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *TopologyMapWidgetDefinitionCustom) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopologyMapWidgetDefinitionCustom) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *TopologyMapWidgetDefinitionCustom) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *TopologyMapWidgetDefinitionCustom) SetTitle(v string) {
	o.Title = &v
}

// GetTitleAlign returns the TitleAlign field value if set, zero value otherwise.
func (o *TopologyMapWidgetDefinitionCustom) GetTitleAlign() WidgetTextAlign {
	if o == nil || o.TitleAlign == nil {
		var ret WidgetTextAlign
		return ret
	}
	return *o.TitleAlign
}

// GetTitleAlignOk returns a tuple with the TitleAlign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopologyMapWidgetDefinitionCustom) GetTitleAlignOk() (*WidgetTextAlign, bool) {
	if o == nil || o.TitleAlign == nil {
		return nil, false
	}
	return o.TitleAlign, true
}

// HasTitleAlign returns a boolean if a field has been set.
func (o *TopologyMapWidgetDefinitionCustom) HasTitleAlign() bool {
	return o != nil && o.TitleAlign != nil
}

// SetTitleAlign gets a reference to the given WidgetTextAlign and assigns it to the TitleAlign field.
func (o *TopologyMapWidgetDefinitionCustom) SetTitleAlign(v WidgetTextAlign) {
	o.TitleAlign = &v
}

// GetTitleSize returns the TitleSize field value if set, zero value otherwise.
func (o *TopologyMapWidgetDefinitionCustom) GetTitleSize() string {
	if o == nil || o.TitleSize == nil {
		var ret string
		return ret
	}
	return *o.TitleSize
}

// GetTitleSizeOk returns a tuple with the TitleSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopologyMapWidgetDefinitionCustom) GetTitleSizeOk() (*string, bool) {
	if o == nil || o.TitleSize == nil {
		return nil, false
	}
	return o.TitleSize, true
}

// HasTitleSize returns a boolean if a field has been set.
func (o *TopologyMapWidgetDefinitionCustom) HasTitleSize() bool {
	return o != nil && o.TitleSize != nil
}

// SetTitleSize gets a reference to the given string and assigns it to the TitleSize field.
func (o *TopologyMapWidgetDefinitionCustom) SetTitleSize(v string) {
	o.TitleSize = &v
}

// GetType returns the Type field value.
func (o *TopologyMapWidgetDefinitionCustom) GetType() TopologyMapWidgetDefinitionType {
	if o == nil {
		var ret TopologyMapWidgetDefinitionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TopologyMapWidgetDefinitionCustom) GetTypeOk() (*TopologyMapWidgetDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *TopologyMapWidgetDefinitionCustom) SetType(v TopologyMapWidgetDefinitionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TopologyMapWidgetDefinitionCustom) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CustomLinks != nil {
		toSerialize["custom_links"] = o.CustomLinks
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	toSerialize["layout"] = o.Layout
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
func (o *TopologyMapWidgetDefinitionCustom) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CustomLinks []WidgetCustomLink               `json:"custom_links,omitempty"`
		Data        *TopologyMapInlineData           `json:"data,omitempty"`
		Description *string                          `json:"description,omitempty"`
		Layout      *TopologyMapLayout               `json:"layout"`
		Requests    *[]TopologyRequestCustom         `json:"requests"`
		Time        *WidgetTime                      `json:"time,omitempty"`
		Title       *string                          `json:"title,omitempty"`
		TitleAlign  *WidgetTextAlign                 `json:"title_align,omitempty"`
		TitleSize   *string                          `json:"title_size,omitempty"`
		Type        *TopologyMapWidgetDefinitionType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Layout == nil {
		return fmt.Errorf("required field layout missing")
	}
	if all.Requests == nil {
		return fmt.Errorf("required field requests missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	o.CustomLinks = all.CustomLinks
	if all.Data != nil && all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Data = all.Data
	o.Description = all.Description
	if all.Layout.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Layout = *all.Layout
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
