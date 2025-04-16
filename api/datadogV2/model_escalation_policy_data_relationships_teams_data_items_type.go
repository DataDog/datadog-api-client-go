// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EscalationPolicyDataRelationshipsTeamsDataItemsType Indicates that the resource is of type `teams`.
type EscalationPolicyDataRelationshipsTeamsDataItemsType string

// List of EscalationPolicyDataRelationshipsTeamsDataItemsType.
const (
	ESCALATIONPOLICYDATARELATIONSHIPSTEAMSDATAITEMSTYPE_TEAMS EscalationPolicyDataRelationshipsTeamsDataItemsType = "teams"
)

var allowedEscalationPolicyDataRelationshipsTeamsDataItemsTypeEnumValues = []EscalationPolicyDataRelationshipsTeamsDataItemsType{
	ESCALATIONPOLICYDATARELATIONSHIPSTEAMSDATAITEMSTYPE_TEAMS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EscalationPolicyDataRelationshipsTeamsDataItemsType) GetAllowedValues() []EscalationPolicyDataRelationshipsTeamsDataItemsType {
	return allowedEscalationPolicyDataRelationshipsTeamsDataItemsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EscalationPolicyDataRelationshipsTeamsDataItemsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EscalationPolicyDataRelationshipsTeamsDataItemsType(value)
	return nil
}

// NewEscalationPolicyDataRelationshipsTeamsDataItemsTypeFromValue returns a pointer to a valid EscalationPolicyDataRelationshipsTeamsDataItemsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEscalationPolicyDataRelationshipsTeamsDataItemsTypeFromValue(v string) (*EscalationPolicyDataRelationshipsTeamsDataItemsType, error) {
	ev := EscalationPolicyDataRelationshipsTeamsDataItemsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EscalationPolicyDataRelationshipsTeamsDataItemsType: valid values are %v", v, allowedEscalationPolicyDataRelationshipsTeamsDataItemsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EscalationPolicyDataRelationshipsTeamsDataItemsType) IsValid() bool {
	for _, existing := range allowedEscalationPolicyDataRelationshipsTeamsDataItemsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EscalationPolicyDataRelationshipsTeamsDataItemsType value.
func (v EscalationPolicyDataRelationshipsTeamsDataItemsType) Ptr() *EscalationPolicyDataRelationshipsTeamsDataItemsType {
	return &v
}
