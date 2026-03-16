// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeploymentRulesOptionsResponse - Options returned in deployment rule responses representing either faulty deployment detection or monitor options. Faulty deployment detection responses always include `excluded_resources`, making the two variants unambiguous.
type DeploymentRulesOptionsResponse struct {
	DeploymentRuleOptionsFaultyDeploymentDetectionResponse *DeploymentRuleOptionsFaultyDeploymentDetectionResponse
	DeploymentRuleOptionsMonitor                           *DeploymentRuleOptionsMonitor

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// DeploymentRuleOptionsFaultyDeploymentDetectionResponseAsDeploymentRulesOptionsResponse is a convenience function that returns DeploymentRuleOptionsFaultyDeploymentDetectionResponse wrapped in DeploymentRulesOptionsResponse.
func DeploymentRuleOptionsFaultyDeploymentDetectionResponseAsDeploymentRulesOptionsResponse(v *DeploymentRuleOptionsFaultyDeploymentDetectionResponse) DeploymentRulesOptionsResponse {
	return DeploymentRulesOptionsResponse{DeploymentRuleOptionsFaultyDeploymentDetectionResponse: v}
}

// DeploymentRuleOptionsMonitorAsDeploymentRulesOptionsResponse is a convenience function that returns DeploymentRuleOptionsMonitor wrapped in DeploymentRulesOptionsResponse.
func DeploymentRuleOptionsMonitorAsDeploymentRulesOptionsResponse(v *DeploymentRuleOptionsMonitor) DeploymentRulesOptionsResponse {
	return DeploymentRulesOptionsResponse{DeploymentRuleOptionsMonitor: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *DeploymentRulesOptionsResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DeploymentRuleOptionsFaultyDeploymentDetectionResponse
	err = datadog.Unmarshal(data, &obj.DeploymentRuleOptionsFaultyDeploymentDetectionResponse)
	if err == nil {
		if obj.DeploymentRuleOptionsFaultyDeploymentDetectionResponse != nil && obj.DeploymentRuleOptionsFaultyDeploymentDetectionResponse.UnparsedObject == nil {
			jsonDeploymentRuleOptionsFaultyDeploymentDetectionResponse, _ := datadog.Marshal(obj.DeploymentRuleOptionsFaultyDeploymentDetectionResponse)
			if string(jsonDeploymentRuleOptionsFaultyDeploymentDetectionResponse) == "{}" { // empty struct
				obj.DeploymentRuleOptionsFaultyDeploymentDetectionResponse = nil
			} else {
				match++
			}
		} else {
			obj.DeploymentRuleOptionsFaultyDeploymentDetectionResponse = nil
		}
	} else {
		obj.DeploymentRuleOptionsFaultyDeploymentDetectionResponse = nil
	}

	// try to unmarshal data into DeploymentRuleOptionsMonitor
	err = datadog.Unmarshal(data, &obj.DeploymentRuleOptionsMonitor)
	if err == nil {
		if obj.DeploymentRuleOptionsMonitor != nil && obj.DeploymentRuleOptionsMonitor.UnparsedObject == nil {
			jsonDeploymentRuleOptionsMonitor, _ := datadog.Marshal(obj.DeploymentRuleOptionsMonitor)
			if string(jsonDeploymentRuleOptionsMonitor) == "{}" { // empty struct
				obj.DeploymentRuleOptionsMonitor = nil
			} else {
				match++
			}
		} else {
			obj.DeploymentRuleOptionsMonitor = nil
		}
	} else {
		obj.DeploymentRuleOptionsMonitor = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.DeploymentRuleOptionsFaultyDeploymentDetectionResponse = nil
		obj.DeploymentRuleOptionsMonitor = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj DeploymentRulesOptionsResponse) MarshalJSON() ([]byte, error) {
	if obj.DeploymentRuleOptionsFaultyDeploymentDetectionResponse != nil {
		return datadog.Marshal(&obj.DeploymentRuleOptionsFaultyDeploymentDetectionResponse)
	}

	if obj.DeploymentRuleOptionsMonitor != nil {
		return datadog.Marshal(&obj.DeploymentRuleOptionsMonitor)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *DeploymentRulesOptionsResponse) GetActualInstance() interface{} {
	if obj.DeploymentRuleOptionsFaultyDeploymentDetectionResponse != nil {
		return obj.DeploymentRuleOptionsFaultyDeploymentDetectionResponse
	}

	if obj.DeploymentRuleOptionsMonitor != nil {
		return obj.DeploymentRuleOptionsMonitor
	}

	// all schemas are nil
	return nil
}
