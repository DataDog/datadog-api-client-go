// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeploymentGateEvaluationDataType JSON:API type for a deployment gate evaluation.
type DeploymentGateEvaluationDataType string

// List of DeploymentGateEvaluationDataType.
const (
	DEPLOYMENTGATEEVALUATIONDATATYPE_DEPLOYMENT_GATE_EVALUATION DeploymentGateEvaluationDataType = "deployment_gate_evaluation"
)

var allowedDeploymentGateEvaluationDataTypeEnumValues = []DeploymentGateEvaluationDataType{
	DEPLOYMENTGATEEVALUATIONDATATYPE_DEPLOYMENT_GATE_EVALUATION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DeploymentGateEvaluationDataType) GetAllowedValues() []DeploymentGateEvaluationDataType {
	return allowedDeploymentGateEvaluationDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DeploymentGateEvaluationDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DeploymentGateEvaluationDataType(value)
	return nil
}

// NewDeploymentGateEvaluationDataTypeFromValue returns a pointer to a valid DeploymentGateEvaluationDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDeploymentGateEvaluationDataTypeFromValue(v string) (*DeploymentGateEvaluationDataType, error) {
	ev := DeploymentGateEvaluationDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DeploymentGateEvaluationDataType: valid values are %v", v, allowedDeploymentGateEvaluationDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DeploymentGateEvaluationDataType) IsValid() bool {
	for _, existing := range allowedDeploymentGateEvaluationDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeploymentGateEvaluationDataType value.
func (v DeploymentGateEvaluationDataType) Ptr() *DeploymentGateEvaluationDataType {
	return &v
}
