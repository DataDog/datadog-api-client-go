// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// QueryValueBackendDefinition Test-only generated schema for the dashboard versioning proof of concept.
type QueryValueBackendDefinition struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Autoscale *bool `json:"autoscale,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	CustomLinks []WidgetCustomLink `json:"custom_links,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	CustomUnit *string `json:"custom_unit,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Description *string `json:"description,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Notes *string `json:"notes,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Precision *float64 `json:"precision,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Requests []QueryValueBackendRequest `json:"requests"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TextAlign *WidgetTextAlign `json:"text_align,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Time *WidgetTime `json:"time,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TimeseriesBackground *TimeseriesBackground `json:"timeseries_background,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Title *string `json:"title,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TitleAlign *WidgetTextAlign `json:"title_align,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TitleSize *string `json:"title_size,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type QueryValueBackendDefinitionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewQueryValueBackendDefinition instantiates a new QueryValueBackendDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewQueryValueBackendDefinition(requests []QueryValueBackendRequest, typeVar QueryValueBackendDefinitionType) *QueryValueBackendDefinition {
	this := QueryValueBackendDefinition{}
	this.Requests = requests
	this.Type = typeVar
	return &this
}

// NewQueryValueBackendDefinitionWithDefaults instantiates a new QueryValueBackendDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewQueryValueBackendDefinitionWithDefaults() *QueryValueBackendDefinition {
	this := QueryValueBackendDefinition{}
	return &this
}

// GetAutoscale returns the Autoscale field value if set, zero value otherwise.
func (o *QueryValueBackendDefinition) GetAutoscale() bool {
	if o == nil || o.Autoscale == nil {
		var ret bool
		return ret
	}
	return *o.Autoscale
}

// GetAutoscaleOk returns a tuple with the Autoscale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryValueBackendDefinition) GetAutoscaleOk() (*bool, bool) {
	if o == nil || o.Autoscale == nil {
		return nil, false
	}
	return o.Autoscale, true
}

// HasAutoscale returns a boolean if a field has been set.
func (o *QueryValueBackendDefinition) HasAutoscale() bool {
	return o != nil && o.Autoscale != nil
}

// SetAutoscale gets a reference to the given bool and assigns it to the Autoscale field.
func (o *QueryValueBackendDefinition) SetAutoscale(v bool) {
	o.Autoscale = &v
}

// GetCustomLinks returns the CustomLinks field value if set, zero value otherwise.
func (o *QueryValueBackendDefinition) GetCustomLinks() []WidgetCustomLink {
	if o == nil || o.CustomLinks == nil {
		var ret []WidgetCustomLink
		return ret
	}
	return o.CustomLinks
}

// GetCustomLinksOk returns a tuple with the CustomLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryValueBackendDefinition) GetCustomLinksOk() (*[]WidgetCustomLink, bool) {
	if o == nil || o.CustomLinks == nil {
		return nil, false
	}
	return &o.CustomLinks, true
}

// HasCustomLinks returns a boolean if a field has been set.
func (o *QueryValueBackendDefinition) HasCustomLinks() bool {
	return o != nil && o.CustomLinks != nil
}

// SetCustomLinks gets a reference to the given []WidgetCustomLink and assigns it to the CustomLinks field.
func (o *QueryValueBackendDefinition) SetCustomLinks(v []WidgetCustomLink) {
	o.CustomLinks = v
}

// GetCustomUnit returns the CustomUnit field value if set, zero value otherwise.
func (o *QueryValueBackendDefinition) GetCustomUnit() string {
	if o == nil || o.CustomUnit == nil {
		var ret string
		return ret
	}
	return *o.CustomUnit
}

// GetCustomUnitOk returns a tuple with the CustomUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryValueBackendDefinition) GetCustomUnitOk() (*string, bool) {
	if o == nil || o.CustomUnit == nil {
		return nil, false
	}
	return o.CustomUnit, true
}

// HasCustomUnit returns a boolean if a field has been set.
func (o *QueryValueBackendDefinition) HasCustomUnit() bool {
	return o != nil && o.CustomUnit != nil
}

// SetCustomUnit gets a reference to the given string and assigns it to the CustomUnit field.
func (o *QueryValueBackendDefinition) SetCustomUnit(v string) {
	o.CustomUnit = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *QueryValueBackendDefinition) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryValueBackendDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *QueryValueBackendDefinition) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *QueryValueBackendDefinition) SetDescription(v string) {
	o.Description = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *QueryValueBackendDefinition) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryValueBackendDefinition) GetNotesOk() (*string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *QueryValueBackendDefinition) HasNotes() bool {
	return o != nil && o.Notes != nil
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *QueryValueBackendDefinition) SetNotes(v string) {
	o.Notes = &v
}

// GetPrecision returns the Precision field value if set, zero value otherwise.
func (o *QueryValueBackendDefinition) GetPrecision() float64 {
	if o == nil || o.Precision == nil {
		var ret float64
		return ret
	}
	return *o.Precision
}

// GetPrecisionOk returns a tuple with the Precision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryValueBackendDefinition) GetPrecisionOk() (*float64, bool) {
	if o == nil || o.Precision == nil {
		return nil, false
	}
	return o.Precision, true
}

// HasPrecision returns a boolean if a field has been set.
func (o *QueryValueBackendDefinition) HasPrecision() bool {
	return o != nil && o.Precision != nil
}

// SetPrecision gets a reference to the given float64 and assigns it to the Precision field.
func (o *QueryValueBackendDefinition) SetPrecision(v float64) {
	o.Precision = &v
}

// GetRequests returns the Requests field value.
func (o *QueryValueBackendDefinition) GetRequests() []QueryValueBackendRequest {
	if o == nil {
		var ret []QueryValueBackendRequest
		return ret
	}
	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value
// and a boolean to check if the value has been set.
func (o *QueryValueBackendDefinition) GetRequestsOk() (*[]QueryValueBackendRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requests, true
}

// SetRequests sets field value.
func (o *QueryValueBackendDefinition) SetRequests(v []QueryValueBackendRequest) {
	o.Requests = v
}

// GetTextAlign returns the TextAlign field value if set, zero value otherwise.
func (o *QueryValueBackendDefinition) GetTextAlign() WidgetTextAlign {
	if o == nil || o.TextAlign == nil {
		var ret WidgetTextAlign
		return ret
	}
	return *o.TextAlign
}

// GetTextAlignOk returns a tuple with the TextAlign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryValueBackendDefinition) GetTextAlignOk() (*WidgetTextAlign, bool) {
	if o == nil || o.TextAlign == nil {
		return nil, false
	}
	return o.TextAlign, true
}

// HasTextAlign returns a boolean if a field has been set.
func (o *QueryValueBackendDefinition) HasTextAlign() bool {
	return o != nil && o.TextAlign != nil
}

// SetTextAlign gets a reference to the given WidgetTextAlign and assigns it to the TextAlign field.
func (o *QueryValueBackendDefinition) SetTextAlign(v WidgetTextAlign) {
	o.TextAlign = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *QueryValueBackendDefinition) GetTime() WidgetTime {
	if o == nil || o.Time == nil {
		var ret WidgetTime
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryValueBackendDefinition) GetTimeOk() (*WidgetTime, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *QueryValueBackendDefinition) HasTime() bool {
	return o != nil && o.Time != nil
}

// SetTime gets a reference to the given WidgetTime and assigns it to the Time field.
func (o *QueryValueBackendDefinition) SetTime(v WidgetTime) {
	o.Time = &v
}

// GetTimeseriesBackground returns the TimeseriesBackground field value if set, zero value otherwise.
func (o *QueryValueBackendDefinition) GetTimeseriesBackground() TimeseriesBackground {
	if o == nil || o.TimeseriesBackground == nil {
		var ret TimeseriesBackground
		return ret
	}
	return *o.TimeseriesBackground
}

// GetTimeseriesBackgroundOk returns a tuple with the TimeseriesBackground field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryValueBackendDefinition) GetTimeseriesBackgroundOk() (*TimeseriesBackground, bool) {
	if o == nil || o.TimeseriesBackground == nil {
		return nil, false
	}
	return o.TimeseriesBackground, true
}

// HasTimeseriesBackground returns a boolean if a field has been set.
func (o *QueryValueBackendDefinition) HasTimeseriesBackground() bool {
	return o != nil && o.TimeseriesBackground != nil
}

// SetTimeseriesBackground gets a reference to the given TimeseriesBackground and assigns it to the TimeseriesBackground field.
func (o *QueryValueBackendDefinition) SetTimeseriesBackground(v TimeseriesBackground) {
	o.TimeseriesBackground = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *QueryValueBackendDefinition) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryValueBackendDefinition) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *QueryValueBackendDefinition) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *QueryValueBackendDefinition) SetTitle(v string) {
	o.Title = &v
}

// GetTitleAlign returns the TitleAlign field value if set, zero value otherwise.
func (o *QueryValueBackendDefinition) GetTitleAlign() WidgetTextAlign {
	if o == nil || o.TitleAlign == nil {
		var ret WidgetTextAlign
		return ret
	}
	return *o.TitleAlign
}

// GetTitleAlignOk returns a tuple with the TitleAlign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryValueBackendDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool) {
	if o == nil || o.TitleAlign == nil {
		return nil, false
	}
	return o.TitleAlign, true
}

// HasTitleAlign returns a boolean if a field has been set.
func (o *QueryValueBackendDefinition) HasTitleAlign() bool {
	return o != nil && o.TitleAlign != nil
}

// SetTitleAlign gets a reference to the given WidgetTextAlign and assigns it to the TitleAlign field.
func (o *QueryValueBackendDefinition) SetTitleAlign(v WidgetTextAlign) {
	o.TitleAlign = &v
}

// GetTitleSize returns the TitleSize field value if set, zero value otherwise.
func (o *QueryValueBackendDefinition) GetTitleSize() string {
	if o == nil || o.TitleSize == nil {
		var ret string
		return ret
	}
	return *o.TitleSize
}

// GetTitleSizeOk returns a tuple with the TitleSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryValueBackendDefinition) GetTitleSizeOk() (*string, bool) {
	if o == nil || o.TitleSize == nil {
		return nil, false
	}
	return o.TitleSize, true
}

// HasTitleSize returns a boolean if a field has been set.
func (o *QueryValueBackendDefinition) HasTitleSize() bool {
	return o != nil && o.TitleSize != nil
}

// SetTitleSize gets a reference to the given string and assigns it to the TitleSize field.
func (o *QueryValueBackendDefinition) SetTitleSize(v string) {
	o.TitleSize = &v
}

// GetType returns the Type field value.
func (o *QueryValueBackendDefinition) GetType() QueryValueBackendDefinitionType {
	if o == nil {
		var ret QueryValueBackendDefinitionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *QueryValueBackendDefinition) GetTypeOk() (*QueryValueBackendDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *QueryValueBackendDefinition) SetType(v QueryValueBackendDefinitionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o QueryValueBackendDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Autoscale != nil {
		toSerialize["autoscale"] = o.Autoscale
	}
	if o.CustomLinks != nil {
		toSerialize["custom_links"] = o.CustomLinks
	}
	if o.CustomUnit != nil {
		toSerialize["custom_unit"] = o.CustomUnit
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
	}
	if o.Precision != nil {
		toSerialize["precision"] = o.Precision
	}
	toSerialize["requests"] = o.Requests
	if o.TextAlign != nil {
		toSerialize["text_align"] = o.TextAlign
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	if o.TimeseriesBackground != nil {
		toSerialize["timeseries_background"] = o.TimeseriesBackground
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
func (o *QueryValueBackendDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Autoscale            *bool                            `json:"autoscale,omitempty"`
		CustomLinks          []WidgetCustomLink               `json:"custom_links,omitempty"`
		CustomUnit           *string                          `json:"custom_unit,omitempty"`
		Description          *string                          `json:"description,omitempty"`
		Notes                *string                          `json:"notes,omitempty"`
		Precision            *float64                         `json:"precision,omitempty"`
		Requests             *[]QueryValueBackendRequest      `json:"requests"`
		TextAlign            *WidgetTextAlign                 `json:"text_align,omitempty"`
		Time                 *WidgetTime                      `json:"time,omitempty"`
		TimeseriesBackground *TimeseriesBackground            `json:"timeseries_background,omitempty"`
		Title                *string                          `json:"title,omitempty"`
		TitleAlign           *WidgetTextAlign                 `json:"title_align,omitempty"`
		TitleSize            *string                          `json:"title_size,omitempty"`
		Type                 *QueryValueBackendDefinitionType `json:"type"`
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
	o.Autoscale = all.Autoscale
	o.CustomLinks = all.CustomLinks
	o.CustomUnit = all.CustomUnit
	o.Description = all.Description
	o.Notes = all.Notes
	o.Precision = all.Precision
	o.Requests = *all.Requests
	if all.TextAlign != nil && !all.TextAlign.IsValid() {
		hasInvalidField = true
	} else {
		o.TextAlign = all.TextAlign
	}
	o.Time = all.Time
	if all.TimeseriesBackground != nil && all.TimeseriesBackground.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TimeseriesBackground = all.TimeseriesBackground
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
