// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringIntegrationConfigUpdateAttributes - Fields to update on the entity context sync configuration. All fields other than the integration type are optional.
type SecurityMonitoringIntegrationConfigUpdateAttributes struct {
	SecurityMonitoringGoogleWorkspaceIntegrationConfigUpdateAttributes *SecurityMonitoringGoogleWorkspaceIntegrationConfigUpdateAttributes
	SecurityMonitoringOktaIntegrationConfigUpdateAttributes            *SecurityMonitoringOktaIntegrationConfigUpdateAttributes
	SecurityMonitoringEntraIdIntegrationConfigUpdateAttributes         *SecurityMonitoringEntraIdIntegrationConfigUpdateAttributes
	SecurityMonitoringCrowdStrikeIntegrationConfigUpdateAttributes     *SecurityMonitoringCrowdStrikeIntegrationConfigUpdateAttributes
	SecurityMonitoringSentinelOneIntegrationConfigUpdateAttributes     *SecurityMonitoringSentinelOneIntegrationConfigUpdateAttributes

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// SecurityMonitoringGoogleWorkspaceIntegrationConfigUpdateAttributesAsSecurityMonitoringIntegrationConfigUpdateAttributes is a convenience function that returns SecurityMonitoringGoogleWorkspaceIntegrationConfigUpdateAttributes wrapped in SecurityMonitoringIntegrationConfigUpdateAttributes.
func SecurityMonitoringGoogleWorkspaceIntegrationConfigUpdateAttributesAsSecurityMonitoringIntegrationConfigUpdateAttributes(v *SecurityMonitoringGoogleWorkspaceIntegrationConfigUpdateAttributes) SecurityMonitoringIntegrationConfigUpdateAttributes {
	return SecurityMonitoringIntegrationConfigUpdateAttributes{SecurityMonitoringGoogleWorkspaceIntegrationConfigUpdateAttributes: v}
}

// SecurityMonitoringOktaIntegrationConfigUpdateAttributesAsSecurityMonitoringIntegrationConfigUpdateAttributes is a convenience function that returns SecurityMonitoringOktaIntegrationConfigUpdateAttributes wrapped in SecurityMonitoringIntegrationConfigUpdateAttributes.
func SecurityMonitoringOktaIntegrationConfigUpdateAttributesAsSecurityMonitoringIntegrationConfigUpdateAttributes(v *SecurityMonitoringOktaIntegrationConfigUpdateAttributes) SecurityMonitoringIntegrationConfigUpdateAttributes {
	return SecurityMonitoringIntegrationConfigUpdateAttributes{SecurityMonitoringOktaIntegrationConfigUpdateAttributes: v}
}

// SecurityMonitoringEntraIdIntegrationConfigUpdateAttributesAsSecurityMonitoringIntegrationConfigUpdateAttributes is a convenience function that returns SecurityMonitoringEntraIdIntegrationConfigUpdateAttributes wrapped in SecurityMonitoringIntegrationConfigUpdateAttributes.
func SecurityMonitoringEntraIdIntegrationConfigUpdateAttributesAsSecurityMonitoringIntegrationConfigUpdateAttributes(v *SecurityMonitoringEntraIdIntegrationConfigUpdateAttributes) SecurityMonitoringIntegrationConfigUpdateAttributes {
	return SecurityMonitoringIntegrationConfigUpdateAttributes{SecurityMonitoringEntraIdIntegrationConfigUpdateAttributes: v}
}

// SecurityMonitoringCrowdStrikeIntegrationConfigUpdateAttributesAsSecurityMonitoringIntegrationConfigUpdateAttributes is a convenience function that returns SecurityMonitoringCrowdStrikeIntegrationConfigUpdateAttributes wrapped in SecurityMonitoringIntegrationConfigUpdateAttributes.
func SecurityMonitoringCrowdStrikeIntegrationConfigUpdateAttributesAsSecurityMonitoringIntegrationConfigUpdateAttributes(v *SecurityMonitoringCrowdStrikeIntegrationConfigUpdateAttributes) SecurityMonitoringIntegrationConfigUpdateAttributes {
	return SecurityMonitoringIntegrationConfigUpdateAttributes{SecurityMonitoringCrowdStrikeIntegrationConfigUpdateAttributes: v}
}

// SecurityMonitoringSentinelOneIntegrationConfigUpdateAttributesAsSecurityMonitoringIntegrationConfigUpdateAttributes is a convenience function that returns SecurityMonitoringSentinelOneIntegrationConfigUpdateAttributes wrapped in SecurityMonitoringIntegrationConfigUpdateAttributes.
func SecurityMonitoringSentinelOneIntegrationConfigUpdateAttributesAsSecurityMonitoringIntegrationConfigUpdateAttributes(v *SecurityMonitoringSentinelOneIntegrationConfigUpdateAttributes) SecurityMonitoringIntegrationConfigUpdateAttributes {
	return SecurityMonitoringIntegrationConfigUpdateAttributes{SecurityMonitoringSentinelOneIntegrationConfigUpdateAttributes: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *SecurityMonitoringIntegrationConfigUpdateAttributes) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SecurityMonitoringGoogleWorkspaceIntegrationConfigUpdateAttributes
	err = datadog.Unmarshal(data, &obj.SecurityMonitoringGoogleWorkspaceIntegrationConfigUpdateAttributes)
	if err == nil {
		if obj.SecurityMonitoringGoogleWorkspaceIntegrationConfigUpdateAttributes != nil && obj.SecurityMonitoringGoogleWorkspaceIntegrationConfigUpdateAttributes.UnparsedObject == nil {
			jsonSecurityMonitoringGoogleWorkspaceIntegrationConfigUpdateAttributes, _ := datadog.Marshal(obj.SecurityMonitoringGoogleWorkspaceIntegrationConfigUpdateAttributes)
			if string(jsonSecurityMonitoringGoogleWorkspaceIntegrationConfigUpdateAttributes) == "{}" { // empty struct
				obj.SecurityMonitoringGoogleWorkspaceIntegrationConfigUpdateAttributes = nil
			} else {
				match++
			}
		} else {
			obj.SecurityMonitoringGoogleWorkspaceIntegrationConfigUpdateAttributes = nil
		}
	} else {
		obj.SecurityMonitoringGoogleWorkspaceIntegrationConfigUpdateAttributes = nil
	}

	// try to unmarshal data into SecurityMonitoringOktaIntegrationConfigUpdateAttributes
	err = datadog.Unmarshal(data, &obj.SecurityMonitoringOktaIntegrationConfigUpdateAttributes)
	if err == nil {
		if obj.SecurityMonitoringOktaIntegrationConfigUpdateAttributes != nil && obj.SecurityMonitoringOktaIntegrationConfigUpdateAttributes.UnparsedObject == nil {
			jsonSecurityMonitoringOktaIntegrationConfigUpdateAttributes, _ := datadog.Marshal(obj.SecurityMonitoringOktaIntegrationConfigUpdateAttributes)
			if string(jsonSecurityMonitoringOktaIntegrationConfigUpdateAttributes) == "{}" { // empty struct
				obj.SecurityMonitoringOktaIntegrationConfigUpdateAttributes = nil
			} else {
				match++
			}
		} else {
			obj.SecurityMonitoringOktaIntegrationConfigUpdateAttributes = nil
		}
	} else {
		obj.SecurityMonitoringOktaIntegrationConfigUpdateAttributes = nil
	}

	// try to unmarshal data into SecurityMonitoringEntraIdIntegrationConfigUpdateAttributes
	err = datadog.Unmarshal(data, &obj.SecurityMonitoringEntraIdIntegrationConfigUpdateAttributes)
	if err == nil {
		if obj.SecurityMonitoringEntraIdIntegrationConfigUpdateAttributes != nil && obj.SecurityMonitoringEntraIdIntegrationConfigUpdateAttributes.UnparsedObject == nil {
			jsonSecurityMonitoringEntraIdIntegrationConfigUpdateAttributes, _ := datadog.Marshal(obj.SecurityMonitoringEntraIdIntegrationConfigUpdateAttributes)
			if string(jsonSecurityMonitoringEntraIdIntegrationConfigUpdateAttributes) == "{}" { // empty struct
				obj.SecurityMonitoringEntraIdIntegrationConfigUpdateAttributes = nil
			} else {
				match++
			}
		} else {
			obj.SecurityMonitoringEntraIdIntegrationConfigUpdateAttributes = nil
		}
	} else {
		obj.SecurityMonitoringEntraIdIntegrationConfigUpdateAttributes = nil
	}

	// try to unmarshal data into SecurityMonitoringCrowdStrikeIntegrationConfigUpdateAttributes
	err = datadog.Unmarshal(data, &obj.SecurityMonitoringCrowdStrikeIntegrationConfigUpdateAttributes)
	if err == nil {
		if obj.SecurityMonitoringCrowdStrikeIntegrationConfigUpdateAttributes != nil && obj.SecurityMonitoringCrowdStrikeIntegrationConfigUpdateAttributes.UnparsedObject == nil {
			jsonSecurityMonitoringCrowdStrikeIntegrationConfigUpdateAttributes, _ := datadog.Marshal(obj.SecurityMonitoringCrowdStrikeIntegrationConfigUpdateAttributes)
			if string(jsonSecurityMonitoringCrowdStrikeIntegrationConfigUpdateAttributes) == "{}" { // empty struct
				obj.SecurityMonitoringCrowdStrikeIntegrationConfigUpdateAttributes = nil
			} else {
				match++
			}
		} else {
			obj.SecurityMonitoringCrowdStrikeIntegrationConfigUpdateAttributes = nil
		}
	} else {
		obj.SecurityMonitoringCrowdStrikeIntegrationConfigUpdateAttributes = nil
	}

	// try to unmarshal data into SecurityMonitoringSentinelOneIntegrationConfigUpdateAttributes
	err = datadog.Unmarshal(data, &obj.SecurityMonitoringSentinelOneIntegrationConfigUpdateAttributes)
	if err == nil {
		if obj.SecurityMonitoringSentinelOneIntegrationConfigUpdateAttributes != nil && obj.SecurityMonitoringSentinelOneIntegrationConfigUpdateAttributes.UnparsedObject == nil {
			jsonSecurityMonitoringSentinelOneIntegrationConfigUpdateAttributes, _ := datadog.Marshal(obj.SecurityMonitoringSentinelOneIntegrationConfigUpdateAttributes)
			if string(jsonSecurityMonitoringSentinelOneIntegrationConfigUpdateAttributes) == "{}" { // empty struct
				obj.SecurityMonitoringSentinelOneIntegrationConfigUpdateAttributes = nil
			} else {
				match++
			}
		} else {
			obj.SecurityMonitoringSentinelOneIntegrationConfigUpdateAttributes = nil
		}
	} else {
		obj.SecurityMonitoringSentinelOneIntegrationConfigUpdateAttributes = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.SecurityMonitoringGoogleWorkspaceIntegrationConfigUpdateAttributes = nil
		obj.SecurityMonitoringOktaIntegrationConfigUpdateAttributes = nil
		obj.SecurityMonitoringEntraIdIntegrationConfigUpdateAttributes = nil
		obj.SecurityMonitoringCrowdStrikeIntegrationConfigUpdateAttributes = nil
		obj.SecurityMonitoringSentinelOneIntegrationConfigUpdateAttributes = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj SecurityMonitoringIntegrationConfigUpdateAttributes) MarshalJSON() ([]byte, error) {
	if obj.SecurityMonitoringGoogleWorkspaceIntegrationConfigUpdateAttributes != nil {
		return datadog.Marshal(&obj.SecurityMonitoringGoogleWorkspaceIntegrationConfigUpdateAttributes)
	}

	if obj.SecurityMonitoringOktaIntegrationConfigUpdateAttributes != nil {
		return datadog.Marshal(&obj.SecurityMonitoringOktaIntegrationConfigUpdateAttributes)
	}

	if obj.SecurityMonitoringEntraIdIntegrationConfigUpdateAttributes != nil {
		return datadog.Marshal(&obj.SecurityMonitoringEntraIdIntegrationConfigUpdateAttributes)
	}

	if obj.SecurityMonitoringCrowdStrikeIntegrationConfigUpdateAttributes != nil {
		return datadog.Marshal(&obj.SecurityMonitoringCrowdStrikeIntegrationConfigUpdateAttributes)
	}

	if obj.SecurityMonitoringSentinelOneIntegrationConfigUpdateAttributes != nil {
		return datadog.Marshal(&obj.SecurityMonitoringSentinelOneIntegrationConfigUpdateAttributes)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *SecurityMonitoringIntegrationConfigUpdateAttributes) GetActualInstance() interface{} {
	if obj.SecurityMonitoringGoogleWorkspaceIntegrationConfigUpdateAttributes != nil {
		return obj.SecurityMonitoringGoogleWorkspaceIntegrationConfigUpdateAttributes
	}

	if obj.SecurityMonitoringOktaIntegrationConfigUpdateAttributes != nil {
		return obj.SecurityMonitoringOktaIntegrationConfigUpdateAttributes
	}

	if obj.SecurityMonitoringEntraIdIntegrationConfigUpdateAttributes != nil {
		return obj.SecurityMonitoringEntraIdIntegrationConfigUpdateAttributes
	}

	if obj.SecurityMonitoringCrowdStrikeIntegrationConfigUpdateAttributes != nil {
		return obj.SecurityMonitoringCrowdStrikeIntegrationConfigUpdateAttributes
	}

	if obj.SecurityMonitoringSentinelOneIntegrationConfigUpdateAttributes != nil {
		return obj.SecurityMonitoringSentinelOneIntegrationConfigUpdateAttributes
	}

	// all schemas are nil
	return nil
}
