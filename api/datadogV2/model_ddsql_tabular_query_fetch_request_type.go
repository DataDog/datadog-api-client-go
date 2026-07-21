// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DdsqlTabularQueryFetchRequestType JSON:API resource type for a DDSQL tabular query fetch request.
type DdsqlTabularQueryFetchRequestType string

// List of DdsqlTabularQueryFetchRequestType.
const (
	DDSQLTABULARQUERYFETCHREQUESTTYPE_DDSQL_QUERY_FETCH_REQUEST DdsqlTabularQueryFetchRequestType = "ddsql_query_fetch_request"
)

var allowedDdsqlTabularQueryFetchRequestTypeEnumValues = []DdsqlTabularQueryFetchRequestType{
	DDSQLTABULARQUERYFETCHREQUESTTYPE_DDSQL_QUERY_FETCH_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DdsqlTabularQueryFetchRequestType) GetAllowedValues() []DdsqlTabularQueryFetchRequestType {
	return allowedDdsqlTabularQueryFetchRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DdsqlTabularQueryFetchRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DdsqlTabularQueryFetchRequestType(value)
	return nil
}

// NewDdsqlTabularQueryFetchRequestTypeFromValue returns a pointer to a valid DdsqlTabularQueryFetchRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDdsqlTabularQueryFetchRequestTypeFromValue(v string) (*DdsqlTabularQueryFetchRequestType, error) {
	ev := DdsqlTabularQueryFetchRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DdsqlTabularQueryFetchRequestType: valid values are %v", v, allowedDdsqlTabularQueryFetchRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DdsqlTabularQueryFetchRequestType) IsValid() bool {
	for _, existing := range allowedDdsqlTabularQueryFetchRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DdsqlTabularQueryFetchRequestType value.
func (v DdsqlTabularQueryFetchRequestType) Ptr() *DdsqlTabularQueryFetchRequestType {
	return &v
}
