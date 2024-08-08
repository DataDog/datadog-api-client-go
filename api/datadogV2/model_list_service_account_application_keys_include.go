// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListServiceAccountApplicationKeysInclude The definition of ListServiceAccountApplicationKeysInclude object.
type ListServiceAccountApplicationKeysInclude string

// List of ListServiceAccountApplicationKeysInclude.
const (
	LISTSERVICEACCOUNTAPPLICATIONKEYSINCLUDE_LEAK_INFORMATION ListServiceAccountApplicationKeysInclude = "leak_information"
	LISTSERVICEACCOUNTAPPLICATIONKEYSINCLUDE_OWNED_BY         ListServiceAccountApplicationKeysInclude = "owned_by"
)

var allowedListServiceAccountApplicationKeysIncludeEnumValues = []ListServiceAccountApplicationKeysInclude{
	LISTSERVICEACCOUNTAPPLICATIONKEYSINCLUDE_LEAK_INFORMATION,
	LISTSERVICEACCOUNTAPPLICATIONKEYSINCLUDE_OWNED_BY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ListServiceAccountApplicationKeysInclude) GetAllowedValues() []ListServiceAccountApplicationKeysInclude {
	return allowedListServiceAccountApplicationKeysIncludeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ListServiceAccountApplicationKeysInclude) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ListServiceAccountApplicationKeysInclude(value)
	return nil
}

// NewListServiceAccountApplicationKeysIncludeFromValue returns a pointer to a valid ListServiceAccountApplicationKeysInclude
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewListServiceAccountApplicationKeysIncludeFromValue(v string) (*ListServiceAccountApplicationKeysInclude, error) {
	ev := ListServiceAccountApplicationKeysInclude(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ListServiceAccountApplicationKeysInclude: valid values are %v", v, allowedListServiceAccountApplicationKeysIncludeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ListServiceAccountApplicationKeysInclude) IsValid() bool {
	for _, existing := range allowedListServiceAccountApplicationKeysIncludeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ListServiceAccountApplicationKeysInclude value.
func (v ListServiceAccountApplicationKeysInclude) Ptr() *ListServiceAccountApplicationKeysInclude {
	return &v
}
