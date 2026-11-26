// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ManageStatusBackendDefinition Test-only generated schema for the dashboard versioning proof of concept.
type ManageStatusBackendDefinition struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	ColorPreference *ManageStatusBackendDefinitionColorPreference `json:"color_preference,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Count *float64 `json:"count,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Description *string `json:"description,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DisplayFormat *ManageStatusBackendDefinitionDisplayFormat `json:"display_format,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	HideZeroCounts *bool `json:"hide_zero_counts,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	LastTriggeredFormat *ManageStatusBackendDefinitionLastTriggeredFormat `json:"last_triggered_format,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Notes *string `json:"notes,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Query string `json:"query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ShowInvestigation *bool `json:"show_investigation,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ShowLastTriggered *bool `json:"show_last_triggered,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ShowPriority *bool `json:"show_priority,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ShowStatus *bool `json:"show_status,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Sort *string `json:"sort,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Start *float64 `json:"start,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	SummaryType *ManageStatusBackendDefinitionSummaryType `json:"summary_type,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Title *string `json:"title,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TitleAlign *WidgetTextAlign `json:"title_align,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TitleSize *string `json:"title_size,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type ManageStatusBackendDefinitionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewManageStatusBackendDefinition instantiates a new ManageStatusBackendDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewManageStatusBackendDefinition(query string, typeVar ManageStatusBackendDefinitionType) *ManageStatusBackendDefinition {
	this := ManageStatusBackendDefinition{}
	this.Query = query
	this.Type = typeVar
	return &this
}

// NewManageStatusBackendDefinitionWithDefaults instantiates a new ManageStatusBackendDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewManageStatusBackendDefinitionWithDefaults() *ManageStatusBackendDefinition {
	this := ManageStatusBackendDefinition{}
	return &this
}

// GetColorPreference returns the ColorPreference field value if set, zero value otherwise.
func (o *ManageStatusBackendDefinition) GetColorPreference() ManageStatusBackendDefinitionColorPreference {
	if o == nil || o.ColorPreference == nil {
		var ret ManageStatusBackendDefinitionColorPreference
		return ret
	}
	return *o.ColorPreference
}

// GetColorPreferenceOk returns a tuple with the ColorPreference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageStatusBackendDefinition) GetColorPreferenceOk() (*ManageStatusBackendDefinitionColorPreference, bool) {
	if o == nil || o.ColorPreference == nil {
		return nil, false
	}
	return o.ColorPreference, true
}

// HasColorPreference returns a boolean if a field has been set.
func (o *ManageStatusBackendDefinition) HasColorPreference() bool {
	return o != nil && o.ColorPreference != nil
}

// SetColorPreference gets a reference to the given ManageStatusBackendDefinitionColorPreference and assigns it to the ColorPreference field.
func (o *ManageStatusBackendDefinition) SetColorPreference(v ManageStatusBackendDefinitionColorPreference) {
	o.ColorPreference = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ManageStatusBackendDefinition) GetCount() float64 {
	if o == nil || o.Count == nil {
		var ret float64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageStatusBackendDefinition) GetCountOk() (*float64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ManageStatusBackendDefinition) HasCount() bool {
	return o != nil && o.Count != nil
}

// SetCount gets a reference to the given float64 and assigns it to the Count field.
func (o *ManageStatusBackendDefinition) SetCount(v float64) {
	o.Count = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ManageStatusBackendDefinition) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageStatusBackendDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ManageStatusBackendDefinition) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ManageStatusBackendDefinition) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayFormat returns the DisplayFormat field value if set, zero value otherwise.
func (o *ManageStatusBackendDefinition) GetDisplayFormat() ManageStatusBackendDefinitionDisplayFormat {
	if o == nil || o.DisplayFormat == nil {
		var ret ManageStatusBackendDefinitionDisplayFormat
		return ret
	}
	return *o.DisplayFormat
}

