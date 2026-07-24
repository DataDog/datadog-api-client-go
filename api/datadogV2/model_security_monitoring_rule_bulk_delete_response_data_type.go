// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringRuleBulkDeleteResponseDataType The resource type for a bulk delete response.
type SecurityMonitoringRuleBulkDeleteResponseDataType string

// List of SecurityMonitoringRuleBulkDeleteResponseDataType.
const (
	SECURITYMONITORINGRULEBULKDELETERESPONSEDATATYPE_BULK_DELETE_RESPONSE SecurityMonitoringRuleBulkDeleteResponseDataType = "bulk_delete_response"
)

var allowedSecurityMonitoringRuleBulkDeleteResponseDataTypeEnumValues = []SecurityMonitoringRuleBulkDeleteResponseDataType{
	SECURITYMONITORINGRULEBULKDELETERESPONSEDATATYPE_BULK_DELETE_RESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringRuleBulkDeleteResponseDataType) GetAllowedValues() []SecurityMonitoringRuleBulkDeleteResponseDataType {
	return allowedSecurityMonitoringRuleBulkDeleteResponseDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringRuleBulkDeleteResponseDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringRuleBulkDeleteResponseDataType(value)
	return nil
}

// NewSecurityMonitoringRuleBulkDeleteResponseDataTypeFromValue returns a pointer to a valid SecurityMonitoringRuleBulkDeleteResponseDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringRuleBulkDeleteResponseDataTypeFromValue(v string) (*SecurityMonitoringRuleBulkDeleteResponseDataType, error) {
	ev := SecurityMonitoringRuleBulkDeleteResponseDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringRuleBulkDeleteResponseDataType: valid values are %v", v, allowedSecurityMonitoringRuleBulkDeleteResponseDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringRuleBulkDeleteResponseDataType) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringRuleBulkDeleteResponseDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringRuleBulkDeleteResponseDataType value.
func (v SecurityMonitoringRuleBulkDeleteResponseDataType) Ptr() *SecurityMonitoringRuleBulkDeleteResponseDataType {
	return &v
}
