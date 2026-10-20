// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
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
	match := 0
	// try to unmarshal data into SecurityMonitoringContentPackLogsDetails
	err = datadog.Unmarshal(data, &obj.SecurityMonitoringContentPackLogsDetails)
	if err == nil {
		if obj.SecurityMonitoringContentPackLogsDetails != nil && obj.SecurityMonitoringContentPackLogsDetails.UnparsedObject == nil {
			jsonSecurityMonitoringContentPackLogsDetails, _ := datadog.Marshal(obj.SecurityMonitoringContentPackLogsDetails)
			if string(jsonSecurityMonitoringContentPackLogsDetails) == "{}" { // empty struct
				obj.SecurityMonitoringContentPackLogsDetails = nil
			} else {
				match++
			}
		} else {
			obj.SecurityMonitoringContentPackLogsDetails = nil
		}
	} else {
		obj.SecurityMonitoringContentPackLogsDetails = nil
	}

	// try to unmarshal data into SecurityMonitoringContentPackThreatIntelDetails
	err = datadog.Unmarshal(data, &obj.SecurityMonitoringContentPackThreatIntelDetails)
	if err == nil {
		if obj.SecurityMonitoringContentPackThreatIntelDetails != nil && obj.SecurityMonitoringContentPackThreatIntelDetails.UnparsedObject == nil {
			jsonSecurityMonitoringContentPackThreatIntelDetails, _ := datadog.Marshal(obj.SecurityMonitoringContentPackThreatIntelDetails)
			if string(jsonSecurityMonitoringContentPackThreatIntelDetails) == "{}" { // empty struct
				obj.SecurityMonitoringContentPackThreatIntelDetails = nil
			} else {
				match++
			}
		} else {
			obj.SecurityMonitoringContentPackThreatIntelDetails = nil
		}
	} else {
		obj.SecurityMonitoringContentPackThreatIntelDetails = nil
	}

	// try to unmarshal data into SecurityMonitoringContentPackEntityDetails
	err = datadog.Unmarshal(data, &obj.SecurityMonitoringContentPackEntityDetails)
	if err == nil {
		if obj.SecurityMonitoringContentPackEntityDetails != nil && obj.SecurityMonitoringContentPackEntityDetails.UnparsedObject == nil {
			jsonSecurityMonitoringContentPackEntityDetails, _ := datadog.Marshal(obj.SecurityMonitoringContentPackEntityDetails)
			if string(jsonSecurityMonitoringContentPackEntityDetails) == "{}" { // empty struct
				obj.SecurityMonitoringContentPackEntityDetails = nil
			} else {
				match++
			}
		} else {
			obj.SecurityMonitoringContentPackEntityDetails = nil
		}
	} else {
		obj.SecurityMonitoringContentPackEntityDetails = nil
	}

	// try to unmarshal data into SecurityMonitoringContentPackAuditDetails
	err = datadog.Unmarshal(data, &obj.SecurityMonitoringContentPackAuditDetails)
	if err == nil {
		if obj.SecurityMonitoringContentPackAuditDetails != nil && obj.SecurityMonitoringContentPackAuditDetails.UnparsedObject == nil {
			jsonSecurityMonitoringContentPackAuditDetails, _ := datadog.Marshal(obj.SecurityMonitoringContentPackAuditDetails)
			if string(jsonSecurityMonitoringContentPackAuditDetails) == "{}" { // empty struct
				obj.SecurityMonitoringContentPackAuditDetails = nil
			} else {
				match++
			}
		} else {
			obj.SecurityMonitoringContentPackAuditDetails = nil
		}
	} else {
		obj.SecurityMonitoringContentPackAuditDetails = nil
	}

	// try to unmarshal data into SecurityMonitoringContentPackAppSecDetails
	err = datadog.Unmarshal(data, &obj.SecurityMonitoringContentPackAppSecDetails)
	if err == nil {
		if obj.SecurityMonitoringContentPackAppSecDetails != nil && obj.SecurityMonitoringContentPackAppSecDetails.UnparsedObject == nil {
			jsonSecurityMonitoringContentPackAppSecDetails, _ := datadog.Marshal(obj.SecurityMonitoringContentPackAppSecDetails)
			if string(jsonSecurityMonitoringContentPackAppSecDetails) == "{}" { // empty struct
				obj.SecurityMonitoringContentPackAppSecDetails = nil
			} else {
				match++
			}
		} else {
			obj.SecurityMonitoringContentPackAppSecDetails = nil
		}
	} else {
		obj.SecurityMonitoringContentPackAppSecDetails = nil
	}

	// try to unmarshal data into SecurityMonitoringContentPackVulnerabilityDetails
	err = datadog.Unmarshal(data, &obj.SecurityMonitoringContentPackVulnerabilityDetails)
	if err == nil {
		if obj.SecurityMonitoringContentPackVulnerabilityDetails != nil && obj.SecurityMonitoringContentPackVulnerabilityDetails.UnparsedObject == nil {
			jsonSecurityMonitoringContentPackVulnerabilityDetails, _ := datadog.Marshal(obj.SecurityMonitoringContentPackVulnerabilityDetails)
			if string(jsonSecurityMonitoringContentPackVulnerabilityDetails) == "{}" { // empty struct
				obj.SecurityMonitoringContentPackVulnerabilityDetails = nil
			} else {
				match++
			}
		} else {
			obj.SecurityMonitoringContentPackVulnerabilityDetails = nil
		}
	} else {
		obj.SecurityMonitoringContentPackVulnerabilityDetails = nil
	}

	// try to unmarshal data into SecurityMonitoringContentPackOnboardingDetails
	err = datadog.Unmarshal(data, &obj.SecurityMonitoringContentPackOnboardingDetails)
	if err == nil {
		if obj.SecurityMonitoringContentPackOnboardingDetails != nil && obj.SecurityMonitoringContentPackOnboardingDetails.UnparsedObject == nil {
			jsonSecurityMonitoringContentPackOnboardingDetails, _ := datadog.Marshal(obj.SecurityMonitoringContentPackOnboardingDetails)
			if string(jsonSecurityMonitoringContentPackOnboardingDetails) == "{}" { // empty struct
				obj.SecurityMonitoringContentPackOnboardingDetails = nil
			} else {
				match++
			}
		} else {
			obj.SecurityMonitoringContentPackOnboardingDetails = nil
		}
	} else {
		obj.SecurityMonitoringContentPackOnboardingDetails = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.SecurityMonitoringContentPackLogsDetails = nil
		obj.SecurityMonitoringContentPackThreatIntelDetails = nil
		obj.SecurityMonitoringContentPackEntityDetails = nil
		obj.SecurityMonitoringContentPackAuditDetails = nil
		obj.SecurityMonitoringContentPackAppSecDetails = nil
		obj.SecurityMonitoringContentPackVulnerabilityDetails = nil
		obj.SecurityMonitoringContentPackOnboardingDetails = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
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
