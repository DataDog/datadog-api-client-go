// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SavedFilterConditionOperator The comparison operator used by a saved filter condition.
type SavedFilterConditionOperator string

// List of SavedFilterConditionOperator.
const (
	SAVEDFILTERCONDITIONOPERATOR_LT          SavedFilterConditionOperator = "LT"
	SAVEDFILTERCONDITIONOPERATOR_LTE         SavedFilterConditionOperator = "LTE"
	SAVEDFILTERCONDITIONOPERATOR_GT          SavedFilterConditionOperator = "GT"
	SAVEDFILTERCONDITIONOPERATOR_GTE         SavedFilterConditionOperator = "GTE"
	SAVEDFILTERCONDITIONOPERATOR_MATCHES     SavedFilterConditionOperator = "MATCHES"
	SAVEDFILTERCONDITIONOPERATOR_NOT_MATCHES SavedFilterConditionOperator = "NOT_MATCHES"
	SAVEDFILTERCONDITIONOPERATOR_ONE_OF      SavedFilterConditionOperator = "ONE_OF"
	SAVEDFILTERCONDITIONOPERATOR_NOT_ONE_OF  SavedFilterConditionOperator = "NOT_ONE_OF"
	SAVEDFILTERCONDITIONOPERATOR_IS_NULL     SavedFilterConditionOperator = "IS_NULL"
	SAVEDFILTERCONDITIONOPERATOR_EQUALS      SavedFilterConditionOperator = "EQUALS"
	SAVEDFILTERCONDITIONOPERATOR_SEMVER_EQ   SavedFilterConditionOperator = "SEMVER_EQ"
	SAVEDFILTERCONDITIONOPERATOR_SEMVER_NEQ  SavedFilterConditionOperator = "SEMVER_NEQ"
	SAVEDFILTERCONDITIONOPERATOR_SEMVER_LT   SavedFilterConditionOperator = "SEMVER_LT"
	SAVEDFILTERCONDITIONOPERATOR_SEMVER_LTE  SavedFilterConditionOperator = "SEMVER_LTE"
	SAVEDFILTERCONDITIONOPERATOR_SEMVER_GT   SavedFilterConditionOperator = "SEMVER_GT"
	SAVEDFILTERCONDITIONOPERATOR_SEMVER_GTE  SavedFilterConditionOperator = "SEMVER_GTE"
)

var allowedSavedFilterConditionOperatorEnumValues = []SavedFilterConditionOperator{
	SAVEDFILTERCONDITIONOPERATOR_LT,
	SAVEDFILTERCONDITIONOPERATOR_LTE,
	SAVEDFILTERCONDITIONOPERATOR_GT,
	SAVEDFILTERCONDITIONOPERATOR_GTE,
	SAVEDFILTERCONDITIONOPERATOR_MATCHES,
	SAVEDFILTERCONDITIONOPERATOR_NOT_MATCHES,
	SAVEDFILTERCONDITIONOPERATOR_ONE_OF,
	SAVEDFILTERCONDITIONOPERATOR_NOT_ONE_OF,
	SAVEDFILTERCONDITIONOPERATOR_IS_NULL,
	SAVEDFILTERCONDITIONOPERATOR_EQUALS,
	SAVEDFILTERCONDITIONOPERATOR_SEMVER_EQ,
	SAVEDFILTERCONDITIONOPERATOR_SEMVER_NEQ,
	SAVEDFILTERCONDITIONOPERATOR_SEMVER_LT,
	SAVEDFILTERCONDITIONOPERATOR_SEMVER_LTE,
	SAVEDFILTERCONDITIONOPERATOR_SEMVER_GT,
	SAVEDFILTERCONDITIONOPERATOR_SEMVER_GTE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SavedFilterConditionOperator) GetAllowedValues() []SavedFilterConditionOperator {
	return allowedSavedFilterConditionOperatorEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SavedFilterConditionOperator) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SavedFilterConditionOperator(value)
	return nil
}

// NewSavedFilterConditionOperatorFromValue returns a pointer to a valid SavedFilterConditionOperator
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSavedFilterConditionOperatorFromValue(v string) (*SavedFilterConditionOperator, error) {
	ev := SavedFilterConditionOperator(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SavedFilterConditionOperator: valid values are %v", v, allowedSavedFilterConditionOperatorEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SavedFilterConditionOperator) IsValid() bool {
	for _, existing := range allowedSavedFilterConditionOperatorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SavedFilterConditionOperator value.
func (v SavedFilterConditionOperator) Ptr() *SavedFilterConditionOperator {
	return &v
}
