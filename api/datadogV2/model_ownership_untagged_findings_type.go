// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OwnershipUntaggedFindingsType The type of the ownership untagged findings resource. The value should always be `ownership_untagged_findings`.
type OwnershipUntaggedFindingsType string

// List of OwnershipUntaggedFindingsType.
const (
	OWNERSHIPUNTAGGEDFINDINGSTYPE_OWNERSHIP_UNTAGGED_FINDINGS OwnershipUntaggedFindingsType = "ownership_untagged_findings"
)

var allowedOwnershipUntaggedFindingsTypeEnumValues = []OwnershipUntaggedFindingsType{
	OWNERSHIPUNTAGGEDFINDINGSTYPE_OWNERSHIP_UNTAGGED_FINDINGS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OwnershipUntaggedFindingsType) GetAllowedValues() []OwnershipUntaggedFindingsType {
	return allowedOwnershipUntaggedFindingsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OwnershipUntaggedFindingsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OwnershipUntaggedFindingsType(value)
	return nil
}

// NewOwnershipUntaggedFindingsTypeFromValue returns a pointer to a valid OwnershipUntaggedFindingsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOwnershipUntaggedFindingsTypeFromValue(v string) (*OwnershipUntaggedFindingsType, error) {
	ev := OwnershipUntaggedFindingsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OwnershipUntaggedFindingsType: valid values are %v", v, allowedOwnershipUntaggedFindingsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OwnershipUntaggedFindingsType) IsValid() bool {
	for _, existing := range allowedOwnershipUntaggedFindingsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OwnershipUntaggedFindingsType value.
func (v OwnershipUntaggedFindingsType) Ptr() *OwnershipUntaggedFindingsType {
	return &v
}
