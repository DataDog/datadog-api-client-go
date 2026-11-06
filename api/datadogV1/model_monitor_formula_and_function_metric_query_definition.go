// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorFormulaAndFunctionMetricQueryDefinition A formula and functions metrics query. The `aggregator` field is required for `cost alert` monitors and not allowed for metric or query alert monitors.
type MonitorFormulaAndFunctionMetricQueryDefinition struct {
	// Aggregation methods for metric queries.
	Aggregator *MonitorFormulaAndFunctionCostAggregator `json:"aggregator,omitempty"`
	// Data source for metrics queries.
	DataSource MonitorFormulaAndFunctionMetricQueryDataSource `json:"data_source"`
	// Name of the query for use in formulas.
	Name string `json:"name"`
	// The metrics query definition.
	Query string `json:"query"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMonitorFormulaAndFunctionMetricQueryDefinition instantiates a new MonitorFormulaAndFunctionMetricQueryDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorFormulaAndFunctionMetricQueryDefinition(dataSource MonitorFormulaAndFunctionMetricQueryDataSource, name string, query string) *MonitorFormulaAndFunctionMetricQueryDefinition {
	this := MonitorFormulaAndFunctionMetricQueryDefinition{}
	this.DataSource = dataSource
	this.Name = name
	this.Query = query
	return &this
}

// NewMonitorFormulaAndFunctionMetricQueryDefinitionWithDefaults instantiates a new MonitorFormulaAndFunctionMetricQueryDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorFormulaAndFunctionMetricQueryDefinitionWithDefaults() *MonitorFormulaAndFunctionMetricQueryDefinition {
	this := MonitorFormulaAndFunctionMetricQueryDefinition{}
	return &this
}

// GetAggregator returns the Aggregator field value if set, zero value otherwise.
func (o *MonitorFormulaAndFunctionMetricQueryDefinition) GetAggregator() MonitorFormulaAndFunctionCostAggregator {
	if o == nil || o.Aggregator == nil {
		var ret MonitorFormulaAndFunctionCostAggregator
		return ret
	}
	return *o.Aggregator
}

// GetAggregatorOk returns a tuple with the Aggregator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionMetricQueryDefinition) GetAggregatorOk() (*MonitorFormulaAndFunctionCostAggregator, bool) {
	if o == nil || o.Aggregator == nil {
		return nil, false
	}
	return o.Aggregator, true
}

// HasAggregator returns a boolean if a field has been set.
func (o *MonitorFormulaAndFunctionMetricQueryDefinition) HasAggregator() bool {
	return o != nil && o.Aggregator != nil
}

// SetAggregator gets a reference to the given MonitorFormulaAndFunctionCostAggregator and assigns it to the Aggregator field.
func (o *MonitorFormulaAndFunctionMetricQueryDefinition) SetAggregator(v MonitorFormulaAndFunctionCostAggregator) {
	o.Aggregator = &v
}

// GetDataSource returns the DataSource field value.
func (o *MonitorFormulaAndFunctionMetricQueryDefinition) GetDataSource() MonitorFormulaAndFunctionMetricQueryDataSource {
	if o == nil {
		var ret MonitorFormulaAndFunctionMetricQueryDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionMetricQueryDefinition) GetDataSourceOk() (*MonitorFormulaAndFunctionMetricQueryDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *MonitorFormulaAndFunctionMetricQueryDefinition) SetDataSource(v MonitorFormulaAndFunctionMetricQueryDataSource) {
	o.DataSource = v
}

// GetName returns the Name field value.
func (o *MonitorFormulaAndFunctionMetricQueryDefinition) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionMetricQueryDefinition) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *MonitorFormulaAndFunctionMetricQueryDefinition) SetName(v string) {
	o.Name = v
}

// GetQuery returns the Query field value.
func (o *MonitorFormulaAndFunctionMetricQueryDefinition) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionMetricQueryDefinition) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *MonitorFormulaAndFunctionMetricQueryDefinition) SetQuery(v string) {
	o.Query = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonitorFormulaAndFunctionMetricQueryDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Aggregator != nil {
		toSerialize["aggregator"] = o.Aggregator
	}
	toSerialize["data_source"] = o.DataSource
	toSerialize["name"] = o.Name
	toSerialize["query"] = o.Query

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonitorFormulaAndFunctionMetricQueryDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregator *MonitorFormulaAndFunctionCostAggregator        `json:"aggregator,omitempty"`
		DataSource *MonitorFormulaAndFunctionMetricQueryDataSource `json:"data_source"`
		Name       *string                                         `json:"name"`
		Query      *string                                         `json:"query"`
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
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"aggregator", "data_source", "name", "query"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Aggregator != nil && !all.Aggregator.IsValid() {
		hasInvalidField = true
	} else {
		o.Aggregator = all.Aggregator
	}
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.Name = *all.Name
	o.Query = *all.Query

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
