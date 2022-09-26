// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"
)

// SecurityMonitoringStandardRuleQueryCreate Query for matching rule.
type SecurityMonitoringStandardRuleQueryCreate struct {
	// The aggregation type.
	Aggregation *SecurityMonitoringRuleQueryAggregation `json:"aggregation,omitempty"`
	// Field for which the cardinality is measured. Sent as an array.
	DistinctFields []string `json:"distinctFields,omitempty"`
	// Fields to group by.
	GroupByFields []string `json:"groupByFields,omitempty"`
	// The target field to aggregate over when using the sum or max
	// aggregations.
	Metric *string `json:"metric,omitempty"`
	// Group of target fields to aggregate over when using the new value aggregations.
	Metrics []string `json:"metrics,omitempty"`
	// Name of the query.
	Name *string `json:"name,omitempty"`
	// Query to run on logs.
	Query string `json:"query"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}

// NewSecurityMonitoringStandardRuleQueryCreate instantiates a new SecurityMonitoringStandardRuleQueryCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringStandardRuleQueryCreate(query string) *SecurityMonitoringStandardRuleQueryCreate {
	this := SecurityMonitoringStandardRuleQueryCreate{}
	this.Query = query
	return &this
}

// NewSecurityMonitoringStandardRuleQueryCreateWithDefaults instantiates a new SecurityMonitoringStandardRuleQueryCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringStandardRuleQueryCreateWithDefaults() *SecurityMonitoringStandardRuleQueryCreate {
	this := SecurityMonitoringStandardRuleQueryCreate{}
	return &this
}

// GetAggregation returns the Aggregation field value if set, zero value otherwise.
func (o *SecurityMonitoringStandardRuleQueryCreate) GetAggregation() SecurityMonitoringRuleQueryAggregation {
	if o == nil || o.Aggregation == nil {
		var ret SecurityMonitoringRuleQueryAggregation
		return ret
	}
	return *o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringStandardRuleQueryCreate) GetAggregationOk() (*SecurityMonitoringRuleQueryAggregation, bool) {
	if o == nil || o.Aggregation == nil {
		return nil, false
	}
	return o.Aggregation, true
}

// HasAggregation returns a boolean if a field has been set.
func (o *SecurityMonitoringStandardRuleQueryCreate) HasAggregation() bool {
	if o != nil && o.Aggregation != nil {
		return true
	}

	return false
}

// SetAggregation gets a reference to the given SecurityMonitoringRuleQueryAggregation and assigns it to the Aggregation field.
func (o *SecurityMonitoringStandardRuleQueryCreate) SetAggregation(v SecurityMonitoringRuleQueryAggregation) {
	o.Aggregation = &v
}

// GetDistinctFields returns the DistinctFields field value if set, zero value otherwise.
func (o *SecurityMonitoringStandardRuleQueryCreate) GetDistinctFields() []string {
	if o == nil || o.DistinctFields == nil {
		var ret []string
		return ret
	}
	return o.DistinctFields
}

// GetDistinctFieldsOk returns a tuple with the DistinctFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringStandardRuleQueryCreate) GetDistinctFieldsOk() (*[]string, bool) {
	if o == nil || o.DistinctFields == nil {
		return nil, false
	}
	return &o.DistinctFields, true
}

// HasDistinctFields returns a boolean if a field has been set.
func (o *SecurityMonitoringStandardRuleQueryCreate) HasDistinctFields() bool {
	if o != nil && o.DistinctFields != nil {
		return true
	}

	return false
}

// SetDistinctFields gets a reference to the given []string and assigns it to the DistinctFields field.
func (o *SecurityMonitoringStandardRuleQueryCreate) SetDistinctFields(v []string) {
	o.DistinctFields = v
}

// GetGroupByFields returns the GroupByFields field value if set, zero value otherwise.
func (o *SecurityMonitoringStandardRuleQueryCreate) GetGroupByFields() []string {
	if o == nil || o.GroupByFields == nil {
		var ret []string
		return ret
	}
	return o.GroupByFields
}

// GetGroupByFieldsOk returns a tuple with the GroupByFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringStandardRuleQueryCreate) GetGroupByFieldsOk() (*[]string, bool) {
	if o == nil || o.GroupByFields == nil {
		return nil, false
	}
	return &o.GroupByFields, true
}

// HasGroupByFields returns a boolean if a field has been set.
func (o *SecurityMonitoringStandardRuleQueryCreate) HasGroupByFields() bool {
	if o != nil && o.GroupByFields != nil {
		return true
	}

	return false
}

// SetGroupByFields gets a reference to the given []string and assigns it to the GroupByFields field.
func (o *SecurityMonitoringStandardRuleQueryCreate) SetGroupByFields(v []string) {
	o.GroupByFields = v
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *SecurityMonitoringStandardRuleQueryCreate) GetMetric() string {
	if o == nil || o.Metric == nil {
		var ret string
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringStandardRuleQueryCreate) GetMetricOk() (*string, bool) {
	if o == nil || o.Metric == nil {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *SecurityMonitoringStandardRuleQueryCreate) HasMetric() bool {
	if o != nil && o.Metric != nil {
		return true
	}

	return false
}

// SetMetric gets a reference to the given string and assigns it to the Metric field.
func (o *SecurityMonitoringStandardRuleQueryCreate) SetMetric(v string) {
	o.Metric = &v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *SecurityMonitoringStandardRuleQueryCreate) GetMetrics() []string {
	if o == nil || o.Metrics == nil {
		var ret []string
		return ret
	}
	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringStandardRuleQueryCreate) GetMetricsOk() (*[]string, bool) {
	if o == nil || o.Metrics == nil {
		return nil, false
	}
	return &o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *SecurityMonitoringStandardRuleQueryCreate) HasMetrics() bool {
	if o != nil && o.Metrics != nil {
		return true
	}

	return false
}

// SetMetrics gets a reference to the given []string and assigns it to the Metrics field.
func (o *SecurityMonitoringStandardRuleQueryCreate) SetMetrics(v []string) {
	o.Metrics = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SecurityMonitoringStandardRuleQueryCreate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringStandardRuleQueryCreate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SecurityMonitoringStandardRuleQueryCreate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SecurityMonitoringStandardRuleQueryCreate) SetName(v string) {
	o.Name = &v
}

// GetQuery returns the Query field value.
func (o *SecurityMonitoringStandardRuleQueryCreate) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringStandardRuleQueryCreate) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *SecurityMonitoringStandardRuleQueryCreate) SetQuery(v string) {
	o.Query = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringStandardRuleQueryCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Aggregation != nil {
		toSerialize["aggregation"] = o.Aggregation
	}
	if o.DistinctFields != nil {
		toSerialize["distinctFields"] = o.DistinctFields
	}
	if o.GroupByFields != nil {
		toSerialize["groupByFields"] = o.GroupByFields
	}
	if o.Metric != nil {
		toSerialize["metric"] = o.Metric
	}
	if o.Metrics != nil {
		toSerialize["metrics"] = o.Metrics
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	toSerialize["query"] = o.Query

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringStandardRuleQueryCreate) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		Query *string `json:"query"`
	}{}
	all := struct {
		Aggregation    *SecurityMonitoringRuleQueryAggregation `json:"aggregation,omitempty"`
		DistinctFields []string                                `json:"distinctFields,omitempty"`
		GroupByFields  []string                                `json:"groupByFields,omitempty"`
		Metric         *string                                 `json:"metric,omitempty"`
		Metrics        []string                                `json:"metrics,omitempty"`
		Name           *string                                 `json:"name,omitempty"`
		Query          string                                  `json:"query"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Query == nil {
		return fmt.Errorf("Required field query missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if v := all.Aggregation; v != nil && !v.IsValid() {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.Aggregation = all.Aggregation
	o.DistinctFields = all.DistinctFields
	o.GroupByFields = all.GroupByFields
	o.Metric = all.Metric
	o.Metrics = all.Metrics
	o.Name = all.Name
	o.Query = all.Query
	return nil
}
