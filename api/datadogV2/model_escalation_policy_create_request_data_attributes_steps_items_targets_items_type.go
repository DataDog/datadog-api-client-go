// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EscalationPolicyCreateRequestDataAttributesStepsItemsTargetsItemsType Specifies the type of escalation target (example `users`, `schedules`, or `teams`).
type EscalationPolicyCreateRequestDataAttributesStepsItemsTargetsItemsType string

// List of EscalationPolicyCreateRequestDataAttributesStepsItemsTargetsItemsType.
const (
	ESCALATIONPOLICYCREATEREQUESTDATAATTRIBUTESSTEPSITEMSTARGETSITEMSTYPE_USERS     EscalationPolicyCreateRequestDataAttributesStepsItemsTargetsItemsType = "users"
	ESCALATIONPOLICYCREATEREQUESTDATAATTRIBUTESSTEPSITEMSTARGETSITEMSTYPE_SCHEDULES EscalationPolicyCreateRequestDataAttributesStepsItemsTargetsItemsType = "schedules"
	ESCALATIONPOLICYCREATEREQUESTDATAATTRIBUTESSTEPSITEMSTARGETSITEMSTYPE_TEAMS     EscalationPolicyCreateRequestDataAttributesStepsItemsTargetsItemsType = "teams"
)

var allowedEscalationPolicyCreateRequestDataAttributesStepsItemsTargetsItemsTypeEnumValues = []EscalationPolicyCreateRequestDataAttributesStepsItemsTargetsItemsType{
	ESCALATIONPOLICYCREATEREQUESTDATAATTRIBUTESSTEPSITEMSTARGETSITEMSTYPE_USERS,
	ESCALATIONPOLICYCREATEREQUESTDATAATTRIBUTESSTEPSITEMSTARGETSITEMSTYPE_SCHEDULES,
	ESCALATIONPOLICYCREATEREQUESTDATAATTRIBUTESSTEPSITEMSTARGETSITEMSTYPE_TEAMS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EscalationPolicyCreateRequestDataAttributesStepsItemsTargetsItemsType) GetAllowedValues() []EscalationPolicyCreateRequestDataAttributesStepsItemsTargetsItemsType {
	return allowedEscalationPolicyCreateRequestDataAttributesStepsItemsTargetsItemsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EscalationPolicyCreateRequestDataAttributesStepsItemsTargetsItemsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EscalationPolicyCreateRequestDataAttributesStepsItemsTargetsItemsType(value)
	return nil
}

// NewEscalationPolicyCreateRequestDataAttributesStepsItemsTargetsItemsTypeFromValue returns a pointer to a valid EscalationPolicyCreateRequestDataAttributesStepsItemsTargetsItemsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEscalationPolicyCreateRequestDataAttributesStepsItemsTargetsItemsTypeFromValue(v string) (*EscalationPolicyCreateRequestDataAttributesStepsItemsTargetsItemsType, error) {
	ev := EscalationPolicyCreateRequestDataAttributesStepsItemsTargetsItemsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EscalationPolicyCreateRequestDataAttributesStepsItemsTargetsItemsType: valid values are %v", v, allowedEscalationPolicyCreateRequestDataAttributesStepsItemsTargetsItemsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EscalationPolicyCreateRequestDataAttributesStepsItemsTargetsItemsType) IsValid() bool {
	for _, existing := range allowedEscalationPolicyCreateRequestDataAttributesStepsItemsTargetsItemsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EscalationPolicyCreateRequestDataAttributesStepsItemsTargetsItemsType value.
func (v EscalationPolicyCreateRequestDataAttributesStepsItemsTargetsItemsType) Ptr() *EscalationPolicyCreateRequestDataAttributesStepsItemsTargetsItemsType {
	return &v
}
