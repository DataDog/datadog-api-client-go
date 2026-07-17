// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

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
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = datadog.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup.")
	}
	// check if the discriminator value is 'CROWDSTRIKE'
	if jsonDict["integration_type"] == "CROWDSTRIKE" {
		// try to unmarshal JSON data into SecurityMonitoringCrowdStrikeIntegrationConfigUpdateAttributes
		err = datadog.Unmarshal(data, &obj.SecurityMonitoringCrowdStrikeIntegrationConfigUpdateAttributes)
		if err == nil {
			return nil // data stored in obj.SecurityMonitoringCrowdStrikeIntegrationConfigUpdateAttributes, return on the first match
		} else {
			obj.SecurityMonitoringCrowdStrikeIntegrationConfigUpdateAttributes = nil
			return fmt.Errorf("failed to unmarshal SecurityMonitoringIntegrationConfigUpdateAttributes as SecurityMonitoringCrowdStrikeIntegrationConfigUpdateAttributes: %s", err.Error())
		}
	}
	// check if the discriminator value is 'ENTRA_ID'
	if jsonDict["integration_type"] == "ENTRA_ID" {
		// try to unmarshal JSON data into SecurityMonitoringEntraIdIntegrationConfigUpdateAttributes
		err = datadog.Unmarshal(data, &obj.SecurityMonitoringEntraIdIntegrationConfigUpdateAttributes)
		if err == nil {
			return nil // data stored in obj.SecurityMonitoringEntraIdIntegrationConfigUpdateAttributes, return on the first match
		} else {
			obj.SecurityMonitoringEntraIdIntegrationConfigUpdateAttributes = nil
			return fmt.Errorf("failed to unmarshal SecurityMonitoringIntegrationConfigUpdateAttributes as SecurityMonitoringEntraIdIntegrationConfigUpdateAttributes: %s", err.Error())
		}
	}
	// check if the discriminator value is 'GOOGLE_WORKSPACE'
	if jsonDict["integration_type"] == "GOOGLE_WORKSPACE" {
		// try to unmarshal JSON data into SecurityMonitoringGoogleWorkspaceIntegrationConfigUpdateAttributes
		err = datadog.Unmarshal(data, &obj.SecurityMonitoringGoogleWorkspaceIntegrationConfigUpdateAttributes)
		if err == nil {
			return nil // data stored in obj.SecurityMonitoringGoogleWorkspaceIntegrationConfigUpdateAttributes, return on the first match
		} else {
			obj.SecurityMonitoringGoogleWorkspaceIntegrationConfigUpdateAttributes = nil
			return fmt.Errorf("failed to unmarshal SecurityMonitoringIntegrationConfigUpdateAttributes as SecurityMonitoringGoogleWorkspaceIntegrationConfigUpdateAttributes: %s", err.Error())
		}
	}
	// check if the discriminator value is 'OKTA'
	if jsonDict["integration_type"] == "OKTA" {
		// try to unmarshal JSON data into SecurityMonitoringOktaIntegrationConfigUpdateAttributes
		err = datadog.Unmarshal(data, &obj.SecurityMonitoringOktaIntegrationConfigUpdateAttributes)
		if err == nil {
			return nil // data stored in obj.SecurityMonitoringOktaIntegrationConfigUpdateAttributes, return on the first match
		} else {
			obj.SecurityMonitoringOktaIntegrationConfigUpdateAttributes = nil
			return fmt.Errorf("failed to unmarshal SecurityMonitoringIntegrationConfigUpdateAttributes as SecurityMonitoringOktaIntegrationConfigUpdateAttributes: %s", err.Error())
		}
	}
	// check if the discriminator value is 'SENTINELONE'
	if jsonDict["integration_type"] == "SENTINELONE" {
		// try to unmarshal JSON data into SecurityMonitoringSentinelOneIntegrationConfigUpdateAttributes
		err = datadog.Unmarshal(data, &obj.SecurityMonitoringSentinelOneIntegrationConfigUpdateAttributes)
		if err == nil {
			return nil // data stored in obj.SecurityMonitoringSentinelOneIntegrationConfigUpdateAttributes, return on the first match
		} else {
			obj.SecurityMonitoringSentinelOneIntegrationConfigUpdateAttributes = nil
			return fmt.Errorf("failed to unmarshal SecurityMonitoringIntegrationConfigUpdateAttributes as SecurityMonitoringSentinelOneIntegrationConfigUpdateAttributes: %s", err.Error())
		}
	}
	return nil
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
