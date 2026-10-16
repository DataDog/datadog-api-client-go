// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringIntegrationCredentialsValidateAttributes - The credentials to validate against the external entity source.
type SecurityMonitoringIntegrationCredentialsValidateAttributes struct {
	SecurityMonitoringGoogleWorkspaceIntegrationCredentialsValidateAttributes *SecurityMonitoringGoogleWorkspaceIntegrationCredentialsValidateAttributes
	SecurityMonitoringOktaIntegrationCredentialsValidateAttributes            *SecurityMonitoringOktaIntegrationCredentialsValidateAttributes
	SecurityMonitoringEntraIdIntegrationCredentialsValidateAttributes         *SecurityMonitoringEntraIdIntegrationCredentialsValidateAttributes
	SecurityMonitoringCrowdStrikeIntegrationCredentialsValidateAttributes     *SecurityMonitoringCrowdStrikeIntegrationCredentialsValidateAttributes
	SecurityMonitoringSentinelOneIntegrationCredentialsValidateAttributes     *SecurityMonitoringSentinelOneIntegrationCredentialsValidateAttributes

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// SecurityMonitoringGoogleWorkspaceIntegrationCredentialsValidateAttributesAsSecurityMonitoringIntegrationCredentialsValidateAttributes is a convenience function that returns SecurityMonitoringGoogleWorkspaceIntegrationCredentialsValidateAttributes wrapped in SecurityMonitoringIntegrationCredentialsValidateAttributes.
func SecurityMonitoringGoogleWorkspaceIntegrationCredentialsValidateAttributesAsSecurityMonitoringIntegrationCredentialsValidateAttributes(v *SecurityMonitoringGoogleWorkspaceIntegrationCredentialsValidateAttributes) SecurityMonitoringIntegrationCredentialsValidateAttributes {
	return SecurityMonitoringIntegrationCredentialsValidateAttributes{SecurityMonitoringGoogleWorkspaceIntegrationCredentialsValidateAttributes: v}
}

// SecurityMonitoringOktaIntegrationCredentialsValidateAttributesAsSecurityMonitoringIntegrationCredentialsValidateAttributes is a convenience function that returns SecurityMonitoringOktaIntegrationCredentialsValidateAttributes wrapped in SecurityMonitoringIntegrationCredentialsValidateAttributes.
func SecurityMonitoringOktaIntegrationCredentialsValidateAttributesAsSecurityMonitoringIntegrationCredentialsValidateAttributes(v *SecurityMonitoringOktaIntegrationCredentialsValidateAttributes) SecurityMonitoringIntegrationCredentialsValidateAttributes {
	return SecurityMonitoringIntegrationCredentialsValidateAttributes{SecurityMonitoringOktaIntegrationCredentialsValidateAttributes: v}
}

// SecurityMonitoringEntraIdIntegrationCredentialsValidateAttributesAsSecurityMonitoringIntegrationCredentialsValidateAttributes is a convenience function that returns SecurityMonitoringEntraIdIntegrationCredentialsValidateAttributes wrapped in SecurityMonitoringIntegrationCredentialsValidateAttributes.
func SecurityMonitoringEntraIdIntegrationCredentialsValidateAttributesAsSecurityMonitoringIntegrationCredentialsValidateAttributes(v *SecurityMonitoringEntraIdIntegrationCredentialsValidateAttributes) SecurityMonitoringIntegrationCredentialsValidateAttributes {
	return SecurityMonitoringIntegrationCredentialsValidateAttributes{SecurityMonitoringEntraIdIntegrationCredentialsValidateAttributes: v}
}

// SecurityMonitoringCrowdStrikeIntegrationCredentialsValidateAttributesAsSecurityMonitoringIntegrationCredentialsValidateAttributes is a convenience function that returns SecurityMonitoringCrowdStrikeIntegrationCredentialsValidateAttributes wrapped in SecurityMonitoringIntegrationCredentialsValidateAttributes.
func SecurityMonitoringCrowdStrikeIntegrationCredentialsValidateAttributesAsSecurityMonitoringIntegrationCredentialsValidateAttributes(v *SecurityMonitoringCrowdStrikeIntegrationCredentialsValidateAttributes) SecurityMonitoringIntegrationCredentialsValidateAttributes {
	return SecurityMonitoringIntegrationCredentialsValidateAttributes{SecurityMonitoringCrowdStrikeIntegrationCredentialsValidateAttributes: v}
}

// SecurityMonitoringSentinelOneIntegrationCredentialsValidateAttributesAsSecurityMonitoringIntegrationCredentialsValidateAttributes is a convenience function that returns SecurityMonitoringSentinelOneIntegrationCredentialsValidateAttributes wrapped in SecurityMonitoringIntegrationCredentialsValidateAttributes.
func SecurityMonitoringSentinelOneIntegrationCredentialsValidateAttributesAsSecurityMonitoringIntegrationCredentialsValidateAttributes(v *SecurityMonitoringSentinelOneIntegrationCredentialsValidateAttributes) SecurityMonitoringIntegrationCredentialsValidateAttributes {
	return SecurityMonitoringIntegrationCredentialsValidateAttributes{SecurityMonitoringSentinelOneIntegrationCredentialsValidateAttributes: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *SecurityMonitoringIntegrationCredentialsValidateAttributes) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SecurityMonitoringGoogleWorkspaceIntegrationCredentialsValidateAttributes
	err = datadog.Unmarshal(data, &obj.SecurityMonitoringGoogleWorkspaceIntegrationCredentialsValidateAttributes)
	if err == nil {
		if obj.SecurityMonitoringGoogleWorkspaceIntegrationCredentialsValidateAttributes != nil && obj.SecurityMonitoringGoogleWorkspaceIntegrationCredentialsValidateAttributes.UnparsedObject == nil {
			jsonSecurityMonitoringGoogleWorkspaceIntegrationCredentialsValidateAttributes, _ := datadog.Marshal(obj.SecurityMonitoringGoogleWorkspaceIntegrationCredentialsValidateAttributes)
			if string(jsonSecurityMonitoringGoogleWorkspaceIntegrationCredentialsValidateAttributes) == "{}" { // empty struct
				obj.SecurityMonitoringGoogleWorkspaceIntegrationCredentialsValidateAttributes = nil
			} else {
				match++
			}
		} else {
			obj.SecurityMonitoringGoogleWorkspaceIntegrationCredentialsValidateAttributes = nil
		}
	} else {
		obj.SecurityMonitoringGoogleWorkspaceIntegrationCredentialsValidateAttributes = nil
	}

	// try to unmarshal data into SecurityMonitoringOktaIntegrationCredentialsValidateAttributes
	err = datadog.Unmarshal(data, &obj.SecurityMonitoringOktaIntegrationCredentialsValidateAttributes)
	if err == nil {
		if obj.SecurityMonitoringOktaIntegrationCredentialsValidateAttributes != nil && obj.SecurityMonitoringOktaIntegrationCredentialsValidateAttributes.UnparsedObject == nil {
			jsonSecurityMonitoringOktaIntegrationCredentialsValidateAttributes, _ := datadog.Marshal(obj.SecurityMonitoringOktaIntegrationCredentialsValidateAttributes)
			if string(jsonSecurityMonitoringOktaIntegrationCredentialsValidateAttributes) == "{}" { // empty struct
				obj.SecurityMonitoringOktaIntegrationCredentialsValidateAttributes = nil
			} else {
				match++
			}
		} else {
			obj.SecurityMonitoringOktaIntegrationCredentialsValidateAttributes = nil
		}
	} else {
		obj.SecurityMonitoringOktaIntegrationCredentialsValidateAttributes = nil
	}

	// try to unmarshal data into SecurityMonitoringEntraIdIntegrationCredentialsValidateAttributes
	err = datadog.Unmarshal(data, &obj.SecurityMonitoringEntraIdIntegrationCredentialsValidateAttributes)
	if err == nil {
		if obj.SecurityMonitoringEntraIdIntegrationCredentialsValidateAttributes != nil && obj.SecurityMonitoringEntraIdIntegrationCredentialsValidateAttributes.UnparsedObject == nil {
			jsonSecurityMonitoringEntraIdIntegrationCredentialsValidateAttributes, _ := datadog.Marshal(obj.SecurityMonitoringEntraIdIntegrationCredentialsValidateAttributes)
			if string(jsonSecurityMonitoringEntraIdIntegrationCredentialsValidateAttributes) == "{}" { // empty struct
				obj.SecurityMonitoringEntraIdIntegrationCredentialsValidateAttributes = nil
			} else {
				match++
			}
		} else {
			obj.SecurityMonitoringEntraIdIntegrationCredentialsValidateAttributes = nil
		}
	} else {
		obj.SecurityMonitoringEntraIdIntegrationCredentialsValidateAttributes = nil
	}

	// try to unmarshal data into SecurityMonitoringCrowdStrikeIntegrationCredentialsValidateAttributes
	err = datadog.Unmarshal(data, &obj.SecurityMonitoringCrowdStrikeIntegrationCredentialsValidateAttributes)
	if err == nil {
		if obj.SecurityMonitoringCrowdStrikeIntegrationCredentialsValidateAttributes != nil && obj.SecurityMonitoringCrowdStrikeIntegrationCredentialsValidateAttributes.UnparsedObject == nil {
			jsonSecurityMonitoringCrowdStrikeIntegrationCredentialsValidateAttributes, _ := datadog.Marshal(obj.SecurityMonitoringCrowdStrikeIntegrationCredentialsValidateAttributes)
			if string(jsonSecurityMonitoringCrowdStrikeIntegrationCredentialsValidateAttributes) == "{}" { // empty struct
				obj.SecurityMonitoringCrowdStrikeIntegrationCredentialsValidateAttributes = nil
			} else {
				match++
			}
		} else {
			obj.SecurityMonitoringCrowdStrikeIntegrationCredentialsValidateAttributes = nil
		}
	} else {
		obj.SecurityMonitoringCrowdStrikeIntegrationCredentialsValidateAttributes = nil
	}

	// try to unmarshal data into SecurityMonitoringSentinelOneIntegrationCredentialsValidateAttributes
	err = datadog.Unmarshal(data, &obj.SecurityMonitoringSentinelOneIntegrationCredentialsValidateAttributes)
	if err == nil {
		if obj.SecurityMonitoringSentinelOneIntegrationCredentialsValidateAttributes != nil && obj.SecurityMonitoringSentinelOneIntegrationCredentialsValidateAttributes.UnparsedObject == nil {
			jsonSecurityMonitoringSentinelOneIntegrationCredentialsValidateAttributes, _ := datadog.Marshal(obj.SecurityMonitoringSentinelOneIntegrationCredentialsValidateAttributes)
			if string(jsonSecurityMonitoringSentinelOneIntegrationCredentialsValidateAttributes) == "{}" { // empty struct
				obj.SecurityMonitoringSentinelOneIntegrationCredentialsValidateAttributes = nil
			} else {
				match++
			}
		} else {
			obj.SecurityMonitoringSentinelOneIntegrationCredentialsValidateAttributes = nil
		}
	} else {
		obj.SecurityMonitoringSentinelOneIntegrationCredentialsValidateAttributes = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.SecurityMonitoringGoogleWorkspaceIntegrationCredentialsValidateAttributes = nil
		obj.SecurityMonitoringOktaIntegrationCredentialsValidateAttributes = nil
		obj.SecurityMonitoringEntraIdIntegrationCredentialsValidateAttributes = nil
		obj.SecurityMonitoringCrowdStrikeIntegrationCredentialsValidateAttributes = nil
		obj.SecurityMonitoringSentinelOneIntegrationCredentialsValidateAttributes = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj SecurityMonitoringIntegrationCredentialsValidateAttributes) MarshalJSON() ([]byte, error) {
	if obj.SecurityMonitoringGoogleWorkspaceIntegrationCredentialsValidateAttributes != nil {
		return datadog.Marshal(&obj.SecurityMonitoringGoogleWorkspaceIntegrationCredentialsValidateAttributes)
	}

	if obj.SecurityMonitoringOktaIntegrationCredentialsValidateAttributes != nil {
		return datadog.Marshal(&obj.SecurityMonitoringOktaIntegrationCredentialsValidateAttributes)
	}

	if obj.SecurityMonitoringEntraIdIntegrationCredentialsValidateAttributes != nil {
		return datadog.Marshal(&obj.SecurityMonitoringEntraIdIntegrationCredentialsValidateAttributes)
	}

	if obj.SecurityMonitoringCrowdStrikeIntegrationCredentialsValidateAttributes != nil {
		return datadog.Marshal(&obj.SecurityMonitoringCrowdStrikeIntegrationCredentialsValidateAttributes)
	}

	if obj.SecurityMonitoringSentinelOneIntegrationCredentialsValidateAttributes != nil {
		return datadog.Marshal(&obj.SecurityMonitoringSentinelOneIntegrationCredentialsValidateAttributes)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *SecurityMonitoringIntegrationCredentialsValidateAttributes) GetActualInstance() interface{} {
	if obj.SecurityMonitoringGoogleWorkspaceIntegrationCredentialsValidateAttributes != nil {
		return obj.SecurityMonitoringGoogleWorkspaceIntegrationCredentialsValidateAttributes
	}

	if obj.SecurityMonitoringOktaIntegrationCredentialsValidateAttributes != nil {
		return obj.SecurityMonitoringOktaIntegrationCredentialsValidateAttributes
	}

	if obj.SecurityMonitoringEntraIdIntegrationCredentialsValidateAttributes != nil {
		return obj.SecurityMonitoringEntraIdIntegrationCredentialsValidateAttributes
	}

	if obj.SecurityMonitoringCrowdStrikeIntegrationCredentialsValidateAttributes != nil {
		return obj.SecurityMonitoringCrowdStrikeIntegrationCredentialsValidateAttributes
	}

	if obj.SecurityMonitoringSentinelOneIntegrationCredentialsValidateAttributes != nil {
		return obj.SecurityMonitoringSentinelOneIntegrationCredentialsValidateAttributes
	}

	// all schemas are nil
	return nil
}
