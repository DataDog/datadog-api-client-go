// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SloBackendDefinition Test-only generated schema for the dashboard versioning proof of concept.
type SloBackendDefinition struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	AdditionalQueryFilters *string `json:"additional_query_filters,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Description *string `json:"description,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	GlobalTimeTarget *string `json:"global_time_target,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Notes *string `json:"notes,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ShowErrorBudget bool `json:"show_error_budget"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	SloId datadog.NullableString `json:"slo_id"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Time *WidgetTime `json:"time,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TimeWindows []SloWidgetTimeWindow `json:"time_windows"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Title *string `json:"title,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TitleAlign *WidgetTextAlign `json:"title_align,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TitleSize *string `json:"title_size,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type SloBackendDefinitionType `json:"type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ViewMode SloWidgetViewMode `json:"view_mode"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ViewType SloBackendDefinitionViewType `json:"view_type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewSloBackendDefinition instantiates a new SloBackendDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSloBackendDefinition(showErrorBudget bool, sloId datadog.NullableString, timeWindows []SloWidgetTimeWindow, typeVar SloBackendDefinitionType, viewMode SloWidgetViewMode, viewType SloBackendDefinitionViewType) *SloBackendDefinition {
	this := SloBackendDefinition{}
	this.ShowErrorBudget = showErrorBudget
	this.SloId = sloId
	this.TimeWindows = timeWindows
	this.Type = typeVar
	this.ViewMode = viewMode
	this.ViewType = viewType
	return &this
}

// NewSloBackendDefinitionWithDefaults instantiates a new SloBackendDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSloBackendDefinitionWithDefaults() *SloBackendDefinition {
	this := SloBackendDefinition{}
	return &this
}

// GetAdditionalQueryFilters returns the AdditionalQueryFilters field value if set, zero value otherwise.
func (o *SloBackendDefinition) GetAdditionalQueryFilters() string {
	if o == nil || o.AdditionalQueryFilters == nil {
		var ret string
		return ret
	}
	return *o.AdditionalQueryFilters
}

// GetAdditionalQueryFiltersOk returns a tuple with the AdditionalQueryFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SloBackendDefinition) GetAdditionalQueryFiltersOk() (*string, bool) {
	if o == nil || o.AdditionalQueryFilters == nil {
		return nil, false
	}
	return o.AdditionalQueryFilters, true
}

// HasAdditionalQueryFilters returns a boolean if a field has been set.
func (o *SloBackendDefinition) HasAdditionalQueryFilters() bool {
	return o != nil && o.AdditionalQueryFilters != nil
}

// SetAdditionalQueryFilters gets a reference to the given string and assigns it to the AdditionalQueryFilters field.
func (o *SloBackendDefinition) SetAdditionalQueryFilters(v string) {
	o.AdditionalQueryFilters = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SloBackendDefinition) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SloBackendDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SloBackendDefinition) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SloBackendDefinition) SetDescription(v string) {
	o.Description = &v
}

// GetGlobalTimeTarget returns the GlobalTimeTarget field value if set, zero value otherwise.
func (o *SloBackendDefinition) GetGlobalTimeTarget() string {
	if o == nil || o.GlobalTimeTarget == nil {
		var ret string
		return ret
	}
	return *o.GlobalTimeTarget
}

// GetGlobalTimeTargetOk returns a tuple with the GlobalTimeTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SloBackendDefinition) GetGlobalTimeTargetOk() (*string, bool) {
	if o == nil || o.GlobalTimeTarget == nil {
		return nil, false
	}
	return o.GlobalTimeTarget, true
}

// HasGlobalTimeTarget returns a boolean if a field has been set.
func (o *SloBackendDefinition) HasGlobalTimeTarget() bool {
	return o != nil && o.GlobalTimeTarget != nil
}

// SetGlobalTimeTarget gets a reference to the given string and assigns it to the GlobalTimeTarget field.
func (o *SloBackendDefinition) SetGlobalTimeTarget(v string) {
	o.GlobalTimeTarget = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *SloBackendDefinition) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SloBackendDefinition) GetNotesOk() (*string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *SloBackendDefinition) HasNotes() bool {
	return o != nil && o.Notes != nil
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *SloBackendDefinition) SetNotes(v string) {
	o.Notes = &v
}

