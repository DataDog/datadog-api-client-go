// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnCallPageLinkRelationshipType How the parent resource relates to the child resource.
type OnCallPageLinkRelationshipType string

// List of OnCallPageLinkRelationshipType.
const (
	ONCALLPAGELINKRELATIONSHIPTYPE_RELATES_TO   OnCallPageLinkRelationshipType = "relates_to"
	ONCALLPAGELINKRELATIONSHIPTYPE_CAUSES       OnCallPageLinkRelationshipType = "causes"
	ONCALLPAGELINKRELATIONSHIPTYPE_BLOCKS       OnCallPageLinkRelationshipType = "blocks"
	ONCALLPAGELINKRELATIONSHIPTYPE_DUPLICATES   OnCallPageLinkRelationshipType = "duplicates"
	ONCALLPAGELINKRELATIONSHIPTYPE_PARENT_OF    OnCallPageLinkRelationshipType = "parent_of"
	ONCALLPAGELINKRELATIONSHIPTYPE_SUCCESSOR_OF OnCallPageLinkRelationshipType = "successor_of"
	ONCALLPAGELINKRELATIONSHIPTYPE_ESCALATES_TO OnCallPageLinkRelationshipType = "escalates_to"
	ONCALLPAGELINKRELATIONSHIPTYPE_PAGES        OnCallPageLinkRelationshipType = "pages"
)

var allowedOnCallPageLinkRelationshipTypeEnumValues = []OnCallPageLinkRelationshipType{
	ONCALLPAGELINKRELATIONSHIPTYPE_RELATES_TO,
	ONCALLPAGELINKRELATIONSHIPTYPE_CAUSES,
	ONCALLPAGELINKRELATIONSHIPTYPE_BLOCKS,
	ONCALLPAGELINKRELATIONSHIPTYPE_DUPLICATES,
	ONCALLPAGELINKRELATIONSHIPTYPE_PARENT_OF,
	ONCALLPAGELINKRELATIONSHIPTYPE_SUCCESSOR_OF,
	ONCALLPAGELINKRELATIONSHIPTYPE_ESCALATES_TO,
	ONCALLPAGELINKRELATIONSHIPTYPE_PAGES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OnCallPageLinkRelationshipType) GetAllowedValues() []OnCallPageLinkRelationshipType {
	return allowedOnCallPageLinkRelationshipTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OnCallPageLinkRelationshipType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OnCallPageLinkRelationshipType(value)
	return nil
}

// NewOnCallPageLinkRelationshipTypeFromValue returns a pointer to a valid OnCallPageLinkRelationshipType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOnCallPageLinkRelationshipTypeFromValue(v string) (*OnCallPageLinkRelationshipType, error) {
	ev := OnCallPageLinkRelationshipType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OnCallPageLinkRelationshipType: valid values are %v", v, allowedOnCallPageLinkRelationshipTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OnCallPageLinkRelationshipType) IsValid() bool {
	for _, existing := range allowedOnCallPageLinkRelationshipTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OnCallPageLinkRelationshipType value.
func (v OnCallPageLinkRelationshipType) Ptr() *OnCallPageLinkRelationshipType {
	return &v
}
