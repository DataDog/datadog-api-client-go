// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeploymentGatesBitsDeploymentAnalysisRuleType The type identifier for a Bits Deployment Analysis rule.
type DeploymentGatesBitsDeploymentAnalysisRuleType string

// List of DeploymentGatesBitsDeploymentAnalysisRuleType.
const (
	DEPLOYMENTGATESBITSDEPLOYMENTANALYSISRULETYPE_BITS_DEPLOYMENT_ANALYSIS DeploymentGatesBitsDeploymentAnalysisRuleType = "bits_deployment_analysis"
)

var allowedDeploymentGatesBitsDeploymentAnalysisRuleTypeEnumValues = []DeploymentGatesBitsDeploymentAnalysisRuleType{
	DEPLOYMENTGATESBITSDEPLOYMENTANALYSISRULETYPE_BITS_DEPLOYMENT_ANALYSIS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DeploymentGatesBitsDeploymentAnalysisRuleType) GetAllowedValues() []DeploymentGatesBitsDeploymentAnalysisRuleType {
	return allowedDeploymentGatesBitsDeploymentAnalysisRuleTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DeploymentGatesBitsDeploymentAnalysisRuleType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DeploymentGatesBitsDeploymentAnalysisRuleType(value)
	return nil
}

// NewDeploymentGatesBitsDeploymentAnalysisRuleTypeFromValue returns a pointer to a valid DeploymentGatesBitsDeploymentAnalysisRuleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDeploymentGatesBitsDeploymentAnalysisRuleTypeFromValue(v string) (*DeploymentGatesBitsDeploymentAnalysisRuleType, error) {
	ev := DeploymentGatesBitsDeploymentAnalysisRuleType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DeploymentGatesBitsDeploymentAnalysisRuleType: valid values are %v", v, allowedDeploymentGatesBitsDeploymentAnalysisRuleTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DeploymentGatesBitsDeploymentAnalysisRuleType) IsValid() bool {
	for _, existing := range allowedDeploymentGatesBitsDeploymentAnalysisRuleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeploymentGatesBitsDeploymentAnalysisRuleType value.
func (v DeploymentGatesBitsDeploymentAnalysisRuleType) Ptr() *DeploymentGatesBitsDeploymentAnalysisRuleType {
	return &v
}
