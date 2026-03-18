// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GuidedTableStorageType Storage tier to target for an events-platform query in a guided table.
type GuidedTableStorageType string

// List of GuidedTableStorageType.
const (
	GUIDEDTABLESTORAGETYPE_LIVE             GuidedTableStorageType = "live"
	GUIDEDTABLESTORAGETYPE_HOT              GuidedTableStorageType = "hot"
	GUIDEDTABLESTORAGETYPE_ONLINE_ARCHIVES  GuidedTableStorageType = "online_archives"
	GUIDEDTABLESTORAGETYPE_DRIVELINE        GuidedTableStorageType = "driveline"
	GUIDEDTABLESTORAGETYPE_FLEX_TIER        GuidedTableStorageType = "flex_tier"
	GUIDEDTABLESTORAGETYPE_CASE_INSENSITIVE GuidedTableStorageType = "case_insensitive"
	GUIDEDTABLESTORAGETYPE_CLOUD_PREM       GuidedTableStorageType = "cloud_prem"
)

var allowedGuidedTableStorageTypeEnumValues = []GuidedTableStorageType{
	GUIDEDTABLESTORAGETYPE_LIVE,
	GUIDEDTABLESTORAGETYPE_HOT,
	GUIDEDTABLESTORAGETYPE_ONLINE_ARCHIVES,
	GUIDEDTABLESTORAGETYPE_DRIVELINE,
	GUIDEDTABLESTORAGETYPE_FLEX_TIER,
	GUIDEDTABLESTORAGETYPE_CASE_INSENSITIVE,
	GUIDEDTABLESTORAGETYPE_CLOUD_PREM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GuidedTableStorageType) GetAllowedValues() []GuidedTableStorageType {
	return allowedGuidedTableStorageTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GuidedTableStorageType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GuidedTableStorageType(value)
	return nil
}

// NewGuidedTableStorageTypeFromValue returns a pointer to a valid GuidedTableStorageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGuidedTableStorageTypeFromValue(v string) (*GuidedTableStorageType, error) {
	ev := GuidedTableStorageType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GuidedTableStorageType: valid values are %v", v, allowedGuidedTableStorageTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GuidedTableStorageType) IsValid() bool {
	for _, existing := range allowedGuidedTableStorageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GuidedTableStorageType value.
func (v GuidedTableStorageType) Ptr() *GuidedTableStorageType {
	return &v
}
