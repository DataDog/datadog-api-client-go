// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

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
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = datadog.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup.")
	}
	// check if the discriminator value is 'CROWDSTRIKE'
	if jsonDict["integration_type"] == "CROWDSTRIKE" {
		// try to unmarshal JSON data into SecurityMonitoringCrowdStrikeIntegrationCredentialsValidateAttributes
		err = datadog.Unmarshal(data, &obj.SecurityMonitoringCrowdStrikeIntegrationCredentialsValidateAttributes)
		if err == nil {
			return nil // data stored in obj.SecurityMonitoringCrowdStrikeIntegrationCredentialsValidateAttributes, return on the first match
		} else {
			obj.SecurityMonitoringCrowdStrikeIntegrationCredentialsValidateAttributes = nil
			return fmt.Errorf("failed to unmarshal SecurityMonitoringIntegrationCredentialsValidateAttributes as SecurityMonitoringCrowdStrikeIntegrationCredentialsValidateAttributes: %s", err.Error())
		}
	}
	// check if the discriminator value is 'ENTRA_ID'
	if jsonDict["integration_type"] == "ENTRA_ID" {
		// try to unmarshal JSON data into SecurityMonitoringEntraIdIntegrationCredentialsValidateAttributes
		err = datadog.Unmarshal(data, &obj.SecurityMonitoringEntraIdIntegrationCredentialsValidateAttributes)
		if err == nil {
			return nil // data stored in obj.SecurityMonitoringEntraIdIntegrationCredentialsValidateAttributes, return on the first match
		} else {
			obj.SecurityMonitoringEntraIdIntegrationCredentialsValidateAttributes = nil
			return fmt.Errorf("failed to unmarshal SecurityMonitoringIntegrationCredentialsValidateAttributes as SecurityMonitoringEntraIdIntegrationCredentialsValidateAttributes: %s", err.Error())
		}
	}
	// check if the discriminator value is 'GOOGLE_WORKSPACE'
	if jsonDict["integration_type"] == "GOOGLE_WORKSPACE" {
		// try to unmarshal JSON data into SecurityMonitoringGoogleWorkspaceIntegrationCredentialsValidateAttributes
		err = datadog.Unmarshal(data, &obj.SecurityMonitoringGoogleWorkspaceIntegrationCredentialsValidateAttributes)
		if err == nil {
			return nil // data stored in obj.SecurityMonitoringGoogleWorkspaceIntegrationCredentialsValidateAttributes, return on the first match
		} else {
			obj.SecurityMonitoringGoogleWorkspaceIntegrationCredentialsValidateAttributes = nil
			return fmt.Errorf("failed to unmarshal SecurityMonitoringIntegrationCredentialsValidateAttributes as SecurityMonitoringGoogleWorkspaceIntegrationCredentialsValidateAttributes: %s", err.Error())
		}
	}
	// check if the discriminator value is 'OKTA'
	if jsonDict["integration_type"] == "OKTA" {
		// try to unmarshal JSON data into SecurityMonitoringOktaIntegrationCredentialsValidateAttributes
		err = datadog.Unmarshal(data, &obj.SecurityMonitoringOktaIntegrationCredentialsValidateAttributes)
		if err == nil {
			return nil // data stored in obj.SecurityMonitoringOktaIntegrationCredentialsValidateAttributes, return on the first match
		} else {
			obj.SecurityMonitoringOktaIntegrationCredentialsValidateAttributes = nil
			return fmt.Errorf("failed to unmarshal SecurityMonitoringIntegrationCredentialsValidateAttributes as SecurityMonitoringOktaIntegrationCredentialsValidateAttributes: %s", err.Error())
		}
	}
	// check if the discriminator value is 'SENTINELONE'
	if jsonDict["integration_type"] == "SENTINELONE" {
		// try to unmarshal JSON data into SecurityMonitoringSentinelOneIntegrationCredentialsValidateAttributes
		err = datadog.Unmarshal(data, &obj.SecurityMonitoringSentinelOneIntegrationCredentialsValidateAttributes)
		if err == nil {
			return nil // data stored in obj.SecurityMonitoringSentinelOneIntegrationCredentialsValidateAttributes, return on the first match
		} else {
			obj.SecurityMonitoringSentinelOneIntegrationCredentialsValidateAttributes = nil
			return fmt.Errorf("failed to unmarshal SecurityMonitoringIntegrationCredentialsValidateAttributes as SecurityMonitoringSentinelOneIntegrationCredentialsValidateAttributes: %s", err.Error())
		}
	}
	return nil
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