// GetShowErrorBudget returns the ShowErrorBudget field value.
func (o *SloBackendDefinition) GetShowErrorBudget() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.ShowErrorBudget
}

// GetShowErrorBudgetOk returns a tuple with the ShowErrorBudget field value
// and a boolean to check if the value has been set.
func (o *SloBackendDefinition) GetShowErrorBudgetOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShowErrorBudget, true
}

// SetShowErrorBudget sets field value.
func (o *SloBackendDefinition) SetShowErrorBudget(v bool) {
	o.ShowErrorBudget = v
}

// GetSloId returns the SloId field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *SloBackendDefinition) GetSloId() string {
	if o == nil || o.SloId.Get() == nil {
		var ret string
		return ret
	}
	return *o.SloId.Get()
}

// GetSloIdOk returns a tuple with the SloId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SloBackendDefinition) GetSloIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SloId.Get(), o.SloId.IsSet()
}

// SetSloId sets field value.
func (o *SloBackendDefinition) SetSloId(v string) {
	o.SloId.Set(&v)
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *SloBackendDefinition) GetTime() WidgetTime {
	if o == nil || o.Time == nil {
		var ret WidgetTime
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SloBackendDefinition) GetTimeOk() (*WidgetTime, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *SloBackendDefinition) HasTime() bool {
	return o != nil && o.Time != nil
}

// SetTime gets a reference to the given WidgetTime and assigns it to the Time field.
func (o *SloBackendDefinition) SetTime(v WidgetTime) {
	o.Time = &v
}

// GetTimeWindows returns the TimeWindows field value.
func (o *SloBackendDefinition) GetTimeWindows() []SloWidgetTimeWindow {
	if o == nil {
		var ret []SloWidgetTimeWindow
		return ret
	}
	return o.TimeWindows
}

// GetTimeWindowsOk returns a tuple with the TimeWindows field value
// and a boolean to check if the value has been set.
func (o *SloBackendDefinition) GetTimeWindowsOk() (*[]SloWidgetTimeWindow, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeWindows, true
}

// SetTimeWindows sets field value.
func (o *SloBackendDefinition) SetTimeWindows(v []SloWidgetTimeWindow) {
	o.TimeWindows = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *SloBackendDefinition) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SloBackendDefinition) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *SloBackendDefinition) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *SloBackendDefinition) SetTitle(v string) {
	o.Title = &v
}

// GetTitleAlign returns the TitleAlign field value if set, zero value otherwise.
func (o *SloBackendDefinition) GetTitleAlign() WidgetTextAlign {
	if o == nil || o.TitleAlign == nil {
		var ret WidgetTextAlign
		return ret
	}
	return *o.TitleAlign
}

// GetTitleAlignOk returns a tuple with the TitleAlign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SloBackendDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool) {
	if o == nil || o.TitleAlign == nil {
		return nil, false
	}
	return o.TitleAlign, true
}

// HasTitleAlign returns a boolean if a field has been set.
func (o *SloBackendDefinition) HasTitleAlign() bool {
	return o != nil && o.TitleAlign != nil
}

// SetTitleAlign gets a reference to the given WidgetTextAlign and assigns it to the TitleAlign field.
func (o *SloBackendDefinition) SetTitleAlign(v WidgetTextAlign) {
	o.TitleAlign = &v
}

// GetTitleSize returns the TitleSize field value if set, zero value otherwise.
func (o *SloBackendDefinition) GetTitleSize() string {
	if o == nil || o.TitleSize == nil {
		var ret string
		return ret
	}
	return *o.TitleSize
}

// GetTitleSizeOk returns a tuple with the TitleSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SloBackendDefinition) GetTitleSizeOk() (*string, bool) {
	if o == nil || o.TitleSize == nil {
		return nil, false
	}
	return o.TitleSize, true
}

// HasTitleSize returns a boolean if a field has been set.
func (o *SloBackendDefinition) HasTitleSize() bool {
	return o != nil && o.TitleSize != nil
}

// SetTitleSize gets a reference to the given string and assigns it to the TitleSize field.
func (o *SloBackendDefinition) SetTitleSize(v string) {
	o.TitleSize = &v
}