// GetDisplayFormatOk returns a tuple with the DisplayFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageStatusBackendDefinition) GetDisplayFormatOk() (*ManageStatusBackendDefinitionDisplayFormat, bool) {
	if o == nil || o.DisplayFormat == nil {
		return nil, false
	}
	return o.DisplayFormat, true
}

// HasDisplayFormat returns a boolean if a field has been set.
func (o *ManageStatusBackendDefinition) HasDisplayFormat() bool {
	return o != nil && o.DisplayFormat != nil
}

// SetDisplayFormat gets a reference to the given ManageStatusBackendDefinitionDisplayFormat and assigns it to the DisplayFormat field.
func (o *ManageStatusBackendDefinition) SetDisplayFormat(v ManageStatusBackendDefinitionDisplayFormat) {
	o.DisplayFormat = &v
}

// GetHideZeroCounts returns the HideZeroCounts field value if set, zero value otherwise.
func (o *ManageStatusBackendDefinition) GetHideZeroCounts() bool {
	if o == nil || o.HideZeroCounts == nil {
		var ret bool
		return ret
	}
	return *o.HideZeroCounts
}

// GetHideZeroCountsOk returns a tuple with the HideZeroCounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageStatusBackendDefinition) GetHideZeroCountsOk() (*bool, bool) {
	if o == nil || o.HideZeroCounts == nil {
		return nil, false
	}
	return o.HideZeroCounts, true
}

// HasHideZeroCounts returns a boolean if a field has been set.
func (o *ManageStatusBackendDefinition) HasHideZeroCounts() bool {
	return o != nil && o.HideZeroCounts != nil
}

// SetHideZeroCounts gets a reference to the given bool and assigns it to the HideZeroCounts field.
func (o *ManageStatusBackendDefinition) SetHideZeroCounts(v bool) {
	o.HideZeroCounts = &v
}

// GetLastTriggeredFormat returns the LastTriggeredFormat field value if set, zero value otherwise.
func (o *ManageStatusBackendDefinition) GetLastTriggeredFormat() ManageStatusBackendDefinitionLastTriggeredFormat {
	if o == nil || o.LastTriggeredFormat == nil {
		var ret ManageStatusBackendDefinitionLastTriggeredFormat
		return ret
	}
	return *o.LastTriggeredFormat
}

// GetLastTriggeredFormatOk returns a tuple with the LastTriggeredFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageStatusBackendDefinition) GetLastTriggeredFormatOk() (*ManageStatusBackendDefinitionLastTriggeredFormat, bool) {
	if o == nil || o.LastTriggeredFormat == nil {
		return nil, false
	}
	return o.LastTriggeredFormat, true
}

// HasLastTriggeredFormat returns a boolean if a field has been set.
func (o *ManageStatusBackendDefinition) HasLastTriggeredFormat() bool {
	return o != nil && o.LastTriggeredFormat != nil
}

// SetLastTriggeredFormat gets a reference to the given ManageStatusBackendDefinitionLastTriggeredFormat and assigns it to the LastTriggeredFormat field.
func (o *ManageStatusBackendDefinition) SetLastTriggeredFormat(v ManageStatusBackendDefinitionLastTriggeredFormat) {
	o.LastTriggeredFormat = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *ManageStatusBackendDefinition) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageStatusBackendDefinition) GetNotesOk() (*string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *ManageStatusBackendDefinition) HasNotes() bool {
	return o != nil && o.Notes != nil
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *ManageStatusBackendDefinition) SetNotes(v string) {
	o.Notes = &v
}

// GetQuery returns the Query field value.
func (o *ManageStatusBackendDefinition) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *ManageStatusBackendDefinition) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *ManageStatusBackendDefinition) SetQuery(v string) {
	o.Query = v
}

// GetShowInvestigation returns the ShowInvestigation field value if set, zero value otherwise.
func (o *ManageStatusBackendDefinition) GetShowInvestigation() bool {
	if o == nil || o.ShowInvestigation == nil {
		var ret bool
		return ret
	}
	return *o.ShowInvestigation
}

