// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AlertValueDefinition Test-only generated schema for the dashboard versioning proof of concept.
type AlertValueDefinition struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	AlertId string `json:"alert_id"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Description *string `json:"description,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Notes *string `json:"notes,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Precision *float64 `json:"precision,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TextAlign *WidgetTextAlign `json:"text_align,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TextSize *string `json:"text_size,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Time *WidgetTime `json:"time,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Title *string `json:"title,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TitleAlign *WidgetTextAlign `json:"title_align,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TitleSize *string `json:"title_size,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type AlertValueDefinitionType `json:"type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Unit *string `json:"unit,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewAlertValueDefinition instantiates a new AlertValueDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAlertValueDefinition(alertId string, typeVar AlertValueDefinitionType) *AlertValueDefinition {
	this := AlertValueDefinition{}
	this.AlertId = alertId
	this.Type = typeVar
	return &this
}

// NewAlertValueDefinitionWithDefaults instantiates a new AlertValueDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAlertValueDefinitionWithDefaults() *AlertValueDefinition {
	this := AlertValueDefinition{}
	return &this
}

// GetAlertId returns the AlertId field value.
func (o *AlertValueDefinition) GetAlertId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AlertId
}

// GetAlertIdOk returns a tuple with the AlertId field value
// and a boolean to check if the value has been set.
func (o *AlertValueDefinition) GetAlertIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlertId, true
}

// SetAlertId sets field value.
func (o *AlertValueDefinition) SetAlertId(v string) {
	o.AlertId = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AlertValueDefinition) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertValueDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AlertValueDefinition) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AlertValueDefinition) SetDescription(v string) {
	o.Description = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *AlertValueDefinition) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertValueDefinition) GetNotesOk() (*string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *AlertValueDefinition) HasNotes() bool {
	return o != nil && o.Notes != nil
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *AlertValueDefinition) SetNotes(v string) {
	o.Notes = &v
}

// GetPrecision returns the Precision field value if set, zero value otherwise.
func (o *AlertValueDefinition) GetPrecision() float64 {
	if o == nil || o.Precision == nil {
		var ret float64
		return ret
	}
	return *o.Precision
}

// GetPrecisionOk returns a tuple with the Precision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertValueDefinition) GetPrecisionOk() (*float64, bool) {
	if o == nil || o.Precision == nil {
		return nil, false
	}
	return o.Precision, true
}

// HasPrecision returns a boolean if a field has been set.
func (o *AlertValueDefinition) HasPrecision() bool {
	return o != nil && o.Precision != nil
}

// SetPrecision gets a reference to the given float64 and assigns it to the Precision field.
func (o *AlertValueDefinition) SetPrecision(v float64) {
	o.Precision = &v
}

// GetTextAlign returns the TextAlign field value if set, zero value otherwise.
func (o *AlertValueDefinition) GetTextAlign() WidgetTextAlign {
	if o == nil || o.TextAlign == nil {
		var ret WidgetTextAlign
		return ret
	}
	return *o.TextAlign
}

// GetTextAlignOk returns a tuple with the TextAlign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertValueDefinition) GetTextAlignOk() (*WidgetTextAlign, bool) {
	if o == nil || o.TextAlign == nil {
		return nil, false
	}
	return o.TextAlign, true
}

// HasTextAlign returns a boolean if a field has been set.
func (o *AlertValueDefinition) HasTextAlign() bool {
	return o != nil && o.TextAlign != nil
}

// SetTextAlign gets a reference to the given WidgetTextAlign and assigns it to the TextAlign field.
func (o *AlertValueDefinition) SetTextAlign(v WidgetTextAlign) {
	o.TextAlign = &v
}

// GetTextSize returns the TextSize field value if set, zero value otherwise.
func (o *AlertValueDefinition) GetTextSize() string {
	if o == nil || o.TextSize == nil {
		var ret string
		return ret
	}
	return *o.TextSize
}

// GetTextSizeOk returns a tuple with the TextSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertValueDefinition) GetTextSizeOk() (*string, bool) {
	if o == nil || o.TextSize == nil {
		return nil, false
	}
	return o.TextSize, true
}

// HasTextSize returns a boolean if a field has been set.
func (o *AlertValueDefinition) HasTextSize() bool {
	return o != nil && o.TextSize != nil
}

