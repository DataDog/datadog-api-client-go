// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaDataQualityMetricsQueryAnomalyOptions Test-only generated schema for the dashboard versioning proof of concept.
type FormulaDataQualityMetricsQueryAnomalyOptions struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Direction *FormulaDataQualityMetricsQueryAnomalyOptionsDirection `json:"direction,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	RuleId *float64 `json:"rule_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewFormulaDataQualityMetricsQueryAnomalyOptions instantiates a new FormulaDataQualityMetricsQueryAnomalyOptions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormulaDataQualityMetricsQueryAnomalyOptions() *FormulaDataQualityMetricsQueryAnomalyOptions {
	this := FormulaDataQualityMetricsQueryAnomalyOptions{}
	return &this
}

// NewFormulaDataQualityMetricsQueryAnomalyOptionsWithDefaults instantiates a new FormulaDataQualityMetricsQueryAnomalyOptions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormulaDataQualityMetricsQueryAnomalyOptionsWithDefaults() *FormulaDataQualityMetricsQueryAnomalyOptions {
	this := FormulaDataQualityMetricsQueryAnomalyOptions{}
	return &this
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *FormulaDataQualityMetricsQueryAnomalyOptions) GetDirection() FormulaDataQualityMetricsQueryAnomalyOptionsDirection {
	if o == nil || o.Direction == nil {
		var ret FormulaDataQualityMetricsQueryAnomalyOptionsDirection
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaDataQualityMetricsQueryAnomalyOptions) GetDirectionOk() (*FormulaDataQualityMetricsQueryAnomalyOptionsDirection, bool) {
	if o == nil || o.Direction == nil {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *FormulaDataQualityMetricsQueryAnomalyOptions) HasDirection() bool {
	return o != nil && o.Direction != nil
}

// SetDirection gets a reference to the given FormulaDataQualityMetricsQueryAnomalyOptionsDirection and assigns it to the Direction field.
func (o *FormulaDataQualityMetricsQueryAnomalyOptions) SetDirection(v FormulaDataQualityMetricsQueryAnomalyOptionsDirection) {
	o.Direction = &v
}

// GetRuleId returns the RuleId field value if set, zero value otherwise.
func (o *FormulaDataQualityMetricsQueryAnomalyOptions) GetRuleId() float64 {
	if o == nil || o.RuleId == nil {
		var ret float64
		return ret
	}
	return *o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaDataQualityMetricsQueryAnomalyOptions) GetRuleIdOk() (*float64, bool) {
	if o == nil || o.RuleId == nil {
		return nil, false
	}
	return o.RuleId, true
}

// HasRuleId returns a boolean if a field has been set.
func (o *FormulaDataQualityMetricsQueryAnomalyOptions) HasRuleId() bool {
	return o != nil && o.RuleId != nil
}

// SetRuleId gets a reference to the given float64 and assigns it to the RuleId field.
func (o *FormulaDataQualityMetricsQueryAnomalyOptions) SetRuleId(v float64) {
	o.RuleId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormulaDataQualityMetricsQueryAnomalyOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Direction != nil {
		toSerialize["direction"] = o.Direction
	}
	if o.RuleId != nil {
		toSerialize["rule_id"] = o.RuleId
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormulaDataQualityMetricsQueryAnomalyOptions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Direction *FormulaDataQualityMetricsQueryAnomalyOptionsDirection `json:"direction,omitempty"`
		RuleId    *float64                                               `json:"rule_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	hasInvalidField := false
	if all.Direction != nil && !all.Direction.IsValid() {
		hasInvalidField = true
	} else {
		o.Direction = all.Direction
	}
	o.RuleId = all.RuleId

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
