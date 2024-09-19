// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.


package datadogV2

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"

)


// AbbreviatedTeamType The definition of `AbbreviatedTeamType` object.
type AbbreviatedTeamType string

// List of AbbreviatedTeamType.
const (
	ABBREVIATEDTEAMTYPE_TEAM AbbreviatedTeamType = "team"
)

var allowedAbbreviatedTeamTypeEnumValues = []AbbreviatedTeamType{
	ABBREVIATEDTEAMTYPE_TEAM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AbbreviatedTeamType) GetAllowedValues() []AbbreviatedTeamType {
	return allowedAbbreviatedTeamTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AbbreviatedTeamType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AbbreviatedTeamType(value)
	return nil
}

// NewAbbreviatedTeamTypeFromValue returns a pointer to a valid AbbreviatedTeamType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAbbreviatedTeamTypeFromValue(v string) (*AbbreviatedTeamType, error) {
	ev := AbbreviatedTeamType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AbbreviatedTeamType: valid values are %v", v, allowedAbbreviatedTeamTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AbbreviatedTeamType) IsValid() bool {
	for _, existing := range allowedAbbreviatedTeamTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AbbreviatedTeamType value.
func (v AbbreviatedTeamType) Ptr() *AbbreviatedTeamType {
	return &v
}
