// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMPermanentRetentionFilterID The identifier of a permanent RUM retention filter.
type RUMPermanentRetentionFilterID string

// List of RUMPermanentRetentionFilterID.
const (
	RUMPERMANENTRETENTIONFILTERID_RUM_APM_FLAT_SAMPLING  RUMPermanentRetentionFilterID = "rum_apm_flat_sampling"
	RUMPERMANENTRETENTIONFILTERID_SYNTHETICS_SESSIONS    RUMPermanentRetentionFilterID = "synthetics_sessions"
	RUMPERMANENTRETENTIONFILTERID_FORCED_REPLAY_SESSIONS RUMPermanentRetentionFilterID = "forced_replay_sessions"
)

var allowedRUMPermanentRetentionFilterIDEnumValues = []RUMPermanentRetentionFilterID{
	RUMPERMANENTRETENTIONFILTERID_RUM_APM_FLAT_SAMPLING,
	RUMPERMANENTRETENTIONFILTERID_SYNTHETICS_SESSIONS,
	RUMPERMANENTRETENTIONFILTERID_FORCED_REPLAY_SESSIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RUMPermanentRetentionFilterID) GetAllowedValues() []RUMPermanentRetentionFilterID {
	return allowedRUMPermanentRetentionFilterIDEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RUMPermanentRetentionFilterID) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMPermanentRetentionFilterID(value)
	return nil
}

// NewRUMPermanentRetentionFilterIDFromValue returns a pointer to a valid RUMPermanentRetentionFilterID
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRUMPermanentRetentionFilterIDFromValue(v string) (*RUMPermanentRetentionFilterID, error) {
	ev := RUMPermanentRetentionFilterID(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RUMPermanentRetentionFilterID: valid values are %v", v, allowedRUMPermanentRetentionFilterIDEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RUMPermanentRetentionFilterID) IsValid() bool {
	for _, existing := range allowedRUMPermanentRetentionFilterIDEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMPermanentRetentionFilterID value.
func (v RUMPermanentRetentionFilterID) Ptr() *RUMPermanentRetentionFilterID {
	return &v
}
