// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringContentPackStateDetails - Type-specific details for a content pack state. The set of fields present depends
// on the content pack's `type`. When Cloud SIEM is inactive for the requesting organization, `onboarding` is returned instead of the content pack's usual type, such as `logs` or `vulnerability`.`
type SecurityMonitoringContentPackStateDetails struct {
	SecurityMonitoringContentPackLogsDetails          *SecurityMonitoringContentPackLogsDetails
	SecurityMonitoringContentPackThreatIntelDetails   *SecurityMonitoringContentPackThreatIntelDetails
	SecurityMonitoringContentPackEntityDetails        *SecurityMonitoringContentPackEntityDetails
	SecurityMonitoringContentPackAuditDetails         *SecurityMonitoringContentPackAuditDetails
	SecurityMonitoringContentPackAppSecDetails        *SecurityMonitoringContentPackAppSecDetails
	SecurityMonitoringContentPackVulnerabilityDetails *SecurityMonitoringContentPackVulnerabilityDetails
	SecurityMonitoringContentPackOnboardingDetails    *SecurityMonitoringContentPackOnboardingDetails

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// SecurityMonitoringContentPackLogsDetailsAsSecurityMonitoringContentPackStateDetails is a convenience function that returns SecurityMonitoringContentPackLogsDetails wrapped in SecurityMonitoringContentPackStateDetails.
func SecurityMonitoringContentPackLogsDetailsAsSecurityMonitoringContentPackStateDetails(v *SecurityMonitoringContentPackLogsDetails) SecurityMonitoringContentPackStateDetails {
	return SecurityMonitoringContentPackStateDetails{SecurityMonitoringContentPackLogsDetails: v}
}

// SecurityMonitoringContentPackThreatIntelDetailsAsSecurityMonitoringContentPackStateDetails is a convenience function that returns SecurityMonitoringContentPackThreatIntelDetails wrapped in SecurityMonitoringContentPackStateDetails.
func SecurityMonitoringContentPackThreatIntelDetailsAsSecurityMonitoringContentPackStateDetails(v *SecurityMonitoringContentPackThreatIntelDetails) SecurityMonitoringContentPackStateDetails {
	return SecurityMonitoringContentPackStateDetails{SecurityMonitoringContentPackThreatIntelDetails: v}
}

// SecurityMonitoringContentPackEntityDetailsAsSecurityMonitoringContentPackStateDetails is a convenience function that returns SecurityMonitoringContentPackEntityDetails wrapped in SecurityMonitoringContentPackStateDetails.
func SecurityMonitoringContentPackEntityDetailsAsSecurityMonitoringContentPackStateDetails(v *SecurityMonitoringContentPackEntityDetails) SecurityMonitoringContentPackStateDetails {
	return SecurityMonitoringContentPackStateDetails{SecurityMonitoringContentPackEntityDetails: v}
}

// SecurityMonitoringContentPackAuditDetailsAsSecurityMonitoringContentPackStateDetails is a convenience function that returns SecurityMonitoringContentPackAuditDetails wrapped in SecurityMonitoringContentPackStateDetails.
func SecurityMonitoringContentPackAuditDetailsAsSecurityMonitoringContentPackStateDetails(v *SecurityMonitoringContentPackAuditDetails) SecurityMonitoringContentPackStateDetails {
	return SecurityMonitoringContentPackStateDetails{SecurityMonitoringContentPackAuditDetails: v}
}

// SecurityMonitoringContentPackAppSecDetailsAsSecurityMonitoringContentPackStateDetails is a convenience function that returns SecurityMonitoringContentPackAppSecDetails wrapped in SecurityMonitoringContentPackStateDetails.
func SecurityMonitoringContentPackAppSecDetailsAsSecurityMonitoringContentPackStateDetails(v *SecurityMonitoringContentPackAppSecDetails) SecurityMonitoringContentPackStateDetails {
	return SecurityMonitoringContentPackStateDetails{SecurityMonitoringContentPackAppSecDetails: v}
}

// SecurityMonitoringContentPackVulnerabilityDetailsAsSecurityMonitoringContentPackStateDetails is a convenience function that returns SecurityMonitoringContentPackVulnerabilityDetails wrapped in SecurityMonitoringContentPackStateDetails.
func SecurityMonitoringContentPackVulnerabilityDetailsAsSecurityMonitoringContentPackStateDetails(v *SecurityMonitoringContentPackVulnerabilityDetails) SecurityMonitoringContentPackStateDetails {
	return SecurityMonitoringContentPackStateDetails{SecurityMonitoringContentPackVulnerabilityDetails: v}
}

// SecurityMonitoringContentPackOnboardingDetailsAsSecurityMonitoringContentPackStateDetails is a convenience function that returns SecurityMonitoringContentPackOnboardingDetails wrapped in SecurityMonitoringContentPackStateDetails.
func SecurityMonitoringContentPackOnboardingDetailsAsSecurityMonitoringContentPackStateDetails(v *SecurityMonitoringContentPackOnboardingDetails) SecurityMonitoringContentPackStateDetails {
	return SecurityMonitoringContentPackStateDetails{SecurityMonitoringContentPackOnboardingDetails: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *SecurityMonitoringContentPackStateDetails) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = datadog.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup.")
	}
	// check if the discriminator value is 'appsec'
	if jsonDict["type"] == "appsec" {
		// try to unmarshal JSON data into SecurityMonitoringContentPackAppSecDetails
		err = datadog.Unmarshal(data, &obj.SecurityMonitoringContentPackAppSecDetails)
		if err == nil {
			return nil // data stored in obj.SecurityMonitoringContentPackAppSecDetails, return on the first match
		} else {
			obj.SecurityMonitoringContentPackAppSecDetails = nil
			return fmt.Errorf("failed to unmarshal SecurityMonitoringContentPackStateDetails as SecurityMonitoringContentPackAppSecDetails: %s", err.Error())
		}
	}
	// check if the discriminator value is 'audit'
	if jsonDict["type"] == "audit" {
		// try to unmarshal JSON data into SecurityMonitoringContentPackAuditDetails
		err = datadog.Unmarshal(data, &obj.SecurityMonitoringContentPackAuditDetails)
		if err == nil {
			return nil // data stored in obj.SecurityMonitoringContentPackAuditDetails, return on the first match
		} else {
			obj.SecurityMonitoringContentPackAuditDetails = nil
			return fmt.Errorf("failed to unmarshal SecurityMonitoringContentPackStateDetails as SecurityMonitoringContentPackAuditDetails: %s", err.Error())
		}
	}
	// check if the discriminator value is 'entity'
	if jsonDict["type"] == "entity" {
		// try to unmarshal JSON data into SecurityMonitoringContentPackEntityDetails
		err = datadog.Unmarshal(data, &obj.SecurityMonitoringContentPackEntityDetails)
		if err == nil {
			return nil // data stored in obj.SecurityMonitoringContentPackEntityDetails, return on the first match
		} else {
			obj.SecurityMonitoringContentPackEntityDetails = nil
			return fmt.Errorf("failed to unmarshal SecurityMonitoringContentPackStateDetails as SecurityMonitoringContentPackEntityDetails: %s", err.Error())
		}
	}
	// check if the discriminator value is 'logs'
	if jsonDict["type"] == "logs" {
		// try to unmarshal JSON data into SecurityMonitoringContentPackLogsDetails
		err = datadog.Unmarshal(data, &obj.SecurityMonitoringContentPackLogsDetails)
		if err == nil {
			return nil // data stored in obj.SecurityMonitoringContentPackLogsDetails, return on the first match
		} else {
			obj.SecurityMonitoringContentPackLogsDetails = nil
			return fmt.Errorf("failed to unmarshal SecurityMonitoringContentPackStateDetails as SecurityMonitoringContentPackLogsDetails: %s", err.Error())
		}
	}
	// check if the discriminator value is 'onboarding'
	if jsonDict["type"] == "onboarding" {
		// try to unmarshal JSON data into SecurityMonitoringContentPackOnboardingDetails
		err = datadog.Unmarshal(data, &obj.SecurityMonitoringContentPackOnboardingDetails)
		if err == nil {
			return nil // data stored in obj.SecurityMonitoringContentPackOnboardingDetails, return on the first match
		} else {
			obj.SecurityMonitoringContentPackOnboardingDetails = nil
			return fmt.Errorf("failed to unmarshal SecurityMonitoringContentPackStateDetails as SecurityMonitoringContentPackOnboardingDetails: %s", err.Error())
		}
	}
	// check if the discriminator value is 'threat_intel'
	if jsonDict["type"] == "threat_intel" {
		// try to unmarshal JSON data into SecurityMonitoringContentPackThreatIntelDetails
		err = datadog.Unmarshal(data, &obj.SecurityMonitoringContentPackThreatIntelDetails)
		if err == nil {
			return nil // data stored in obj.SecurityMonitoringContentPackThreatIntelDetails, return on the first match
		} else {
			obj.SecurityMonitoringContentPackThreatIntelDetails = nil
			return fmt.Errorf("failed to unmarshal SecurityMonitoringContentPackStateDetails as SecurityMonitoringContentPackThreatIntelDetails: %s", err.Error())
		}
	}
	// check if the discriminator value is 'vulnerability'
	if jsonDict["type"] == "vulnerability" {
		// try to unmarshal JSON data into SecurityMonitoringContentPackVulnerabilityDetails
		err = datadog.Unmarshal(data, &obj.SecurityMonitoringContentPackVulnerabilityDetails)
		if err == nil {
			return nil // data stored in obj.SecurityMonitoringContentPackVulnerabilityDetails, return on the first match
		} else {
			obj.SecurityMonitoringContentPackVulnerabilityDetails = nil
			return fmt.Errorf("failed to unmarshal SecurityMonitoringContentPackStateDetails as SecurityMonitoringContentPackVulnerabilityDetails: %s", err.Error())
		}
	}
	return nil
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj SecurityMonitoringContentPackStateDetails) MarshalJSON() ([]byte, error) {
	if obj.SecurityMonitoringContentPackLogsDetails != nil {
		return datadog.Marshal(&obj.SecurityMonitoringContentPackLogsDetails)
	}

	if obj.SecurityMonitoringContentPackThreatIntelDetails != nil {
		return datadog.Marshal(&obj.SecurityMonitoringContentPackThreatIntelDetails)
	}

	if obj.SecurityMonitoringContentPackEntityDetails != nil {
		return datadog.Marshal(&obj.SecurityMonitoringContentPackEntityDetails)
	}

	if obj.SecurityMonitoringContentPackAuditDetails != nil {
		return datadog.Marshal(&obj.SecurityMonitoringContentPackAuditDetails)
	}

	if obj.SecurityMonitoringContentPackAppSecDetails != nil {
		return datadog.Marshal(&obj.SecurityMonitoringContentPackAppSecDetails)
	}

	if obj.SecurityMonitoringContentPackVulnerabilityDetails != nil {
		return datadog.Marshal(&obj.SecurityMonitoringContentPackVulnerabilityDetails)
	}

	if obj.SecurityMonitoringContentPackOnboardingDetails != nil {
		return datadog.Marshal(&obj.SecurityMonitoringContentPackOnboardingDetails)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *SecurityMonitoringContentPackStateDetails) GetActualInstance() interface{} {
	if obj.SecurityMonitoringContentPackLogsDetails != nil {
		return obj.SecurityMonitoringContentPackLogsDetails
	}

	if obj.SecurityMonitoringContentPackThreatIntelDetails != nil {
		return obj.SecurityMonitoringContentPackThreatIntelDetails
	}

	if obj.SecurityMonitoringContentPackEntityDetails != nil {
		return obj.SecurityMonitoringContentPackEntityDetails
	}

	if obj.SecurityMonitoringContentPackAuditDetails != nil {
		return obj.SecurityMonitoringContentPackAuditDetails
	}

	if obj.SecurityMonitoringContentPackAppSecDetails != nil {
		return obj.SecurityMonitoringContentPackAppSecDetails
	}

	if obj.SecurityMonitoringContentPackVulnerabilityDetails != nil {
		return obj.SecurityMonitoringContentPackVulnerabilityDetails
	}

	if obj.SecurityMonitoringContentPackOnboardingDetails != nil {
		return obj.SecurityMonitoringContentPackOnboardingDetails
	}

	// all schemas are nil
	return nil
}
