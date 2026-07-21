// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DdsqlTabularQueryResponseType JSON:API resource type for a DDSQL tabular query response.
type DdsqlTabularQueryResponseType string

// List of DdsqlTabularQueryResponseType.
const (
	DDSQLTABULARQUERYRESPONSETYPE_DDSQL_QUERY_RESPONSE DdsqlTabularQueryResponseType = "ddsql_query_response"
)

var allowedDdsqlTabularQueryResponseTypeEnumValues = []DdsqlTabularQueryResponseType{
	DDSQLTABULARQUERYRESPONSETYPE_DDSQL_QUERY_RESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DdsqlTabularQueryResponseType) GetAllowedValues() []DdsqlTabularQueryResponseType {
	return allowedDdsqlTabularQueryResponseTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DdsqlTabularQueryResponseType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DdsqlTabularQueryResponseType(value)
	return nil
}

// NewDdsqlTabularQueryResponseTypeFromValue returns a pointer to a valid DdsqlTabularQueryResponseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDdsqlTabularQueryResponseTypeFromValue(v string) (*DdsqlTabularQueryResponseType, error) {
	ev := DdsqlTabularQueryResponseType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DdsqlTabularQueryResponseType: valid values are %v", v, allowedDdsqlTabularQueryResponseTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DdsqlTabularQueryResponseType) IsValid() bool {
	for _, existing := range allowedDdsqlTabularQueryResponseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DdsqlTabularQueryResponseType value.
func (v DdsqlTabularQueryResponseType) Ptr() *DdsqlTabularQueryResponseType {
	return &v
}
