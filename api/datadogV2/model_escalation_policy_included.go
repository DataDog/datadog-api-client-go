// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EscalationPolicyIncluded - Represents included related resources when retrieving an escalation policy, such as teams, steps, or targets.
type EscalationPolicyIncluded struct {
	TeamReference        *TeamReference
	EscalationPolicyStep *EscalationPolicyStep
	UserTarget           *UserTarget
	ScheduleTarget       *ScheduleTarget
	TeamTarget           *TeamTarget

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// TeamReferenceAsEscalationPolicyIncluded is a convenience function that returns TeamReference wrapped in EscalationPolicyIncluded.
func TeamReferenceAsEscalationPolicyIncluded(v *TeamReference) EscalationPolicyIncluded {
	return EscalationPolicyIncluded{TeamReference: v}
}

// EscalationPolicyStepAsEscalationPolicyIncluded is a convenience function that returns EscalationPolicyStep wrapped in EscalationPolicyIncluded.
func EscalationPolicyStepAsEscalationPolicyIncluded(v *EscalationPolicyStep) EscalationPolicyIncluded {
	return EscalationPolicyIncluded{EscalationPolicyStep: v}
}

// UserTargetAsEscalationPolicyIncluded is a convenience function that returns UserTarget wrapped in EscalationPolicyIncluded.
func UserTargetAsEscalationPolicyIncluded(v *UserTarget) EscalationPolicyIncluded {
	return EscalationPolicyIncluded{UserTarget: v}
}

// ScheduleTargetAsEscalationPolicyIncluded is a convenience function that returns ScheduleTarget wrapped in EscalationPolicyIncluded.
func ScheduleTargetAsEscalationPolicyIncluded(v *ScheduleTarget) EscalationPolicyIncluded {
	return EscalationPolicyIncluded{ScheduleTarget: v}
}

// TeamTargetAsEscalationPolicyIncluded is a convenience function that returns TeamTarget wrapped in EscalationPolicyIncluded.
func TeamTargetAsEscalationPolicyIncluded(v *TeamTarget) EscalationPolicyIncluded {
	return EscalationPolicyIncluded{TeamTarget: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *EscalationPolicyIncluded) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into TeamReference
	err = datadog.Unmarshal(data, &obj.TeamReference)
	if err == nil {
		if obj.TeamReference != nil && obj.TeamReference.UnparsedObject == nil {
			jsonTeamReference, _ := datadog.Marshal(obj.TeamReference)
			if string(jsonTeamReference) == "{}" { // empty struct
				obj.TeamReference = nil
			} else {
				match++
			}
		} else {
			obj.TeamReference = nil
		}
	} else {
		obj.TeamReference = nil
	}

	// try to unmarshal data into EscalationPolicyStep
	err = datadog.Unmarshal(data, &obj.EscalationPolicyStep)
	if err == nil {
		if obj.EscalationPolicyStep != nil && obj.EscalationPolicyStep.UnparsedObject == nil {
			jsonEscalationPolicyStep, _ := datadog.Marshal(obj.EscalationPolicyStep)
			if string(jsonEscalationPolicyStep) == "{}" { // empty struct
				obj.EscalationPolicyStep = nil
			} else {
				match++
			}
		} else {
			obj.EscalationPolicyStep = nil
		}
	} else {
		obj.EscalationPolicyStep = nil
	}

	// try to unmarshal data into UserTarget
	err = datadog.Unmarshal(data, &obj.UserTarget)
	if err == nil {
		if obj.UserTarget != nil && obj.UserTarget.UnparsedObject == nil {
			jsonUserTarget, _ := datadog.Marshal(obj.UserTarget)
			if string(jsonUserTarget) == "{}" { // empty struct
				obj.UserTarget = nil
			} else {
				match++
			}
		} else {
			obj.UserTarget = nil
		}
	} else {
		obj.UserTarget = nil
	}

	// try to unmarshal data into ScheduleTarget
	err = datadog.Unmarshal(data, &obj.ScheduleTarget)
	if err == nil {
		if obj.ScheduleTarget != nil && obj.ScheduleTarget.UnparsedObject == nil {
			jsonScheduleTarget, _ := datadog.Marshal(obj.ScheduleTarget)
			if string(jsonScheduleTarget) == "{}" { // empty struct
				obj.ScheduleTarget = nil
			} else {
				match++
			}
		} else {
			obj.ScheduleTarget = nil
		}
	} else {
		obj.ScheduleTarget = nil
	}

	// try to unmarshal data into TeamTarget
	err = datadog.Unmarshal(data, &obj.TeamTarget)
	if err == nil {
		if obj.TeamTarget != nil && obj.TeamTarget.UnparsedObject == nil {
			jsonTeamTarget, _ := datadog.Marshal(obj.TeamTarget)
			if string(jsonTeamTarget) == "{}" { // empty struct
				obj.TeamTarget = nil
			} else {
				match++
			}
		} else {
			obj.TeamTarget = nil
		}
	} else {
		obj.TeamTarget = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.TeamReference = nil
		obj.EscalationPolicyStep = nil
		obj.UserTarget = nil
		obj.ScheduleTarget = nil
		obj.TeamTarget = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj EscalationPolicyIncluded) MarshalJSON() ([]byte, error) {
	if obj.TeamReference != nil {
		return datadog.Marshal(&obj.TeamReference)
	}

	if obj.EscalationPolicyStep != nil {
		return datadog.Marshal(&obj.EscalationPolicyStep)
	}

	if obj.UserTarget != nil {
		return datadog.Marshal(&obj.UserTarget)
	}

	if obj.ScheduleTarget != nil {
		return datadog.Marshal(&obj.ScheduleTarget)
	}

	if obj.TeamTarget != nil {
		return datadog.Marshal(&obj.TeamTarget)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *EscalationPolicyIncluded) GetActualInstance() interface{} {
	if obj.TeamReference != nil {
		return obj.TeamReference
	}

	if obj.EscalationPolicyStep != nil {
		return obj.EscalationPolicyStep
	}

	if obj.UserTarget != nil {
		return obj.UserTarget
	}

	if obj.ScheduleTarget != nil {
		return obj.ScheduleTarget
	}

	if obj.TeamTarget != nil {
		return obj.TeamTarget
	}

	// all schemas are nil
	return nil
}
