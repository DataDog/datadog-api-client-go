// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetricsLikeQueryFragmentMetrics Test-only generated schema for the dashboard versioning proof of concept.
type MetricsLikeQueryFragmentMetrics struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Alias *string `json:"alias,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource MetricsLikeQueryFragmentMetricsDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Filter *string `json:"filter,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	MetricName string `json:"metric_name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewMetricsLikeQueryFragmentMetrics instantiates a new MetricsLikeQueryFragmentMetrics object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMetricsLikeQueryFragmentMetrics(dataSource MetricsLikeQueryFragmentMetricsDataSource, metricName string, name string) *MetricsLikeQueryFragmentMetrics {
	this := MetricsLikeQueryFragmentMetrics{}
	this.DataSource = dataSource
	this.MetricName = metricName
	this.Name = name
	return &this
}

// NewMetricsLikeQueryFragmentMetricsWithDefaults instantiates a new MetricsLikeQueryFragmentMetrics object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMetricsLikeQueryFragmentMetricsWithDefaults() *MetricsLikeQueryFragmentMetrics {
	this := MetricsLikeQueryFragmentMetrics{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *MetricsLikeQueryFragmentMetrics) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsLikeQueryFragmentMetrics) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *MetricsLikeQueryFragmentMetrics) HasAlias() bool {
	return o != nil && o.Alias != nil
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *MetricsLikeQueryFragmentMetrics) SetAlias(v string) {
	o.Alias = &v
}

// GetDataSource returns the DataSource field value.
func (o *MetricsLikeQueryFragmentMetrics) GetDataSource() MetricsLikeQueryFragmentMetricsDataSource {
	if o == nil {
		var ret MetricsLikeQueryFragmentMetricsDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *MetricsLikeQueryFragmentMetrics) GetDataSourceOk() (*MetricsLikeQueryFragmentMetricsDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *MetricsLikeQueryFragmentMetrics) SetDataSource(v MetricsLikeQueryFragmentMetricsDataSource) {
	o.DataSource = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *MetricsLikeQueryFragmentMetrics) GetFilter() string {
	if o == nil || o.Filter == nil {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsLikeQueryFragmentMetrics) GetFilterOk() (*string, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *MetricsLikeQueryFragmentMetrics) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *MetricsLikeQueryFragmentMetrics) SetFilter(v string) {
	o.Filter = &v
}

// GetMetricName returns the MetricName field value.
func (o *MetricsLikeQueryFragmentMetrics) GetMetricName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.MetricName
}

// GetMetricNameOk returns a tuple with the MetricName field value
// and a boolean to check if the value has been set.
func (o *MetricsLikeQueryFragmentMetrics) GetMetricNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricName, true
}

// SetMetricName sets field value.
func (o *MetricsLikeQueryFragmentMetrics) SetMetricName(v string) {
	o.MetricName = v
}

// GetName returns the Name field value.
func (o *MetricsLikeQueryFragmentMetrics) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MetricsLikeQueryFragmentMetrics) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *MetricsLikeQueryFragmentMetrics) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MetricsLikeQueryFragmentMetrics) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Alias != nil {
		toSerialize["alias"] = o.Alias
	}
	toSerialize["data_source"] = o.DataSource
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	toSerialize["metric_name"] = o.MetricName
	toSerialize["name"] = o.Name
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MetricsLikeQueryFragmentMetrics) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Alias      *string                                    `json:"alias,omitempty"`
		DataSource *MetricsLikeQueryFragmentMetricsDataSource `json:"data_source"`
		Filter     *string                                    `json:"filter,omitempty"`
		MetricName *string                                    `json:"metric_name"`
		Name       *string                                    `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
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
	o.Alias = all.Alias
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.Filter = all.Filter
	o.MetricName = *all.MetricName
	o.Name = *all.Name

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