// GetType returns the Type field value.
func (o *SloBackendDefinition) GetType() SloBackendDefinitionType {
	if o == nil {
		var ret SloBackendDefinitionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SloBackendDefinition) GetTypeOk() (*SloBackendDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SloBackendDefinition) SetType(v SloBackendDefinitionType) {
	o.Type = v
}

// GetViewMode returns the ViewMode field value.
func (o *SloBackendDefinition) GetViewMode() SloWidgetViewMode {
	if o == nil {
		var ret SloWidgetViewMode
		return ret
	}
	return o.ViewMode
}

// GetViewModeOk returns a tuple with the ViewMode field value
// and a boolean to check if the value has been set.
func (o *SloBackendDefinition) GetViewModeOk() (*SloWidgetViewMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ViewMode, true
}

// SetViewMode sets field value.
func (o *SloBackendDefinition) SetViewMode(v SloWidgetViewMode) {
	o.ViewMode = v
}

// GetViewType returns the ViewType field value.
func (o *SloBackendDefinition) GetViewType() SloBackendDefinitionViewType {
	if o == nil {
		var ret SloBackendDefinitionViewType
		return ret
	}
	return o.ViewType
}

// GetViewTypeOk returns a tuple with the ViewType field value
// and a boolean to check if the value has been set.
func (o *SloBackendDefinition) GetViewTypeOk() (*SloBackendDefinitionViewType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ViewType, true
}

// SetViewType sets field value.
func (o *SloBackendDefinition) SetViewType(v SloBackendDefinitionViewType) {
	o.ViewType = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SloBackendDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AdditionalQueryFilters != nil {
		toSerialize["additional_query_filters"] = o.AdditionalQueryFilters
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.GlobalTimeTarget != nil {
		toSerialize["global_time_target"] = o.GlobalTimeTarget
	}
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
	}
	toSerialize["show_error_budget"] = o.ShowErrorBudget
	toSerialize["slo_id"] = o.SloId.Get()
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	toSerialize["time_windows"] = o.TimeWindows
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
	toSerialize["view_mode"] = o.ViewMode
	toSerialize["view_type"] = o.ViewType
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SloBackendDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AdditionalQueryFilters *string                       `json:"additional_query_filters,omitempty"`
		Description            *string                       `json:"description,omitempty"`
		GlobalTimeTarget       *string                       `json:"global_time_target,omitempty"`
		Notes                  *string                       `json:"notes,omitempty"`
		ShowErrorBudget        *bool                         `json:"show_error_budget"`
		SloId                  datadog.NullableString        `json:"slo_id"`
		Time                   *WidgetTime                   `json:"time,omitempty"`
		TimeWindows            *[]SloWidgetTimeWindow        `json:"time_windows"`
		Title                  *string                       `json:"title,omitempty"`
		TitleAlign             *WidgetTextAlign              `json:"title_align,omitempty"`
		TitleSize              *string                       `json:"title_size,omitempty"`
		Type                   *SloBackendDefinitionType     `json:"type"`
		ViewMode               *SloWidgetViewMode            `json:"view_mode"`
		ViewType               *SloBackendDefinitionViewType `json:"view_type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ShowErrorBudget == nil {
		return fmt.Errorf("required field show_error_budget missing")
	}
	if !all.SloId.IsSet() {
		return fmt.Errorf("required field slo_id missing")
	}
	if all.TimeWindows == nil {
		return fmt.Errorf("required field time_windows missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.ViewMode == nil {
		return fmt.Errorf("required field view_mode missing")
	}
	if all.ViewType == nil {
		return fmt.Errorf("required field view_type missing")
	}

	hasInvalidField := false
	o.AdditionalQueryFilters = all.AdditionalQueryFilters
	o.Description = all.Description
	o.GlobalTimeTarget = all.GlobalTimeTarget
	o.Notes = all.Notes
	o.ShowErrorBudget = *all.ShowErrorBudget
	o.SloId = all.SloId
	o.Time = all.Time
	o.TimeWindows = *all.TimeWindows
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
	if !all.ViewMode.IsValid() {
		hasInvalidField = true
	} else {
		o.ViewMode = *all.ViewMode
	}
	if !all.ViewType.IsValid() {
		hasInvalidField = true
	} else {
		o.ViewType = *all.ViewType
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
