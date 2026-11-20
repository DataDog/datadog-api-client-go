// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetricsHistogramRequestQuery Test-only generated schema for the dashboard versioning proof of concept.
type MetricsHistogramRequestQuery struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Aggregator *ScalarAggregator `json:"aggregator,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ByocCluster *string `json:"byoc_cluster,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	CrossOrgUuids []string `json:"cross_org_uuids,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource FormulaMetricsDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	HistogramMode *HistogramMode `json:"histogram_mode,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Query string `json:"query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	SemanticMode *FormulaSemanticMode `json:"semantic_mode,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewMetricsHistogramRequestQuery instantiates a new MetricsHistogramRequestQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMetricsHistogramRequestQuery(dataSource FormulaMetricsDataSource, name string, query string) *MetricsHistogramRequestQuery {
	this := MetricsHistogramRequestQuery{}
	this.DataSource = dataSource
	this.Name = name
	this.Query = query
	return &this
}

// NewMetricsHistogramRequestQueryWithDefaults instantiates a new MetricsHistogramRequestQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMetricsHistogramRequestQueryWithDefaults() *MetricsHistogramRequestQuery {
	this := MetricsHistogramRequestQuery{}
	return &this
}

// GetAggregator returns the Aggregator field value if set, zero value otherwise.
func (o *MetricsHistogramRequestQuery) GetAggregator() ScalarAggregator {
	if o == nil || o.Aggregator == nil {
		var ret ScalarAggregator
		return ret
	}
	return *o.Aggregator
}

// GetAggregatorOk returns a tuple with the Aggregator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsHistogramRequestQuery) GetAggregatorOk() (*ScalarAggregator, bool) {
	if o == nil || o.Aggregator == nil {
		return nil, false
	}
	return o.Aggregator, true
}

// HasAggregator returns a boolean if a field has been set.
func (o *MetricsHistogramRequestQuery) HasAggregator() bool {
	return o != nil && o.Aggregator != nil
}

// SetAggregator gets a reference to the given ScalarAggregator and assigns it to the Aggregator field.
func (o *MetricsHistogramRequestQuery) SetAggregator(v ScalarAggregator) {
	o.Aggregator = &v
}

// GetByocCluster returns the ByocCluster field value if set, zero value otherwise.
func (o *MetricsHistogramRequestQuery) GetByocCluster() string {
	if o == nil || o.ByocCluster == nil {
		var ret string
		return ret
	}
	return *o.ByocCluster
}

// GetByocClusterOk returns a tuple with the ByocCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsHistogramRequestQuery) GetByocClusterOk() (*string, bool) {
	if o == nil || o.ByocCluster == nil {
		return nil, false
	}
	return o.ByocCluster, true
}

// HasByocCluster returns a boolean if a field has been set.
func (o *MetricsHistogramRequestQuery) HasByocCluster() bool {
	return o != nil && o.ByocCluster != nil
}

// SetByocCluster gets a reference to the given string and assigns it to the ByocCluster field.
func (o *MetricsHistogramRequestQuery) SetByocCluster(v string) {
	o.ByocCluster = &v
}

// GetCrossOrgUuids returns the CrossOrgUuids field value if set, zero value otherwise.
func (o *MetricsHistogramRequestQuery) GetCrossOrgUuids() []string {
	if o == nil || o.CrossOrgUuids == nil {
		var ret []string
		return ret
	}
	return o.CrossOrgUuids
}

// GetCrossOrgUuidsOk returns a tuple with the CrossOrgUuids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsHistogramRequestQuery) GetCrossOrgUuidsOk() (*[]string, bool) {
	if o == nil || o.CrossOrgUuids == nil {
		return nil, false
	}
	return &o.CrossOrgUuids, true
}

// HasCrossOrgUuids returns a boolean if a field has been set.
func (o *MetricsHistogramRequestQuery) HasCrossOrgUuids() bool {
	return o != nil && o.CrossOrgUuids != nil
}

// SetCrossOrgUuids gets a reference to the given []string and assigns it to the CrossOrgUuids field.
func (o *MetricsHistogramRequestQuery) SetCrossOrgUuids(v []string) {
	o.CrossOrgUuids = v
}

