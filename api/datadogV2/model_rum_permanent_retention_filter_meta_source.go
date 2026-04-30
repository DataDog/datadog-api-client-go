// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumPermanentRetentionFilterMetaSource The source of the last update to a permanent retention filter.
type RumPermanentRetentionFilterMetaSource string

// List of RumPermanentRetentionFilterMetaSource.
const (
	RUMPERMANENTRETENTIONFILTERMETASOURCE_DEFAULT   RumPermanentRetentionFilterMetaSource = "default"
	RUMPERMANENTRETENTIONFILTERMETASOURCE_UI        RumPermanentRetentionFilterMetaSource = "ui"
	RUMPERMANENTRETENTIONFILTERMETASOURCE_TERRAFORM RumPermanentRetentionFilterMetaSource = "terraform"
)

var allowedRumPermanentRetentionFilterMetaSourceEnumValues = []RumPermanentRetentionFilterMetaSource{
	RUMPERMANENTRETENTIONFILTERMETASOURCE_DEFAULT,
	RUMPERMANENTRETENTIONFILTERMETASOURCE_UI,
	RUMPERMANENTRETENTIONFILTERMETASOURCE_TERRAFORM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumPermanentRetentionFilterMetaSource) GetAllowedValues() []RumPermanentRetentionFilterMetaSource {
	return allowedRumPermanentRetentionFilterMetaSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumPermanentRetentionFilterMetaSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumPermanentRetentionFilterMetaSource(value)
	return nil
}

// NewRumPermanentRetentionFilterMetaSourceFromValue returns a pointer to a valid RumPermanentRetentionFilterMetaSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumPermanentRetentionFilterMetaSourceFromValue(v string) (*RumPermanentRetentionFilterMetaSource, error) {
	ev := RumPermanentRetentionFilterMetaSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumPermanentRetentionFilterMetaSource: valid values are %v", v, allowedRumPermanentRetentionFilterMetaSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumPermanentRetentionFilterMetaSource) IsValid() bool {
	for _, existing := range allowedRumPermanentRetentionFilterMetaSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumPermanentRetentionFilterMetaSource value.
func (v RumPermanentRetentionFilterMetaSource) Ptr() *RumPermanentRetentionFilterMetaSource {
	return &v
}
