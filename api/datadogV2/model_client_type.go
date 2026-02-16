// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ClientType The client type for action filtering.
type ClientType string

// List of ClientType.
const (
	CLIENTTYPE_WORKFLOWS   ClientType = "workflows"
	CLIENTTYPE_APP_BUILDER ClientType = "app_builder"
	CLIENTTYPE_ACTIONS_API ClientType = "actions_api"
)

var allowedClientTypeEnumValues = []ClientType{
	CLIENTTYPE_WORKFLOWS,
	CLIENTTYPE_APP_BUILDER,
	CLIENTTYPE_ACTIONS_API,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ClientType) GetAllowedValues() []ClientType {
	return allowedClientTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ClientType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ClientType(value)
	return nil
}

// NewClientTypeFromValue returns a pointer to a valid ClientType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewClientTypeFromValue(v string) (*ClientType, error) {
	ev := ClientType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ClientType: valid values are %v", v, allowedClientTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ClientType) IsValid() bool {
	for _, existing := range allowedClientTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ClientType value.
func (v ClientType) Ptr() *ClientType {
	return &v
}
