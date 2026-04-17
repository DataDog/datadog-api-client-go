// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FlagSuggestionAction The type of change action for a suggestion.
type FlagSuggestionAction string

// List of FlagSuggestionAction.
const (
	FLAGSUGGESTIONACTION_CREATED    FlagSuggestionAction = "created"
	FLAGSUGGESTIONACTION_UPDATED    FlagSuggestionAction = "updated"
	FLAGSUGGESTIONACTION_DELETED    FlagSuggestionAction = "deleted"
	FLAGSUGGESTIONACTION_ARCHIVED   FlagSuggestionAction = "archived"
	FLAGSUGGESTIONACTION_UNARCHIVED FlagSuggestionAction = "unarchived"
	FLAGSUGGESTIONACTION_STARTED    FlagSuggestionAction = "started"
	FLAGSUGGESTIONACTION_STOPPED    FlagSuggestionAction = "stopped"
	FLAGSUGGESTIONACTION_PAUSED     FlagSuggestionAction = "paused"
	FLAGSUGGESTIONACTION_UNPAUSED   FlagSuggestionAction = "unpaused"
)

var allowedFlagSuggestionActionEnumValues = []FlagSuggestionAction{
	FLAGSUGGESTIONACTION_CREATED,
	FLAGSUGGESTIONACTION_UPDATED,
	FLAGSUGGESTIONACTION_DELETED,
	FLAGSUGGESTIONACTION_ARCHIVED,
	FLAGSUGGESTIONACTION_UNARCHIVED,
	FLAGSUGGESTIONACTION_STARTED,
	FLAGSUGGESTIONACTION_STOPPED,
	FLAGSUGGESTIONACTION_PAUSED,
	FLAGSUGGESTIONACTION_UNPAUSED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FlagSuggestionAction) GetAllowedValues() []FlagSuggestionAction {
	return allowedFlagSuggestionActionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FlagSuggestionAction) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FlagSuggestionAction(value)
	return nil
}

// NewFlagSuggestionActionFromValue returns a pointer to a valid FlagSuggestionAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFlagSuggestionActionFromValue(v string) (*FlagSuggestionAction, error) {
	ev := FlagSuggestionAction(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FlagSuggestionAction: valid values are %v", v, allowedFlagSuggestionActionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FlagSuggestionAction) IsValid() bool {
	for _, existing := range allowedFlagSuggestionActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FlagSuggestionAction value.
func (v FlagSuggestionAction) Ptr() *FlagSuggestionAction {
	return &v
}