// GetDataSource returns the DataSource field value.
func (o *MetricsHistogramRequestQuery) GetDataSource() FormulaMetricsDataSource {
	if o == nil {
		var ret FormulaMetricsDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *MetricsHistogramRequestQuery) GetDataSourceOk() (*FormulaMetricsDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *MetricsHistogramRequestQuery) SetDataSource(v FormulaMetricsDataSource) {
	o.DataSource = v
}

// GetHistogramMode returns the HistogramMode field value if set, zero value otherwise.
func (o *MetricsHistogramRequestQuery) GetHistogramMode() HistogramMode {
	if o == nil || o.HistogramMode == nil {
		var ret HistogramMode
		return ret
	}
	return *o.HistogramMode
}

// GetHistogramModeOk returns a tuple with the HistogramMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsHistogramRequestQuery) GetHistogramModeOk() (*HistogramMode, bool) {
	if o == nil || o.HistogramMode == nil {
		return nil, false
	}
	return o.HistogramMode, true
}

// HasHistogramMode returns a boolean if a field has been set.
func (o *MetricsHistogramRequestQuery) HasHistogramMode() bool {
	return o != nil && o.HistogramMode != nil
}

// SetHistogramMode gets a reference to the given HistogramMode and assigns it to the HistogramMode field.
func (o *MetricsHistogramRequestQuery) SetHistogramMode(v HistogramMode) {
	o.HistogramMode = &v
}

// GetName returns the Name field value.
func (o *MetricsHistogramRequestQuery) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MetricsHistogramRequestQuery) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *MetricsHistogramRequestQuery) SetName(v string) {
	o.Name = v
}

// GetQuery returns the Query field value.
func (o *MetricsHistogramRequestQuery) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *MetricsHistogramRequestQuery) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *MetricsHistogramRequestQuery) SetQuery(v string) {
	o.Query = v
}

// GetSemanticMode returns the SemanticMode field value if set, zero value otherwise.
func (o *MetricsHistogramRequestQuery) GetSemanticMode() FormulaSemanticMode {
	if o == nil || o.SemanticMode == nil {
		var ret FormulaSemanticMode
		return ret
	}
	return *o.SemanticMode
}

// GetSemanticModeOk returns a tuple with the SemanticMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsHistogramRequestQuery) GetSemanticModeOk() (*FormulaSemanticMode, bool) {
	if o == nil || o.SemanticMode == nil {
		return nil, false
	}
	return o.SemanticMode, true
}

// HasSemanticMode returns a boolean if a field has been set.
func (o *MetricsHistogramRequestQuery) HasSemanticMode() bool {
	return o != nil && o.SemanticMode != nil
}

// SetSemanticMode gets a reference to the given FormulaSemanticMode and assigns it to the SemanticMode field.
func (o *MetricsHistogramRequestQuery) SetSemanticMode(v FormulaSemanticMode) {
	o.SemanticMode = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MetricsHistogramRequestQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Aggregator != nil {
		toSerialize["aggregator"] = o.Aggregator
	}
	if o.ByocCluster != nil {
		toSerialize["byoc_cluster"] = o.ByocCluster
	}
	if o.CrossOrgUuids != nil {
		toSerialize["cross_org_uuids"] = o.CrossOrgUuids
	}
	toSerialize["data_source"] = o.DataSource
	if o.HistogramMode != nil {
		toSerialize["histogram_mode"] = o.HistogramMode
	}
	toSerialize["name"] = o.Name
	toSerialize["query"] = o.Query
	if o.SemanticMode != nil {
		toSerialize["semantic_mode"] = o.SemanticMode
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MetricsHistogramRequestQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregator    *ScalarAggregator         `json:"aggregator,omitempty"`
		ByocCluster   *string                   `json:"byoc_cluster,omitempty"`
		CrossOrgUuids []string                  `json:"cross_org_uuids,omitempty"`
		DataSource    *FormulaMetricsDataSource `json:"data_source"`
		HistogramMode *HistogramMode            `json:"histogram_mode,omitempty"`
		Name          *string                   `json:"name"`
		Query         *string                   `json:"query"`
		SemanticMode  *FormulaSemanticMode      `json:"semantic_mode,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field data_source missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}

	hasInvalidField := false
	if all.Aggregator != nil && !all.Aggregator.IsValid() {
		hasInvalidField = true
	} else {
		o.Aggregator = all.Aggregator
	}
	o.ByocCluster = all.ByocCluster
	o.CrossOrgUuids = all.CrossOrgUuids
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	if all.HistogramMode != nil && !all.HistogramMode.IsValid() {
		hasInvalidField = true
	} else {
		o.HistogramMode = all.HistogramMode
	}
	o.Name = *all.Name
	o.Query = *all.Query
	if all.SemanticMode != nil && !all.SemanticMode.IsValid() {
		hasInvalidField = true
	} else {
		o.SemanticMode = all.SemanticMode
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
