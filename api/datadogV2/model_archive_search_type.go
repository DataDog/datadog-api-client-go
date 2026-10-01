// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ArchiveSearchType Archive Search resource type.
type ArchiveSearchType string

// List of ArchiveSearchType.
const (
	ARCHIVESEARCHTYPE_ARCHIVE_SEARCH ArchiveSearchType = "archive_search"
)

var allowedArchiveSearchTypeEnumValues = []ArchiveSearchType{
	ARCHIVESEARCHTYPE_ARCHIVE_SEARCH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ArchiveSearchType) GetAllowedValues() []ArchiveSearchType {
	return allowedArchiveSearchTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ArchiveSearchType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ArchiveSearchType(value)
	return nil
}

// NewArchiveSearchTypeFromValue returns a pointer to a valid ArchiveSearchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewArchiveSearchTypeFromValue(v string) (*ArchiveSearchType, error) {
	ev := ArchiveSearchType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ArchiveSearchType: valid values are %v", v, allowedArchiveSearchTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ArchiveSearchType) IsValid() bool {
	for _, existing := range allowedArchiveSearchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ArchiveSearchType value.
func (v ArchiveSearchType) Ptr() *ArchiveSearchType {
	return &v
}
