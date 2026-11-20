// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// Hour24Format Test-only generated schema for the dashboard versioning proof of concept.
type Hour24Format string

// List of Hour24Format.
const (
	HOUR24FORMAT_ZERO         Hour24Format = "0"
	HOUR24FORMAT_ONE          Hour24Format = "1"
	HOUR24FORMAT_TWO          Hour24Format = "2"
	HOUR24FORMAT_THREE        Hour24Format = "3"
	HOUR24FORMAT_FOUR         Hour24Format = "4"
	HOUR24FORMAT_FIVE         Hour24Format = "5"
	HOUR24FORMAT_SIX          Hour24Format = "6"
	HOUR24FORMAT_SEVEN        Hour24Format = "7"
	HOUR24FORMAT_EIGHT        Hour24Format = "8"
	HOUR24FORMAT_NINE         Hour24Format = "9"
	HOUR24FORMAT_TEN          Hour24Format = "10"
	HOUR24FORMAT_ELEVEN       Hour24Format = "11"
	HOUR24FORMAT_TWELVE       Hour24Format = "12"
	HOUR24FORMAT_THIRTEEN     Hour24Format = "13"
	HOUR24FORMAT_FOURTEEN     Hour24Format = "14"
	HOUR24FORMAT_FIFTEEN      Hour24Format = "15"
	HOUR24FORMAT_SIXTEEN      Hour24Format = "16"
	HOUR24FORMAT_SEVENTEEN    Hour24Format = "17"
	HOUR24FORMAT_EIGHTEEN     Hour24Format = "18"
	HOUR24FORMAT_NINETEEN     Hour24Format = "19"
	HOUR24FORMAT_TWENTY       Hour24Format = "20"
	HOUR24FORMAT_TWENTY_ONE   Hour24Format = "21"
	HOUR24FORMAT_TWENTY_TWO   Hour24Format = "22"
	HOUR24FORMAT_TWENTY_THREE Hour24Format = "23"
)

var allowedHour24FormatEnumValues = []Hour24Format{
	HOUR24FORMAT_ZERO,
	HOUR24FORMAT_ONE,
	HOUR24FORMAT_TWO,
	HOUR24FORMAT_THREE,
	HOUR24FORMAT_FOUR,
	HOUR24FORMAT_FIVE,
	HOUR24FORMAT_SIX,
	HOUR24FORMAT_SEVEN,
	HOUR24FORMAT_EIGHT,
	HOUR24FORMAT_NINE,
	HOUR24FORMAT_TEN,
	HOUR24FORMAT_ELEVEN,
	HOUR24FORMAT_TWELVE,
	HOUR24FORMAT_THIRTEEN,
	HOUR24FORMAT_FOURTEEN,
	HOUR24FORMAT_FIFTEEN,
	HOUR24FORMAT_SIXTEEN,
	HOUR24FORMAT_SEVENTEEN,
	HOUR24FORMAT_EIGHTEEN,
	HOUR24FORMAT_NINETEEN,
	HOUR24FORMAT_TWENTY,
	HOUR24FORMAT_TWENTY_ONE,
	HOUR24FORMAT_TWENTY_TWO,
	HOUR24FORMAT_TWENTY_THREE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *Hour24Format) GetAllowedValues() []Hour24Format {
	return allowedHour24FormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *Hour24Format) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = Hour24Format(value)
	return nil
}

// NewHour24FormatFromValue returns a pointer to a valid Hour24Format
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHour24FormatFromValue(v string) (*Hour24Format, error) {
	ev := Hour24Format(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for Hour24Format: valid values are %v", v, allowedHour24FormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v Hour24Format) IsValid() bool {
	for _, existing := range allowedHour24FormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Hour24Format value.
func (v Hour24Format) Ptr() *Hour24Format {
	return &v
}
