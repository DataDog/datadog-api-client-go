// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeploymentGateRuleEvaluationDataType JSON:API type for a deployment gate rule evaluation.
type DeploymentGateRuleEvaluationDataType string

// List of DeploymentGateRuleEvaluationDataType.
const (
	DEPLOYMENTGATERULEEVALUATIONDATATYPE_DEPLOYMENT_GATE_RULE_EVALUATION DeploymentGateRuleEvaluationDataType = "deployment_gate_rule_evaluation"
)

var allowedDeploymentGateRuleEvaluationDataTypeEnumValues = []DeploymentGateRuleEvaluationDataType{
	DEPLOYMENTGATERULEEVALUATIONDATATYPE_DEPLOYMENT_GATE_RULE_EVALUATION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DeploymentGateRuleEvaluationDataType) GetAllowedValues() []DeploymentGateRuleEvaluationDataType {
	return allowedDeploymentGateRuleEvaluationDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DeploymentGateRuleEvaluationDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DeploymentGateRuleEvaluationDataType(value)
	return nil
}

// NewDeploymentGateRuleEvaluationDataTypeFromValue returns a pointer to a valid DeploymentGateRuleEvaluationDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDeploymentGateRuleEvaluationDataTypeFromValue(v string) (*DeploymentGateRuleEvaluationDataType, error) {
	ev := DeploymentGateRuleEvaluationDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DeploymentGateRuleEvaluationDataType: valid values are %v", v, allowedDeploymentGateRuleEvaluationDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DeploymentGateRuleEvaluationDataType) IsValid() bool {
	for _, existing := range allowedDeploymentGateRuleEvaluationDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeploymentGateRuleEvaluationDataType value.
func (v DeploymentGateRuleEvaluationDataType) Ptr() *DeploymentGateRuleEvaluationDataType {
	return &v
}
