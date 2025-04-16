// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EscalationPolicyUpdateRequestDataRelationshipsTeamsDataItemsType Indicates that the resource is of type `teams`.
type EscalationPolicyUpdateRequestDataRelationshipsTeamsDataItemsType string

// List of EscalationPolicyUpdateRequestDataRelationshipsTeamsDataItemsType.
const (
	ESCALATIONPOLICYUPDATEREQUESTDATARELATIONSHIPSTEAMSDATAITEMSTYPE_TEAMS EscalationPolicyUpdateRequestDataRelationshipsTeamsDataItemsType = "teams"
)

var allowedEscalationPolicyUpdateRequestDataRelationshipsTeamsDataItemsTypeEnumValues = []EscalationPolicyUpdateRequestDataRelationshipsTeamsDataItemsType{
	ESCALATIONPOLICYUPDATEREQUESTDATARELATIONSHIPSTEAMSDATAITEMSTYPE_TEAMS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EscalationPolicyUpdateRequestDataRelationshipsTeamsDataItemsType) GetAllowedValues() []EscalationPolicyUpdateRequestDataRelationshipsTeamsDataItemsType {
	return allowedEscalationPolicyUpdateRequestDataRelationshipsTeamsDataItemsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EscalationPolicyUpdateRequestDataRelationshipsTeamsDataItemsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EscalationPolicyUpdateRequestDataRelationshipsTeamsDataItemsType(value)
	return nil
}

// NewEscalationPolicyUpdateRequestDataRelationshipsTeamsDataItemsTypeFromValue returns a pointer to a valid EscalationPolicyUpdateRequestDataRelationshipsTeamsDataItemsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEscalationPolicyUpdateRequestDataRelationshipsTeamsDataItemsTypeFromValue(v string) (*EscalationPolicyUpdateRequestDataRelationshipsTeamsDataItemsType, error) {
	ev := EscalationPolicyUpdateRequestDataRelationshipsTeamsDataItemsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EscalationPolicyUpdateRequestDataRelationshipsTeamsDataItemsType: valid values are %v", v, allowedEscalationPolicyUpdateRequestDataRelationshipsTeamsDataItemsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EscalationPolicyUpdateRequestDataRelationshipsTeamsDataItemsType) IsValid() bool {
	for _, existing := range allowedEscalationPolicyUpdateRequestDataRelationshipsTeamsDataItemsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EscalationPolicyUpdateRequestDataRelationshipsTeamsDataItemsType value.
func (v EscalationPolicyUpdateRequestDataRelationshipsTeamsDataItemsType) Ptr() *EscalationPolicyUpdateRequestDataRelationshipsTeamsDataItemsType {
	return &v
}
