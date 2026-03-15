// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ThreatIntelIndicatorType The type of threat indicator.
type ThreatIntelIndicatorType string

// List of ThreatIntelIndicatorType.
const (
	THREATINTELINDICATORTYPE_IP_ADDRESS ThreatIntelIndicatorType = "ip_address"
	THREATINTELINDICATORTYPE_DOMAIN     ThreatIntelIndicatorType = "domain"
	THREATINTELINDICATORTYPE_SHA256     ThreatIntelIndicatorType = "sha256"
)

var allowedThreatIntelIndicatorTypeEnumValues = []ThreatIntelIndicatorType{
	THREATINTELINDICATORTYPE_IP_ADDRESS,
	THREATINTELINDICATORTYPE_DOMAIN,
	THREATINTELINDICATORTYPE_SHA256,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ThreatIntelIndicatorType) GetAllowedValues() []ThreatIntelIndicatorType {
	return allowedThreatIntelIndicatorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ThreatIntelIndicatorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ThreatIntelIndicatorType(value)
	return nil
}

// NewThreatIntelIndicatorTypeFromValue returns a pointer to a valid ThreatIntelIndicatorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewThreatIntelIndicatorTypeFromValue(v string) (*ThreatIntelIndicatorType, error) {
	ev := ThreatIntelIndicatorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ThreatIntelIndicatorType: valid values are %v", v, allowedThreatIntelIndicatorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ThreatIntelIndicatorType) IsValid() bool {
	for _, existing := range allowedThreatIntelIndicatorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ThreatIntelIndicatorType value.
func (v ThreatIntelIndicatorType) Ptr() *ThreatIntelIndicatorType {
	return &v
}
