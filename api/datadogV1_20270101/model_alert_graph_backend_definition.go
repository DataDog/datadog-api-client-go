// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AlertGraphBackendDefinition Test-only generated schema for the dashboard versioning proof of concept.
type AlertGraphBackendDefinition struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	AlertId string `json:"alert_id"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Description *string `json:"description,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Notes *string `json:"notes,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Time *WidgetTime `json:"time,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Title *string `json:"title,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TitleAlign *WidgetTextAlign `json:"title_align,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TitleSize *string `json:"title_size,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type AlertGraphBackendDefinitionType `json:"type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	VizType AlertGraphVizType `json:"viz_type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewAlertGraphBackendDefinition instantiates a new AlertGraphBackendDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAlertGraphBackendDefinition(alertId string, typeVar AlertGraphBackendDefinitionType, vizType AlertGraphVizType) *AlertGraphBackendDefinition {
	this := AlertGraphBackendDefinition{}
	this.AlertId = alertId
	this.Type = typeVar
	this.VizType = vizType
	return &this
}

// NewAlertGraphBackendDefinitionWithDefaults instantiates a new AlertGraphBackendDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAlertGraphBackendDefinitionWithDefaults() *AlertGraphBackendDefinition {
	this := AlertGraphBackendDefinition{}
	return &this
}

// GetAlertId returns the AlertId field value.
func (o *AlertGraphBackendDefinition) GetAlertId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AlertId
}

// GetAlertIdOk returns a tuple with the AlertId field value
// and a boolean to check if the value has been set.
func (o *AlertGraphBackendDefinition) GetAlertIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlertId, true
}

// SetAlertId sets field value.
func (o *AlertGraphBackendDefinition) SetAlertId(v string) {
	o.AlertId = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AlertGraphBackendDefinition) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertGraphBackendDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AlertGraphBackendDefinition) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AlertGraphBackendDefinition) SetDescription(v string) {
	o.Description = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *AlertGraphBackendDefinition) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertGraphBackendDefinition) GetNotesOk() (*string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *AlertGraphBackendDefinition) HasNotes() bool {
	return o != nil && o.Notes != nil
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *AlertGraphBackendDefinition) SetNotes(v string) {
	o.Notes = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *AlertGraphBackendDefinition) GetTime() WidgetTime {
	if o == nil || o.Time == nil {
		var ret WidgetTime
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertGraphBackendDefinition) GetTimeOk() (*WidgetTime, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *AlertGraphBackendDefinition) HasTime() bool {
	return o != nil && o.Time != nil
}

// SetTime gets a reference to the given WidgetTime and assigns it to the Time field.
func (o *AlertGraphBackendDefinition) SetTime(v WidgetTime) {
	o.Time = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AlertGraphBackendDefinition) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertGraphBackendDefinition) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AlertGraphBackendDefinition) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AlertGraphBackendDefinition) SetTitle(v string) {
	o.Title = &v
}

// GetTitleAlign returns the TitleAlign field value if set, zero value otherwise.
func (o *AlertGraphBackendDefinition) GetTitleAlign() WidgetTextAlign {
	if o == nil || o.TitleAlign == nil {
		var ret WidgetTextAlign
		return ret
	}
	return *o.TitleAlign
}

// GetTitleAlignOk returns a tuple with the TitleAlign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertGraphBackendDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool) {
	if o == nil || o.TitleAlign == nil {
		return nil, false
	}
	return o.TitleAlign, true
}

// HasTitleAlign returns a boolean if a field has been set.
func (o *AlertGraphBackendDefinition) HasTitleAlign() bool {
	return o != nil && o.TitleAlign != nil
}

// SetTitleAlign gets a reference to the given WidgetTextAlign and assigns it to the TitleAlign field.
func (o *AlertGraphBackendDefinition) SetTitleAlign(v WidgetTextAlign) {
	o.TitleAlign = &v
}

// GetTitleSize returns the TitleSize field value if set, zero value otherwise.
func (o *AlertGraphBackendDefinition) GetTitleSize() string {
	if o == nil || o.TitleSize == nil {
		var ret string
		return ret
	}
	return *o.TitleSize
}

// GetTitleSizeOk returns a tuple with the TitleSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertGraphBackendDefinition) GetTitleSizeOk() (*string, bool) {
	if o == nil || o.TitleSize == nil {
		return nil, false
	}
	return o.TitleSize, true
}

// HasTitleSize returns a boolean if a field has been set.
func (o *AlertGraphBackendDefinition) HasTitleSize() bool {
	return o != nil && o.TitleSize != nil
}

// SetTitleSize gets a reference to the given string and assigns it to the TitleSize field.
func (o *AlertGraphBackendDefinition) SetTitleSize(v string) {
	o.TitleSize = &v
}

// GetType returns the Type field value.
func (o *AlertGraphBackendDefinition) GetType() AlertGraphBackendDefinitionType {
	if o == nil {
		var ret AlertGraphBackendDefinitionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AlertGraphBackendDefinition) GetTypeOk() (*AlertGraphBackendDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AlertGraphBackendDefinition) SetType(v AlertGraphBackendDefinitionType) {
	o.Type = v
}

// GetVizType returns the VizType field value.
func (o *AlertGraphBackendDefinition) GetVizType() AlertGraphVizType {
	if o == nil {
		var ret AlertGraphVizType
		return ret
	}
	return o.VizType
}

// GetVizTypeOk returns a tuple with the VizType field value
// and a boolean to check if the value has been set.
func (o *AlertGraphBackendDefinition) GetVizTypeOk() (*AlertGraphVizType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VizType, true
}

// SetVizType sets field value.
func (o *AlertGraphBackendDefinition) SetVizType(v AlertGraphVizType) {
	o.VizType = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AlertGraphBackendDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["alert_id"] = o.AlertId
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
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
	toSerialize["viz_type"] = o.VizType
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AlertGraphBackendDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AlertId     *string                          `json:"alert_id"`
		Description *string                          `json:"description,omitempty"`
		Notes       *string                          `json:"notes,omitempty"`
		Time        *WidgetTime                      `json:"time,omitempty"`
		Title       *string                          `json:"title,omitempty"`
		TitleAlign  *WidgetTextAlign                 `json:"title_align,omitempty"`
		TitleSize   *string                          `json:"title_size,omitempty"`
		Type        *AlertGraphBackendDefinitionType `json:"type"`
		VizType     *AlertGraphVizType               `json:"viz_type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AlertId == nil {
		return fmt.Errorf("required field alert_id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.VizType == nil {
		return fmt.Errorf("required field viz_type missing")
	}

	hasInvalidField := false
	o.AlertId = *all.AlertId
	o.Description = all.Description
	o.Notes = all.Notes
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
	if !all.VizType.IsValid() {
		hasInvalidField = true
	} else {
		o.VizType = *all.VizType
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
