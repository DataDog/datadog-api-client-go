// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GuidedTableMetricsQuery A metrics (or cloud cost) query fragment used as source data for a guided table. Group-by and time/space aggregation are defined separately.
type GuidedTableMetricsQuery struct {
	// Display alias for the query.
	Alias *string `json:"alias,omitempty"`
	// Metrics data source.
	DataSource GuidedTableMetricsQueryDataSource `json:"data_source"`
	// Filter string to apply to the metric query.
	Filter *string `json:"filter,omitempty"`
	// Name of the metric to query.
	MetricName string `json:"metric_name"`
	// Variable name used to reference this query in columns and formulas.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGuidedTableMetricsQuery instantiates a new GuidedTableMetricsQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGuidedTableMetricsQuery(dataSource GuidedTableMetricsQueryDataSource, metricName string, name string) *GuidedTableMetricsQuery {
	this := GuidedTableMetricsQuery{}
	this.DataSource = dataSource
	this.MetricName = metricName
	this.Name = name
	return &this
}

// NewGuidedTableMetricsQueryWithDefaults instantiates a new GuidedTableMetricsQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGuidedTableMetricsQueryWithDefaults() *GuidedTableMetricsQuery {
	this := GuidedTableMetricsQuery{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *GuidedTableMetricsQuery) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidedTableMetricsQuery) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *GuidedTableMetricsQuery) HasAlias() bool {
	return o != nil && o.Alias != nil
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *GuidedTableMetricsQuery) SetAlias(v string) {
	o.Alias = &v
}

// GetDataSource returns the DataSource field value.
func (o *GuidedTableMetricsQuery) GetDataSource() GuidedTableMetricsQueryDataSource {
	if o == nil {
		var ret GuidedTableMetricsQueryDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *GuidedTableMetricsQuery) GetDataSourceOk() (*GuidedTableMetricsQueryDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *GuidedTableMetricsQuery) SetDataSource(v GuidedTableMetricsQueryDataSource) {
	o.DataSource = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *GuidedTableMetricsQuery) GetFilter() string {
	if o == nil || o.Filter == nil {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidedTableMetricsQuery) GetFilterOk() (*string, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *GuidedTableMetricsQuery) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *GuidedTableMetricsQuery) SetFilter(v string) {
	o.Filter = &v
}

// GetMetricName returns the MetricName field value.
func (o *GuidedTableMetricsQuery) GetMetricName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.MetricName
}

// GetMetricNameOk returns a tuple with the MetricName field value
// and a boolean to check if the value has been set.
func (o *GuidedTableMetricsQuery) GetMetricNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricName, true
}

// SetMetricName sets field value.
func (o *GuidedTableMetricsQuery) SetMetricName(v string) {
	o.MetricName = v
}

// GetName returns the Name field value.
func (o *GuidedTableMetricsQuery) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GuidedTableMetricsQuery) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *GuidedTableMetricsQuery) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GuidedTableMetricsQuery) MarshalJSON() ([]byte, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GuidedTableMetricsQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Alias      *string                            `json:"alias,omitempty"`
		DataSource *GuidedTableMetricsQueryDataSource `json:"data_source"`
		Filter     *string                            `json:"filter,omitempty"`
		MetricName *string                            `json:"metric_name"`
		Name       *string                            `json:"name"`
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
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"alias", "data_source", "filter", "metric_name", "name"})
	} else {
		return err
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

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
