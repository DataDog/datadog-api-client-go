// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/goccy/go-json"
)

// GCSFallbackDestinationType Type of the GCS archive destination.
type GCSFallbackDestinationType string

// List of GCSFallbackDestinationType.
const (
	GCSFALLBACKDESTINATIONTYPE_GCS GCSFallbackDestinationType = "gcs"
)

var allowedGCSFallbackDestinationTypeEnumValues = []GCSFallbackDestinationType{
	GCSFALLBACKDESTINATIONTYPE_GCS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GCSFallbackDestinationType) GetAllowedValues() []GCSFallbackDestinationType {
	return allowedGCSFallbackDestinationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GCSFallbackDestinationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GCSFallbackDestinationType(value)
	return nil
}

// NewGCSFallbackDestinationTypeFromValue returns a pointer to a valid GCSFallbackDestinationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGCSFallbackDestinationTypeFromValue(v string) (*GCSFallbackDestinationType, error) {
	ev := GCSFallbackDestinationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GCSFallbackDestinationType: valid values are %v", v, allowedGCSFallbackDestinationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GCSFallbackDestinationType) IsValid() bool {
	for _, existing := range allowedGCSFallbackDestinationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GCSFallbackDestinationType value.
func (v GCSFallbackDestinationType) Ptr() *GCSFallbackDestinationType {
	return &v
}
