// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// JourneyMapDefinition Test-only generated schema for the dashboard versioning proof of concept.
type JourneyMapDefinition struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Description *string `json:"description,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Notes *string `json:"notes,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Requests []interface{} `json:"requests"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Time *WidgetTime `json:"time,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Title *string `json:"title,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TitleAlign *WidgetTextAlign `json:"title_align,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TitleSize *string `json:"title_size,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type JourneyMapDefinitionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewJourneyMapDefinition instantiates a new JourneyMapDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewJourneyMapDefinition(requests []interface{}, typeVar JourneyMapDefinitionType) *JourneyMapDefinition {
	this := JourneyMapDefinition{}
	this.Requests = requests
	this.Type = typeVar
	return &this
}

// NewJourneyMapDefinitionWithDefaults instantiates a new JourneyMapDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewJourneyMapDefinitionWithDefaults() *JourneyMapDefinition {
	this := JourneyMapDefinition{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *JourneyMapDefinition) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JourneyMapDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *JourneyMapDefinition) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *JourneyMapDefinition) SetDescription(v string) {
	o.Description = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *JourneyMapDefinition) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JourneyMapDefinition) GetNotesOk() (*string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *JourneyMapDefinition) HasNotes() bool {
	return o != nil && o.Notes != nil
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *JourneyMapDefinition) SetNotes(v string) {
	o.Notes = &v
}

// GetRequests returns the Requests field value.
func (o *JourneyMapDefinition) GetRequests() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}
	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value
// and a boolean to check if the value has been set.
func (o *JourneyMapDefinition) GetRequestsOk() (*[]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requests, true
}

// SetRequests sets field value.
func (o *JourneyMapDefinition) SetRequests(v []interface{}) {
	o.Requests = v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *JourneyMapDefinition) GetTime() WidgetTime {
	if o == nil || o.Time == nil {
		var ret WidgetTime
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JourneyMapDefinition) GetTimeOk() (*WidgetTime, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *JourneyMapDefinition) HasTime() bool {
	return o != nil && o.Time != nil
}

// SetTime gets a reference to the given WidgetTime and assigns it to the Time field.
func (o *JourneyMapDefinition) SetTime(v WidgetTime) {
	o.Time = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *JourneyMapDefinition) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JourneyMapDefinition) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *JourneyMapDefinition) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *JourneyMapDefinition) SetTitle(v string) {
	o.Title = &v
}

// GetTitleAlign returns the TitleAlign field value if set, zero value otherwise.
func (o *JourneyMapDefinition) GetTitleAlign() WidgetTextAlign {
	if o == nil || o.TitleAlign == nil {
		var ret WidgetTextAlign
		return ret
	}
	return *o.TitleAlign
}

// GetTitleAlignOk returns a tuple with the TitleAlign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JourneyMapDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool) {
	if o == nil || o.TitleAlign == nil {
		return nil, false
	}
	return o.TitleAlign, true
}

// HasTitleAlign returns a boolean if a field has been set.
func (o *JourneyMapDefinition) HasTitleAlign() bool {
	return o != nil && o.TitleAlign != nil
}

// SetTitleAlign gets a reference to the given WidgetTextAlign and assigns it to the TitleAlign field.
func (o *JourneyMapDefinition) SetTitleAlign(v WidgetTextAlign) {
	o.TitleAlign = &v
}

// GetTitleSize returns the TitleSize field value if set, zero value otherwise.
func (o *JourneyMapDefinition) GetTitleSize() string {
	if o == nil || o.TitleSize == nil {
		var ret string
		return ret
	}
	return *o.TitleSize
}

// GetTitleSizeOk returns a tuple with the TitleSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JourneyMapDefinition) GetTitleSizeOk() (*string, bool) {
	if o == nil || o.TitleSize == nil {
		return nil, false
	}
	return o.TitleSize, true
}

// HasTitleSize returns a boolean if a field has been set.
func (o *JourneyMapDefinition) HasTitleSize() bool {
	return o != nil && o.TitleSize != nil
}

// SetTitleSize gets a reference to the given string and assigns it to the TitleSize field.
func (o *JourneyMapDefinition) SetTitleSize(v string) {
	o.TitleSize = &v
}

// GetType returns the Type field value.
func (o *JourneyMapDefinition) GetType() JourneyMapDefinitionType {
	if o == nil {
		var ret JourneyMapDefinitionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *JourneyMapDefinition) GetTypeOk() (*JourneyMapDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *JourneyMapDefinition) SetType(v JourneyMapDefinitionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o JourneyMapDefinition) MarshalJSON() ([]byte, error) {
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
func (o *JourneyMapDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string                   `json:"description,omitempty"`
		Notes       *string                   `json:"notes,omitempty"`
		Requests    *[]interface{}            `json:"requests"`
		Time        *WidgetTime               `json:"time,omitempty"`
		Title       *string                   `json:"title,omitempty"`
		TitleAlign  *WidgetTextAlign          `json:"title_align,omitempty"`
		TitleSize   *string                   `json:"title_size,omitempty"`
		Type        *JourneyMapDefinitionType `json:"type"`
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
