// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetricDatasetMetricScalar Test-only generated schema for the dashboard versioning proof of concept.
type MetricDatasetMetricScalar struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Aggregation MetricDatasetMetricScalarAggregation `json:"aggregation"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource MetricDatasetMetricScalarDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	MetricName string `json:"metric_name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	QueryFilter *string `json:"query_filter,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TagNames []string `json:"tag_names,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TimeWindow *TimeWindow `json:"time_window,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewMetricDatasetMetricScalar instantiates a new MetricDatasetMetricScalar object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMetricDatasetMetricScalar(aggregation MetricDatasetMetricScalarAggregation, dataSource MetricDatasetMetricScalarDataSource, metricName string, name string) *MetricDatasetMetricScalar {
	this := MetricDatasetMetricScalar{}
	this.Aggregation = aggregation
	this.DataSource = dataSource
	this.MetricName = metricName
	this.Name = name
	return &this
}

// NewMetricDatasetMetricScalarWithDefaults instantiates a new MetricDatasetMetricScalar object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMetricDatasetMetricScalarWithDefaults() *MetricDatasetMetricScalar {
	this := MetricDatasetMetricScalar{}
	return &this
}

// GetAggregation returns the Aggregation field value.
func (o *MetricDatasetMetricScalar) GetAggregation() MetricDatasetMetricScalarAggregation {
	if o == nil {
		var ret MetricDatasetMetricScalarAggregation
		return ret
	}
	return o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value
// and a boolean to check if the value has been set.
func (o *MetricDatasetMetricScalar) GetAggregationOk() (*MetricDatasetMetricScalarAggregation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aggregation, true
}

// SetAggregation sets field value.
func (o *MetricDatasetMetricScalar) SetAggregation(v MetricDatasetMetricScalarAggregation) {
	o.Aggregation = v
}

// GetDataSource returns the DataSource field value.
func (o *MetricDatasetMetricScalar) GetDataSource() MetricDatasetMetricScalarDataSource {
	if o == nil {
		var ret MetricDatasetMetricScalarDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *MetricDatasetMetricScalar) GetDataSourceOk() (*MetricDatasetMetricScalarDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *MetricDatasetMetricScalar) SetDataSource(v MetricDatasetMetricScalarDataSource) {
	o.DataSource = v
}

// GetMetricName returns the MetricName field value.
func (o *MetricDatasetMetricScalar) GetMetricName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.MetricName
}

// GetMetricNameOk returns a tuple with the MetricName field value
// and a boolean to check if the value has been set.
func (o *MetricDatasetMetricScalar) GetMetricNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricName, true
}

// SetMetricName sets field value.
func (o *MetricDatasetMetricScalar) SetMetricName(v string) {
	o.MetricName = v
}

// GetName returns the Name field value.
func (o *MetricDatasetMetricScalar) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MetricDatasetMetricScalar) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *MetricDatasetMetricScalar) SetName(v string) {
	o.Name = v
}

// GetQueryFilter returns the QueryFilter field value if set, zero value otherwise.
func (o *MetricDatasetMetricScalar) GetQueryFilter() string {
	if o == nil || o.QueryFilter == nil {
		var ret string
		return ret
	}
	return *o.QueryFilter
}

// GetQueryFilterOk returns a tuple with the QueryFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDatasetMetricScalar) GetQueryFilterOk() (*string, bool) {
	if o == nil || o.QueryFilter == nil {
		return nil, false
	}
	return o.QueryFilter, true
}

// HasQueryFilter returns a boolean if a field has been set.
func (o *MetricDatasetMetricScalar) HasQueryFilter() bool {
	return o != nil && o.QueryFilter != nil
}

// SetQueryFilter gets a reference to the given string and assigns it to the QueryFilter field.
func (o *MetricDatasetMetricScalar) SetQueryFilter(v string) {
	o.QueryFilter = &v
}

// GetTagNames returns the TagNames field value if set, zero value otherwise.
func (o *MetricDatasetMetricScalar) GetTagNames() []string {
	if o == nil || o.TagNames == nil {
		var ret []string
		return ret
	}
	return o.TagNames
}

// GetTagNamesOk returns a tuple with the TagNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDatasetMetricScalar) GetTagNamesOk() (*[]string, bool) {
	if o == nil || o.TagNames == nil {
		return nil, false
	}
	return &o.TagNames, true
}

// HasTagNames returns a boolean if a field has been set.
func (o *MetricDatasetMetricScalar) HasTagNames() bool {
	return o != nil && o.TagNames != nil
}

// SetTagNames gets a reference to the given []string and assigns it to the TagNames field.
func (o *MetricDatasetMetricScalar) SetTagNames(v []string) {
	o.TagNames = v
}

// GetTimeWindow returns the TimeWindow field value if set, zero value otherwise.
func (o *MetricDatasetMetricScalar) GetTimeWindow() TimeWindow {
	if o == nil || o.TimeWindow == nil {
		var ret TimeWindow
		return ret
	}
	return *o.TimeWindow
}

// GetTimeWindowOk returns a tuple with the TimeWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDatasetMetricScalar) GetTimeWindowOk() (*TimeWindow, bool) {
	if o == nil || o.TimeWindow == nil {
		return nil, false
	}
	return o.TimeWindow, true
}

// HasTimeWindow returns a boolean if a field has been set.
func (o *MetricDatasetMetricScalar) HasTimeWindow() bool {
	return o != nil && o.TimeWindow != nil
}

// SetTimeWindow gets a reference to the given TimeWindow and assigns it to the TimeWindow field.
func (o *MetricDatasetMetricScalar) SetTimeWindow(v TimeWindow) {
	o.TimeWindow = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MetricDatasetMetricScalar) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["aggregation"] = o.Aggregation
	toSerialize["data_source"] = o.DataSource
	toSerialize["metric_name"] = o.MetricName
	toSerialize["name"] = o.Name
	if o.QueryFilter != nil {
		toSerialize["query_filter"] = o.QueryFilter
	}
	if o.TagNames != nil {
		toSerialize["tag_names"] = o.TagNames
	}
	if o.TimeWindow != nil {
		toSerialize["time_window"] = o.TimeWindow
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MetricDatasetMetricScalar) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregation *MetricDatasetMetricScalarAggregation `json:"aggregation"`
		DataSource  *MetricDatasetMetricScalarDataSource  `json:"data_source"`
		MetricName  *string                               `json:"metric_name"`
		Name        *string                               `json:"name"`
		QueryFilter *string                               `json:"query_filter,omitempty"`
		TagNames    []string                              `json:"tag_names,omitempty"`
		TimeWindow  *TimeWindow                           `json:"time_window,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Aggregation == nil {
		return fmt.Errorf("required field aggregation missing")
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field data_source missing")
	}
	if all.MetricName == nil {
		return fmt.Errorf("required field metric_name missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}

	hasInvalidField := false
	if !all.Aggregation.IsValid() {
		hasInvalidField = true
	} else {
		o.Aggregation = *all.Aggregation
	}
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.MetricName = *all.MetricName
	o.Name = *all.Name
	o.QueryFilter = all.QueryFilter
	o.TagNames = all.TagNames
	if all.TimeWindow != nil && all.TimeWindow.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TimeWindow = all.TimeWindow

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
