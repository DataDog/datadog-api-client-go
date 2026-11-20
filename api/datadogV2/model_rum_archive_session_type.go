// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMArchiveSessionType JSON:API resource type.
type RUMArchiveSessionType string

// List of RUMArchiveSessionType.
const (
	RUMARCHIVESESSIONTYPE_RUM_ARCHIVE_SESSION RUMArchiveSessionType = "rum_archive_session"
)

var allowedRUMArchiveSessionTypeEnumValues = []RUMArchiveSessionType{
	RUMARCHIVESESSIONTYPE_RUM_ARCHIVE_SESSION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RUMArchiveSessionType) GetAllowedValues() []RUMArchiveSessionType {
	return allowedRUMArchiveSessionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RUMArchiveSessionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMArchiveSessionType(value)
	return nil
}

// NewRUMArchiveSessionTypeFromValue returns a pointer to a valid RUMArchiveSessionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRUMArchiveSessionTypeFromValue(v string) (*RUMArchiveSessionType, error) {
	ev := RUMArchiveSessionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RUMArchiveSessionType: valid values are %v", v, allowedRUMArchiveSessionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RUMArchiveSessionType) IsValid() bool {
	for _, existing := range allowedRUMArchiveSessionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMArchiveSessionType value.
func (v RUMArchiveSessionType) Ptr() *RUMArchiveSessionType {
	return &v
}
