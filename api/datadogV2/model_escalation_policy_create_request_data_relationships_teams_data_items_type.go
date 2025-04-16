// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EscalationPolicyCreateRequestDataRelationshipsTeamsDataItemsType Indicates that the resource is of type `teams`.
type EscalationPolicyCreateRequestDataRelationshipsTeamsDataItemsType string

// List of EscalationPolicyCreateRequestDataRelationshipsTeamsDataItemsType.
const (
	ESCALATIONPOLICYCREATEREQUESTDATARELATIONSHIPSTEAMSDATAITEMSTYPE_TEAMS EscalationPolicyCreateRequestDataRelationshipsTeamsDataItemsType = "teams"
)

var allowedEscalationPolicyCreateRequestDataRelationshipsTeamsDataItemsTypeEnumValues = []EscalationPolicyCreateRequestDataRelationshipsTeamsDataItemsType{
	ESCALATIONPOLICYCREATEREQUESTDATARELATIONSHIPSTEAMSDATAITEMSTYPE_TEAMS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EscalationPolicyCreateRequestDataRelationshipsTeamsDataItemsType) GetAllowedValues() []EscalationPolicyCreateRequestDataRelationshipsTeamsDataItemsType {
	return allowedEscalationPolicyCreateRequestDataRelationshipsTeamsDataItemsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EscalationPolicyCreateRequestDataRelationshipsTeamsDataItemsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EscalationPolicyCreateRequestDataRelationshipsTeamsDataItemsType(value)
	return nil
}

// NewEscalationPolicyCreateRequestDataRelationshipsTeamsDataItemsTypeFromValue returns a pointer to a valid EscalationPolicyCreateRequestDataRelationshipsTeamsDataItemsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEscalationPolicyCreateRequestDataRelationshipsTeamsDataItemsTypeFromValue(v string) (*EscalationPolicyCreateRequestDataRelationshipsTeamsDataItemsType, error) {
	ev := EscalationPolicyCreateRequestDataRelationshipsTeamsDataItemsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EscalationPolicyCreateRequestDataRelationshipsTeamsDataItemsType: valid values are %v", v, allowedEscalationPolicyCreateRequestDataRelationshipsTeamsDataItemsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EscalationPolicyCreateRequestDataRelationshipsTeamsDataItemsType) IsValid() bool {
	for _, existing := range allowedEscalationPolicyCreateRequestDataRelationshipsTeamsDataItemsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EscalationPolicyCreateRequestDataRelationshipsTeamsDataItemsType value.
func (v EscalationPolicyCreateRequestDataRelationshipsTeamsDataItemsType) Ptr() *EscalationPolicyCreateRequestDataRelationshipsTeamsDataItemsType {
	return &v
}
