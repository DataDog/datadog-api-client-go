// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EmailFormatType Specifies the format of the e-mail that is sent for On-Call notifications
type EmailFormatType string

// List of EmailFormatType.
const (
	EMAILFORMATTYPE_HTML EmailFormatType = "html"
	EMAILFORMATTYPE_TEXT EmailFormatType = "text"
)

var allowedEmailFormatTypeEnumValues = []EmailFormatType{
	EMAILFORMATTYPE_HTML,
	EMAILFORMATTYPE_TEXT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EmailFormatType) GetAllowedValues() []EmailFormatType {
	return allowedEmailFormatTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EmailFormatType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EmailFormatType(value)
	return nil
}

// NewEmailFormatTypeFromValue returns a pointer to a valid EmailFormatType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEmailFormatTypeFromValue(v string) (*EmailFormatType, error) {
	ev := EmailFormatType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EmailFormatType: valid values are %v", v, allowedEmailFormatTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EmailFormatType) IsValid() bool {
	for _, existing := range allowedEmailFormatTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EmailFormatType value.
func (v EmailFormatType) Ptr() *EmailFormatType {
	return &v
}
