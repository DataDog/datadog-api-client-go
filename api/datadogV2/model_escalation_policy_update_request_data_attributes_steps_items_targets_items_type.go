// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EscalationPolicyUpdateRequestDataAttributesStepsItemsTargetsItemsType Specifies the type of escalation target (example `users`, `schedules`, or `teams`).
type EscalationPolicyUpdateRequestDataAttributesStepsItemsTargetsItemsType string

// List of EscalationPolicyUpdateRequestDataAttributesStepsItemsTargetsItemsType.
const (
	ESCALATIONPOLICYUPDATEREQUESTDATAATTRIBUTESSTEPSITEMSTARGETSITEMSTYPE_USERS     EscalationPolicyUpdateRequestDataAttributesStepsItemsTargetsItemsType = "users"
	ESCALATIONPOLICYUPDATEREQUESTDATAATTRIBUTESSTEPSITEMSTARGETSITEMSTYPE_SCHEDULES EscalationPolicyUpdateRequestDataAttributesStepsItemsTargetsItemsType = "schedules"
	ESCALATIONPOLICYUPDATEREQUESTDATAATTRIBUTESSTEPSITEMSTARGETSITEMSTYPE_TEAMS     EscalationPolicyUpdateRequestDataAttributesStepsItemsTargetsItemsType = "teams"
)

var allowedEscalationPolicyUpdateRequestDataAttributesStepsItemsTargetsItemsTypeEnumValues = []EscalationPolicyUpdateRequestDataAttributesStepsItemsTargetsItemsType{
	ESCALATIONPOLICYUPDATEREQUESTDATAATTRIBUTESSTEPSITEMSTARGETSITEMSTYPE_USERS,
	ESCALATIONPOLICYUPDATEREQUESTDATAATTRIBUTESSTEPSITEMSTARGETSITEMSTYPE_SCHEDULES,
	ESCALATIONPOLICYUPDATEREQUESTDATAATTRIBUTESSTEPSITEMSTARGETSITEMSTYPE_TEAMS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EscalationPolicyUpdateRequestDataAttributesStepsItemsTargetsItemsType) GetAllowedValues() []EscalationPolicyUpdateRequestDataAttributesStepsItemsTargetsItemsType {
	return allowedEscalationPolicyUpdateRequestDataAttributesStepsItemsTargetsItemsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EscalationPolicyUpdateRequestDataAttributesStepsItemsTargetsItemsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EscalationPolicyUpdateRequestDataAttributesStepsItemsTargetsItemsType(value)
	return nil
}

// NewEscalationPolicyUpdateRequestDataAttributesStepsItemsTargetsItemsTypeFromValue returns a pointer to a valid EscalationPolicyUpdateRequestDataAttributesStepsItemsTargetsItemsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEscalationPolicyUpdateRequestDataAttributesStepsItemsTargetsItemsTypeFromValue(v string) (*EscalationPolicyUpdateRequestDataAttributesStepsItemsTargetsItemsType, error) {
	ev := EscalationPolicyUpdateRequestDataAttributesStepsItemsTargetsItemsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EscalationPolicyUpdateRequestDataAttributesStepsItemsTargetsItemsType: valid values are %v", v, allowedEscalationPolicyUpdateRequestDataAttributesStepsItemsTargetsItemsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EscalationPolicyUpdateRequestDataAttributesStepsItemsTargetsItemsType) IsValid() bool {
	for _, existing := range allowedEscalationPolicyUpdateRequestDataAttributesStepsItemsTargetsItemsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EscalationPolicyUpdateRequestDataAttributesStepsItemsTargetsItemsType value.
func (v EscalationPolicyUpdateRequestDataAttributesStepsItemsTargetsItemsType) Ptr() *EscalationPolicyUpdateRequestDataAttributesStepsItemsTargetsItemsType {
	return &v
}
