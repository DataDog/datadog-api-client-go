// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMRetentionFilterMetaSource The source of the last update to the retention filter. The value is `unknown` if the request doesn't specify a source.
type RUMRetentionFilterMetaSource string

// List of RUMRetentionFilterMetaSource.
const (
	RUMRETENTIONFILTERMETASOURCE_DEFAULT   RUMRetentionFilterMetaSource = "default"
	RUMRETENTIONFILTERMETASOURCE_UI        RUMRetentionFilterMetaSource = "ui"
	RUMRETENTIONFILTERMETASOURCE_TERRAFORM RUMRetentionFilterMetaSource = "terraform"
	RUMRETENTIONFILTERMETASOURCE_UNKNOWN   RUMRetentionFilterMetaSource = "unknown"
)

var allowedRUMRetentionFilterMetaSourceEnumValues = []RUMRetentionFilterMetaSource{
	RUMRETENTIONFILTERMETASOURCE_DEFAULT,
	RUMRETENTIONFILTERMETASOURCE_UI,
	RUMRETENTIONFILTERMETASOURCE_TERRAFORM,
	RUMRETENTIONFILTERMETASOURCE_UNKNOWN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RUMRetentionFilterMetaSource) GetAllowedValues() []RUMRetentionFilterMetaSource {
	return allowedRUMRetentionFilterMetaSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RUMRetentionFilterMetaSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMRetentionFilterMetaSource(value)
	return nil
}

// NewRUMRetentionFilterMetaSourceFromValue returns a pointer to a valid RUMRetentionFilterMetaSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRUMRetentionFilterMetaSourceFromValue(v string) (*RUMRetentionFilterMetaSource, error) {
	ev := RUMRetentionFilterMetaSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RUMRetentionFilterMetaSource: valid values are %v", v, allowedRUMRetentionFilterMetaSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RUMRetentionFilterMetaSource) IsValid() bool {
	for _, existing := range allowedRUMRetentionFilterMetaSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMRetentionFilterMetaSource value.
func (v RUMRetentionFilterMetaSource) Ptr() *RUMRetentionFilterMetaSource {
	return &v
}
