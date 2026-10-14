// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsTargetingRuleOperator Required with attribute and value for an inline condition; omit when saved_filter_id is set.
type ExperimentsTargetingRuleOperator string

// List of ExperimentsTargetingRuleOperator.
const (
	EXPERIMENTSTARGETINGRULEOPERATOR_LT          ExperimentsTargetingRuleOperator = "LT"
	EXPERIMENTSTARGETINGRULEOPERATOR_LTE         ExperimentsTargetingRuleOperator = "LTE"
	EXPERIMENTSTARGETINGRULEOPERATOR_GT          ExperimentsTargetingRuleOperator = "GT"
	EXPERIMENTSTARGETINGRULEOPERATOR_GTE         ExperimentsTargetingRuleOperator = "GTE"
	EXPERIMENTSTARGETINGRULEOPERATOR_MATCHES     ExperimentsTargetingRuleOperator = "MATCHES"
	EXPERIMENTSTARGETINGRULEOPERATOR_NOT_MATCHES ExperimentsTargetingRuleOperator = "NOT_MATCHES"
	EXPERIMENTSTARGETINGRULEOPERATOR_ONE_OF      ExperimentsTargetingRuleOperator = "ONE_OF"
	EXPERIMENTSTARGETINGRULEOPERATOR_NOT_ONE_OF  ExperimentsTargetingRuleOperator = "NOT_ONE_OF"
	EXPERIMENTSTARGETINGRULEOPERATOR_IS_NULL     ExperimentsTargetingRuleOperator = "IS_NULL"
	EXPERIMENTSTARGETINGRULEOPERATOR_EQUALS      ExperimentsTargetingRuleOperator = "EQUALS"
	EXPERIMENTSTARGETINGRULEOPERATOR_SEMVER_EQ   ExperimentsTargetingRuleOperator = "SEMVER_EQ"
	EXPERIMENTSTARGETINGRULEOPERATOR_SEMVER_NEQ  ExperimentsTargetingRuleOperator = "SEMVER_NEQ"
	EXPERIMENTSTARGETINGRULEOPERATOR_SEMVER_LT   ExperimentsTargetingRuleOperator = "SEMVER_LT"
	EXPERIMENTSTARGETINGRULEOPERATOR_SEMVER_LTE  ExperimentsTargetingRuleOperator = "SEMVER_LTE"
	EXPERIMENTSTARGETINGRULEOPERATOR_SEMVER_GT   ExperimentsTargetingRuleOperator = "SEMVER_GT"
	EXPERIMENTSTARGETINGRULEOPERATOR_SEMVER_GTE  ExperimentsTargetingRuleOperator = "SEMVER_GTE"
)

var allowedExperimentsTargetingRuleOperatorEnumValues = []ExperimentsTargetingRuleOperator{
	EXPERIMENTSTARGETINGRULEOPERATOR_LT,
	EXPERIMENTSTARGETINGRULEOPERATOR_LTE,
	EXPERIMENTSTARGETINGRULEOPERATOR_GT,
	EXPERIMENTSTARGETINGRULEOPERATOR_GTE,
	EXPERIMENTSTARGETINGRULEOPERATOR_MATCHES,
	EXPERIMENTSTARGETINGRULEOPERATOR_NOT_MATCHES,
	EXPERIMENTSTARGETINGRULEOPERATOR_ONE_OF,
	EXPERIMENTSTARGETINGRULEOPERATOR_NOT_ONE_OF,
	EXPERIMENTSTARGETINGRULEOPERATOR_IS_NULL,
	EXPERIMENTSTARGETINGRULEOPERATOR_EQUALS,
	EXPERIMENTSTARGETINGRULEOPERATOR_SEMVER_EQ,
	EXPERIMENTSTARGETINGRULEOPERATOR_SEMVER_NEQ,
	EXPERIMENTSTARGETINGRULEOPERATOR_SEMVER_LT,
	EXPERIMENTSTARGETINGRULEOPERATOR_SEMVER_LTE,
	EXPERIMENTSTARGETINGRULEOPERATOR_SEMVER_GT,
	EXPERIMENTSTARGETINGRULEOPERATOR_SEMVER_GTE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsTargetingRuleOperator) GetAllowedValues() []ExperimentsTargetingRuleOperator {
	return allowedExperimentsTargetingRuleOperatorEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsTargetingRuleOperator) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsTargetingRuleOperator(value)
	return nil
}

// NewExperimentsTargetingRuleOperatorFromValue returns a pointer to a valid ExperimentsTargetingRuleOperator
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsTargetingRuleOperatorFromValue(v string) (*ExperimentsTargetingRuleOperator, error) {
	ev := ExperimentsTargetingRuleOperator(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsTargetingRuleOperator: valid values are %v", v, allowedExperimentsTargetingRuleOperatorEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsTargetingRuleOperator) IsValid() bool {
	for _, existing := range allowedExperimentsTargetingRuleOperatorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsTargetingRuleOperator value.
func (v ExperimentsTargetingRuleOperator) Ptr() *ExperimentsTargetingRuleOperator {
	return &v
}
