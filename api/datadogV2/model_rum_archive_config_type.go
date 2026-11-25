// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMArchiveConfigType JSON:API resource type.
type RUMArchiveConfigType string

// List of RUMArchiveConfigType.
const (
	RUMARCHIVECONFIGTYPE_RUM_ARCHIVE_CONFIG RUMArchiveConfigType = "rum_archive_config"
)

var allowedRUMArchiveConfigTypeEnumValues = []RUMArchiveConfigType{
	RUMARCHIVECONFIGTYPE_RUM_ARCHIVE_CONFIG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RUMArchiveConfigType) GetAllowedValues() []RUMArchiveConfigType {
	return allowedRUMArchiveConfigTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RUMArchiveConfigType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMArchiveConfigType(value)
	return nil
}

// NewRUMArchiveConfigTypeFromValue returns a pointer to a valid RUMArchiveConfigType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRUMArchiveConfigTypeFromValue(v string) (*RUMArchiveConfigType, error) {
	ev := RUMArchiveConfigType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RUMArchiveConfigType: valid values are %v", v, allowedRUMArchiveConfigTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RUMArchiveConfigType) IsValid() bool {
	for _, existing := range allowedRUMArchiveConfigTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMArchiveConfigType value.
func (v RUMArchiveConfigType) Ptr() *RUMArchiveConfigType {
	return &v
}
