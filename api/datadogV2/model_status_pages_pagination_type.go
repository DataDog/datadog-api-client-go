// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StatusPagesPaginationType
type StatusPagesPaginationType string

// List of StatusPagesPaginationType.
const (
	STATUSPAGESPAGINATIONTYPE_OFFSET_LIMIT StatusPagesPaginationType = "offset_limit"
)

var allowedStatusPagesPaginationTypeEnumValues = []StatusPagesPaginationType{
	STATUSPAGESPAGINATIONTYPE_OFFSET_LIMIT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *StatusPagesPaginationType) GetAllowedValues() []StatusPagesPaginationType {
	return allowedStatusPagesPaginationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *StatusPagesPaginationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = StatusPagesPaginationType(value)
	return nil
}

// NewStatusPagesPaginationTypeFromValue returns a pointer to a valid StatusPagesPaginationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewStatusPagesPaginationTypeFromValue(v string) (*StatusPagesPaginationType, error) {
	ev := StatusPagesPaginationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for StatusPagesPaginationType: valid values are %v", v, allowedStatusPagesPaginationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v StatusPagesPaginationType) IsValid() bool {
	for _, existing := range allowedStatusPagesPaginationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StatusPagesPaginationType value.
func (v StatusPagesPaginationType) Ptr() *StatusPagesPaginationType {
	return &v
}
