// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringContentPackAppSecDetailsType Type for Application Security content pack details.
type SecurityMonitoringContentPackAppSecDetailsType string

// List of SecurityMonitoringContentPackAppSecDetailsType.
const (
	SECURITYMONITORINGCONTENTPACKAPPSECDETAILSTYPE_APPSEC SecurityMonitoringContentPackAppSecDetailsType = "appsec"
)

var allowedSecurityMonitoringContentPackAppSecDetailsTypeEnumValues = []SecurityMonitoringContentPackAppSecDetailsType{
	SECURITYMONITORINGCONTENTPACKAPPSECDETAILSTYPE_APPSEC,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringContentPackAppSecDetailsType) GetAllowedValues() []SecurityMonitoringContentPackAppSecDetailsType {
	return allowedSecurityMonitoringContentPackAppSecDetailsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringContentPackAppSecDetailsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringContentPackAppSecDetailsType(value)
	return nil
}

// NewSecurityMonitoringContentPackAppSecDetailsTypeFromValue returns a pointer to a valid SecurityMonitoringContentPackAppSecDetailsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringContentPackAppSecDetailsTypeFromValue(v string) (*SecurityMonitoringContentPackAppSecDetailsType, error) {
	ev := SecurityMonitoringContentPackAppSecDetailsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringContentPackAppSecDetailsType: valid values are %v", v, allowedSecurityMonitoringContentPackAppSecDetailsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringContentPackAppSecDetailsType) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringContentPackAppSecDetailsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringContentPackAppSecDetailsType value.
func (v SecurityMonitoringContentPackAppSecDetailsType) Ptr() *SecurityMonitoringContentPackAppSecDetailsType {
	return &v
}
