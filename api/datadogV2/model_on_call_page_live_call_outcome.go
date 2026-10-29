// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnCallPageLiveCallOutcome The outcome of a live call that has ended.
type OnCallPageLiveCallOutcome string

// List of OnCallPageLiveCallOutcome.
const (
	ONCALLPAGELIVECALLOUTCOME_SUCCESS_CALL_ANSWERED         OnCallPageLiveCallOutcome = "success_call_answered"
	ONCALLPAGELIVECALLOUTCOME_FAILURE_NO_RESPONDER          OnCallPageLiveCallOutcome = "failure_no_responder"
	ONCALLPAGELIVECALLOUTCOME_FAILURE_CALLER_ABANDONED      OnCallPageLiveCallOutcome = "failure_caller_abandoned"
	ONCALLPAGELIVECALLOUTCOME_FAILURE_OUTSIDE_SUPPORT_HOURS OnCallPageLiveCallOutcome = "failure_outside_support_hours"
	ONCALLPAGELIVECALLOUTCOME_FAILURE_INTERNAL_ERROR        OnCallPageLiveCallOutcome = "failure_internal_error"
)

var allowedOnCallPageLiveCallOutcomeEnumValues = []OnCallPageLiveCallOutcome{
	ONCALLPAGELIVECALLOUTCOME_SUCCESS_CALL_ANSWERED,
	ONCALLPAGELIVECALLOUTCOME_FAILURE_NO_RESPONDER,
	ONCALLPAGELIVECALLOUTCOME_FAILURE_CALLER_ABANDONED,
	ONCALLPAGELIVECALLOUTCOME_FAILURE_OUTSIDE_SUPPORT_HOURS,
	ONCALLPAGELIVECALLOUTCOME_FAILURE_INTERNAL_ERROR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OnCallPageLiveCallOutcome) GetAllowedValues() []OnCallPageLiveCallOutcome {
	return allowedOnCallPageLiveCallOutcomeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OnCallPageLiveCallOutcome) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OnCallPageLiveCallOutcome(value)
	return nil
}

// NewOnCallPageLiveCallOutcomeFromValue returns a pointer to a valid OnCallPageLiveCallOutcome
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOnCallPageLiveCallOutcomeFromValue(v string) (*OnCallPageLiveCallOutcome, error) {
	ev := OnCallPageLiveCallOutcome(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OnCallPageLiveCallOutcome: valid values are %v", v, allowedOnCallPageLiveCallOutcomeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OnCallPageLiveCallOutcome) IsValid() bool {
	for _, existing := range allowedOnCallPageLiveCallOutcomeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OnCallPageLiveCallOutcome value.
func (v OnCallPageLiveCallOutcome) Ptr() *OnCallPageLiveCallOutcome {
	return &v
}
