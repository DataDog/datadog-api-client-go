// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageQuotaType The JSON:API resource type for a usage quota.
type UsageQuotaType string

// List of UsageQuotaType.
const (
	USAGEQUOTATYPE_QUOTAS UsageQuotaType = "quotas"
)

var allowedUsageQuotaTypeEnumValues = []UsageQuotaType{
	USAGEQUOTATYPE_QUOTAS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *UsageQuotaType) GetAllowedValues() []UsageQuotaType {
	return allowedUsageQuotaTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *UsageQuotaType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = UsageQuotaType(value)
	return nil
}

// NewUsageQuotaTypeFromValue returns a pointer to a valid UsageQuotaType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewUsageQuotaTypeFromValue(v string) (*UsageQuotaType, error) {
	ev := UsageQuotaType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for UsageQuotaType: valid values are %v", v, allowedUsageQuotaTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v UsageQuotaType) IsValid() bool {
	for _, existing := range allowedUsageQuotaTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UsageQuotaType value.
func (v UsageQuotaType) Ptr() *UsageQuotaType {
	return &v
}
