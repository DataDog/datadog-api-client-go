// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringIntegrationTypeCrowdStrike The source type for a CrowdStrike entity context sync.
type SecurityMonitoringIntegrationTypeCrowdStrike string

// List of SecurityMonitoringIntegrationTypeCrowdStrike.
const (
	SECURITYMONITORINGINTEGRATIONTYPECROWDSTRIKE_CROWDSTRIKE SecurityMonitoringIntegrationTypeCrowdStrike = "CROWDSTRIKE"
)

var allowedSecurityMonitoringIntegrationTypeCrowdStrikeEnumValues = []SecurityMonitoringIntegrationTypeCrowdStrike{
	SECURITYMONITORINGINTEGRATIONTYPECROWDSTRIKE_CROWDSTRIKE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringIntegrationTypeCrowdStrike) GetAllowedValues() []SecurityMonitoringIntegrationTypeCrowdStrike {
	return allowedSecurityMonitoringIntegrationTypeCrowdStrikeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringIntegrationTypeCrowdStrike) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringIntegrationTypeCrowdStrike(value)
	return nil
}

// NewSecurityMonitoringIntegrationTypeCrowdStrikeFromValue returns a pointer to a valid SecurityMonitoringIntegrationTypeCrowdStrike
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringIntegrationTypeCrowdStrikeFromValue(v string) (*SecurityMonitoringIntegrationTypeCrowdStrike, error) {
	ev := SecurityMonitoringIntegrationTypeCrowdStrike(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringIntegrationTypeCrowdStrike: valid values are %v", v, allowedSecurityMonitoringIntegrationTypeCrowdStrikeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringIntegrationTypeCrowdStrike) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringIntegrationTypeCrowdStrikeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringIntegrationTypeCrowdStrike value.
func (v SecurityMonitoringIntegrationTypeCrowdStrike) Ptr() *SecurityMonitoringIntegrationTypeCrowdStrike {
	return &v
}
