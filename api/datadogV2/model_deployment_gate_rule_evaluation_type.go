// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeploymentGateRuleEvaluationType Type of deployment gate rule.
type DeploymentGateRuleEvaluationType string

// List of DeploymentGateRuleEvaluationType.
const (
	DEPLOYMENTGATERULEEVALUATIONTYPE_MONITOR                     DeploymentGateRuleEvaluationType = "monitor"
	DEPLOYMENTGATERULEEVALUATIONTYPE_FAULTY_DEPLOYMENT_DETECTION DeploymentGateRuleEvaluationType = "faulty_deployment_detection"
)

var allowedDeploymentGateRuleEvaluationTypeEnumValues = []DeploymentGateRuleEvaluationType{
	DEPLOYMENTGATERULEEVALUATIONTYPE_MONITOR,
	DEPLOYMENTGATERULEEVALUATIONTYPE_FAULTY_DEPLOYMENT_DETECTION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DeploymentGateRuleEvaluationType) GetAllowedValues() []DeploymentGateRuleEvaluationType {
	return allowedDeploymentGateRuleEvaluationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DeploymentGateRuleEvaluationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DeploymentGateRuleEvaluationType(value)
	return nil
}

// NewDeploymentGateRuleEvaluationTypeFromValue returns a pointer to a valid DeploymentGateRuleEvaluationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDeploymentGateRuleEvaluationTypeFromValue(v string) (*DeploymentGateRuleEvaluationType, error) {
	ev := DeploymentGateRuleEvaluationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DeploymentGateRuleEvaluationType: valid values are %v", v, allowedDeploymentGateRuleEvaluationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DeploymentGateRuleEvaluationType) IsValid() bool {
	for _, existing := range allowedDeploymentGateRuleEvaluationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeploymentGateRuleEvaluationType value.
func (v DeploymentGateRuleEvaluationType) Ptr() *DeploymentGateRuleEvaluationType {
	return &v
}
