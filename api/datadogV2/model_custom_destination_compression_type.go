// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/goccy/go-json"
)

// CustomDestinationCompressionType The compression method used for the custom destination.
type CustomDestinationCompressionType string

// List of CustomDestinationCompressionType.
const (
	CUSTOMDESTINATIONCOMPRESSIONTYPE_GZIP_COMPRESSION CustomDestinationCompressionType = "GZIP"
	CUSTOMDESTINATIONCOMPRESSIONTYPE_NO_COMPRESSION   CustomDestinationCompressionType = "NONE"
)

var allowedCustomDestinationCompressionTypeEnumValues = []CustomDestinationCompressionType{
	CUSTOMDESTINATIONCOMPRESSIONTYPE_GZIP_COMPRESSION,
	CUSTOMDESTINATIONCOMPRESSIONTYPE_NO_COMPRESSION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CustomDestinationCompressionType) GetAllowedValues() []CustomDestinationCompressionType {
	return allowedCustomDestinationCompressionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CustomDestinationCompressionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CustomDestinationCompressionType(value)
	return nil
}

// NewCustomDestinationCompressionTypeFromValue returns a pointer to a valid CustomDestinationCompressionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCustomDestinationCompressionTypeFromValue(v string) (*CustomDestinationCompressionType, error) {
	ev := CustomDestinationCompressionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CustomDestinationCompressionType: valid values are %v", v, allowedCustomDestinationCompressionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CustomDestinationCompressionType) IsValid() bool {
	for _, existing := range allowedCustomDestinationCompressionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomDestinationCompressionType value.
func (v CustomDestinationCompressionType) Ptr() *CustomDestinationCompressionType {
	return &v
}
