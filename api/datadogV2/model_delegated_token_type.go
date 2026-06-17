// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DelegatedTokenType The resource type for a delegated token.
type DelegatedTokenType string

// List of DelegatedTokenType.
const (
	DELEGATEDTOKENTYPE_TOKEN DelegatedTokenType = "token"
)

var allowedDelegatedTokenTypeEnumValues = []DelegatedTokenType{
	DELEGATEDTOKENTYPE_TOKEN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DelegatedTokenType) GetAllowedValues() []DelegatedTokenType {
	return allowedDelegatedTokenTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DelegatedTokenType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DelegatedTokenType(value)
	return nil
}

// NewDelegatedTokenTypeFromValue returns a pointer to a valid DelegatedTokenType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDelegatedTokenTypeFromValue(v string) (*DelegatedTokenType, error) {
	ev := DelegatedTokenType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DelegatedTokenType: valid values are %v", v, allowedDelegatedTokenTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DelegatedTokenType) IsValid() bool {
	for _, existing := range allowedDelegatedTokenTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DelegatedTokenType value.
func (v DelegatedTokenType) Ptr() *DelegatedTokenType {
	return &v
}
