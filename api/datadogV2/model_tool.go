// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// Tool The vulnerability tool.
type Tool string

// List of Tool.
const (
	TOOL_IAST  Tool = "IAST"
	TOOL_SCA   Tool = "SCA"
	TOOL_INFRA Tool = "Infra"
)

var allowedToolEnumValues = []Tool{
	TOOL_IAST,
	TOOL_SCA,
	TOOL_INFRA,
}

// GetAllowedValues reeturns the list of possible values.
func (v *Tool) GetAllowedValues() []Tool {
	return allowedToolEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *Tool) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = Tool(value)
	return nil
}

// NewToolFromValue returns a pointer to a valid Tool
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewToolFromValue(v string) (*Tool, error) {
	ev := Tool(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for Tool: valid values are %v", v, allowedToolEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v Tool) IsValid() bool {
	for _, existing := range allowedToolEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Tool value.
func (v Tool) Ptr() *Tool {
	return &v
}