// GetShowInvestigationOk returns a tuple with the ShowInvestigation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageStatusBackendDefinition) GetShowInvestigationOk() (*bool, bool) {
	if o == nil || o.ShowInvestigation == nil {
		return nil, false
	}
	return o.ShowInvestigation, true
}

// HasShowInvestigation returns a boolean if a field has been set.
func (o *ManageStatusBackendDefinition) HasShowInvestigation() bool {
	return o != nil && o.ShowInvestigation != nil
}

// SetShowInvestigation gets a reference to the given bool and assigns it to the ShowInvestigation field.
func (o *ManageStatusBackendDefinition) SetShowInvestigation(v bool) {
	o.ShowInvestigation = &v
}

// GetShowLastTriggered returns the ShowLastTriggered field value if set, zero value otherwise.
func (o *ManageStatusBackendDefinition) GetShowLastTriggered() bool {
	if o == nil || o.ShowLastTriggered == nil {
		var ret bool
		return ret
	}
	return *o.ShowLastTriggered
}

// GetShowLastTriggeredOk returns a tuple with the ShowLastTriggered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageStatusBackendDefinition) GetShowLastTriggeredOk() (*bool, bool) {
	if o == nil || o.ShowLastTriggered == nil {
		return nil, false
	}
	return o.ShowLastTriggered, true
}

// HasShowLastTriggered returns a boolean if a field has been set.
func (o *ManageStatusBackendDefinition) HasShowLastTriggered() bool {
	return o != nil && o.ShowLastTriggered != nil
}

// SetShowLastTriggered gets a reference to the given bool and assigns it to the ShowLastTriggered field.
func (o *ManageStatusBackendDefinition) SetShowLastTriggered(v bool) {
	o.ShowLastTriggered = &v
}

// GetShowPriority returns the ShowPriority field value if set, zero value otherwise.
func (o *ManageStatusBackendDefinition) GetShowPriority() bool {
	if o == nil || o.ShowPriority == nil {
		var ret bool
		return ret
	}
	return *o.ShowPriority
}

// GetShowPriorityOk returns a tuple with the ShowPriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageStatusBackendDefinition) GetShowPriorityOk() (*bool, bool) {
	if o == nil || o.ShowPriority == nil {
		return nil, false
	}
	return o.ShowPriority, true
}

// HasShowPriority returns a boolean if a field has been set.
func (o *ManageStatusBackendDefinition) HasShowPriority() bool {
	return o != nil && o.ShowPriority != nil
}

// SetShowPriority gets a reference to the given bool and assigns it to the ShowPriority field.
func (o *ManageStatusBackendDefinition) SetShowPriority(v bool) {
	o.ShowPriority = &v
}

// GetShowStatus returns the ShowStatus field value if set, zero value otherwise.
func (o *ManageStatusBackendDefinition) GetShowStatus() bool {
	if o == nil || o.ShowStatus == nil {
		var ret bool
		return ret
	}
	return *o.ShowStatus
}

// GetShowStatusOk returns a tuple with the ShowStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageStatusBackendDefinition) GetShowStatusOk() (*bool, bool) {
	if o == nil || o.ShowStatus == nil {
		return nil, false
	}
	return o.ShowStatus, true
}

// HasShowStatus returns a boolean if a field has been set.
func (o *ManageStatusBackendDefinition) HasShowStatus() bool {
	return o != nil && o.ShowStatus != nil
}

