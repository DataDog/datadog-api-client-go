// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamsOwnershipMappingBatchOperationOp Whether this operation adds a new mapping or removes an existing one.
type TeamsOwnershipMappingBatchOperationOp string

// List of TeamsOwnershipMappingBatchOperationOp.
const (
	TEAMSOWNERSHIPMAPPINGBATCHOPERATIONOP_ADD    TeamsOwnershipMappingBatchOperationOp = "add"
	TEAMSOWNERSHIPMAPPINGBATCHOPERATIONOP_REMOVE TeamsOwnershipMappingBatchOperationOp = "remove"
)

var allowedTeamsOwnershipMappingBatchOperationOpEnumValues = []TeamsOwnershipMappingBatchOperationOp{
	TEAMSOWNERSHIPMAPPINGBATCHOPERATIONOP_ADD,
	TEAMSOWNERSHIPMAPPINGBATCHOPERATIONOP_REMOVE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TeamsOwnershipMappingBatchOperationOp) GetAllowedValues() []TeamsOwnershipMappingBatchOperationOp {
	return allowedTeamsOwnershipMappingBatchOperationOpEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TeamsOwnershipMappingBatchOperationOp) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TeamsOwnershipMappingBatchOperationOp(value)
	return nil
}

// NewTeamsOwnershipMappingBatchOperationOpFromValue returns a pointer to a valid TeamsOwnershipMappingBatchOperationOp
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTeamsOwnershipMappingBatchOperationOpFromValue(v string) (*TeamsOwnershipMappingBatchOperationOp, error) {
	ev := TeamsOwnershipMappingBatchOperationOp(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TeamsOwnershipMappingBatchOperationOp: valid values are %v", v, allowedTeamsOwnershipMappingBatchOperationOpEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TeamsOwnershipMappingBatchOperationOp) IsValid() bool {
	for _, existing := range allowedTeamsOwnershipMappingBatchOperationOpEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TeamsOwnershipMappingBatchOperationOp value.
func (v TeamsOwnershipMappingBatchOperationOp) Ptr() *TeamsOwnershipMappingBatchOperationOp {
	return &v
}
