// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SankeyDefinition Test-only generated schema for the dashboard versioning proof of concept.
type SankeyDefinition struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Description *string `json:"description,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Notes *string `json:"notes,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Requests []interface{} `json:"requests"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ShowOtherLinks *bool `json:"show_other_links,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	SortNodes *bool `json:"sort_nodes,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Time *WidgetTime `json:"time,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Title *string `json:"title,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TitleAlign *WidgetTextAlign `json:"title_align,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TitleSize *string `json:"title_size,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type SankeyDefinitionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewSankeyDefinition instantiates a new SankeyDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSankeyDefinition(requests []interface{}, typeVar SankeyDefinitionType) *SankeyDefinition {
	this := SankeyDefinition{}
	this.Requests = requests
	this.Type = typeVar
	return &this
}

// NewSankeyDefinitionWithDefaults instantiates a new SankeyDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSankeyDefinitionWithDefaults() *SankeyDefinition {
	this := SankeyDefinition{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SankeyDefinition) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SankeyDefinition) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SankeyDefinition) SetDescription(v string) {
	o.Description = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *SankeyDefinition) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyDefinition) GetNotesOk() (*string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *SankeyDefinition) HasNotes() bool {
	return o != nil && o.Notes != nil
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *SankeyDefinition) SetNotes(v string) {
	o.Notes = &v
}

// GetRequests returns the Requests field value.
func (o *SankeyDefinition) GetRequests() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}
	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value
// and a boolean to check if the value has been set.
func (o *SankeyDefinition) GetRequestsOk() (*[]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requests, true
}

// SetRequests sets field value.
func (o *SankeyDefinition) SetRequests(v []interface{}) {
	o.Requests = v
}

// GetShowOtherLinks returns the ShowOtherLinks field value if set, zero value otherwise.
func (o *SankeyDefinition) GetShowOtherLinks() bool {
	if o == nil || o.ShowOtherLinks == nil {
		var ret bool
		return ret
	}
	return *o.ShowOtherLinks
}

// GetShowOtherLinksOk returns a tuple with the ShowOtherLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyDefinition) GetShowOtherLinksOk() (*bool, bool) {
	if o == nil || o.ShowOtherLinks == nil {
		return nil, false
	}
	return o.ShowOtherLinks, true
}

// HasShowOtherLinks returns a boolean if a field has been set.
func (o *SankeyDefinition) HasShowOtherLinks() bool {
	return o != nil && o.ShowOtherLinks != nil
}

// SetShowOtherLinks gets a reference to the given bool and assigns it to the ShowOtherLinks field.
func (o *SankeyDefinition) SetShowOtherLinks(v bool) {
	o.ShowOtherLinks = &v
}

// GetSortNodes returns the SortNodes field value if set, zero value otherwise.
func (o *SankeyDefinition) GetSortNodes() bool {
	if o == nil || o.SortNodes == nil {
		var ret bool
		return ret
	}
	return *o.SortNodes
}

// GetSortNodesOk returns a tuple with the SortNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyDefinition) GetSortNodesOk() (*bool, bool) {
	if o == nil || o.SortNodes == nil {
		return nil, false
	}
	return o.SortNodes, true
}

// HasSortNodes returns a boolean if a field has been set.
func (o *SankeyDefinition) HasSortNodes() bool {
	return o != nil && o.SortNodes != nil
}

// SetSortNodes gets a reference to the given bool and assigns it to the SortNodes field.
func (o *SankeyDefinition) SetSortNodes(v bool) {
	o.SortNodes = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *SankeyDefinition) GetTime() WidgetTime {
	if o == nil || o.Time == nil {
		var ret WidgetTime
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyDefinition) GetTimeOk() (*WidgetTime, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *SankeyDefinition) HasTime() bool {
	return o != nil && o.Time != nil
}

// SetTime gets a reference to the given WidgetTime and assigns it to the Time field.
func (o *SankeyDefinition) SetTime(v WidgetTime) {
	o.Time = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *SankeyDefinition) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyDefinition) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *SankeyDefinition) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *SankeyDefinition) SetTitle(v string) {
	o.Title = &v
}

// GetTitleAlign returns the TitleAlign field value if set, zero value otherwise.
func (o *SankeyDefinition) GetTitleAlign() WidgetTextAlign {
	if o == nil || o.TitleAlign == nil {
		var ret WidgetTextAlign
		return ret
	}
	return *o.TitleAlign
}

// GetTitleAlignOk returns a tuple with the TitleAlign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool) {
	if o == nil || o.TitleAlign == nil {
		return nil, false
	}
	return o.TitleAlign, true
}

// HasTitleAlign returns a boolean if a field has been set.
func (o *SankeyDefinition) HasTitleAlign() bool {
	return o != nil && o.TitleAlign != nil
}

// SetTitleAlign gets a reference to the given WidgetTextAlign and assigns it to the TitleAlign field.
func (o *SankeyDefinition) SetTitleAlign(v WidgetTextAlign) {
	o.TitleAlign = &v
}

// GetTitleSize returns the TitleSize field value if set, zero value otherwise.
func (o *SankeyDefinition) GetTitleSize() string {
	if o == nil || o.TitleSize == nil {
		var ret string
		return ret
	}
	return *o.TitleSize
}

// GetTitleSizeOk returns a tuple with the TitleSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyDefinition) GetTitleSizeOk() (*string, bool) {
	if o == nil || o.TitleSize == nil {
		return nil, false
	}
	return o.TitleSize, true
}

// HasTitleSize returns a boolean if a field has been set.
func (o *SankeyDefinition) HasTitleSize() bool {
	return o != nil && o.TitleSize != nil
}

// SetTitleSize gets a reference to the given string and assigns it to the TitleSize field.
func (o *SankeyDefinition) SetTitleSize(v string) {
	o.TitleSize = &v
}

// GetType returns the Type field value.
func (o *SankeyDefinition) GetType() SankeyDefinitionType {
	if o == nil {
		var ret SankeyDefinitionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SankeyDefinition) GetTypeOk() (*SankeyDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SankeyDefinition) SetType(v SankeyDefinitionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SankeyDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
	}
	toSerialize["requests"] = o.Requests
	if o.ShowOtherLinks != nil {
		toSerialize["show_other_links"] = o.ShowOtherLinks
	}
	if o.SortNodes != nil {
		toSerialize["sort_nodes"] = o.SortNodes
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
func (o *SankeyDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description    *string               `json:"description,omitempty"`
		Notes          *string               `json:"notes,omitempty"`
		Requests       *[]interface{}        `json:"requests"`
		ShowOtherLinks *bool                 `json:"show_other_links,omitempty"`
		SortNodes      *bool                 `json:"sort_nodes,omitempty"`
		Time           *WidgetTime           `json:"time,omitempty"`
		Title          *string               `json:"title,omitempty"`
		TitleAlign     *WidgetTextAlign      `json:"title_align,omitempty"`
		TitleSize      *string               `json:"title_size,omitempty"`
		Type           *SankeyDefinitionType `json:"type"`
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
	o.Notes = all.Notes
	o.Requests = *all.Requests
	o.ShowOtherLinks = all.ShowOtherLinks
	o.SortNodes = all.SortNodes
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
