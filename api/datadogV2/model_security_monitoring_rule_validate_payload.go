// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringRuleValidatePayload - Validate a rule.
type SecurityMonitoringRuleValidatePayload struct {
	SecurityMonitoringStandardRulePayload *SecurityMonitoringStandardRulePayload
	SecurityMonitoringSignalRulePayload   *SecurityMonitoringSignalRulePayload
	CloudConfigurationRulePayload         *CloudConfigurationRulePayload

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// SecurityMonitoringStandardRulePayloadAsSecurityMonitoringRuleValidatePayload is a convenience function that returns SecurityMonitoringStandardRulePayload wrapped in SecurityMonitoringRuleValidatePayload.
func SecurityMonitoringStandardRulePayloadAsSecurityMonitoringRuleValidatePayload(v *SecurityMonitoringStandardRulePayload) SecurityMonitoringRuleValidatePayload {
	return SecurityMonitoringRuleValidatePayload{SecurityMonitoringStandardRulePayload: v}
}

// SecurityMonitoringSignalRulePayloadAsSecurityMonitoringRuleValidatePayload is a convenience function that returns SecurityMonitoringSignalRulePayload wrapped in SecurityMonitoringRuleValidatePayload.
func SecurityMonitoringSignalRulePayloadAsSecurityMonitoringRuleValidatePayload(v *SecurityMonitoringSignalRulePayload) SecurityMonitoringRuleValidatePayload {
	return SecurityMonitoringRuleValidatePayload{SecurityMonitoringSignalRulePayload: v}
}

// CloudConfigurationRulePayloadAsSecurityMonitoringRuleValidatePayload is a convenience function that returns CloudConfigurationRulePayload wrapped in SecurityMonitoringRuleValidatePayload.
func CloudConfigurationRulePayloadAsSecurityMonitoringRuleValidatePayload(v *CloudConfigurationRulePayload) SecurityMonitoringRuleValidatePayload {
	return SecurityMonitoringRuleValidatePayload{CloudConfigurationRulePayload: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *SecurityMonitoringRuleValidatePayload) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SecurityMonitoringStandardRulePayload
	err = datadog.Unmarshal(data, &obj.SecurityMonitoringStandardRulePayload)
	if err == nil {
		if obj.SecurityMonitoringStandardRulePayload != nil && obj.SecurityMonitoringStandardRulePayload.UnparsedObject == nil {
			jsonSecurityMonitoringStandardRulePayload, _ := datadog.Marshal(obj.SecurityMonitoringStandardRulePayload)
			if string(jsonSecurityMonitoringStandardRulePayload) == "{}" { // empty struct
				obj.SecurityMonitoringStandardRulePayload = nil
			} else {
				match++
			}
		} else {
			obj.SecurityMonitoringStandardRulePayload = nil
		}
	} else {
		obj.SecurityMonitoringStandardRulePayload = nil
	}

	// try to unmarshal data into SecurityMonitoringSignalRulePayload
	err = datadog.Unmarshal(data, &obj.SecurityMonitoringSignalRulePayload)
	if err == nil {
		if obj.SecurityMonitoringSignalRulePayload != nil && obj.SecurityMonitoringSignalRulePayload.UnparsedObject == nil {
			jsonSecurityMonitoringSignalRulePayload, _ := datadog.Marshal(obj.SecurityMonitoringSignalRulePayload)
			if string(jsonSecurityMonitoringSignalRulePayload) == "{}" { // empty struct
				obj.SecurityMonitoringSignalRulePayload = nil
			} else {
				match++
			}
		} else {
			obj.SecurityMonitoringSignalRulePayload = nil
		}
	} else {
		obj.SecurityMonitoringSignalRulePayload = nil
	}

	// try to unmarshal data into CloudConfigurationRulePayload
	err = datadog.Unmarshal(data, &obj.CloudConfigurationRulePayload)
	if err == nil {
		if obj.CloudConfigurationRulePayload != nil && obj.CloudConfigurationRulePayload.UnparsedObject == nil {
			jsonCloudConfigurationRulePayload, _ := datadog.Marshal(obj.CloudConfigurationRulePayload)
			if string(jsonCloudConfigurationRulePayload) == "{}" { // empty struct
				obj.CloudConfigurationRulePayload = nil
			} else {
				match++
			}
		} else {
			obj.CloudConfigurationRulePayload = nil
		}
	} else {
		obj.CloudConfigurationRulePayload = nil
	}

	if match > 1 {
		// more than one variant matched, so tell them apart by the items of the arrays they hold
		if obj.SecurityMonitoringStandardRulePayload != nil {
			allCasesUnparsed := len(obj.SecurityMonitoringStandardRulePayload.Cases) > 0
			for _, item := range obj.SecurityMonitoringStandardRulePayload.Cases {
				if item.UnparsedObject == nil {
					allCasesUnparsed = false
					break
				}
			}
			allQueriesUnparsed := len(obj.SecurityMonitoringStandardRulePayload.Queries) > 0
			for _, item := range obj.SecurityMonitoringStandardRulePayload.Queries {
				if item.UnparsedObject == nil {
					allQueriesUnparsed = false
					break
				}
			}
			if allCasesUnparsed || allQueriesUnparsed {
				obj.SecurityMonitoringStandardRulePayload = nil
				match--
			}
		}
		if obj.SecurityMonitoringSignalRulePayload != nil {
			allCasesUnparsed := len(obj.SecurityMonitoringSignalRulePayload.Cases) > 0
			for _, item := range obj.SecurityMonitoringSignalRulePayload.Cases {
				if item.UnparsedObject == nil {
					allCasesUnparsed = false
					break
				}
			}
			allQueriesUnparsed := len(obj.SecurityMonitoringSignalRulePayload.Queries) > 0
			for _, item := range obj.SecurityMonitoringSignalRulePayload.Queries {
				if item.UnparsedObject == nil {
					allQueriesUnparsed = false
					break
				}
			}
			if allCasesUnparsed || allQueriesUnparsed {
				obj.SecurityMonitoringSignalRulePayload = nil
				match--
			}
		}
		if obj.CloudConfigurationRulePayload != nil {
			allCasesUnparsed := len(obj.CloudConfigurationRulePayload.Cases) > 0
			for _, item := range obj.CloudConfigurationRulePayload.Cases {
				if item.UnparsedObject == nil {
					allCasesUnparsed = false
					break
				}
			}
			if allCasesUnparsed {
				obj.CloudConfigurationRulePayload = nil
				match--
			}
		}
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.SecurityMonitoringStandardRulePayload = nil
		obj.SecurityMonitoringSignalRulePayload = nil
		obj.CloudConfigurationRulePayload = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj SecurityMonitoringRuleValidatePayload) MarshalJSON() ([]byte, error) {
	if obj.SecurityMonitoringStandardRulePayload != nil {
		return datadog.Marshal(&obj.SecurityMonitoringStandardRulePayload)
	}

	if obj.SecurityMonitoringSignalRulePayload != nil {
		return datadog.Marshal(&obj.SecurityMonitoringSignalRulePayload)
	}

	if obj.CloudConfigurationRulePayload != nil {
		return datadog.Marshal(&obj.CloudConfigurationRulePayload)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *SecurityMonitoringRuleValidatePayload) GetActualInstance() interface{} {
	if obj.SecurityMonitoringStandardRulePayload != nil {
		return obj.SecurityMonitoringStandardRulePayload
	}

	if obj.SecurityMonitoringSignalRulePayload != nil {
		return obj.SecurityMonitoringSignalRulePayload
	}

	if obj.CloudConfigurationRulePayload != nil {
		return obj.CloudConfigurationRulePayload
	}

	// all schemas are nil
	return nil
}
