// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMRetentionQuotaReachedAction The action to take when the session quota is reached.
type RUMRetentionQuotaReachedAction string

// List of RUMRetentionQuotaReachedAction.
const (
	RUMRETENTIONQUOTAREACHEDACTION_STOP     RUMRetentionQuotaReachedAction = "stop"
	RUMRETENTIONQUOTAREACHEDACTION_SLOWDOWN RUMRetentionQuotaReachedAction = "slowdown"
)

var allowedRUMRetentionQuotaReachedActionEnumValues = []RUMRetentionQuotaReachedAction{
	RUMRETENTIONQUOTAREACHEDACTION_STOP,
	RUMRETENTIONQUOTAREACHEDACTION_SLOWDOWN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RUMRetentionQuotaReachedAction) GetAllowedValues() []RUMRetentionQuotaReachedAction {
	return allowedRUMRetentionQuotaReachedActionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RUMRetentionQuotaReachedAction) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMRetentionQuotaReachedAction(value)
	return nil
}

// NewRUMRetentionQuotaReachedActionFromValue returns a pointer to a valid RUMRetentionQuotaReachedAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRUMRetentionQuotaReachedActionFromValue(v string) (*RUMRetentionQuotaReachedAction, error) {
	ev := RUMRetentionQuotaReachedAction(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RUMRetentionQuotaReachedAction: valid values are %v", v, allowedRUMRetentionQuotaReachedActionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RUMRetentionQuotaReachedAction) IsValid() bool {
	for _, existing := range allowedRUMRetentionQuotaReachedActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMRetentionQuotaReachedAction value.
func (v RUMRetentionQuotaReachedAction) Ptr() *RUMRetentionQuotaReachedAction {
	return &v
}
