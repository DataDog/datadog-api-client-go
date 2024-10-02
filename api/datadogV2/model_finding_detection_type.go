// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FindingDetectionType The detection type of the finding.
type FindingDetectionType string

// List of FindingDetectionType.
const (
	FINDINGDETECTIONTYPE_MISCONFIGURATION FindingDetectionType = "misconfiguration"
	FINDINGDETECTIONTYPE_ATTACK_PATH      FindingDetectionType = "attack_path"
	FINDINGDETECTIONTYPE_IDENTITY_RISK    FindingDetectionType = "identity_risk"
	FINDINGDETECTIONTYPE_API_SECURITY     FindingDetectionType = "api_security"
)

var allowedFindingDetectionTypeEnumValues = []FindingDetectionType{
	FINDINGDETECTIONTYPE_MISCONFIGURATION,
	FINDINGDETECTIONTYPE_ATTACK_PATH,
	FINDINGDETECTIONTYPE_IDENTITY_RISK,
	FINDINGDETECTIONTYPE_API_SECURITY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FindingDetectionType) GetAllowedValues() []FindingDetectionType {
	return allowedFindingDetectionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FindingDetectionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FindingDetectionType(value)
	return nil
}

// NewFindingDetectionTypeFromValue returns a pointer to a valid FindingDetectionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFindingDetectionTypeFromValue(v string) (*FindingDetectionType, error) {
	ev := FindingDetectionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FindingDetectionType: valid values are %v", v, allowedFindingDetectionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FindingDetectionType) IsValid() bool {
	for _, existing := range allowedFindingDetectionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FindingDetectionType value.
func (v FindingDetectionType) Ptr() *FindingDetectionType {
	return &v
}
