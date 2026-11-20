// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetricDatasetCloudCostScalar Test-only generated schema for the dashboard versioning proof of concept.
type MetricDatasetCloudCostScalar struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Aggregation MetricDatasetCloudCostScalarAggregation `json:"aggregation"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource MetricDatasetCloudCostScalarDataSource `json:"data_source"`
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

// NewMetricDatasetCloudCostScalar instantiates a new MetricDatasetCloudCostScalar object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMetricDatasetCloudCostScalar(aggregation MetricDatasetCloudCostScalarAggregation, dataSource MetricDatasetCloudCostScalarDataSource, metricName string, name string) *MetricDatasetCloudCostScalar {
	this := MetricDatasetCloudCostScalar{}
	this.Aggregation = aggregation
	this.DataSource = dataSource
	this.MetricName = metricName
	this.Name = name
	return &this
}

// NewMetricDatasetCloudCostScalarWithDefaults instantiates a new MetricDatasetCloudCostScalar object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMetricDatasetCloudCostScalarWithDefaults() *MetricDatasetCloudCostScalar {
	this := MetricDatasetCloudCostScalar{}
	return &this
}

// GetAggregation returns the Aggregation field value.
func (o *MetricDatasetCloudCostScalar) GetAggregation() MetricDatasetCloudCostScalarAggregation {
	if o == nil {
		var ret MetricDatasetCloudCostScalarAggregation
		return ret
	}
	return o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value
// and a boolean to check if the value has been set.
func (o *MetricDatasetCloudCostScalar) GetAggregationOk() (*MetricDatasetCloudCostScalarAggregation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aggregation, true
}

// SetAggregation sets field value.
func (o *MetricDatasetCloudCostScalar) SetAggregation(v MetricDatasetCloudCostScalarAggregation) {
	o.Aggregation = v
}

// GetDataSource returns the DataSource field value.
func (o *MetricDatasetCloudCostScalar) GetDataSource() MetricDatasetCloudCostScalarDataSource {
	if o == nil {
		var ret MetricDatasetCloudCostScalarDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *MetricDatasetCloudCostScalar) GetDataSourceOk() (*MetricDatasetCloudCostScalarDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *MetricDatasetCloudCostScalar) SetDataSource(v MetricDatasetCloudCostScalarDataSource) {
	o.DataSource = v
}

// GetMetricName returns the MetricName field value.
func (o *MetricDatasetCloudCostScalar) GetMetricName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.MetricName
}

// GetMetricNameOk returns a tuple with the MetricName field value
// and a boolean to check if the value has been set.
func (o *MetricDatasetCloudCostScalar) GetMetricNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricName, true
}

// SetMetricName sets field value.
func (o *MetricDatasetCloudCostScalar) SetMetricName(v string) {
	o.MetricName = v
}

// GetName returns the Name field value.
func (o *MetricDatasetCloudCostScalar) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MetricDatasetCloudCostScalar) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *MetricDatasetCloudCostScalar) SetName(v string) {
	o.Name = v
}

// GetQueryFilter returns the QueryFilter field value if set, zero value otherwise.
func (o *MetricDatasetCloudCostScalar) GetQueryFilter() string {
	if o == nil || o.QueryFilter == nil {
		var ret string
		return ret
	}
	return *o.QueryFilter
}

// GetQueryFilterOk returns a tuple with the QueryFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDatasetCloudCostScalar) GetQueryFilterOk() (*string, bool) {
	if o == nil || o.QueryFilter == nil {
		return nil, false
	}
	return o.QueryFilter, true
}

// HasQueryFilter returns a boolean if a field has been set.
func (o *MetricDatasetCloudCostScalar) HasQueryFilter() bool {
	return o != nil && o.QueryFilter != nil
}

// SetQueryFilter gets a reference to the given string and assigns it to the QueryFilter field.
func (o *MetricDatasetCloudCostScalar) SetQueryFilter(v string) {
	o.QueryFilter = &v
}

// GetTagNames returns the TagNames field value if set, zero value otherwise.
func (o *MetricDatasetCloudCostScalar) GetTagNames() []string {
	if o == nil || o.TagNames == nil {
		var ret []string
		return ret
	}
	return o.TagNames
}

// GetTagNamesOk returns a tuple with the TagNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDatasetCloudCostScalar) GetTagNamesOk() (*[]string, bool) {
	if o == nil || o.TagNames == nil {
		return nil, false
	}
	return &o.TagNames, true
}

// HasTagNames returns a boolean if a field has been set.
func (o *MetricDatasetCloudCostScalar) HasTagNames() bool {
	return o != nil && o.TagNames != nil
}

// SetTagNames gets a reference to the given []string and assigns it to the TagNames field.
func (o *MetricDatasetCloudCostScalar) SetTagNames(v []string) {
	o.TagNames = v
}

// GetTimeWindow returns the TimeWindow field value if set, zero value otherwise.
func (o *MetricDatasetCloudCostScalar) GetTimeWindow() TimeWindow {
	if o == nil || o.TimeWindow == nil {
		var ret TimeWindow
		return ret
	}
	return *o.TimeWindow
}

// GetTimeWindowOk returns a tuple with the TimeWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDatasetCloudCostScalar) GetTimeWindowOk() (*TimeWindow, bool) {
	if o == nil || o.TimeWindow == nil {
		return nil, false
	}
	return o.TimeWindow, true
}

// HasTimeWindow returns a boolean if a field has been set.
func (o *MetricDatasetCloudCostScalar) HasTimeWindow() bool {
	return o != nil && o.TimeWindow != nil
}

// SetTimeWindow gets a reference to the given TimeWindow and assigns it to the TimeWindow field.
func (o *MetricDatasetCloudCostScalar) SetTimeWindow(v TimeWindow) {
	o.TimeWindow = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MetricDatasetCloudCostScalar) MarshalJSON() ([]byte, error) {
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
func (o *MetricDatasetCloudCostScalar) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregation *MetricDatasetCloudCostScalarAggregation `json:"aggregation"`
		DataSource  *MetricDatasetCloudCostScalarDataSource  `json:"data_source"`
		MetricName  *string                                  `json:"metric_name"`
		Name        *string                                  `json:"name"`
		QueryFilter *string                                  `json:"query_filter,omitempty"`
		TagNames    []string                                 `json:"tag_names,omitempty"`
		TimeWindow  *TimeWindow                              `json:"time_window,omitempty"`
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
