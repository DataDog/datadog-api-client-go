// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnCallPageCommentType Indicates that the resource is of type `comments`.
type OnCallPageCommentType string

// List of OnCallPageCommentType.
const (
	ONCALLPAGECOMMENTTYPE_COMMENTS OnCallPageCommentType = "comments"
)

var allowedOnCallPageCommentTypeEnumValues = []OnCallPageCommentType{
	ONCALLPAGECOMMENTTYPE_COMMENTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OnCallPageCommentType) GetAllowedValues() []OnCallPageCommentType {
	return allowedOnCallPageCommentTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OnCallPageCommentType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OnCallPageCommentType(value)
	return nil
}

// NewOnCallPageCommentTypeFromValue returns a pointer to a valid OnCallPageCommentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOnCallPageCommentTypeFromValue(v string) (*OnCallPageCommentType, error) {
	ev := OnCallPageCommentType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OnCallPageCommentType: valid values are %v", v, allowedOnCallPageCommentTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OnCallPageCommentType) IsValid() bool {
	for _, existing := range allowedOnCallPageCommentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OnCallPageCommentType value.
func (v OnCallPageCommentType) Ptr() *OnCallPageCommentType {
	return &v
}
