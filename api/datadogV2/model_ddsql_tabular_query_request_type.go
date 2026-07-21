// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DdsqlTabularQueryRequestType JSON:API resource type for a DDSQL tabular query request.
type DdsqlTabularQueryRequestType string

// List of DdsqlTabularQueryRequestType.
const (
	DDSQLTABULARQUERYREQUESTTYPE_DDSQL_QUERY_REQUEST DdsqlTabularQueryRequestType = "ddsql_query_request"
)

var allowedDdsqlTabularQueryRequestTypeEnumValues = []DdsqlTabularQueryRequestType{
	DDSQLTABULARQUERYREQUESTTYPE_DDSQL_QUERY_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DdsqlTabularQueryRequestType) GetAllowedValues() []DdsqlTabularQueryRequestType {
	return allowedDdsqlTabularQueryRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DdsqlTabularQueryRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DdsqlTabularQueryRequestType(value)
	return nil
}

// NewDdsqlTabularQueryRequestTypeFromValue returns a pointer to a valid DdsqlTabularQueryRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDdsqlTabularQueryRequestTypeFromValue(v string) (*DdsqlTabularQueryRequestType, error) {
	ev := DdsqlTabularQueryRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DdsqlTabularQueryRequestType: valid values are %v", v, allowedDdsqlTabularQueryRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DdsqlTabularQueryRequestType) IsValid() bool {
	for _, existing := range allowedDdsqlTabularQueryRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DdsqlTabularQueryRequestType value.
func (v DdsqlTabularQueryRequestType) Ptr() *DdsqlTabularQueryRequestType {
	return &v
}
