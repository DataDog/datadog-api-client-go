// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumRetentionFilterTrack The product track a RUM retention filter applies to.
type RumRetentionFilterTrack string

// List of RumRetentionFilterTrack.
const (
	RUMRETENTIONFILTERTRACK_RUM               RumRetentionFilterTrack = "rum"
	RUMRETENTIONFILTERTRACK_PRODUCT_ANALYTICS RumRetentionFilterTrack = "product_analytics"
)

var allowedRumRetentionFilterTrackEnumValues = []RumRetentionFilterTrack{
	RUMRETENTIONFILTERTRACK_RUM,
	RUMRETENTIONFILTERTRACK_PRODUCT_ANALYTICS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumRetentionFilterTrack) GetAllowedValues() []RumRetentionFilterTrack {
	return allowedRumRetentionFilterTrackEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumRetentionFilterTrack) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumRetentionFilterTrack(value)
	return nil
}

// NewRumRetentionFilterTrackFromValue returns a pointer to a valid RumRetentionFilterTrack
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumRetentionFilterTrackFromValue(v string) (*RumRetentionFilterTrack, error) {
	ev := RumRetentionFilterTrack(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumRetentionFilterTrack: valid values are %v", v, allowedRumRetentionFilterTrackEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumRetentionFilterTrack) IsValid() bool {
	for _, existing := range allowedRumRetentionFilterTrackEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumRetentionFilterTrack value.
func (v RumRetentionFilterTrack) Ptr() *RumRetentionFilterTrack {
	return &v
}
