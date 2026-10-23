// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnCallPageLinkEntityType The type of a resource linked to a page.
type OnCallPageLinkEntityType string

// List of OnCallPageLinkEntityType.
const (
	ONCALLPAGELINKENTITYTYPE_CASE               OnCallPageLinkEntityType = "case"
	ONCALLPAGELINKENTITYTYPE_INCIDENT           OnCallPageLinkEntityType = "incident"
	ONCALLPAGELINKENTITYTYPE_PAGE               OnCallPageLinkEntityType = "page"
	ONCALLPAGELINKENTITYTYPE_AGENT_CONVERSATION OnCallPageLinkEntityType = "agent_conversation"
)

var allowedOnCallPageLinkEntityTypeEnumValues = []OnCallPageLinkEntityType{
	ONCALLPAGELINKENTITYTYPE_CASE,
	ONCALLPAGELINKENTITYTYPE_INCIDENT,
	ONCALLPAGELINKENTITYTYPE_PAGE,
	ONCALLPAGELINKENTITYTYPE_AGENT_CONVERSATION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OnCallPageLinkEntityType) GetAllowedValues() []OnCallPageLinkEntityType {
	return allowedOnCallPageLinkEntityTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OnCallPageLinkEntityType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OnCallPageLinkEntityType(value)
	return nil
}

// NewOnCallPageLinkEntityTypeFromValue returns a pointer to a valid OnCallPageLinkEntityType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOnCallPageLinkEntityTypeFromValue(v string) (*OnCallPageLinkEntityType, error) {
	ev := OnCallPageLinkEntityType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OnCallPageLinkEntityType: valid values are %v", v, allowedOnCallPageLinkEntityTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OnCallPageLinkEntityType) IsValid() bool {
	for _, existing := range allowedOnCallPageLinkEntityTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OnCallPageLinkEntityType value.
func (v OnCallPageLinkEntityType) Ptr() *OnCallPageLinkEntityType {
	return &v
}
