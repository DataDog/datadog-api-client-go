// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaDataQualityMetricsQuery Test-only generated schema for the dashboard versioning proof of concept.
type FormulaDataQualityMetricsQuery struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	AnomalyOptions *FormulaDataQualityMetricsQueryAnomalyOptions `json:"anomaly_options,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource FormulaDataQualityMetricsQueryDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Filter string `json:"filter"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	GroupBy []string `json:"group_by,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Measure string `json:"measure"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Scope *string `json:"scope,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewFormulaDataQualityMetricsQuery instantiates a new FormulaDataQualityMetricsQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormulaDataQualityMetricsQuery(dataSource FormulaDataQualityMetricsQueryDataSource, filter string, measure string, name string) *FormulaDataQualityMetricsQuery {
	this := FormulaDataQualityMetricsQuery{}
	this.DataSource = dataSource
	this.Filter = filter
	this.Measure = measure
	this.Name = name
	return &this
}

// NewFormulaDataQualityMetricsQueryWithDefaults instantiates a new FormulaDataQualityMetricsQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormulaDataQualityMetricsQueryWithDefaults() *FormulaDataQualityMetricsQuery {
	this := FormulaDataQualityMetricsQuery{}
	return &this
}

// GetAnomalyOptions returns the AnomalyOptions field value if set, zero value otherwise.
func (o *FormulaDataQualityMetricsQuery) GetAnomalyOptions() FormulaDataQualityMetricsQueryAnomalyOptions {
	if o == nil || o.AnomalyOptions == nil {
		var ret FormulaDataQualityMetricsQueryAnomalyOptions
		return ret
	}
	return *o.AnomalyOptions
}

// GetAnomalyOptionsOk returns a tuple with the AnomalyOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaDataQualityMetricsQuery) GetAnomalyOptionsOk() (*FormulaDataQualityMetricsQueryAnomalyOptions, bool) {
	if o == nil || o.AnomalyOptions == nil {
		return nil, false
	}
	return o.AnomalyOptions, true
}

// HasAnomalyOptions returns a boolean if a field has been set.
func (o *FormulaDataQualityMetricsQuery) HasAnomalyOptions() bool {
	return o != nil && o.AnomalyOptions != nil
}

// SetAnomalyOptions gets a reference to the given FormulaDataQualityMetricsQueryAnomalyOptions and assigns it to the AnomalyOptions field.
func (o *FormulaDataQualityMetricsQuery) SetAnomalyOptions(v FormulaDataQualityMetricsQueryAnomalyOptions) {
	o.AnomalyOptions = &v
}

// GetDataSource returns the DataSource field value.
func (o *FormulaDataQualityMetricsQuery) GetDataSource() FormulaDataQualityMetricsQueryDataSource {
	if o == nil {
		var ret FormulaDataQualityMetricsQueryDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *FormulaDataQualityMetricsQuery) GetDataSourceOk() (*FormulaDataQualityMetricsQueryDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *FormulaDataQualityMetricsQuery) SetDataSource(v FormulaDataQualityMetricsQueryDataSource) {
	o.DataSource = v
}

// GetFilter returns the Filter field value.
func (o *FormulaDataQualityMetricsQuery) GetFilter() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *FormulaDataQualityMetricsQuery) GetFilterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filter, true
}

// SetFilter sets field value.
func (o *FormulaDataQualityMetricsQuery) SetFilter(v string) {
	o.Filter = v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *FormulaDataQualityMetricsQuery) GetGroupBy() []string {
	if o == nil || o.GroupBy == nil {
		var ret []string
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaDataQualityMetricsQuery) GetGroupByOk() (*[]string, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *FormulaDataQualityMetricsQuery) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []string and assigns it to the GroupBy field.
func (o *FormulaDataQualityMetricsQuery) SetGroupBy(v []string) {
	o.GroupBy = v
}

// GetMeasure returns the Measure field value.
func (o *FormulaDataQualityMetricsQuery) GetMeasure() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Measure
}

// GetMeasureOk returns a tuple with the Measure field value
// and a boolean to check if the value has been set.
func (o *FormulaDataQualityMetricsQuery) GetMeasureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Measure, true
}

// SetMeasure sets field value.
func (o *FormulaDataQualityMetricsQuery) SetMeasure(v string) {
	o.Measure = v
}

// GetName returns the Name field value.
func (o *FormulaDataQualityMetricsQuery) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FormulaDataQualityMetricsQuery) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *FormulaDataQualityMetricsQuery) SetName(v string) {
	o.Name = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *FormulaDataQualityMetricsQuery) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaDataQualityMetricsQuery) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *FormulaDataQualityMetricsQuery) HasScope() bool {
	return o != nil && o.Scope != nil
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *FormulaDataQualityMetricsQuery) SetScope(v string) {
	o.Scope = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormulaDataQualityMetricsQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AnomalyOptions != nil {
		toSerialize["anomaly_options"] = o.AnomalyOptions
	}
	toSerialize["data_source"] = o.DataSource
	toSerialize["filter"] = o.Filter
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}
	toSerialize["measure"] = o.Measure
	toSerialize["name"] = o.Name
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormulaDataQualityMetricsQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AnomalyOptions *FormulaDataQualityMetricsQueryAnomalyOptions `json:"anomaly_options,omitempty"`
		DataSource     *FormulaDataQualityMetricsQueryDataSource     `json:"data_source"`
		Filter         *string                                       `json:"filter"`
		GroupBy        []string                                      `json:"group_by,omitempty"`
		Measure        *string                                       `json:"measure"`
		Name           *string                                       `json:"name"`
		Scope          *string                                       `json:"scope,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field data_source missing")
	}
	if all.Filter == nil {
		return fmt.Errorf("required field filter missing")
	}
	if all.Measure == nil {
		return fmt.Errorf("required field measure missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}

	hasInvalidField := false
	if all.AnomalyOptions != nil && all.AnomalyOptions.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AnomalyOptions = all.AnomalyOptions
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.Filter = *all.Filter
	o.GroupBy = all.GroupBy
	o.Measure = *all.Measure
	o.Name = *all.Name
	o.Scope = all.Scope

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
