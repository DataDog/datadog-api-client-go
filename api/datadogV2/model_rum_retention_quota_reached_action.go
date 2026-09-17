// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumRetentionQuotaReachedAction The action to take when the session quota is reached.
type RumRetentionQuotaReachedAction string

// List of RumRetentionQuotaReachedAction.
const (
	RUMRETENTIONQUOTAREACHEDACTION_STOP     RumRetentionQuotaReachedAction = "stop"
	RUMRETENTIONQUOTAREACHEDACTION_SLOWDOWN RumRetentionQuotaReachedAction = "slowdown"
)

var allowedRumRetentionQuotaReachedActionEnumValues = []RumRetentionQuotaReachedAction{
	RUMRETENTIONQUOTAREACHEDACTION_STOP,
	RUMRETENTIONQUOTAREACHEDACTION_SLOWDOWN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumRetentionQuotaReachedAction) GetAllowedValues() []RumRetentionQuotaReachedAction {
	return allowedRumRetentionQuotaReachedActionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumRetentionQuotaReachedAction) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumRetentionQuotaReachedAction(value)
	return nil
}

// NewRumRetentionQuotaReachedActionFromValue returns a pointer to a valid RumRetentionQuotaReachedAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumRetentionQuotaReachedActionFromValue(v string) (*RumRetentionQuotaReachedAction, error) {
	ev := RumRetentionQuotaReachedAction(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumRetentionQuotaReachedAction: valid values are %v", v, allowedRumRetentionQuotaReachedActionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumRetentionQuotaReachedAction) IsValid() bool {
	for _, existing := range allowedRumRetentionQuotaReachedActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumRetentionQuotaReachedAction value.
func (v RumRetentionQuotaReachedAction) Ptr() *RumRetentionQuotaReachedAction {
	return &v
}