// SetTextSize gets a reference to the given string and assigns it to the TextSize field.
func (o *AlertValueDefinition) SetTextSize(v string) {
	o.TextSize = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *AlertValueDefinition) GetTime() WidgetTime {
	if o == nil || o.Time == nil {
		var ret WidgetTime
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertValueDefinition) GetTimeOk() (*WidgetTime, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *AlertValueDefinition) HasTime() bool {
	return o != nil && o.Time != nil
}

// SetTime gets a reference to the given WidgetTime and assigns it to the Time field.
func (o *AlertValueDefinition) SetTime(v WidgetTime) {
	o.Time = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AlertValueDefinition) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertValueDefinition) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AlertValueDefinition) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AlertValueDefinition) SetTitle(v string) {
	o.Title = &v
}

// GetTitleAlign returns the TitleAlign field value if set, zero value otherwise.
func (o *AlertValueDefinition) GetTitleAlign() WidgetTextAlign {
	if o == nil || o.TitleAlign == nil {
		var ret WidgetTextAlign
		return ret
	}
	return *o.TitleAlign
}

// GetTitleAlignOk returns a tuple with the TitleAlign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertValueDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool) {
	if o == nil || o.TitleAlign == nil {
		return nil, false
	}
	return o.TitleAlign, true
}

// HasTitleAlign returns a boolean if a field has been set.
func (o *AlertValueDefinition) HasTitleAlign() bool {
	return o != nil && o.TitleAlign != nil
}

// SetTitleAlign gets a reference to the given WidgetTextAlign and assigns it to the TitleAlign field.
func (o *AlertValueDefinition) SetTitleAlign(v WidgetTextAlign) {
	o.TitleAlign = &v
}

// GetTitleSize returns the TitleSize field value if set, zero value otherwise.
func (o *AlertValueDefinition) GetTitleSize() string {
	if o == nil || o.TitleSize == nil {
		var ret string
		return ret
	}
	return *o.TitleSize
}

// GetTitleSizeOk returns a tuple with the TitleSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertValueDefinition) GetTitleSizeOk() (*string, bool) {
	if o == nil || o.TitleSize == nil {
		return nil, false
	}
	return o.TitleSize, true
}

// HasTitleSize returns a boolean if a field has been set.
func (o *AlertValueDefinition) HasTitleSize() bool {
	return o != nil && o.TitleSize != nil
}

// SetTitleSize gets a reference to the given string and assigns it to the TitleSize field.
func (o *AlertValueDefinition) SetTitleSize(v string) {
	o.TitleSize = &v
}

// GetType returns the Type field value.
func (o *AlertValueDefinition) GetType() AlertValueDefinitionType {
	if o == nil {
		var ret AlertValueDefinitionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AlertValueDefinition) GetTypeOk() (*AlertValueDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AlertValueDefinition) SetType(v AlertValueDefinitionType) {
	o.Type = v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *AlertValueDefinition) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertValueDefinition) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *AlertValueDefinition) HasUnit() bool {
	return o != nil && o.Unit != nil
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *AlertValueDefinition) SetUnit(v string) {
	o.Unit = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AlertValueDefinition) MarshalJSON() ([]byte, error) {
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
	if o.Precision != nil {
		toSerialize["precision"] = o.Precision
	}
	if o.TextAlign != nil {
		toSerialize["text_align"] = o.TextAlign
	}
	if o.TextSize != nil {
		toSerialize["text_size"] = o.TextSize
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
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AlertValueDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AlertId     *string                   `json:"alert_id"`
		Description *string                   `json:"description,omitempty"`
		Notes       *string                   `json:"notes,omitempty"`
		Precision   *float64                  `json:"precision,omitempty"`
		TextAlign   *WidgetTextAlign          `json:"text_align,omitempty"`
		TextSize    *string                   `json:"text_size,omitempty"`
		Time        *WidgetTime               `json:"time,omitempty"`
		Title       *string                   `json:"title,omitempty"`
		TitleAlign  *WidgetTextAlign          `json:"title_align,omitempty"`
		TitleSize   *string                   `json:"title_size,omitempty"`
		Type        *AlertValueDefinitionType `json:"type"`
		Unit        *string                   `json:"unit,omitempty"`
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

	hasInvalidField := false
	o.AlertId = *all.AlertId
	o.Description = all.Description
	o.Notes = all.Notes
	o.Precision = all.Precision
	if all.TextAlign != nil && !all.TextAlign.IsValid() {
		hasInvalidField = true
	} else {
		o.TextAlign = all.TextAlign
	}
	o.TextSize = all.TextSize
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
	o.Unit = all.Unit

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
