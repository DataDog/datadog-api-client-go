// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PersonalAccessTokenType Personal access tokens resource type.
type PersonalAccessTokenType string

// List of PersonalAccessTokenType.
const (
	PERSONALACCESSTOKENTYPE_PERSONAL_ACCESS_TOKENS PersonalAccessTokenType = "personal_access_tokens"
)

var allowedPersonalAccessTokenTypeEnumValues = []PersonalAccessTokenType{
	PERSONALACCESSTOKENTYPE_PERSONAL_ACCESS_TOKENS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PersonalAccessTokenType) GetAllowedValues() []PersonalAccessTokenType {
	return allowedPersonalAccessTokenTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PersonalAccessTokenType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PersonalAccessTokenType(value)
	return nil
}

// NewPersonalAccessTokenTypeFromValue returns a pointer to a valid PersonalAccessTokenType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPersonalAccessTokenTypeFromValue(v string) (*PersonalAccessTokenType, error) {
	ev := PersonalAccessTokenType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PersonalAccessTokenType: valid values are %v", v, allowedPersonalAccessTokenTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PersonalAccessTokenType) IsValid() bool {
	for _, existing := range allowedPersonalAccessTokenTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PersonalAccessTokenType value.
func (v PersonalAccessTokenType) Ptr() *PersonalAccessTokenType {
	return &v
}