// SetShowStatus gets a reference to the given bool and assigns it to the ShowStatus field.
func (o *ManageStatusBackendDefinition) SetShowStatus(v bool) {
	o.ShowStatus = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *ManageStatusBackendDefinition) GetSort() string {
	if o == nil || o.Sort == nil {
		var ret string
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageStatusBackendDefinition) GetSortOk() (*string, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *ManageStatusBackendDefinition) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given string and assigns it to the Sort field.
func (o *ManageStatusBackendDefinition) SetSort(v string) {
	o.Sort = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *ManageStatusBackendDefinition) GetStart() float64 {
	if o == nil || o.Start == nil {
		var ret float64
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageStatusBackendDefinition) GetStartOk() (*float64, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *ManageStatusBackendDefinition) HasStart() bool {
	return o != nil && o.Start != nil
}

// SetStart gets a reference to the given float64 and assigns it to the Start field.
func (o *ManageStatusBackendDefinition) SetStart(v float64) {
	o.Start = &v
}

// GetSummaryType returns the SummaryType field value if set, zero value otherwise.
func (o *ManageStatusBackendDefinition) GetSummaryType() ManageStatusBackendDefinitionSummaryType {
	if o == nil || o.SummaryType == nil {
		var ret ManageStatusBackendDefinitionSummaryType
		return ret
	}
	return *o.SummaryType
}

// GetSummaryTypeOk returns a tuple with the SummaryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageStatusBackendDefinition) GetSummaryTypeOk() (*ManageStatusBackendDefinitionSummaryType, bool) {
	if o == nil || o.SummaryType == nil {
		return nil, false
	}
	return o.SummaryType, true
}

// HasSummaryType returns a boolean if a field has been set.
func (o *ManageStatusBackendDefinition) HasSummaryType() bool {
	return o != nil && o.SummaryType != nil
}

// SetSummaryType gets a reference to the given ManageStatusBackendDefinitionSummaryType and assigns it to the SummaryType field.
func (o *ManageStatusBackendDefinition) SetSummaryType(v ManageStatusBackendDefinitionSummaryType) {
	o.SummaryType = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ManageStatusBackendDefinition) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageStatusBackendDefinition) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ManageStatusBackendDefinition) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ManageStatusBackendDefinition) SetTitle(v string) {
	o.Title = &v
}

// GetTitleAlign returns the TitleAlign field value if set, zero value otherwise.
func (o *ManageStatusBackendDefinition) GetTitleAlign() WidgetTextAlign {
	if o == nil || o.TitleAlign == nil {
		var ret WidgetTextAlign
		return ret
	}
	return *o.TitleAlign
}

// GetTitleAlignOk returns a tuple with the TitleAlign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageStatusBackendDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool) {
	if o == nil || o.TitleAlign == nil {
		return nil, false
	}
	return o.TitleAlign, true
}

// HasTitleAlign returns a boolean if a field has been set.
func (o *ManageStatusBackendDefinition) HasTitleAlign() bool {
	return o != nil && o.TitleAlign != nil
}

// SetTitleAlign gets a reference to the given WidgetTextAlign and assigns it to the TitleAlign field.
func (o *ManageStatusBackendDefinition) SetTitleAlign(v WidgetTextAlign) {
	o.TitleAlign = &v
}

// GetTitleSize returns the TitleSize field value if set, zero value otherwise.
func (o *ManageStatusBackendDefinition) GetTitleSize() string {
	if o == nil || o.TitleSize == nil {
		var ret string
		return ret
	}
	return *o.TitleSize
}

// GetTitleSizeOk returns a tuple with the TitleSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageStatusBackendDefinition) GetTitleSizeOk() (*string, bool) {
	if o == nil || o.TitleSize == nil {
		return nil, false
	}
	return o.TitleSize, true
}

// HasTitleSize returns a boolean if a field has been set.
func (o *ManageStatusBackendDefinition) HasTitleSize() bool {
	return o != nil && o.TitleSize != nil
}

// SetTitleSize gets a reference to the given string and assigns it to the TitleSize field.
func (o *ManageStatusBackendDefinition) SetTitleSize(v string) {
	o.TitleSize = &v
}

