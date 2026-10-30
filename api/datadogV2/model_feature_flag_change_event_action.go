// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FeatureFlagChangeEventAction The kind of change.
type FeatureFlagChangeEventAction string

// List of FeatureFlagChangeEventAction.
const (
	FEATUREFLAGCHANGEEVENTACTION_CREATED    FeatureFlagChangeEventAction = "created"
	FEATUREFLAGCHANGEEVENTACTION_UPDATED    FeatureFlagChangeEventAction = "updated"
	FEATUREFLAGCHANGEEVENTACTION_DELETED    FeatureFlagChangeEventAction = "deleted"
	FEATUREFLAGCHANGEEVENTACTION_ARCHIVED   FeatureFlagChangeEventAction = "archived"
	FEATUREFLAGCHANGEEVENTACTION_UNARCHIVED FeatureFlagChangeEventAction = "unarchived"
	FEATUREFLAGCHANGEEVENTACTION_SCHEDULED  FeatureFlagChangeEventAction = "scheduled"
	FEATUREFLAGCHANGEEVENTACTION_STARTED    FeatureFlagChangeEventAction = "started"
	FEATUREFLAGCHANGEEVENTACTION_PAUSED     FeatureFlagChangeEventAction = "paused"
	FEATUREFLAGCHANGEEVENTACTION_UNPAUSED   FeatureFlagChangeEventAction = "unpaused"
	FEATUREFLAGCHANGEEVENTACTION_STOPPED    FeatureFlagChangeEventAction = "stopped"
)

var allowedFeatureFlagChangeEventActionEnumValues = []FeatureFlagChangeEventAction{
	FEATUREFLAGCHANGEEVENTACTION_CREATED,
	FEATUREFLAGCHANGEEVENTACTION_UPDATED,
	FEATUREFLAGCHANGEEVENTACTION_DELETED,
	FEATUREFLAGCHANGEEVENTACTION_ARCHIVED,
	FEATUREFLAGCHANGEEVENTACTION_UNARCHIVED,
	FEATUREFLAGCHANGEEVENTACTION_SCHEDULED,
	FEATUREFLAGCHANGEEVENTACTION_STARTED,
	FEATUREFLAGCHANGEEVENTACTION_PAUSED,
	FEATUREFLAGCHANGEEVENTACTION_UNPAUSED,
	FEATUREFLAGCHANGEEVENTACTION_STOPPED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FeatureFlagChangeEventAction) GetAllowedValues() []FeatureFlagChangeEventAction {
	return allowedFeatureFlagChangeEventActionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FeatureFlagChangeEventAction) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FeatureFlagChangeEventAction(value)
	return nil
}

// NewFeatureFlagChangeEventActionFromValue returns a pointer to a valid FeatureFlagChangeEventAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFeatureFlagChangeEventActionFromValue(v string) (*FeatureFlagChangeEventAction, error) {
	ev := FeatureFlagChangeEventAction(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FeatureFlagChangeEventAction: valid values are %v", v, allowedFeatureFlagChangeEventActionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FeatureFlagChangeEventAction) IsValid() bool {
	for _, existing := range allowedFeatureFlagChangeEventActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FeatureFlagChangeEventAction value.
func (v FeatureFlagChangeEventAction) Ptr() *FeatureFlagChangeEventAction {
	return &v
}
