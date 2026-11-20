// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaSloQuery Test-only generated schema for the dashboard versioning proof of concept.
type FormulaSloQuery struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	AdditionalQueryFilters *string `json:"additional_query_filters,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ApplyCorrections *bool `json:"apply_corrections,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource FormulaSloQueryDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	GroupMode SloGroupMode `json:"group_mode"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Measure SloMeasure `json:"measure"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	SloId string `json:"slo_id"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	SloQueryType SloQueryType `json:"slo_query_type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	SloTimeWindow *SloTimeWindow `json:"slo_time_window,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewFormulaSloQuery instantiates a new FormulaSloQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormulaSloQuery(dataSource FormulaSloQueryDataSource, groupMode SloGroupMode, measure SloMeasure, name string, sloId string, sloQueryType SloQueryType) *FormulaSloQuery {
	this := FormulaSloQuery{}
	this.DataSource = dataSource
	this.GroupMode = groupMode
	this.Measure = measure
	this.Name = name
	this.SloId = sloId
	this.SloQueryType = sloQueryType
	return &this
}

// NewFormulaSloQueryWithDefaults instantiates a new FormulaSloQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormulaSloQueryWithDefaults() *FormulaSloQuery {
	this := FormulaSloQuery{}
	return &this
}

// GetAdditionalQueryFilters returns the AdditionalQueryFilters field value if set, zero value otherwise.
func (o *FormulaSloQuery) GetAdditionalQueryFilters() string {
	if o == nil || o.AdditionalQueryFilters == nil {
		var ret string
		return ret
	}
	return *o.AdditionalQueryFilters
}

// GetAdditionalQueryFiltersOk returns a tuple with the AdditionalQueryFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaSloQuery) GetAdditionalQueryFiltersOk() (*string, bool) {
	if o == nil || o.AdditionalQueryFilters == nil {
		return nil, false
	}
	return o.AdditionalQueryFilters, true
}

// HasAdditionalQueryFilters returns a boolean if a field has been set.
func (o *FormulaSloQuery) HasAdditionalQueryFilters() bool {
	return o != nil && o.AdditionalQueryFilters != nil
}

// SetAdditionalQueryFilters gets a reference to the given string and assigns it to the AdditionalQueryFilters field.
func (o *FormulaSloQuery) SetAdditionalQueryFilters(v string) {
	o.AdditionalQueryFilters = &v
}

// GetApplyCorrections returns the ApplyCorrections field value if set, zero value otherwise.
func (o *FormulaSloQuery) GetApplyCorrections() bool {
	if o == nil || o.ApplyCorrections == nil {
		var ret bool
		return ret
	}
	return *o.ApplyCorrections
}

// GetApplyCorrectionsOk returns a tuple with the ApplyCorrections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaSloQuery) GetApplyCorrectionsOk() (*bool, bool) {
	if o == nil || o.ApplyCorrections == nil {
		return nil, false
	}
	return o.ApplyCorrections, true
}

// HasApplyCorrections returns a boolean if a field has been set.
func (o *FormulaSloQuery) HasApplyCorrections() bool {
	return o != nil && o.ApplyCorrections != nil
}

// SetApplyCorrections gets a reference to the given bool and assigns it to the ApplyCorrections field.
func (o *FormulaSloQuery) SetApplyCorrections(v bool) {
	o.ApplyCorrections = &v
}

// GetDataSource returns the DataSource field value.
func (o *FormulaSloQuery) GetDataSource() FormulaSloQueryDataSource {
	if o == nil {
		var ret FormulaSloQueryDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *FormulaSloQuery) GetDataSourceOk() (*FormulaSloQueryDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *FormulaSloQuery) SetDataSource(v FormulaSloQueryDataSource) {
	o.DataSource = v
}

// GetGroupMode returns the GroupMode field value.
func (o *FormulaSloQuery) GetGroupMode() SloGroupMode {
	if o == nil {
		var ret SloGroupMode
		return ret
	}
	return o.GroupMode
}

// GetGroupModeOk returns a tuple with the GroupMode field value
// and a boolean to check if the value has been set.
func (o *FormulaSloQuery) GetGroupModeOk() (*SloGroupMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupMode, true
}

// SetGroupMode sets field value.
func (o *FormulaSloQuery) SetGroupMode(v SloGroupMode) {
	o.GroupMode = v
}

// GetMeasure returns the Measure field value.
func (o *FormulaSloQuery) GetMeasure() SloMeasure {
	if o == nil {
		var ret SloMeasure
		return ret
	}
	return o.Measure
}

// GetMeasureOk returns a tuple with the Measure field value
// and a boolean to check if the value has been set.
func (o *FormulaSloQuery) GetMeasureOk() (*SloMeasure, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Measure, true
}

// SetMeasure sets field value.
func (o *FormulaSloQuery) SetMeasure(v SloMeasure) {
	o.Measure = v
}

// GetName returns the Name field value.
func (o *FormulaSloQuery) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FormulaSloQuery) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *FormulaSloQuery) SetName(v string) {
	o.Name = v
}

// GetSloId returns the SloId field value.
func (o *FormulaSloQuery) GetSloId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SloId
}

// GetSloIdOk returns a tuple with the SloId field value
// and a boolean to check if the value has been set.
func (o *FormulaSloQuery) GetSloIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SloId, true
}

// SetSloId sets field value.
func (o *FormulaSloQuery) SetSloId(v string) {
	o.SloId = v
}

// GetSloQueryType returns the SloQueryType field value.
func (o *FormulaSloQuery) GetSloQueryType() SloQueryType {
	if o == nil {
		var ret SloQueryType
		return ret
	}
	return o.SloQueryType
}

// GetSloQueryTypeOk returns a tuple with the SloQueryType field value
// and a boolean to check if the value has been set.
func (o *FormulaSloQuery) GetSloQueryTypeOk() (*SloQueryType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SloQueryType, true
}

// SetSloQueryType sets field value.
func (o *FormulaSloQuery) SetSloQueryType(v SloQueryType) {
	o.SloQueryType = v
}

// GetSloTimeWindow returns the SloTimeWindow field value if set, zero value otherwise.
func (o *FormulaSloQuery) GetSloTimeWindow() SloTimeWindow {
	if o == nil || o.SloTimeWindow == nil {
		var ret SloTimeWindow
		return ret
	}
	return *o.SloTimeWindow
}

// GetSloTimeWindowOk returns a tuple with the SloTimeWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaSloQuery) GetSloTimeWindowOk() (*SloTimeWindow, bool) {
	if o == nil || o.SloTimeWindow == nil {
		return nil, false
	}
	return o.SloTimeWindow, true
}

// HasSloTimeWindow returns a boolean if a field has been set.
func (o *FormulaSloQuery) HasSloTimeWindow() bool {
	return o != nil && o.SloTimeWindow != nil
}

// SetSloTimeWindow gets a reference to the given SloTimeWindow and assigns it to the SloTimeWindow field.
func (o *FormulaSloQuery) SetSloTimeWindow(v SloTimeWindow) {
	o.SloTimeWindow = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormulaSloQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AdditionalQueryFilters != nil {
		toSerialize["additional_query_filters"] = o.AdditionalQueryFilters
	}
	if o.ApplyCorrections != nil {
		toSerialize["apply_corrections"] = o.ApplyCorrections
	}
	toSerialize["data_source"] = o.DataSource
	toSerialize["group_mode"] = o.GroupMode
	toSerialize["measure"] = o.Measure
	toSerialize["name"] = o.Name
	toSerialize["slo_id"] = o.SloId
	toSerialize["slo_query_type"] = o.SloQueryType
	if o.SloTimeWindow != nil {
		toSerialize["slo_time_window"] = o.SloTimeWindow
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormulaSloQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AdditionalQueryFilters *string                    `json:"additional_query_filters,omitempty"`
		ApplyCorrections       *bool                      `json:"apply_corrections,omitempty"`
		DataSource             *FormulaSloQueryDataSource `json:"data_source"`
		GroupMode              *SloGroupMode              `json:"group_mode"`
		Measure                *SloMeasure                `json:"measure"`
		Name                   *string                    `json:"name"`
		SloId                  *string                    `json:"slo_id"`
		SloQueryType           *SloQueryType              `json:"slo_query_type"`
		SloTimeWindow          *SloTimeWindow             `json:"slo_time_window,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field data_source missing")
	}
	if all.GroupMode == nil {
		return fmt.Errorf("required field group_mode missing")
	}
	if all.Measure == nil {
		return fmt.Errorf("required field measure missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.SloId == nil {
		return fmt.Errorf("required field slo_id missing")
	}
	if all.SloQueryType == nil {
		return fmt.Errorf("required field slo_query_type missing")
	}

	hasInvalidField := false
	o.AdditionalQueryFilters = all.AdditionalQueryFilters
	o.ApplyCorrections = all.ApplyCorrections
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	if !all.GroupMode.IsValid() {
		hasInvalidField = true
	} else {
		o.GroupMode = *all.GroupMode
	}
	o.Measure = *all.Measure
	o.Name = *all.Name
	o.SloId = *all.SloId
	if !all.SloQueryType.IsValid() {
		hasInvalidField = true
	} else {
		o.SloQueryType = *all.SloQueryType
	}
	if all.SloTimeWindow != nil && !all.SloTimeWindow.IsValid() {
		hasInvalidField = true
	} else {
		o.SloTimeWindow = all.SloTimeWindow
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
