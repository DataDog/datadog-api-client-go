// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EscalationPolicyCreateRequestDataAttributesStepsItemsAssignment Specifies how this escalation step will assign targets (example `default`).
type EscalationPolicyCreateRequestDataAttributesStepsItemsAssignment string

// List of EscalationPolicyCreateRequestDataAttributesStepsItemsAssignment.
const (
	ESCALATIONPOLICYCREATEREQUESTDATAATTRIBUTESSTEPSITEMSASSIGNMENT_DEFAULT EscalationPolicyCreateRequestDataAttributesStepsItemsAssignment = "default"
)

var allowedEscalationPolicyCreateRequestDataAttributesStepsItemsAssignmentEnumValues = []EscalationPolicyCreateRequestDataAttributesStepsItemsAssignment{
	ESCALATIONPOLICYCREATEREQUESTDATAATTRIBUTESSTEPSITEMSASSIGNMENT_DEFAULT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EscalationPolicyCreateRequestDataAttributesStepsItemsAssignment) GetAllowedValues() []EscalationPolicyCreateRequestDataAttributesStepsItemsAssignment {
	return allowedEscalationPolicyCreateRequestDataAttributesStepsItemsAssignmentEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EscalationPolicyCreateRequestDataAttributesStepsItemsAssignment) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EscalationPolicyCreateRequestDataAttributesStepsItemsAssignment(value)
	return nil
}

// NewEscalationPolicyCreateRequestDataAttributesStepsItemsAssignmentFromValue returns a pointer to a valid EscalationPolicyCreateRequestDataAttributesStepsItemsAssignment
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEscalationPolicyCreateRequestDataAttributesStepsItemsAssignmentFromValue(v string) (*EscalationPolicyCreateRequestDataAttributesStepsItemsAssignment, error) {
	ev := EscalationPolicyCreateRequestDataAttributesStepsItemsAssignment(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EscalationPolicyCreateRequestDataAttributesStepsItemsAssignment: valid values are %v", v, allowedEscalationPolicyCreateRequestDataAttributesStepsItemsAssignmentEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EscalationPolicyCreateRequestDataAttributesStepsItemsAssignment) IsValid() bool {
	for _, existing := range allowedEscalationPolicyCreateRequestDataAttributesStepsItemsAssignmentEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EscalationPolicyCreateRequestDataAttributesStepsItemsAssignment value.
func (v EscalationPolicyCreateRequestDataAttributesStepsItemsAssignment) Ptr() *EscalationPolicyCreateRequestDataAttributesStepsItemsAssignment {
	return &v
}
