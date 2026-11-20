// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMRetentionFilterTrack The product whose events the retention filter applies to.
type RUMRetentionFilterTrack string

// List of RUMRetentionFilterTrack.
const (
	RUMRETENTIONFILTERTRACK_RUM               RUMRetentionFilterTrack = "rum"
	RUMRETENTIONFILTERTRACK_PRODUCT_ANALYTICS RUMRetentionFilterTrack = "product_analytics"
)

var allowedRUMRetentionFilterTrackEnumValues = []RUMRetentionFilterTrack{
	RUMRETENTIONFILTERTRACK_RUM,
	RUMRETENTIONFILTERTRACK_PRODUCT_ANALYTICS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RUMRetentionFilterTrack) GetAllowedValues() []RUMRetentionFilterTrack {
	return allowedRUMRetentionFilterTrackEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RUMRetentionFilterTrack) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMRetentionFilterTrack(value)
	return nil
}

// NewRUMRetentionFilterTrackFromValue returns a pointer to a valid RUMRetentionFilterTrack
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRUMRetentionFilterTrackFromValue(v string) (*RUMRetentionFilterTrack, error) {
	ev := RUMRetentionFilterTrack(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RUMRetentionFilterTrack: valid values are %v", v, allowedRUMRetentionFilterTrackEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RUMRetentionFilterTrack) IsValid() bool {
	for _, existing := range allowedRUMRetentionFilterTrackEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMRetentionFilterTrack value.
func (v RUMRetentionFilterTrack) Ptr() *RUMRetentionFilterTrack {
	return &v
}