// GetType returns the Type field value.
func (o *ManageStatusBackendDefinition) GetType() ManageStatusBackendDefinitionType {
	if o == nil {
		var ret ManageStatusBackendDefinitionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ManageStatusBackendDefinition) GetTypeOk() (*ManageStatusBackendDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ManageStatusBackendDefinition) SetType(v ManageStatusBackendDefinitionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ManageStatusBackendDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ColorPreference != nil {
		toSerialize["color_preference"] = o.ColorPreference
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisplayFormat != nil {
		toSerialize["display_format"] = o.DisplayFormat
	}
	if o.HideZeroCounts != nil {
		toSerialize["hide_zero_counts"] = o.HideZeroCounts
	}
	if o.LastTriggeredFormat != nil {
		toSerialize["last_triggered_format"] = o.LastTriggeredFormat
	}
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
	}
	toSerialize["query"] = o.Query
	if o.ShowInvestigation != nil {
		toSerialize["show_investigation"] = o.ShowInvestigation
	}
	if o.ShowLastTriggered != nil {
		toSerialize["show_last_triggered"] = o.ShowLastTriggered
	}
	if o.ShowPriority != nil {
		toSerialize["show_priority"] = o.ShowPriority
	}
	if o.ShowStatus != nil {
		toSerialize["show_status"] = o.ShowStatus
	}
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}
	if o.Start != nil {
		toSerialize["start"] = o.Start
	}
	if o.SummaryType != nil {
		toSerialize["summary_type"] = o.SummaryType
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
func (o *ManageStatusBackendDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ColorPreference     *ManageStatusBackendDefinitionColorPreference     `json:"color_preference,omitempty"`
		Count               *float64                                          `json:"count,omitempty"`
		Description         *string                                           `json:"description,omitempty"`
		DisplayFormat       *ManageStatusBackendDefinitionDisplayFormat       `json:"display_format,omitempty"`
		HideZeroCounts      *bool                                             `json:"hide_zero_counts,omitempty"`
		LastTriggeredFormat *ManageStatusBackendDefinitionLastTriggeredFormat `json:"last_triggered_format,omitempty"`
		Notes               *string                                           `json:"notes,omitempty"`
		Query               *string                                           `json:"query"`
		ShowInvestigation   *bool                                             `json:"show_investigation,omitempty"`
		ShowLastTriggered   *bool                                             `json:"show_last_triggered,omitempty"`
		ShowPriority        *bool                                             `json:"show_priority,omitempty"`
		ShowStatus          *bool                                             `json:"show_status,omitempty"`
		Sort                *string                                           `json:"sort,omitempty"`
		Start               *float64                                          `json:"start,omitempty"`
		SummaryType         *ManageStatusBackendDefinitionSummaryType         `json:"summary_type,omitempty"`
		Title               *string                                           `json:"title,omitempty"`
		TitleAlign          *WidgetTextAlign                                  `json:"title_align,omitempty"`
		TitleSize           *string                                           `json:"title_size,omitempty"`
		Type                *ManageStatusBackendDefinitionType                `json:"type"`
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
	if all.ColorPreference != nil && !all.ColorPreference.IsValid() {
		hasInvalidField = true
	} else {
		o.ColorPreference = all.ColorPreference
	}
	o.Count = all.Count
	o.Description = all.Description
	if all.DisplayFormat != nil && !all.DisplayFormat.IsValid() {
		hasInvalidField = true
	} else {
		o.DisplayFormat = all.DisplayFormat
	}
	o.HideZeroCounts = all.HideZeroCounts
	if all.LastTriggeredFormat != nil && !all.LastTriggeredFormat.IsValid() {
		hasInvalidField = true
	} else {
		o.LastTriggeredFormat = all.LastTriggeredFormat
	}
	o.Notes = all.Notes
	o.Query = *all.Query
	o.ShowInvestigation = all.ShowInvestigation
	o.ShowLastTriggered = all.ShowLastTriggered
	o.ShowPriority = all.ShowPriority
	o.ShowStatus = all.ShowStatus
	o.Sort = all.Sort
	o.Start = all.Start
	if all.SummaryType != nil && !all.SummaryType.IsValid() {
		hasInvalidField = true
	} else {
		o.SummaryType = all.SummaryType
	}
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
