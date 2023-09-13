// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/goccy/go-json"
)

// S3FallbackDestinationType Type of the S3 archive destination.
type S3FallbackDestinationType string

// List of S3FallbackDestinationType.
const (
	S3FALLBACKDESTINATIONTYPE_S3 S3FallbackDestinationType = "s3"
)

var allowedS3FallbackDestinationTypeEnumValues = []S3FallbackDestinationType{
	S3FALLBACKDESTINATIONTYPE_S3,
}

// GetAllowedValues reeturns the list of possible values.
func (v *S3FallbackDestinationType) GetAllowedValues() []S3FallbackDestinationType {
	return allowedS3FallbackDestinationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *S3FallbackDestinationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = S3FallbackDestinationType(value)
	return nil
}

// NewS3FallbackDestinationTypeFromValue returns a pointer to a valid S3FallbackDestinationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewS3FallbackDestinationTypeFromValue(v string) (*S3FallbackDestinationType, error) {
	ev := S3FallbackDestinationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for S3FallbackDestinationType: valid values are %v", v, allowedS3FallbackDestinationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v S3FallbackDestinationType) IsValid() bool {
	for _, existing := range allowedS3FallbackDestinationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to S3FallbackDestinationType value.
func (v S3FallbackDestinationType) Ptr() *S3FallbackDestinationType {
	return &v
}
