// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardAvailableValuesMetricsQuery Query for available values using the metrics data source.
type DashboardAvailableValuesMetricsQuery struct {
	// Data source for metrics queries.
	DataSource FormulaAndFunctionMetricDataSource `json:"data_source"`
	// The metrics query string.
	Query string `json:"query"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewDashboardAvailableValuesMetricsQuery instantiates a new DashboardAvailableValuesMetricsQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDashboardAvailableValuesMetricsQuery(dataSource FormulaAndFunctionMetricDataSource, query string) *DashboardAvailableValuesMetricsQuery {
	this := DashboardAvailableValuesMetricsQuery{}
	this.DataSource = dataSource
	this.Query = query
	return &this
}

// NewDashboardAvailableValuesMetricsQueryWithDefaults instantiates a new DashboardAvailableValuesMetricsQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDashboardAvailableValuesMetricsQueryWithDefaults() *DashboardAvailableValuesMetricsQuery {
	this := DashboardAvailableValuesMetricsQuery{}
	return &this
}

// GetDataSource returns the DataSource field value.
func (o *DashboardAvailableValuesMetricsQuery) GetDataSource() FormulaAndFunctionMetricDataSource {
	if o == nil {
		var ret FormulaAndFunctionMetricDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *DashboardAvailableValuesMetricsQuery) GetDataSourceOk() (*FormulaAndFunctionMetricDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *DashboardAvailableValuesMetricsQuery) SetDataSource(v FormulaAndFunctionMetricDataSource) {
	o.DataSource = v
}

// GetQuery returns the Query field value.
func (o *DashboardAvailableValuesMetricsQuery) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *DashboardAvailableValuesMetricsQuery) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *DashboardAvailableValuesMetricsQuery) SetQuery(v string) {
	o.Query = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DashboardAvailableValuesMetricsQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data_source"] = o.DataSource
	toSerialize["query"] = o.Query
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DashboardAvailableValuesMetricsQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataSource *FormulaAndFunctionMetricDataSource `json:"data_source"`
		Query      *string                             `json:"query"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field data_source missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}

	hasInvalidField := false
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.Query = *all.Query

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
