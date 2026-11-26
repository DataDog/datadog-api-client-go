// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventTimelineBackendDefinition Test-only generated schema for the dashboard versioning proof of concept.
type EventTimelineBackendDefinition struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Description *string `json:"description,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Notes *string `json:"notes,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Query string `json:"query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TagsExecution *EventTimelineBackendDefinitionTagsExecution `json:"tags_execution,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Time *WidgetTime `json:"time,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Title *string `json:"title,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TitleAlign *WidgetTextAlign `json:"title_align,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TitleSize *string `json:"title_size,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type EventTimelineBackendDefinitionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewEventTimelineBackendDefinition instantiates a new EventTimelineBackendDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEventTimelineBackendDefinition(query string, typeVar EventTimelineBackendDefinitionType) *EventTimelineBackendDefinition {
	this := EventTimelineBackendDefinition{}
	this.Query = query
	this.Type = typeVar
	return &this
}

// NewEventTimelineBackendDefinitionWithDefaults instantiates a new EventTimelineBackendDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEventTimelineBackendDefinitionWithDefaults() *EventTimelineBackendDefinition {
	this := EventTimelineBackendDefinition{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EventTimelineBackendDefinition) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventTimelineBackendDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EventTimelineBackendDefinition) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EventTimelineBackendDefinition) SetDescription(v string) {
	o.Description = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *EventTimelineBackendDefinition) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventTimelineBackendDefinition) GetNotesOk() (*string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *EventTimelineBackendDefinition) HasNotes() bool {
	return o != nil && o.Notes != nil
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *EventTimelineBackendDefinition) SetNotes(v string) {
	o.Notes = &v
}

// GetQuery returns the Query field value.
func (o *EventTimelineBackendDefinition) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *EventTimelineBackendDefinition) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *EventTimelineBackendDefinition) SetQuery(v string) {
	o.Query = v
}

// GetTagsExecution returns the TagsExecution field value if set, zero value otherwise.
func (o *EventTimelineBackendDefinition) GetTagsExecution() EventTimelineBackendDefinitionTagsExecution {
	if o == nil || o.TagsExecution == nil {
		var ret EventTimelineBackendDefinitionTagsExecution
		return ret
	}
	return *o.TagsExecution
}

// GetTagsExecutionOk returns a tuple with the TagsExecution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventTimelineBackendDefinition) GetTagsExecutionOk() (*EventTimelineBackendDefinitionTagsExecution, bool) {
	if o == nil || o.TagsExecution == nil {
		return nil, false
	}
	return o.TagsExecution, true
}

// HasTagsExecution returns a boolean if a field has been set.
func (o *EventTimelineBackendDefinition) HasTagsExecution() bool {
	return o != nil && o.TagsExecution != nil
}

// SetTagsExecution gets a reference to the given EventTimelineBackendDefinitionTagsExecution and assigns it to the TagsExecution field.
func (o *EventTimelineBackendDefinition) SetTagsExecution(v EventTimelineBackendDefinitionTagsExecution) {
	o.TagsExecution = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *EventTimelineBackendDefinition) GetTime() WidgetTime {
	if o == nil || o.Time == nil {
		var ret WidgetTime
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventTimelineBackendDefinition) GetTimeOk() (*WidgetTime, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *EventTimelineBackendDefinition) HasTime() bool {
	return o != nil && o.Time != nil
}

// SetTime gets a reference to the given WidgetTime and assigns it to the Time field.
func (o *EventTimelineBackendDefinition) SetTime(v WidgetTime) {
	o.Time = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *EventTimelineBackendDefinition) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventTimelineBackendDefinition) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *EventTimelineBackendDefinition) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *EventTimelineBackendDefinition) SetTitle(v string) {
	o.Title = &v
}

// GetTitleAlign returns the TitleAlign field value if set, zero value otherwise.
func (o *EventTimelineBackendDefinition) GetTitleAlign() WidgetTextAlign {
	if o == nil || o.TitleAlign == nil {
		var ret WidgetTextAlign
		return ret
	}
	return *o.TitleAlign
}

// GetTitleAlignOk returns a tuple with the TitleAlign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventTimelineBackendDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool) {
	if o == nil || o.TitleAlign == nil {
		return nil, false
	}
	return o.TitleAlign, true
}

// HasTitleAlign returns a boolean if a field has been set.
func (o *EventTimelineBackendDefinition) HasTitleAlign() bool {
	return o != nil && o.TitleAlign != nil
}

// SetTitleAlign gets a reference to the given WidgetTextAlign and assigns it to the TitleAlign field.
func (o *EventTimelineBackendDefinition) SetTitleAlign(v WidgetTextAlign) {
	o.TitleAlign = &v
}

// GetTitleSize returns the TitleSize field value if set, zero value otherwise.
func (o *EventTimelineBackendDefinition) GetTitleSize() string {
	if o == nil || o.TitleSize == nil {
		var ret string
		return ret
	}
	return *o.TitleSize
}

// GetTitleSizeOk returns a tuple with the TitleSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventTimelineBackendDefinition) GetTitleSizeOk() (*string, bool) {
	if o == nil || o.TitleSize == nil {
		return nil, false
	}
	return o.TitleSize, true
}

// HasTitleSize returns a boolean if a field has been set.
func (o *EventTimelineBackendDefinition) HasTitleSize() bool {
	return o != nil && o.TitleSize != nil
}

// SetTitleSize gets a reference to the given string and assigns it to the TitleSize field.
func (o *EventTimelineBackendDefinition) SetTitleSize(v string) {
	o.TitleSize = &v
}

// GetType returns the Type field value.
func (o *EventTimelineBackendDefinition) GetType() EventTimelineBackendDefinitionType {
	if o == nil {
		var ret EventTimelineBackendDefinitionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EventTimelineBackendDefinition) GetTypeOk() (*EventTimelineBackendDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *EventTimelineBackendDefinition) SetType(v EventTimelineBackendDefinitionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EventTimelineBackendDefinition) MarshalJSON() ([]byte, error) {
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
	toSerialize["query"] = o.Query
	if o.TagsExecution != nil {
		toSerialize["tags_execution"] = o.TagsExecution
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
func (o *EventTimelineBackendDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description   *string                                      `json:"description,omitempty"`
		Notes         *string                                      `json:"notes,omitempty"`
		Query         *string                                      `json:"query"`
		TagsExecution *EventTimelineBackendDefinitionTagsExecution `json:"tags_execution,omitempty"`
		Time          *WidgetTime                                  `json:"time,omitempty"`
		Title         *string                                      `json:"title,omitempty"`
		TitleAlign    *WidgetTextAlign                             `json:"title_align,omitempty"`
		TitleSize     *string                                      `json:"title_size,omitempty"`
		Type          *EventTimelineBackendDefinitionType          `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	o.Description = all.Description
	o.Notes = all.Notes
	o.Query = *all.Query
	if all.TagsExecution != nil && !all.TagsExecution.IsValid() {
		hasInvalidField = true
	} else {
		o.TagsExecution = all.TagsExecution
	}
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
