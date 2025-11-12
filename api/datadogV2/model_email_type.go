// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EmailType Indicates that the resource is of type 'emails'.
type EmailType string

// List of EmailType.
const (
	EMAILTYPE_EMAILS EmailType = "emails"
)

var allowedEmailTypeEnumValues = []EmailType{
	EMAILTYPE_EMAILS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EmailType) GetAllowedValues() []EmailType {
	return allowedEmailTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EmailType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EmailType(value)
	return nil
}

// NewEmailTypeFromValue returns a pointer to a valid EmailType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEmailTypeFromValue(v string) (*EmailType, error) {
	ev := EmailType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EmailType: valid values are %v", v, allowedEmailTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EmailType) IsValid() bool {
	for _, existing := range allowedEmailTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EmailType value.
func (v EmailType) Ptr() *EmailType {
	return &v
}
