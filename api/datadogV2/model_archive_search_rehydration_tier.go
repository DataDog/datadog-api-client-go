// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ArchiveSearchRehydrationTier Storage tier the matched logs are rehydrated into.
type ArchiveSearchRehydrationTier string

// List of ArchiveSearchRehydrationTier.
const (
	ARCHIVESEARCHREHYDRATIONTIER_STANDARD ArchiveSearchRehydrationTier = "standard"
	ARCHIVESEARCHREHYDRATIONTIER_FLEX     ArchiveSearchRehydrationTier = "flex"
)

var allowedArchiveSearchRehydrationTierEnumValues = []ArchiveSearchRehydrationTier{
	ARCHIVESEARCHREHYDRATIONTIER_STANDARD,
	ARCHIVESEARCHREHYDRATIONTIER_FLEX,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ArchiveSearchRehydrationTier) GetAllowedValues() []ArchiveSearchRehydrationTier {
	return allowedArchiveSearchRehydrationTierEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ArchiveSearchRehydrationTier) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ArchiveSearchRehydrationTier(value)
	return nil
}

// NewArchiveSearchRehydrationTierFromValue returns a pointer to a valid ArchiveSearchRehydrationTier
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewArchiveSearchRehydrationTierFromValue(v string) (*ArchiveSearchRehydrationTier, error) {
	ev := ArchiveSearchRehydrationTier(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ArchiveSearchRehydrationTier: valid values are %v", v, allowedArchiveSearchRehydrationTierEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ArchiveSearchRehydrationTier) IsValid() bool {
	for _, existing := range allowedArchiveSearchRehydrationTierEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ArchiveSearchRehydrationTier value.
func (v ArchiveSearchRehydrationTier) Ptr() *ArchiveSearchRehydrationTier {
	return &v
}
