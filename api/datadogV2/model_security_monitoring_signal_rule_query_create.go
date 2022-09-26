// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"
)

// SecurityMonitoringSignalRuleQueryCreate Query for matching rule on signals
type SecurityMonitoringSignalRuleQueryCreate struct {
	// The aggregation type.
	Aggregation *SecurityMonitoringRuleQueryAggregation `json:"aggregation,omitempty"`
	// Fields to group by.
	CorrelatedByFields []string `json:"correlatedByFields,omitempty"`
	// Index of the rule query used to retrieve the correlated field.
	CorrelatedQueryIndex *int32 `json:"correlatedQueryIndex,omitempty"`
	// Group of target fields to aggregate over when using the new value aggregations.
	Metrics []string `json:"metrics,omitempty"`
	// Name of the query.
	Name *string `json:"name,omitempty"`
	// Rule ID to match on signals.
	RuleId string `json:"ruleId"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}

// NewSecurityMonitoringSignalRuleQueryCreate instantiates a new SecurityMonitoringSignalRuleQueryCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringSignalRuleQueryCreate(ruleId string) *SecurityMonitoringSignalRuleQueryCreate {
	this := SecurityMonitoringSignalRuleQueryCreate{}
	this.RuleId = ruleId
	return &this
}

// NewSecurityMonitoringSignalRuleQueryCreateWithDefaults instantiates a new SecurityMonitoringSignalRuleQueryCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringSignalRuleQueryCreateWithDefaults() *SecurityMonitoringSignalRuleQueryCreate {
	this := SecurityMonitoringSignalRuleQueryCreate{}
	return &this
}

// GetAggregation returns the Aggregation field value if set, zero value otherwise.
func (o *SecurityMonitoringSignalRuleQueryCreate) GetAggregation() SecurityMonitoringRuleQueryAggregation {
	if o == nil || o.Aggregation == nil {
		var ret SecurityMonitoringRuleQueryAggregation
		return ret
	}
	return *o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalRuleQueryCreate) GetAggregationOk() (*SecurityMonitoringRuleQueryAggregation, bool) {
	if o == nil || o.Aggregation == nil {
		return nil, false
	}
	return o.Aggregation, true
}

// HasAggregation returns a boolean if a field has been set.
func (o *SecurityMonitoringSignalRuleQueryCreate) HasAggregation() bool {
	if o != nil && o.Aggregation != nil {
		return true
	}

	return false
}

// SetAggregation gets a reference to the given SecurityMonitoringRuleQueryAggregation and assigns it to the Aggregation field.
func (o *SecurityMonitoringSignalRuleQueryCreate) SetAggregation(v SecurityMonitoringRuleQueryAggregation) {
	o.Aggregation = &v
}

// GetCorrelatedByFields returns the CorrelatedByFields field value if set, zero value otherwise.
func (o *SecurityMonitoringSignalRuleQueryCreate) GetCorrelatedByFields() []string {
	if o == nil || o.CorrelatedByFields == nil {
		var ret []string
		return ret
	}
	return o.CorrelatedByFields
}

// GetCorrelatedByFieldsOk returns a tuple with the CorrelatedByFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalRuleQueryCreate) GetCorrelatedByFieldsOk() (*[]string, bool) {
	if o == nil || o.CorrelatedByFields == nil {
		return nil, false
	}
	return &o.CorrelatedByFields, true
}

// HasCorrelatedByFields returns a boolean if a field has been set.
func (o *SecurityMonitoringSignalRuleQueryCreate) HasCorrelatedByFields() bool {
	if o != nil && o.CorrelatedByFields != nil {
		return true
	}

	return false
}

// SetCorrelatedByFields gets a reference to the given []string and assigns it to the CorrelatedByFields field.
func (o *SecurityMonitoringSignalRuleQueryCreate) SetCorrelatedByFields(v []string) {
	o.CorrelatedByFields = v
}

// GetCorrelatedQueryIndex returns the CorrelatedQueryIndex field value if set, zero value otherwise.
func (o *SecurityMonitoringSignalRuleQueryCreate) GetCorrelatedQueryIndex() int32 {
	if o == nil || o.CorrelatedQueryIndex == nil {
		var ret int32
		return ret
	}
	return *o.CorrelatedQueryIndex
}

// GetCorrelatedQueryIndexOk returns a tuple with the CorrelatedQueryIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalRuleQueryCreate) GetCorrelatedQueryIndexOk() (*int32, bool) {
	if o == nil || o.CorrelatedQueryIndex == nil {
		return nil, false
	}
	return o.CorrelatedQueryIndex, true
}

// HasCorrelatedQueryIndex returns a boolean if a field has been set.
func (o *SecurityMonitoringSignalRuleQueryCreate) HasCorrelatedQueryIndex() bool {
	if o != nil && o.CorrelatedQueryIndex != nil {
		return true
	}

	return false
}

// SetCorrelatedQueryIndex gets a reference to the given int32 and assigns it to the CorrelatedQueryIndex field.
func (o *SecurityMonitoringSignalRuleQueryCreate) SetCorrelatedQueryIndex(v int32) {
	o.CorrelatedQueryIndex = &v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *SecurityMonitoringSignalRuleQueryCreate) GetMetrics() []string {
	if o == nil || o.Metrics == nil {
		var ret []string
		return ret
	}
	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalRuleQueryCreate) GetMetricsOk() (*[]string, bool) {
	if o == nil || o.Metrics == nil {
		return nil, false
	}
	return &o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *SecurityMonitoringSignalRuleQueryCreate) HasMetrics() bool {
	if o != nil && o.Metrics != nil {
		return true
	}

	return false
}

// SetMetrics gets a reference to the given []string and assigns it to the Metrics field.
func (o *SecurityMonitoringSignalRuleQueryCreate) SetMetrics(v []string) {
	o.Metrics = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SecurityMonitoringSignalRuleQueryCreate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalRuleQueryCreate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SecurityMonitoringSignalRuleQueryCreate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SecurityMonitoringSignalRuleQueryCreate) SetName(v string) {
	o.Name = &v
}

// GetRuleId returns the RuleId field value.
func (o *SecurityMonitoringSignalRuleQueryCreate) GetRuleId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalRuleQueryCreate) GetRuleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleId, true
}

// SetRuleId sets field value.
func (o *SecurityMonitoringSignalRuleQueryCreate) SetRuleId(v string) {
	o.RuleId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringSignalRuleQueryCreate) MarshalJSON() ([]byte, error) {
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
	if o.Metrics != nil {
		toSerialize["metrics"] = o.Metrics
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	toSerialize["ruleId"] = o.RuleId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringSignalRuleQueryCreate) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		RuleId *string `json:"ruleId"`
	}{}
	all := struct {
		Aggregation          *SecurityMonitoringRuleQueryAggregation `json:"aggregation,omitempty"`
		CorrelatedByFields   []string                                `json:"correlatedByFields,omitempty"`
		CorrelatedQueryIndex *int32                                  `json:"correlatedQueryIndex,omitempty"`
		Metrics              []string                                `json:"metrics,omitempty"`
		Name                 *string                                 `json:"name,omitempty"`
		RuleId               string                                  `json:"ruleId"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.RuleId == nil {
		return fmt.Errorf("Required field ruleId missing")
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
	o.CorrelatedByFields = all.CorrelatedByFields
	o.CorrelatedQueryIndex = all.CorrelatedQueryIndex
	o.Metrics = all.Metrics
	o.Name = all.Name
	o.RuleId = all.RuleId
	return nil
}
