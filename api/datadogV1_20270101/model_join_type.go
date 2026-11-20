// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// JoinType Test-only generated schema for the dashboard versioning proof of concept.
type JoinType string

// List of JoinType.
const (
	JOINTYPE_LEFT  JoinType = "left"
	JOINTYPE_INNER JoinType = "inner"
)

var allowedJoinTypeEnumValues = []JoinType{
	JOINTYPE_LEFT,
	JOINTYPE_INNER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *JoinType) GetAllowedValues() []JoinType {
	return allowedJoinTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *JoinType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = JoinType(value)
	return nil
}

// NewJoinTypeFromValue returns a pointer to a valid JoinType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewJoinTypeFromValue(v string) (*JoinType, error) {
	ev := JoinType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for JoinType: valid values are %v", v, allowedJoinTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v JoinType) IsValid() bool {
	for _, existing := range allowedJoinTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to JoinType value.
func (v JoinType) Ptr() *JoinType {
	return &v
}
