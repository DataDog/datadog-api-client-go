// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceSummaryBackendDefinition Test-only generated schema for the dashboard versioning proof of concept.
type ServiceSummaryBackendDefinition struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Description *string `json:"description,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DisplayFormat *ServiceSummaryLayout `json:"display_format,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Env string `json:"env"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Notes *string `json:"notes,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Service string `json:"service"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ShowBreakdown *bool `json:"show_breakdown,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ShowDistribution *bool `json:"show_distribution,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ShowErrors *bool `json:"show_errors,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ShowHits *bool `json:"show_hits,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ShowLatency *bool `json:"show_latency,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ShowResourceList *bool `json:"show_resource_list,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	SizeFormat *ServiceSummarySize `json:"size_format,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	SpanName string `json:"span_name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Time *WidgetTime `json:"time,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Title *string `json:"title,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TitleAlign *WidgetTextAlign `json:"title_align,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TitleSize *string `json:"title_size,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type ServiceSummaryBackendDefinitionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewServiceSummaryBackendDefinition instantiates a new ServiceSummaryBackendDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewServiceSummaryBackendDefinition(env string, service string, spanName string, typeVar ServiceSummaryBackendDefinitionType) *ServiceSummaryBackendDefinition {
	this := ServiceSummaryBackendDefinition{}
	this.Env = env
	this.Service = service
	this.SpanName = spanName
	this.Type = typeVar
	return &this
}

// NewServiceSummaryBackendDefinitionWithDefaults instantiates a new ServiceSummaryBackendDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewServiceSummaryBackendDefinitionWithDefaults() *ServiceSummaryBackendDefinition {
	this := ServiceSummaryBackendDefinition{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ServiceSummaryBackendDefinition) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSummaryBackendDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ServiceSummaryBackendDefinition) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ServiceSummaryBackendDefinition) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayFormat returns the DisplayFormat field value if set, zero value otherwise.
func (o *ServiceSummaryBackendDefinition) GetDisplayFormat() ServiceSummaryLayout {
	if o == nil || o.DisplayFormat == nil {
		var ret ServiceSummaryLayout
		return ret
	}
	return *o.DisplayFormat
}

// GetDisplayFormatOk returns a tuple with the DisplayFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSummaryBackendDefinition) GetDisplayFormatOk() (*ServiceSummaryLayout, bool) {
	if o == nil || o.DisplayFormat == nil {
		return nil, false
	}
	return o.DisplayFormat, true
}

// HasDisplayFormat returns a boolean if a field has been set.
func (o *ServiceSummaryBackendDefinition) HasDisplayFormat() bool {
	return o != nil && o.DisplayFormat != nil
}

// SetDisplayFormat gets a reference to the given ServiceSummaryLayout and assigns it to the DisplayFormat field.
func (o *ServiceSummaryBackendDefinition) SetDisplayFormat(v ServiceSummaryLayout) {
	o.DisplayFormat = &v
}

// GetEnv returns the Env field value.
func (o *ServiceSummaryBackendDefinition) GetEnv() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Env
}

// GetEnvOk returns a tuple with the Env field value
// and a boolean to check if the value has been set.
func (o *ServiceSummaryBackendDefinition) GetEnvOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Env, true
}

// SetEnv sets field value.
func (o *ServiceSummaryBackendDefinition) SetEnv(v string) {
	o.Env = v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *ServiceSummaryBackendDefinition) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSummaryBackendDefinition) GetNotesOk() (*string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *ServiceSummaryBackendDefinition) HasNotes() bool {
	return o != nil && o.Notes != nil
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *ServiceSummaryBackendDefinition) SetNotes(v string) {
	o.Notes = &v
}

// GetService returns the Service field value.
func (o *ServiceSummaryBackendDefinition) GetService() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *ServiceSummaryBackendDefinition) GetServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value.
func (o *ServiceSummaryBackendDefinition) SetService(v string) {
	o.Service = v
}

// GetShowBreakdown returns the ShowBreakdown field value if set, zero value otherwise.
func (o *ServiceSummaryBackendDefinition) GetShowBreakdown() bool {
	if o == nil || o.ShowBreakdown == nil {
		var ret bool
		return ret
	}
	return *o.ShowBreakdown
}

// GetShowBreakdownOk returns a tuple with the ShowBreakdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSummaryBackendDefinition) GetShowBreakdownOk() (*bool, bool) {
	if o == nil || o.ShowBreakdown == nil {
		return nil, false
	}
	return o.ShowBreakdown, true
}

// HasShowBreakdown returns a boolean if a field has been set.
func (o *ServiceSummaryBackendDefinition) HasShowBreakdown() bool {
	return o != nil && o.ShowBreakdown != nil
}

// SetShowBreakdown gets a reference to the given bool and assigns it to the ShowBreakdown field.
func (o *ServiceSummaryBackendDefinition) SetShowBreakdown(v bool) {
	o.ShowBreakdown = &v
}

// GetShowDistribution returns the ShowDistribution field value if set, zero value otherwise.
func (o *ServiceSummaryBackendDefinition) GetShowDistribution() bool {
	if o == nil || o.ShowDistribution == nil {
		var ret bool
		return ret
	}
	return *o.ShowDistribution
}

// GetShowDistributionOk returns a tuple with the ShowDistribution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSummaryBackendDefinition) GetShowDistributionOk() (*bool, bool) {
	if o == nil || o.ShowDistribution == nil {
		return nil, false
	}
	return o.ShowDistribution, true
}

// HasShowDistribution returns a boolean if a field has been set.
func (o *ServiceSummaryBackendDefinition) HasShowDistribution() bool {
	return o != nil && o.ShowDistribution != nil
}

// SetShowDistribution gets a reference to the given bool and assigns it to the ShowDistribution field.
func (o *ServiceSummaryBackendDefinition) SetShowDistribution(v bool) {
	o.ShowDistribution = &v
}

// GetShowErrors returns the ShowErrors field value if set, zero value otherwise.
func (o *ServiceSummaryBackendDefinition) GetShowErrors() bool {
	if o == nil || o.ShowErrors == nil {
		var ret bool
		return ret
	}
	return *o.ShowErrors
}

// GetShowErrorsOk returns a tuple with the ShowErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSummaryBackendDefinition) GetShowErrorsOk() (*bool, bool) {
	if o == nil || o.ShowErrors == nil {
		return nil, false
	}
	return o.ShowErrors, true
}

// HasShowErrors returns a boolean if a field has been set.
func (o *ServiceSummaryBackendDefinition) HasShowErrors() bool {
	return o != nil && o.ShowErrors != nil
}

// SetShowErrors gets a reference to the given bool and assigns it to the ShowErrors field.
func (o *ServiceSummaryBackendDefinition) SetShowErrors(v bool) {
	o.ShowErrors = &v
}

// GetShowHits returns the ShowHits field value if set, zero value otherwise.
func (o *ServiceSummaryBackendDefinition) GetShowHits() bool {
	if o == nil || o.ShowHits == nil {
		var ret bool
		return ret
	}
	return *o.ShowHits
}

// GetShowHitsOk returns a tuple with the ShowHits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSummaryBackendDefinition) GetShowHitsOk() (*bool, bool) {
	if o == nil || o.ShowHits == nil {
		return nil, false
	}
	return o.ShowHits, true
}

// HasShowHits returns a boolean if a field has been set.
func (o *ServiceSummaryBackendDefinition) HasShowHits() bool {
	return o != nil && o.ShowHits != nil
}

// SetShowHits gets a reference to the given bool and assigns it to the ShowHits field.
func (o *ServiceSummaryBackendDefinition) SetShowHits(v bool) {
	o.ShowHits = &v
}

// GetShowLatency returns the ShowLatency field value if set, zero value otherwise.
func (o *ServiceSummaryBackendDefinition) GetShowLatency() bool {
	if o == nil || o.ShowLatency == nil {
		var ret bool
		return ret
	}
	return *o.ShowLatency
}

// GetShowLatencyOk returns a tuple with the ShowLatency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSummaryBackendDefinition) GetShowLatencyOk() (*bool, bool) {
	if o == nil || o.ShowLatency == nil {
		return nil, false
	}
	return o.ShowLatency, true
}

// HasShowLatency returns a boolean if a field has been set.
func (o *ServiceSummaryBackendDefinition) HasShowLatency() bool {
	return o != nil && o.ShowLatency != nil
}

// SetShowLatency gets a reference to the given bool and assigns it to the ShowLatency field.
func (o *ServiceSummaryBackendDefinition) SetShowLatency(v bool) {
	o.ShowLatency = &v
}

// GetShowResourceList returns the ShowResourceList field value if set, zero value otherwise.
func (o *ServiceSummaryBackendDefinition) GetShowResourceList() bool {
	if o == nil || o.ShowResourceList == nil {
		var ret bool
		return ret
	}
	return *o.ShowResourceList
}

// GetShowResourceListOk returns a tuple with the ShowResourceList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSummaryBackendDefinition) GetShowResourceListOk() (*bool, bool) {
	if o == nil || o.ShowResourceList == nil {
		return nil, false
	}
	return o.ShowResourceList, true
}

// HasShowResourceList returns a boolean if a field has been set.
func (o *ServiceSummaryBackendDefinition) HasShowResourceList() bool {
	return o != nil && o.ShowResourceList != nil
}

// SetShowResourceList gets a reference to the given bool and assigns it to the ShowResourceList field.
func (o *ServiceSummaryBackendDefinition) SetShowResourceList(v bool) {
	o.ShowResourceList = &v
}

// GetSizeFormat returns the SizeFormat field value if set, zero value otherwise.
func (o *ServiceSummaryBackendDefinition) GetSizeFormat() ServiceSummarySize {
	if o == nil || o.SizeFormat == nil {
		var ret ServiceSummarySize
		return ret
	}
	return *o.SizeFormat
}

// GetSizeFormatOk returns a tuple with the SizeFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSummaryBackendDefinition) GetSizeFormatOk() (*ServiceSummarySize, bool) {
	if o == nil || o.SizeFormat == nil {
		return nil, false
	}
	return o.SizeFormat, true
}

// HasSizeFormat returns a boolean if a field has been set.
func (o *ServiceSummaryBackendDefinition) HasSizeFormat() bool {
	return o != nil && o.SizeFormat != nil
}

// SetSizeFormat gets a reference to the given ServiceSummarySize and assigns it to the SizeFormat field.
func (o *ServiceSummaryBackendDefinition) SetSizeFormat(v ServiceSummarySize) {
	o.SizeFormat = &v
}

// GetSpanName returns the SpanName field value.
func (o *ServiceSummaryBackendDefinition) GetSpanName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SpanName
}

// GetSpanNameOk returns a tuple with the SpanName field value
// and a boolean to check if the value has been set.
func (o *ServiceSummaryBackendDefinition) GetSpanNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpanName, true
}

// SetSpanName sets field value.
func (o *ServiceSummaryBackendDefinition) SetSpanName(v string) {
	o.SpanName = v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *ServiceSummaryBackendDefinition) GetTime() WidgetTime {
	if o == nil || o.Time == nil {
		var ret WidgetTime
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSummaryBackendDefinition) GetTimeOk() (*WidgetTime, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *ServiceSummaryBackendDefinition) HasTime() bool {
	return o != nil && o.Time != nil
}

// SetTime gets a reference to the given WidgetTime and assigns it to the Time field.
func (o *ServiceSummaryBackendDefinition) SetTime(v WidgetTime) {
	o.Time = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ServiceSummaryBackendDefinition) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSummaryBackendDefinition) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ServiceSummaryBackendDefinition) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ServiceSummaryBackendDefinition) SetTitle(v string) {
	o.Title = &v
}

// GetTitleAlign returns the TitleAlign field value if set, zero value otherwise.
func (o *ServiceSummaryBackendDefinition) GetTitleAlign() WidgetTextAlign {
	if o == nil || o.TitleAlign == nil {
		var ret WidgetTextAlign
		return ret
	}
	return *o.TitleAlign
}

// GetTitleAlignOk returns a tuple with the TitleAlign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSummaryBackendDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool) {
	if o == nil || o.TitleAlign == nil {
		return nil, false
	}
	return o.TitleAlign, true
}

// HasTitleAlign returns a boolean if a field has been set.
func (o *ServiceSummaryBackendDefinition) HasTitleAlign() bool {
	return o != nil && o.TitleAlign != nil
}

// SetTitleAlign gets a reference to the given WidgetTextAlign and assigns it to the TitleAlign field.
func (o *ServiceSummaryBackendDefinition) SetTitleAlign(v WidgetTextAlign) {
	o.TitleAlign = &v
}

// GetTitleSize returns the TitleSize field value if set, zero value otherwise.
func (o *ServiceSummaryBackendDefinition) GetTitleSize() string {
	if o == nil || o.TitleSize == nil {
		var ret string
		return ret
	}
	return *o.TitleSize
}

// GetTitleSizeOk returns a tuple with the TitleSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSummaryBackendDefinition) GetTitleSizeOk() (*string, bool) {
	if o == nil || o.TitleSize == nil {
		return nil, false
	}
	return o.TitleSize, true
}

// HasTitleSize returns a boolean if a field has been set.
func (o *ServiceSummaryBackendDefinition) HasTitleSize() bool {
	return o != nil && o.TitleSize != nil
}

// SetTitleSize gets a reference to the given string and assigns it to the TitleSize field.
func (o *ServiceSummaryBackendDefinition) SetTitleSize(v string) {
	o.TitleSize = &v
}

// GetType returns the Type field value.
func (o *ServiceSummaryBackendDefinition) GetType() ServiceSummaryBackendDefinitionType {
	if o == nil {
		var ret ServiceSummaryBackendDefinitionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ServiceSummaryBackendDefinition) GetTypeOk() (*ServiceSummaryBackendDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ServiceSummaryBackendDefinition) SetType(v ServiceSummaryBackendDefinitionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ServiceSummaryBackendDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisplayFormat != nil {
		toSerialize["display_format"] = o.DisplayFormat
	}
	toSerialize["env"] = o.Env
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
	}
	toSerialize["service"] = o.Service
	if o.ShowBreakdown != nil {
		toSerialize["show_breakdown"] = o.ShowBreakdown
	}
	if o.ShowDistribution != nil {
		toSerialize["show_distribution"] = o.ShowDistribution
	}
	if o.ShowErrors != nil {
		toSerialize["show_errors"] = o.ShowErrors
	}
	if o.ShowHits != nil {
		toSerialize["show_hits"] = o.ShowHits
	}
	if o.ShowLatency != nil {
		toSerialize["show_latency"] = o.ShowLatency
	}
	if o.ShowResourceList != nil {
		toSerialize["show_resource_list"] = o.ShowResourceList
	}
	if o.SizeFormat != nil {
		toSerialize["size_format"] = o.SizeFormat
	}
	toSerialize["span_name"] = o.SpanName
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
func (o *ServiceSummaryBackendDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description      *string                              `json:"description,omitempty"`
		DisplayFormat    *ServiceSummaryLayout                `json:"display_format,omitempty"`
		Env              *string                              `json:"env"`
		Notes            *string                              `json:"notes,omitempty"`
		Service          *string                              `json:"service"`
		ShowBreakdown    *bool                                `json:"show_breakdown,omitempty"`
		ShowDistribution *bool                                `json:"show_distribution,omitempty"`
		ShowErrors       *bool                                `json:"show_errors,omitempty"`
		ShowHits         *bool                                `json:"show_hits,omitempty"`
		ShowLatency      *bool                                `json:"show_latency,omitempty"`
		ShowResourceList *bool                                `json:"show_resource_list,omitempty"`
		SizeFormat       *ServiceSummarySize                  `json:"size_format,omitempty"`
		SpanName         *string                              `json:"span_name"`
		Time             *WidgetTime                          `json:"time,omitempty"`
		Title            *string                              `json:"title,omitempty"`
		TitleAlign       *WidgetTextAlign                     `json:"title_align,omitempty"`
		TitleSize        *string                              `json:"title_size,omitempty"`
		Type             *ServiceSummaryBackendDefinitionType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Env == nil {
		return fmt.Errorf("required field env missing")
	}
	if all.Service == nil {
		return fmt.Errorf("required field service missing")
	}
	if all.SpanName == nil {
		return fmt.Errorf("required field span_name missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	o.Description = all.Description
	if all.DisplayFormat != nil && !all.DisplayFormat.IsValid() {
		hasInvalidField = true
	} else {
		o.DisplayFormat = all.DisplayFormat
	}
	o.Env = *all.Env
	o.Notes = all.Notes
	o.Service = *all.Service
	o.ShowBreakdown = all.ShowBreakdown
	o.ShowDistribution = all.ShowDistribution
	o.ShowErrors = all.ShowErrors
	o.ShowHits = all.ShowHits
	o.ShowLatency = all.ShowLatency
	o.ShowResourceList = all.ShowResourceList
	if all.SizeFormat != nil && !all.SizeFormat.IsValid() {
		hasInvalidField = true
	} else {
		o.SizeFormat = all.SizeFormat
	}
	o.SpanName = *all.SpanName
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
