// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotebookTemplateVariableAvailableValuesQueryMetrics Available values query for the metrics data source.
type NotebookTemplateVariableAvailableValuesQueryMetrics struct {
	// The data source for the query. Must be `metrics`.
	DataSource string `json:"data_source"`
	// The metrics query string.
	Query string `json:"query"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewNotebookTemplateVariableAvailableValuesQueryMetrics instantiates a new NotebookTemplateVariableAvailableValuesQueryMetrics object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNotebookTemplateVariableAvailableValuesQueryMetrics(dataSource string, query string) *NotebookTemplateVariableAvailableValuesQueryMetrics {
	this := NotebookTemplateVariableAvailableValuesQueryMetrics{}
	this.DataSource = dataSource
	this.Query = query
	return &this
}

// NewNotebookTemplateVariableAvailableValuesQueryMetricsWithDefaults instantiates a new NotebookTemplateVariableAvailableValuesQueryMetrics object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNotebookTemplateVariableAvailableValuesQueryMetricsWithDefaults() *NotebookTemplateVariableAvailableValuesQueryMetrics {
	this := NotebookTemplateVariableAvailableValuesQueryMetrics{}
	return &this
}

// GetDataSource returns the DataSource field value.
func (o *NotebookTemplateVariableAvailableValuesQueryMetrics) GetDataSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *NotebookTemplateVariableAvailableValuesQueryMetrics) GetDataSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *NotebookTemplateVariableAvailableValuesQueryMetrics) SetDataSource(v string) {
	o.DataSource = v
}

// GetQuery returns the Query field value.
func (o *NotebookTemplateVariableAvailableValuesQueryMetrics) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *NotebookTemplateVariableAvailableValuesQueryMetrics) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *NotebookTemplateVariableAvailableValuesQueryMetrics) SetQuery(v string) {
	o.Query = v
}

// MarshalJSON serializes the struct using spec logic.
func (o NotebookTemplateVariableAvailableValuesQueryMetrics) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data_source"] = o.DataSource
	toSerialize["query"] = o.Query
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NotebookTemplateVariableAvailableValuesQueryMetrics) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataSource *string `json:"data_source"`
		Query      *string `json:"query"`
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
	o.DataSource = *all.DataSource
	o.Query = *all.Query

	return nil
}
