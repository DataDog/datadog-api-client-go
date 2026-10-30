// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumRetentionFilterMetaSource The source of the last change to a RUM retention filter.
type RumRetentionFilterMetaSource string

// List of RumRetentionFilterMetaSource.
const (
	RUMRETENTIONFILTERMETASOURCE_DEFAULT   RumRetentionFilterMetaSource = "default"
	RUMRETENTIONFILTERMETASOURCE_UI        RumRetentionFilterMetaSource = "ui"
	RUMRETENTIONFILTERMETASOURCE_TERRAFORM RumRetentionFilterMetaSource = "terraform"
	RUMRETENTIONFILTERMETASOURCE_UNKNOWN   RumRetentionFilterMetaSource = "unknown"
)

var allowedRumRetentionFilterMetaSourceEnumValues = []RumRetentionFilterMetaSource{
	RUMRETENTIONFILTERMETASOURCE_DEFAULT,
	RUMRETENTIONFILTERMETASOURCE_UI,
	RUMRETENTIONFILTERMETASOURCE_TERRAFORM,
	RUMRETENTIONFILTERMETASOURCE_UNKNOWN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumRetentionFilterMetaSource) GetAllowedValues() []RumRetentionFilterMetaSource {
	return allowedRumRetentionFilterMetaSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumRetentionFilterMetaSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumRetentionFilterMetaSource(value)
	return nil
}

// NewRumRetentionFilterMetaSourceFromValue returns a pointer to a valid RumRetentionFilterMetaSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumRetentionFilterMetaSourceFromValue(v string) (*RumRetentionFilterMetaSource, error) {
	ev := RumRetentionFilterMetaSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumRetentionFilterMetaSource: valid values are %v", v, allowedRumRetentionFilterMetaSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumRetentionFilterMetaSource) IsValid() bool {
	for _, existing := range allowedRumRetentionFilterMetaSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumRetentionFilterMetaSource value.
func (v RumRetentionFilterMetaSource) Ptr() *RumRetentionFilterMetaSource {
	return &v
}
