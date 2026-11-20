// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FORMSIZES Test-only generated schema for the dashboard versioning proof of concept.
type FORMSIZES string

// List of FORMSIZES.
const (
	FORMSIZES_XS FORMSIZES = "xs"
	FORMSIZES_SM FORMSIZES = "sm"
	FORMSIZES_MD FORMSIZES = "md"
	FORMSIZES_LG FORMSIZES = "lg"
)

var allowedFORMSIZESEnumValues = []FORMSIZES{
	FORMSIZES_XS,
	FORMSIZES_SM,
	FORMSIZES_MD,
	FORMSIZES_LG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FORMSIZES) GetAllowedValues() []FORMSIZES {
	return allowedFORMSIZESEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FORMSIZES) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FORMSIZES(value)
	return nil
}

// NewFORMSIZESFromValue returns a pointer to a valid FORMSIZES
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFORMSIZESFromValue(v string) (*FORMSIZES, error) {
	ev := FORMSIZES(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FORMSIZES: valid values are %v", v, allowedFORMSIZESEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FORMSIZES) IsValid() bool {
	for _, existing := range allowedFORMSIZESEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FORMSIZES value.
func (v FORMSIZES) Ptr() *FORMSIZES {
	return &v
}
