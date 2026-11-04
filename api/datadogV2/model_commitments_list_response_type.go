// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CommitmentsListResponseType Type of the commitments list response resource.
type CommitmentsListResponseType string

// List of CommitmentsListResponseType.
const (
	COMMITMENTSLISTRESPONSETYPE_COMMITMENTS_LIST_RESPONSE CommitmentsListResponseType = "commitments_list_response"
)

var allowedCommitmentsListResponseTypeEnumValues = []CommitmentsListResponseType{
	COMMITMENTSLISTRESPONSETYPE_COMMITMENTS_LIST_RESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CommitmentsListResponseType) GetAllowedValues() []CommitmentsListResponseType {
	return allowedCommitmentsListResponseTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CommitmentsListResponseType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CommitmentsListResponseType(value)
	return nil
}

// NewCommitmentsListResponseTypeFromValue returns a pointer to a valid CommitmentsListResponseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCommitmentsListResponseTypeFromValue(v string) (*CommitmentsListResponseType, error) {
	ev := CommitmentsListResponseType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CommitmentsListResponseType: valid values are %v", v, allowedCommitmentsListResponseTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CommitmentsListResponseType) IsValid() bool {
	for _, existing := range allowedCommitmentsListResponseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CommitmentsListResponseType value.
func (v CommitmentsListResponseType) Ptr() *CommitmentsListResponseType {
	return &v
}
