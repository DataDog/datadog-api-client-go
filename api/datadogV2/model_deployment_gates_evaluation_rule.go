// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeploymentGatesEvaluationRule - A rule to evaluate as part of a deployment gate evaluation.
type DeploymentGatesEvaluationRule struct {
	DeploymentGatesMonitorRule                *DeploymentGatesMonitorRule
	DeploymentGatesFDDRule                    *DeploymentGatesFDDRule
	DeploymentGatesBitsDeploymentAnalysisRule *DeploymentGatesBitsDeploymentAnalysisRule

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// DeploymentGatesMonitorRuleAsDeploymentGatesEvaluationRule is a convenience function that returns DeploymentGatesMonitorRule wrapped in DeploymentGatesEvaluationRule.
func DeploymentGatesMonitorRuleAsDeploymentGatesEvaluationRule(v *DeploymentGatesMonitorRule) DeploymentGatesEvaluationRule {
	return DeploymentGatesEvaluationRule{DeploymentGatesMonitorRule: v}
}

// DeploymentGatesFDDRuleAsDeploymentGatesEvaluationRule is a convenience function that returns DeploymentGatesFDDRule wrapped in DeploymentGatesEvaluationRule.
func DeploymentGatesFDDRuleAsDeploymentGatesEvaluationRule(v *DeploymentGatesFDDRule) DeploymentGatesEvaluationRule {
	return DeploymentGatesEvaluationRule{DeploymentGatesFDDRule: v}
}

// DeploymentGatesBitsDeploymentAnalysisRuleAsDeploymentGatesEvaluationRule is a convenience function that returns DeploymentGatesBitsDeploymentAnalysisRule wrapped in DeploymentGatesEvaluationRule.
func DeploymentGatesBitsDeploymentAnalysisRuleAsDeploymentGatesEvaluationRule(v *DeploymentGatesBitsDeploymentAnalysisRule) DeploymentGatesEvaluationRule {
	return DeploymentGatesEvaluationRule{DeploymentGatesBitsDeploymentAnalysisRule: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *DeploymentGatesEvaluationRule) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DeploymentGatesMonitorRule
	err = datadog.Unmarshal(data, &obj.DeploymentGatesMonitorRule)
	if err == nil {
		if obj.DeploymentGatesMonitorRule != nil && obj.DeploymentGatesMonitorRule.UnparsedObject == nil {
			jsonDeploymentGatesMonitorRule, _ := datadog.Marshal(obj.DeploymentGatesMonitorRule)
			if string(jsonDeploymentGatesMonitorRule) == "{}" { // empty struct
				obj.DeploymentGatesMonitorRule = nil
			} else {
				match++
			}
		} else {
			obj.DeploymentGatesMonitorRule = nil
		}
	} else {
		obj.DeploymentGatesMonitorRule = nil
	}

	// try to unmarshal data into DeploymentGatesFDDRule
	err = datadog.Unmarshal(data, &obj.DeploymentGatesFDDRule)
	if err == nil {
		if obj.DeploymentGatesFDDRule != nil && obj.DeploymentGatesFDDRule.UnparsedObject == nil {
			jsonDeploymentGatesFDDRule, _ := datadog.Marshal(obj.DeploymentGatesFDDRule)
			if string(jsonDeploymentGatesFDDRule) == "{}" { // empty struct
				obj.DeploymentGatesFDDRule = nil
			} else {
				match++
			}
		} else {
			obj.DeploymentGatesFDDRule = nil
		}
	} else {
		obj.DeploymentGatesFDDRule = nil
	}

	// try to unmarshal data into DeploymentGatesBitsDeploymentAnalysisRule
	err = datadog.Unmarshal(data, &obj.DeploymentGatesBitsDeploymentAnalysisRule)
	if err == nil {
		if obj.DeploymentGatesBitsDeploymentAnalysisRule != nil && obj.DeploymentGatesBitsDeploymentAnalysisRule.UnparsedObject == nil {
			jsonDeploymentGatesBitsDeploymentAnalysisRule, _ := datadog.Marshal(obj.DeploymentGatesBitsDeploymentAnalysisRule)
			if string(jsonDeploymentGatesBitsDeploymentAnalysisRule) == "{}" { // empty struct
				obj.DeploymentGatesBitsDeploymentAnalysisRule = nil
			} else {
				match++
			}
		} else {
			obj.DeploymentGatesBitsDeploymentAnalysisRule = nil
		}
	} else {
		obj.DeploymentGatesBitsDeploymentAnalysisRule = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.DeploymentGatesMonitorRule = nil
		obj.DeploymentGatesFDDRule = nil
		obj.DeploymentGatesBitsDeploymentAnalysisRule = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj DeploymentGatesEvaluationRule) MarshalJSON() ([]byte, error) {
	if obj.DeploymentGatesMonitorRule != nil {
		return datadog.Marshal(&obj.DeploymentGatesMonitorRule)
	}

	if obj.DeploymentGatesFDDRule != nil {
		return datadog.Marshal(&obj.DeploymentGatesFDDRule)
	}

	if obj.DeploymentGatesBitsDeploymentAnalysisRule != nil {
		return datadog.Marshal(&obj.DeploymentGatesBitsDeploymentAnalysisRule)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *DeploymentGatesEvaluationRule) GetActualInstance() interface{} {
	if obj.DeploymentGatesMonitorRule != nil {
		return obj.DeploymentGatesMonitorRule
	}

	if obj.DeploymentGatesFDDRule != nil {
		return obj.DeploymentGatesFDDRule
	}

	if obj.DeploymentGatesBitsDeploymentAnalysisRule != nil {
		return obj.DeploymentGatesBitsDeploymentAnalysisRule
	}

	// all schemas are nil
	return nil
}
