// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaMetricsQueryDatadogUsage Test-only generated schema for the dashboard versioning proof of concept.
type FormulaMetricsQueryDatadogUsage struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Aggregator *ScalarAggregator `json:"aggregator,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ByocCluster *string `json:"byoc_cluster,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	CrossOrgUuids []string `json:"cross_org_uuids,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource FormulaMetricsQueryDatadogUsageDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Query string `json:"query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	SemanticMode *FormulaSemanticMode `json:"semantic_mode,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewFormulaMetricsQueryDatadogUsage instantiates a new FormulaMetricsQueryDatadogUsage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormulaMetricsQueryDatadogUsage(dataSource FormulaMetricsQueryDatadogUsageDataSource, name string, query string) *FormulaMetricsQueryDatadogUsage {
	this := FormulaMetricsQueryDatadogUsage{}
	this.DataSource = dataSource
	this.Name = name
	this.Query = query
	return &this
}

// NewFormulaMetricsQueryDatadogUsageWithDefaults instantiates a new FormulaMetricsQueryDatadogUsage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormulaMetricsQueryDatadogUsageWithDefaults() *FormulaMetricsQueryDatadogUsage {
	this := FormulaMetricsQueryDatadogUsage{}
	return &this
}

// GetAggregator returns the Aggregator field value if set, zero value otherwise.
func (o *FormulaMetricsQueryDatadogUsage) GetAggregator() ScalarAggregator {
	if o == nil || o.Aggregator == nil {
		var ret ScalarAggregator
		return ret
	}
	return *o.Aggregator
}

// GetAggregatorOk returns a tuple with the Aggregator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaMetricsQueryDatadogUsage) GetAggregatorOk() (*ScalarAggregator, bool) {
	if o == nil || o.Aggregator == nil {
		return nil, false
	}
	return o.Aggregator, true
}

// HasAggregator returns a boolean if a field has been set.
func (o *FormulaMetricsQueryDatadogUsage) HasAggregator() bool {
	return o != nil && o.Aggregator != nil
}

// SetAggregator gets a reference to the given ScalarAggregator and assigns it to the Aggregator field.
func (o *FormulaMetricsQueryDatadogUsage) SetAggregator(v ScalarAggregator) {
	o.Aggregator = &v
}

// GetByocCluster returns the ByocCluster field value if set, zero value otherwise.
func (o *FormulaMetricsQueryDatadogUsage) GetByocCluster() string {
	if o == nil || o.ByocCluster == nil {
		var ret string
		return ret
	}
	return *o.ByocCluster
}

// GetByocClusterOk returns a tuple with the ByocCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaMetricsQueryDatadogUsage) GetByocClusterOk() (*string, bool) {
	if o == nil || o.ByocCluster == nil {
		return nil, false
	}
	return o.ByocCluster, true
}

// HasByocCluster returns a boolean if a field has been set.
func (o *FormulaMetricsQueryDatadogUsage) HasByocCluster() bool {
	return o != nil && o.ByocCluster != nil
}

// SetByocCluster gets a reference to the given string and assigns it to the ByocCluster field.
func (o *FormulaMetricsQueryDatadogUsage) SetByocCluster(v string) {
	o.ByocCluster = &v
}

// GetCrossOrgUuids returns the CrossOrgUuids field value if set, zero value otherwise.
func (o *FormulaMetricsQueryDatadogUsage) GetCrossOrgUuids() []string {
	if o == nil || o.CrossOrgUuids == nil {
		var ret []string
		return ret
	}
	return o.CrossOrgUuids
}

// GetCrossOrgUuidsOk returns a tuple with the CrossOrgUuids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaMetricsQueryDatadogUsage) GetCrossOrgUuidsOk() (*[]string, bool) {
	if o == nil || o.CrossOrgUuids == nil {
		return nil, false
	}
	return &o.CrossOrgUuids, true
}

// HasCrossOrgUuids returns a boolean if a field has been set.
func (o *FormulaMetricsQueryDatadogUsage) HasCrossOrgUuids() bool {
	return o != nil && o.CrossOrgUuids != nil
}

// SetCrossOrgUuids gets a reference to the given []string and assigns it to the CrossOrgUuids field.
func (o *FormulaMetricsQueryDatadogUsage) SetCrossOrgUuids(v []string) {
	o.CrossOrgUuids = v
}

// GetDataSource returns the DataSource field value.
func (o *FormulaMetricsQueryDatadogUsage) GetDataSource() FormulaMetricsQueryDatadogUsageDataSource {
	if o == nil {
		var ret FormulaMetricsQueryDatadogUsageDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *FormulaMetricsQueryDatadogUsage) GetDataSourceOk() (*FormulaMetricsQueryDatadogUsageDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *FormulaMetricsQueryDatadogUsage) SetDataSource(v FormulaMetricsQueryDatadogUsageDataSource) {
	o.DataSource = v
}

// GetName returns the Name field value.
func (o *FormulaMetricsQueryDatadogUsage) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FormulaMetricsQueryDatadogUsage) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *FormulaMetricsQueryDatadogUsage) SetName(v string) {
	o.Name = v
}

// GetQuery returns the Query field value.
func (o *FormulaMetricsQueryDatadogUsage) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *FormulaMetricsQueryDatadogUsage) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *FormulaMetricsQueryDatadogUsage) SetQuery(v string) {
	o.Query = v
}

// GetSemanticMode returns the SemanticMode field value if set, zero value otherwise.
func (o *FormulaMetricsQueryDatadogUsage) GetSemanticMode() FormulaSemanticMode {
	if o == nil || o.SemanticMode == nil {
		var ret FormulaSemanticMode
		return ret
	}
	return *o.SemanticMode
}

// GetSemanticModeOk returns a tuple with the SemanticMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaMetricsQueryDatadogUsage) GetSemanticModeOk() (*FormulaSemanticMode, bool) {
	if o == nil || o.SemanticMode == nil {
		return nil, false
	}
	return o.SemanticMode, true
}

// HasSemanticMode returns a boolean if a field has been set.
func (o *FormulaMetricsQueryDatadogUsage) HasSemanticMode() bool {
	return o != nil && o.SemanticMode != nil
}

// SetSemanticMode gets a reference to the given FormulaSemanticMode and assigns it to the SemanticMode field.
func (o *FormulaMetricsQueryDatadogUsage) SetSemanticMode(v FormulaSemanticMode) {
	o.SemanticMode = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormulaMetricsQueryDatadogUsage) MarshalJSON() ([]byte, error) {
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
	toSerialize["name"] = o.Name
	toSerialize["query"] = o.Query
	if o.SemanticMode != nil {
		toSerialize["semantic_mode"] = o.SemanticMode
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormulaMetricsQueryDatadogUsage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregator    *ScalarAggregator                          `json:"aggregator,omitempty"`
		ByocCluster   *string                                    `json:"byoc_cluster,omitempty"`
		CrossOrgUuids []string                                   `json:"cross_org_uuids,omitempty"`
		DataSource    *FormulaMetricsQueryDatadogUsageDataSource `json:"data_source"`
		Name          *string                                    `json:"name"`
		Query         *string                                    `json:"query"`
		SemanticMode  *FormulaSemanticMode                       `json:"semantic_mode,omitempty"`
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
