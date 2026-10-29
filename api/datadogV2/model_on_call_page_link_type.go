// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnCallPageLinkType Page link resource type.
type OnCallPageLinkType string

// List of OnCallPageLinkType.
const (
	ONCALLPAGELINKTYPE_LINKS OnCallPageLinkType = "links"
)

var allowedOnCallPageLinkTypeEnumValues = []OnCallPageLinkType{
	ONCALLPAGELINKTYPE_LINKS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OnCallPageLinkType) GetAllowedValues() []OnCallPageLinkType {
	return allowedOnCallPageLinkTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OnCallPageLinkType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OnCallPageLinkType(value)
	return nil
}

// NewOnCallPageLinkTypeFromValue returns a pointer to a valid OnCallPageLinkType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOnCallPageLinkTypeFromValue(v string) (*OnCallPageLinkType, error) {
	ev := OnCallPageLinkType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OnCallPageLinkType: valid values are %v", v, allowedOnCallPageLinkTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OnCallPageLinkType) IsValid() bool {
	for _, existing := range allowedOnCallPageLinkTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OnCallPageLinkType value.
func (v OnCallPageLinkType) Ptr() *OnCallPageLinkType {
	return &v
}
