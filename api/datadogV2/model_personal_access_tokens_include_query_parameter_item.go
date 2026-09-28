// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PersonalAccessTokensIncludeQueryParameterItem Relationship object that should be included in the response.
type PersonalAccessTokensIncludeQueryParameterItem string

// List of PersonalAccessTokensIncludeQueryParameterItem.
const (
	PERSONALACCESSTOKENSINCLUDEQUERYPARAMETERITEM_LEAK_INFORMATION PersonalAccessTokensIncludeQueryParameterItem = "leak_information"
)

var allowedPersonalAccessTokensIncludeQueryParameterItemEnumValues = []PersonalAccessTokensIncludeQueryParameterItem{
	PERSONALACCESSTOKENSINCLUDEQUERYPARAMETERITEM_LEAK_INFORMATION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PersonalAccessTokensIncludeQueryParameterItem) GetAllowedValues() []PersonalAccessTokensIncludeQueryParameterItem {
	return allowedPersonalAccessTokensIncludeQueryParameterItemEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PersonalAccessTokensIncludeQueryParameterItem) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PersonalAccessTokensIncludeQueryParameterItem(value)
	return nil
}

// NewPersonalAccessTokensIncludeQueryParameterItemFromValue returns a pointer to a valid PersonalAccessTokensIncludeQueryParameterItem
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPersonalAccessTokensIncludeQueryParameterItemFromValue(v string) (*PersonalAccessTokensIncludeQueryParameterItem, error) {
	ev := PersonalAccessTokensIncludeQueryParameterItem(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PersonalAccessTokensIncludeQueryParameterItem: valid values are %v", v, allowedPersonalAccessTokensIncludeQueryParameterItemEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PersonalAccessTokensIncludeQueryParameterItem) IsValid() bool {
	for _, existing := range allowedPersonalAccessTokensIncludeQueryParameterItemEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PersonalAccessTokensIncludeQueryParameterItem value.
func (v PersonalAccessTokensIncludeQueryParameterItem) Ptr() *PersonalAccessTokensIncludeQueryParameterItem {
	return &v
}
