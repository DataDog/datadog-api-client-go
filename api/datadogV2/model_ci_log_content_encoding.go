// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CILogContentEncoding HTTP header used to compress the JSON request body.
type CILogContentEncoding string

// List of CILogContentEncoding.
const (
	CILOGCONTENTENCODING_IDENTITY CILogContentEncoding = "identity"
	CILOGCONTENTENCODING_GZIP     CILogContentEncoding = "gzip"
)

var allowedCILogContentEncodingEnumValues = []CILogContentEncoding{
	CILOGCONTENTENCODING_IDENTITY,
	CILOGCONTENTENCODING_GZIP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CILogContentEncoding) GetAllowedValues() []CILogContentEncoding {
	return allowedCILogContentEncodingEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CILogContentEncoding) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CILogContentEncoding(value)
	return nil
}

// NewCILogContentEncodingFromValue returns a pointer to a valid CILogContentEncoding
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCILogContentEncodingFromValue(v string) (*CILogContentEncoding, error) {
	ev := CILogContentEncoding(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CILogContentEncoding: valid values are %v", v, allowedCILogContentEncodingEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CILogContentEncoding) IsValid() bool {
	for _, existing := range allowedCILogContentEncodingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CILogContentEncoding value.
func (v CILogContentEncoding) Ptr() *CILogContentEncoding {
	return &v
}
