// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EscalationPolicyUpdateRequestDataAttributesStepsItemsAssignment Specifies how this escalation step will assign targets (example `default` or `round-robin`).
type EscalationPolicyUpdateRequestDataAttributesStepsItemsAssignment string

// List of EscalationPolicyUpdateRequestDataAttributesStepsItemsAssignment.
const (
	ESCALATIONPOLICYUPDATEREQUESTDATAATTRIBUTESSTEPSITEMSASSIGNMENT_DEFAULT     EscalationPolicyUpdateRequestDataAttributesStepsItemsAssignment = "default"
	ESCALATIONPOLICYUPDATEREQUESTDATAATTRIBUTESSTEPSITEMSASSIGNMENT_ROUND_ROBIN EscalationPolicyUpdateRequestDataAttributesStepsItemsAssignment = "round-robin"
)

var allowedEscalationPolicyUpdateRequestDataAttributesStepsItemsAssignmentEnumValues = []EscalationPolicyUpdateRequestDataAttributesStepsItemsAssignment{
	ESCALATIONPOLICYUPDATEREQUESTDATAATTRIBUTESSTEPSITEMSASSIGNMENT_DEFAULT,
	ESCALATIONPOLICYUPDATEREQUESTDATAATTRIBUTESSTEPSITEMSASSIGNMENT_ROUND_ROBIN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EscalationPolicyUpdateRequestDataAttributesStepsItemsAssignment) GetAllowedValues() []EscalationPolicyUpdateRequestDataAttributesStepsItemsAssignment {
	return allowedEscalationPolicyUpdateRequestDataAttributesStepsItemsAssignmentEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EscalationPolicyUpdateRequestDataAttributesStepsItemsAssignment) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EscalationPolicyUpdateRequestDataAttributesStepsItemsAssignment(value)
	return nil
}

// NewEscalationPolicyUpdateRequestDataAttributesStepsItemsAssignmentFromValue returns a pointer to a valid EscalationPolicyUpdateRequestDataAttributesStepsItemsAssignment
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEscalationPolicyUpdateRequestDataAttributesStepsItemsAssignmentFromValue(v string) (*EscalationPolicyUpdateRequestDataAttributesStepsItemsAssignment, error) {
	ev := EscalationPolicyUpdateRequestDataAttributesStepsItemsAssignment(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EscalationPolicyUpdateRequestDataAttributesStepsItemsAssignment: valid values are %v", v, allowedEscalationPolicyUpdateRequestDataAttributesStepsItemsAssignmentEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EscalationPolicyUpdateRequestDataAttributesStepsItemsAssignment) IsValid() bool {
	for _, existing := range allowedEscalationPolicyUpdateRequestDataAttributesStepsItemsAssignmentEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EscalationPolicyUpdateRequestDataAttributesStepsItemsAssignment value.
func (v EscalationPolicyUpdateRequestDataAttributesStepsItemsAssignment) Ptr() *EscalationPolicyUpdateRequestDataAttributesStepsItemsAssignment {
	return &v
}
