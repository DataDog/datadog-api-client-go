// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

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
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = datadog.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup.")
	}
	// check if the discriminator value is 'CROWDSTRIKE'
	if jsonDict["integration_type"] == "CROWDSTRIKE" {
		// try to unmarshal JSON data into SecurityMonitoringCrowdStrikeIntegrationConfigCreateAttributes
		err = datadog.Unmarshal(data, &obj.SecurityMonitoringCrowdStrikeIntegrationConfigCreateAttributes)
		if err == nil {
			return nil // data stored in obj.SecurityMonitoringCrowdStrikeIntegrationConfigCreateAttributes, return on the first match
		} else {
			obj.SecurityMonitoringCrowdStrikeIntegrationConfigCreateAttributes = nil
			return fmt.Errorf("failed to unmarshal SecurityMonitoringIntegrationConfigCreateAttributes as SecurityMonitoringCrowdStrikeIntegrationConfigCreateAttributes: %s", err.Error())
		}
	}
	// check if the discriminator value is 'ENTRA_ID'
	if jsonDict["integration_type"] == "ENTRA_ID" {
		// try to unmarshal JSON data into SecurityMonitoringEntraIdIntegrationConfigCreateAttributes
		err = datadog.Unmarshal(data, &obj.SecurityMonitoringEntraIdIntegrationConfigCreateAttributes)
		if err == nil {
			return nil // data stored in obj.SecurityMonitoringEntraIdIntegrationConfigCreateAttributes, return on the first match
		} else {
			obj.SecurityMonitoringEntraIdIntegrationConfigCreateAttributes = nil
			return fmt.Errorf("failed to unmarshal SecurityMonitoringIntegrationConfigCreateAttributes as SecurityMonitoringEntraIdIntegrationConfigCreateAttributes: %s", err.Error())
		}
	}
	// check if the discriminator value is 'GOOGLE_WORKSPACE'
	if jsonDict["integration_type"] == "GOOGLE_WORKSPACE" {
		// try to unmarshal JSON data into SecurityMonitoringGoogleWorkspaceIntegrationConfigCreateAttributes
		err = datadog.Unmarshal(data, &obj.SecurityMonitoringGoogleWorkspaceIntegrationConfigCreateAttributes)
		if err == nil {
			return nil // data stored in obj.SecurityMonitoringGoogleWorkspaceIntegrationConfigCreateAttributes, return on the first match
		} else {
			obj.SecurityMonitoringGoogleWorkspaceIntegrationConfigCreateAttributes = nil
			return fmt.Errorf("failed to unmarshal SecurityMonitoringIntegrationConfigCreateAttributes as SecurityMonitoringGoogleWorkspaceIntegrationConfigCreateAttributes: %s", err.Error())
		}
	}
	// check if the discriminator value is 'OKTA'
	if jsonDict["integration_type"] == "OKTA" {
		// try to unmarshal JSON data into SecurityMonitoringOktaIntegrationConfigCreateAttributes
		err = datadog.Unmarshal(data, &obj.SecurityMonitoringOktaIntegrationConfigCreateAttributes)
		if err == nil {
			return nil // data stored in obj.SecurityMonitoringOktaIntegrationConfigCreateAttributes, return on the first match
		} else {
			obj.SecurityMonitoringOktaIntegrationConfigCreateAttributes = nil
			return fmt.Errorf("failed to unmarshal SecurityMonitoringIntegrationConfigCreateAttributes as SecurityMonitoringOktaIntegrationConfigCreateAttributes: %s", err.Error())
		}
	}
	// check if the discriminator value is 'SENTINELONE'
	if jsonDict["integration_type"] == "SENTINELONE" {
		// try to unmarshal JSON data into SecurityMonitoringSentinelOneIntegrationConfigCreateAttributes
		err = datadog.Unmarshal(data, &obj.SecurityMonitoringSentinelOneIntegrationConfigCreateAttributes)
		if err == nil {
			return nil // data stored in obj.SecurityMonitoringSentinelOneIntegrationConfigCreateAttributes, return on the first match
		} else {
			obj.SecurityMonitoringSentinelOneIntegrationConfigCreateAttributes = nil
			return fmt.Errorf("failed to unmarshal SecurityMonitoringIntegrationConfigCreateAttributes as SecurityMonitoringSentinelOneIntegrationConfigCreateAttributes: %s", err.Error())
		}
	}
	return nil
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
