// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// STIXContentEncoding The content encoding applied to the request body.
type STIXContentEncoding string

// List of STIXContentEncoding.
const (
	STIXCONTENTENCODING_GZIP STIXContentEncoding = "gzip"
)

var allowedSTIXContentEncodingEnumValues = []STIXContentEncoding{
	STIXCONTENTENCODING_GZIP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *STIXContentEncoding) GetAllowedValues() []STIXContentEncoding {
	return allowedSTIXContentEncodingEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *STIXContentEncoding) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = STIXContentEncoding(value)
	return nil
}

// NewSTIXContentEncodingFromValue returns a pointer to a valid STIXContentEncoding
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSTIXContentEncodingFromValue(v string) (*STIXContentEncoding, error) {
	ev := STIXContentEncoding(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for STIXContentEncoding: valid values are %v", v, allowedSTIXContentEncodingEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v STIXContentEncoding) IsValid() bool {
	for _, existing := range allowedSTIXContentEncodingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to STIXContentEncoding value.
func (v STIXContentEncoding) Ptr() *STIXContentEncoding {
	return &v
}
