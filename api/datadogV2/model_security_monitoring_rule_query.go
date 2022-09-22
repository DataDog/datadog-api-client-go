// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
)

// SecurityMonitoringRuleQuery Query for matching rule.
type SecurityMonitoringRuleQuery struct {
	// The aggregation type.
	Aggregation *SecurityMonitoringRuleQueryAggregation `json:"aggregation,omitempty"`
	// Fields to group by for Signal Correlation rules.
	CorrelatedByFields []string `json:"correlatedByFields,omitempty"`
	// Index of the rule query used to retrieve the correlated field for Signal Correlation rules.
	CorrelatedQueryIndex *int32 `json:"correlatedQueryIndex,omitempty"`
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
	Query *string `json:"query,omitempty"`
	// Rule ID to match on signals for Signal Correlation rules.
	RuleId *string `json:"ruleId,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}

// NewSecurityMonitoringRuleQuery instantiates a new SecurityMonitoringRuleQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringRuleQuery() *SecurityMonitoringRuleQuery {
	this := SecurityMonitoringRuleQuery{}
	return &this
}

// NewSecurityMonitoringRuleQueryWithDefaults instantiates a new SecurityMonitoringRuleQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringRuleQueryWithDefaults() *SecurityMonitoringRuleQuery {
	this := SecurityMonitoringRuleQuery{}
	return &this
}

// GetAggregation returns the Aggregation field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleQuery) GetAggregation() SecurityMonitoringRuleQueryAggregation {
	if o == nil || o.Aggregation == nil {
		var ret SecurityMonitoringRuleQueryAggregation
		return ret
	}
	return *o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleQuery) GetAggregationOk() (*SecurityMonitoringRuleQueryAggregation, bool) {
	if o == nil || o.Aggregation == nil {
		return nil, false
	}
	return o.Aggregation, true
}

// HasAggregation returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleQuery) HasAggregation() bool {
	if o != nil && o.Aggregation != nil {
		return true
	}

	return false
}

// SetAggregation gets a reference to the given SecurityMonitoringRuleQueryAggregation and assigns it to the Aggregation field.
func (o *SecurityMonitoringRuleQuery) SetAggregation(v SecurityMonitoringRuleQueryAggregation) {
	o.Aggregation = &v
}

// GetCorrelatedByFields returns the CorrelatedByFields field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleQuery) GetCorrelatedByFields() []string {
	if o == nil || o.CorrelatedByFields == nil {
		var ret []string
		return ret
	}
	return o.CorrelatedByFields
}

// GetCorrelatedByFieldsOk returns a tuple with the CorrelatedByFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleQuery) GetCorrelatedByFieldsOk() (*[]string, bool) {
	if o == nil || o.CorrelatedByFields == nil {
		return nil, false
	}
	return &o.CorrelatedByFields, true
}

// HasCorrelatedByFields returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleQuery) HasCorrelatedByFields() bool {
	if o != nil && o.CorrelatedByFields != nil {
		return true
	}

	return false
}

// SetCorrelatedByFields gets a reference to the given []string and assigns it to the CorrelatedByFields field.
func (o *SecurityMonitoringRuleQuery) SetCorrelatedByFields(v []string) {
	o.CorrelatedByFields = v
}

// GetCorrelatedQueryIndex returns the CorrelatedQueryIndex field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleQuery) GetCorrelatedQueryIndex() int32 {
	if o == nil || o.CorrelatedQueryIndex == nil {
		var ret int32
		return ret
	}
	return *o.CorrelatedQueryIndex
}

// GetCorrelatedQueryIndexOk returns a tuple with the CorrelatedQueryIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleQuery) GetCorrelatedQueryIndexOk() (*int32, bool) {
	if o == nil || o.CorrelatedQueryIndex == nil {
		return nil, false
	}
	return o.CorrelatedQueryIndex, true
}

// HasCorrelatedQueryIndex returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleQuery) HasCorrelatedQueryIndex() bool {
	if o != nil && o.CorrelatedQueryIndex != nil {
		return true
	}

	return false
}

// SetCorrelatedQueryIndex gets a reference to the given int32 and assigns it to the CorrelatedQueryIndex field.
func (o *SecurityMonitoringRuleQuery) SetCorrelatedQueryIndex(v int32) {
	o.CorrelatedQueryIndex = &v
}

// GetDistinctFields returns the DistinctFields field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleQuery) GetDistinctFields() []string {
	if o == nil || o.DistinctFields == nil {
		var ret []string
		return ret
	}
	return o.DistinctFields
}

// GetDistinctFieldsOk returns a tuple with the DistinctFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleQuery) GetDistinctFieldsOk() (*[]string, bool) {
	if o == nil || o.DistinctFields == nil {
		return nil, false
	}
	return &o.DistinctFields, true
}

// HasDistinctFields returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleQuery) HasDistinctFields() bool {
	if o != nil && o.DistinctFields != nil {
		return true
	}

	return false
}

// SetDistinctFields gets a reference to the given []string and assigns it to the DistinctFields field.
func (o *SecurityMonitoringRuleQuery) SetDistinctFields(v []string) {
	o.DistinctFields = v
}

// GetGroupByFields returns the GroupByFields field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleQuery) GetGroupByFields() []string {
	if o == nil || o.GroupByFields == nil {
		var ret []string
		return ret
	}
	return o.GroupByFields
}

// GetGroupByFieldsOk returns a tuple with the GroupByFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleQuery) GetGroupByFieldsOk() (*[]string, bool) {
	if o == nil || o.GroupByFields == nil {
		return nil, false
	}
	return &o.GroupByFields, true
}

// HasGroupByFields returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleQuery) HasGroupByFields() bool {
	if o != nil && o.GroupByFields != nil {
		return true
	}

	return false
}

// SetGroupByFields gets a reference to the given []string and assigns it to the GroupByFields field.
func (o *SecurityMonitoringRuleQuery) SetGroupByFields(v []string) {
	o.GroupByFields = v
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleQuery) GetMetric() string {
	if o == nil || o.Metric == nil {
		var ret string
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleQuery) GetMetricOk() (*string, bool) {
	if o == nil || o.Metric == nil {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleQuery) HasMetric() bool {
	if o != nil && o.Metric != nil {
		return true
	}

	return false
}

// SetMetric gets a reference to the given string and assigns it to the Metric field.
func (o *SecurityMonitoringRuleQuery) SetMetric(v string) {
	o.Metric = &v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleQuery) GetMetrics() []string {
	if o == nil || o.Metrics == nil {
		var ret []string
		return ret
	}
	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleQuery) GetMetricsOk() (*[]string, bool) {
	if o == nil || o.Metrics == nil {
		return nil, false
	}
	return &o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleQuery) HasMetrics() bool {
	if o != nil && o.Metrics != nil {
		return true
	}

	return false
}

// SetMetrics gets a reference to the given []string and assigns it to the Metrics field.
func (o *SecurityMonitoringRuleQuery) SetMetrics(v []string) {
	o.Metrics = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleQuery) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleQuery) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleQuery) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SecurityMonitoringRuleQuery) SetName(v string) {
	o.Name = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleQuery) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleQuery) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleQuery) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *SecurityMonitoringRuleQuery) SetQuery(v string) {
	o.Query = &v
}

// GetRuleId returns the RuleId field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleQuery) GetRuleId() string {
	if o == nil || o.RuleId == nil {
		var ret string
		return ret
	}
	return *o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleQuery) GetRuleIdOk() (*string, bool) {
	if o == nil || o.RuleId == nil {
		return nil, false
	}
	return o.RuleId, true
}

// HasRuleId returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleQuery) HasRuleId() bool {
	if o != nil && o.RuleId != nil {
		return true
	}

	return false
}

// SetRuleId gets a reference to the given string and assigns it to the RuleId field.
func (o *SecurityMonitoringRuleQuery) SetRuleId(v string) {
	o.RuleId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringRuleQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Aggregation != nil {
		toSerialize["aggregation"] = o.Aggregation
	}
	if o.CorrelatedByFields != nil {
		toSerialize["correlatedByFields"] = o.CorrelatedByFields
	}
	if o.CorrelatedQueryIndex != nil {
		toSerialize["correlatedQueryIndex"] = o.CorrelatedQueryIndex
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
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.RuleId != nil {
		toSerialize["ruleId"] = o.RuleId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringRuleQuery) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Aggregation          *SecurityMonitoringRuleQueryAggregation `json:"aggregation,omitempty"`
		CorrelatedByFields   []string                                `json:"correlatedByFields,omitempty"`
		CorrelatedQueryIndex *int32                                  `json:"correlatedQueryIndex,omitempty"`
		DistinctFields       []string                                `json:"distinctFields,omitempty"`
		GroupByFields        []string                                `json:"groupByFields,omitempty"`
		Metric               *string                                 `json:"metric,omitempty"`
		Metrics              []string                                `json:"metrics,omitempty"`
		Name                 *string                                 `json:"name,omitempty"`
		Query                *string                                 `json:"query,omitempty"`
		RuleId               *string                                 `json:"ruleId,omitempty"`
	}{}
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
	o.CorrelatedByFields = all.CorrelatedByFields
	o.CorrelatedQueryIndex = all.CorrelatedQueryIndex
	o.DistinctFields = all.DistinctFields
	o.GroupByFields = all.GroupByFields
	o.Metric = all.Metric
	o.Metrics = all.Metrics
	o.Name = all.Name
	o.Query = all.Query
	o.RuleId = all.RuleId
	return nil
}
