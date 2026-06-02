// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumHardcodedRetentionFilterMetaSource The source of the last update to a hardcoded retention filter.
type RumHardcodedRetentionFilterMetaSource string

// List of RumHardcodedRetentionFilterMetaSource.
const (
	RUMHARDCODEDRETENTIONFILTERMETASOURCE_DEFAULT   RumHardcodedRetentionFilterMetaSource = "default"
	RUMHARDCODEDRETENTIONFILTERMETASOURCE_UI        RumHardcodedRetentionFilterMetaSource = "ui"
	RUMHARDCODEDRETENTIONFILTERMETASOURCE_TERRAFORM RumHardcodedRetentionFilterMetaSource = "terraform"
)

var allowedRumHardcodedRetentionFilterMetaSourceEnumValues = []RumHardcodedRetentionFilterMetaSource{
	RUMHARDCODEDRETENTIONFILTERMETASOURCE_DEFAULT,
	RUMHARDCODEDRETENTIONFILTERMETASOURCE_UI,
	RUMHARDCODEDRETENTIONFILTERMETASOURCE_TERRAFORM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumHardcodedRetentionFilterMetaSource) GetAllowedValues() []RumHardcodedRetentionFilterMetaSource {
	return allowedRumHardcodedRetentionFilterMetaSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumHardcodedRetentionFilterMetaSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumHardcodedRetentionFilterMetaSource(value)
	return nil
}

// NewRumHardcodedRetentionFilterMetaSourceFromValue returns a pointer to a valid RumHardcodedRetentionFilterMetaSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumHardcodedRetentionFilterMetaSourceFromValue(v string) (*RumHardcodedRetentionFilterMetaSource, error) {
	ev := RumHardcodedRetentionFilterMetaSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumHardcodedRetentionFilterMetaSource: valid values are %v", v, allowedRumHardcodedRetentionFilterMetaSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumHardcodedRetentionFilterMetaSource) IsValid() bool {
	for _, existing := range allowedRumHardcodedRetentionFilterMetaSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumHardcodedRetentionFilterMetaSource value.
func (v RumHardcodedRetentionFilterMetaSource) Ptr() *RumHardcodedRetentionFilterMetaSource {
	return &v
}
