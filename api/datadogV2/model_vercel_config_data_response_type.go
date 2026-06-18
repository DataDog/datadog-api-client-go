// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// VercelConfigDataResponseType Type identifier for a Vercel configuration resource.
type VercelConfigDataResponseType string

// List of VercelConfigDataResponseType.
const (
	VERCELCONFIGDATARESPONSETYPE_VERCEL_CONFIGURATION VercelConfigDataResponseType = "vercelConfiguration"
)

var allowedVercelConfigDataResponseTypeEnumValues = []VercelConfigDataResponseType{
	VERCELCONFIGDATARESPONSETYPE_VERCEL_CONFIGURATION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *VercelConfigDataResponseType) GetAllowedValues() []VercelConfigDataResponseType {
	return allowedVercelConfigDataResponseTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *VercelConfigDataResponseType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = VercelConfigDataResponseType(value)
	return nil
}

// NewVercelConfigDataResponseTypeFromValue returns a pointer to a valid VercelConfigDataResponseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewVercelConfigDataResponseTypeFromValue(v string) (*VercelConfigDataResponseType, error) {
	ev := VercelConfigDataResponseType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for VercelConfigDataResponseType: valid values are %v", v, allowedVercelConfigDataResponseTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v VercelConfigDataResponseType) IsValid() bool {
	for _, existing := range allowedVercelConfigDataResponseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VercelConfigDataResponseType value.
func (v VercelConfigDataResponseType) Ptr() *VercelConfigDataResponseType {
	return &v
}
