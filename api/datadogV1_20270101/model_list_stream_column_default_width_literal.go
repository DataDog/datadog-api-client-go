// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListStreamColumnDefaultWidthLiteral Test-only generated schema for the dashboard versioning proof of concept.
type ListStreamColumnDefaultWidthLiteral string

// List of ListStreamColumnDefaultWidthLiteral.
const (
	LISTSTREAMCOLUMNDEFAULTWIDTHLITERAL_AUTO        ListStreamColumnDefaultWidthLiteral = "auto"
	LISTSTREAMCOLUMNDEFAULTWIDTHLITERAL_COMPACT     ListStreamColumnDefaultWidthLiteral = "compact"
	LISTSTREAMCOLUMNDEFAULTWIDTHLITERAL_FULL        ListStreamColumnDefaultWidthLiteral = "full"
	LISTSTREAMCOLUMNDEFAULTWIDTHLITERAL_FIT_CONTENT ListStreamColumnDefaultWidthLiteral = "fit-content"
)

var allowedListStreamColumnDefaultWidthLiteralEnumValues = []ListStreamColumnDefaultWidthLiteral{
	LISTSTREAMCOLUMNDEFAULTWIDTHLITERAL_AUTO,
	LISTSTREAMCOLUMNDEFAULTWIDTHLITERAL_COMPACT,
	LISTSTREAMCOLUMNDEFAULTWIDTHLITERAL_FULL,
	LISTSTREAMCOLUMNDEFAULTWIDTHLITERAL_FIT_CONTENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ListStreamColumnDefaultWidthLiteral) GetAllowedValues() []ListStreamColumnDefaultWidthLiteral {
	return allowedListStreamColumnDefaultWidthLiteralEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ListStreamColumnDefaultWidthLiteral) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ListStreamColumnDefaultWidthLiteral(value)
	return nil
}

// NewListStreamColumnDefaultWidthLiteralFromValue returns a pointer to a valid ListStreamColumnDefaultWidthLiteral
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewListStreamColumnDefaultWidthLiteralFromValue(v string) (*ListStreamColumnDefaultWidthLiteral, error) {
	ev := ListStreamColumnDefaultWidthLiteral(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ListStreamColumnDefaultWidthLiteral: valid values are %v", v, allowedListStreamColumnDefaultWidthLiteralEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ListStreamColumnDefaultWidthLiteral) IsValid() bool {
	for _, existing := range allowedListStreamColumnDefaultWidthLiteralEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ListStreamColumnDefaultWidthLiteral value.
func (v ListStreamColumnDefaultWidthLiteral) Ptr() *ListStreamColumnDefaultWidthLiteral {
	return &v
}
