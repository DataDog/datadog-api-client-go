// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringIntegrationConfigCreateAttributes - The attributes of the entity context sync configuration to create.
type SecurityMonitoringIntegrationConfigCreateAttributes struct {
	SecurityMonitoringGoogleWorkspaceIntegrationConfigCreateAttributes *SecurityMonitoringGoogleWorkspaceIntegrationConfigCreateAttributes
	SecurityMonitoringOktaIntegrationConfigCreateAttributes            *SecurityMonitoringOktaIntegrationConfigCreateAttributes
	SecurityMonitoringEntraIdIntegrationConfigCreateAttributes         *SecurityMonitoringEntraIdIntegrationConfigCreateAttributes
	SecurityMonitoringCrowdStrikeIntegrationConfigCreateAttributes     *SecurityMonitoringCrowdStrikeIntegrationConfigCreateAttributes
	SecurityMonitoringSentinelOneIntegrationConfigCreateAttributes     *SecurityMonitoringSentinelOneIntegrationConfigCreateAttributes

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// SecurityMonitoringGoogleWorkspaceIntegrationConfigCreateAttributesAsSecurityMonitoringIntegrationConfigCreateAttributes is a convenience function that returns SecurityMonitoringGoogleWorkspaceIntegrationConfigCreateAttributes wrapped in SecurityMonitoringIntegrationConfigCreateAttributes.
func SecurityMonitoringGoogleWorkspaceIntegrationConfigCreateAttributesAsSecurityMonitoringIntegrationConfigCreateAttributes(v *SecurityMonitoringGoogleWorkspaceIntegrationConfigCreateAttributes) SecurityMonitoringIntegrationConfigCreateAttributes {
	return SecurityMonitoringIntegrationConfigCreateAttributes{SecurityMonitoringGoogleWorkspaceIntegrationConfigCreateAttributes: v}
}

// SecurityMonitoringOktaIntegrationConfigCreateAttributesAsSecurityMonitoringIntegrationConfigCreateAttributes is a convenience function that returns SecurityMonitoringOktaIntegrationConfigCreateAttributes wrapped in SecurityMonitoringIntegrationConfigCreateAttributes.
func SecurityMonitoringOktaIntegrationConfigCreateAttributesAsSecurityMonitoringIntegrationConfigCreateAttributes(v *SecurityMonitoringOktaIntegrationConfigCreateAttributes) SecurityMonitoringIntegrationConfigCreateAttributes {
	return SecurityMonitoringIntegrationConfigCreateAttributes{SecurityMonitoringOktaIntegrationConfigCreateAttributes: v}
}

// SecurityMonitoringEntraIdIntegrationConfigCreateAttributesAsSecurityMonitoringIntegrationConfigCreateAttributes is a convenience function that returns SecurityMonitoringEntraIdIntegrationConfigCreateAttributes wrapped in SecurityMonitoringIntegrationConfigCreateAttributes.
func SecurityMonitoringEntraIdIntegrationConfigCreateAttributesAsSecurityMonitoringIntegrationConfigCreateAttributes(v *SecurityMonitoringEntraIdIntegrationConfigCreateAttributes) SecurityMonitoringIntegrationConfigCreateAttributes {
	return SecurityMonitoringIntegrationConfigCreateAttributes{SecurityMonitoringEntraIdIntegrationConfigCreateAttributes: v}
}

// SecurityMonitoringCrowdStrikeIntegrationConfigCreateAttributesAsSecurityMonitoringIntegrationConfigCreateAttributes is a convenience function that returns SecurityMonitoringCrowdStrikeIntegrationConfigCreateAttributes wrapped in SecurityMonitoringIntegrationConfigCreateAttributes.
func SecurityMonitoringCrowdStrikeIntegrationConfigCreateAttributesAsSecurityMonitoringIntegrationConfigCreateAttributes(v *SecurityMonitoringCrowdStrikeIntegrationConfigCreateAttributes) SecurityMonitoringIntegrationConfigCreateAttributes {
	return SecurityMonitoringIntegrationConfigCreateAttributes{SecurityMonitoringCrowdStrikeIntegrationConfigCreateAttributes: v}
}

// SecurityMonitoringSentinelOneIntegrationConfigCreateAttributesAsSecurityMonitoringIntegrationConfigCreateAttributes is a convenience function that returns SecurityMonitoringSentinelOneIntegrationConfigCreateAttributes wrapped in SecurityMonitoringIntegrationConfigCreateAttributes.
func SecurityMonitoringSentinelOneIntegrationConfigCreateAttributesAsSecurityMonitoringIntegrationConfigCreateAttributes(v *SecurityMonitoringSentinelOneIntegrationConfigCreateAttributes) SecurityMonitoringIntegrationConfigCreateAttributes {
	return SecurityMonitoringIntegrationConfigCreateAttributes{SecurityMonitoringSentinelOneIntegrationConfigCreateAttributes: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *SecurityMonitoringIntegrationConfigCreateAttributes) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SecurityMonitoringGoogleWorkspaceIntegrationConfigCreateAttributes
	err = datadog.Unmarshal(data, &obj.SecurityMonitoringGoogleWorkspaceIntegrationConfigCreateAttributes)
	if err == nil {
		if obj.SecurityMonitoringGoogleWorkspaceIntegrationConfigCreateAttributes != nil && obj.SecurityMonitoringGoogleWorkspaceIntegrationConfigCreateAttributes.UnparsedObject == nil {
			jsonSecurityMonitoringGoogleWorkspaceIntegrationConfigCreateAttributes, _ := datadog.Marshal(obj.SecurityMonitoringGoogleWorkspaceIntegrationConfigCreateAttributes)
			if string(jsonSecurityMonitoringGoogleWorkspaceIntegrationConfigCreateAttributes) == "{}" { // empty struct
				obj.SecurityMonitoringGoogleWorkspaceIntegrationConfigCreateAttributes = nil
			} else {
				match++
			}
		} else {
			obj.SecurityMonitoringGoogleWorkspaceIntegrationConfigCreateAttributes = nil
		}
	} else {
		obj.SecurityMonitoringGoogleWorkspaceIntegrationConfigCreateAttributes = nil
	}

	// try to unmarshal data into SecurityMonitoringOktaIntegrationConfigCreateAttributes
	err = datadog.Unmarshal(data, &obj.SecurityMonitoringOktaIntegrationConfigCreateAttributes)
	if err == nil {
		if obj.SecurityMonitoringOktaIntegrationConfigCreateAttributes != nil && obj.SecurityMonitoringOktaIntegrationConfigCreateAttributes.UnparsedObject == nil {
			jsonSecurityMonitoringOktaIntegrationConfigCreateAttributes, _ := datadog.Marshal(obj.SecurityMonitoringOktaIntegrationConfigCreateAttributes)
			if string(jsonSecurityMonitoringOktaIntegrationConfigCreateAttributes) == "{}" { // empty struct
				obj.SecurityMonitoringOktaIntegrationConfigCreateAttributes = nil
			} else {
				match++
			}
		} else {
			obj.SecurityMonitoringOktaIntegrationConfigCreateAttributes = nil
		}
	} else {
		obj.SecurityMonitoringOktaIntegrationConfigCreateAttributes = nil
	}

	// try to unmarshal data into SecurityMonitoringEntraIdIntegrationConfigCreateAttributes
	err = datadog.Unmarshal(data, &obj.SecurityMonitoringEntraIdIntegrationConfigCreateAttributes)
	if err == nil {
		if obj.SecurityMonitoringEntraIdIntegrationConfigCreateAttributes != nil && obj.SecurityMonitoringEntraIdIntegrationConfigCreateAttributes.UnparsedObject == nil {
			jsonSecurityMonitoringEntraIdIntegrationConfigCreateAttributes, _ := datadog.Marshal(obj.SecurityMonitoringEntraIdIntegrationConfigCreateAttributes)
			if string(jsonSecurityMonitoringEntraIdIntegrationConfigCreateAttributes) == "{}" { // empty struct
				obj.SecurityMonitoringEntraIdIntegrationConfigCreateAttributes = nil
			} else {
				match++
			}
		} else {
			obj.SecurityMonitoringEntraIdIntegrationConfigCreateAttributes = nil
		}
	} else {
		obj.SecurityMonitoringEntraIdIntegrationConfigCreateAttributes = nil
	}

	// try to unmarshal data into SecurityMonitoringCrowdStrikeIntegrationConfigCreateAttributes
	err = datadog.Unmarshal(data, &obj.SecurityMonitoringCrowdStrikeIntegrationConfigCreateAttributes)
	if err == nil {
		if obj.SecurityMonitoringCrowdStrikeIntegrationConfigCreateAttributes != nil && obj.SecurityMonitoringCrowdStrikeIntegrationConfigCreateAttributes.UnparsedObject == nil {
			jsonSecurityMonitoringCrowdStrikeIntegrationConfigCreateAttributes, _ := datadog.Marshal(obj.SecurityMonitoringCrowdStrikeIntegrationConfigCreateAttributes)
			if string(jsonSecurityMonitoringCrowdStrikeIntegrationConfigCreateAttributes) == "{}" { // empty struct
				obj.SecurityMonitoringCrowdStrikeIntegrationConfigCreateAttributes = nil
			} else {
				match++
			}
		} else {
			obj.SecurityMonitoringCrowdStrikeIntegrationConfigCreateAttributes = nil
		}
	} else {
		obj.SecurityMonitoringCrowdStrikeIntegrationConfigCreateAttributes = nil
	}

	// try to unmarshal data into SecurityMonitoringSentinelOneIntegrationConfigCreateAttributes
	err = datadog.Unmarshal(data, &obj.SecurityMonitoringSentinelOneIntegrationConfigCreateAttributes)
	if err == nil {
		if obj.SecurityMonitoringSentinelOneIntegrationConfigCreateAttributes != nil && obj.SecurityMonitoringSentinelOneIntegrationConfigCreateAttributes.UnparsedObject == nil {
			jsonSecurityMonitoringSentinelOneIntegrationConfigCreateAttributes, _ := datadog.Marshal(obj.SecurityMonitoringSentinelOneIntegrationConfigCreateAttributes)
			if string(jsonSecurityMonitoringSentinelOneIntegrationConfigCreateAttributes) == "{}" { // empty struct
				obj.SecurityMonitoringSentinelOneIntegrationConfigCreateAttributes = nil
			} else {
				match++
			}
		} else {
			obj.SecurityMonitoringSentinelOneIntegrationConfigCreateAttributes = nil
		}
	} else {
		obj.SecurityMonitoringSentinelOneIntegrationConfigCreateAttributes = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.SecurityMonitoringGoogleWorkspaceIntegrationConfigCreateAttributes = nil
		obj.SecurityMonitoringOktaIntegrationConfigCreateAttributes = nil
		obj.SecurityMonitoringEntraIdIntegrationConfigCreateAttributes = nil
		obj.SecurityMonitoringCrowdStrikeIntegrationConfigCreateAttributes = nil
		obj.SecurityMonitoringSentinelOneIntegrationConfigCreateAttributes = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj SecurityMonitoringIntegrationConfigCreateAttributes) MarshalJSON() ([]byte, error) {
	if obj.SecurityMonitoringGoogleWorkspaceIntegrationConfigCreateAttributes != nil {
		return datadog.Marshal(&obj.SecurityMonitoringGoogleWorkspaceIntegrationConfigCreateAttributes)
	}

	if obj.SecurityMonitoringOktaIntegrationConfigCreateAttributes != nil {
		return datadog.Marshal(&obj.SecurityMonitoringOktaIntegrationConfigCreateAttributes)
	}

	if obj.SecurityMonitoringEntraIdIntegrationConfigCreateAttributes != nil {
		return datadog.Marshal(&obj.SecurityMonitoringEntraIdIntegrationConfigCreateAttributes)
	}

	if obj.SecurityMonitoringCrowdStrikeIntegrationConfigCreateAttributes != nil {
		return datadog.Marshal(&obj.SecurityMonitoringCrowdStrikeIntegrationConfigCreateAttributes)
	}

	if obj.SecurityMonitoringSentinelOneIntegrationConfigCreateAttributes != nil {
		return datadog.Marshal(&obj.SecurityMonitoringSentinelOneIntegrationConfigCreateAttributes)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *SecurityMonitoringIntegrationConfigCreateAttributes) GetActualInstance() interface{} {
	if obj.SecurityMonitoringGoogleWorkspaceIntegrationConfigCreateAttributes != nil {
		return obj.SecurityMonitoringGoogleWorkspaceIntegrationConfigCreateAttributes
	}

	if obj.SecurityMonitoringOktaIntegrationConfigCreateAttributes != nil {
		return obj.SecurityMonitoringOktaIntegrationConfigCreateAttributes
	}

	if obj.SecurityMonitoringEntraIdIntegrationConfigCreateAttributes != nil {
		return obj.SecurityMonitoringEntraIdIntegrationConfigCreateAttributes
	}

	if obj.SecurityMonitoringCrowdStrikeIntegrationConfigCreateAttributes != nil {
		return obj.SecurityMonitoringCrowdStrikeIntegrationConfigCreateAttributes
	}

	if obj.SecurityMonitoringSentinelOneIntegrationConfigCreateAttributes != nil {
		return obj.SecurityMonitoringSentinelOneIntegrationConfigCreateAttributes
	}

	// all schemas are nil
	return nil
}
