// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FunnelRequestDataType
type FunnelRequestDataType string

// List of FunnelRequestDataType.
const (
	FUNNELREQUESTDATATYPE_FUNNEL_REQUEST FunnelRequestDataType = "funnel_request"
)

var allowedFunnelRequestDataTypeEnumValues = []FunnelRequestDataType{
	FUNNELREQUESTDATATYPE_FUNNEL_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FunnelRequestDataType) GetAllowedValues() []FunnelRequestDataType {
	return allowedFunnelRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FunnelRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FunnelRequestDataType(value)
	return nil
}

// NewFunnelRequestDataTypeFromValue returns a pointer to a valid FunnelRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFunnelRequestDataTypeFromValue(v string) (*FunnelRequestDataType, error) {
	ev := FunnelRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FunnelRequestDataType: valid values are %v", v, allowedFunnelRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FunnelRequestDataType) IsValid() bool {
	for _, existing := range allowedFunnelRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FunnelRequestDataType value.
func (v FunnelRequestDataType) Ptr() *FunnelRequestDataType {
	return &v
}
