// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType Test-only generated schema for the dashboard versioning proof of concept.
type ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType string

// List of ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType.
const (
	SCATTERPLOTBACKENDTABLEREQUESTANYOFANYOFITEMREQUESTTYPE_DDSQL ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType = "ddsql"
)

var allowedScatterplotBackendTableRequestAnyOfAnyOfItemRequestTypeEnumValues = []ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType{
	SCATTERPLOTBACKENDTABLEREQUESTANYOFANYOFITEMREQUESTTYPE_DDSQL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType) GetAllowedValues() []ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType {
	return allowedScatterplotBackendTableRequestAnyOfAnyOfItemRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType(value)
	return nil
}

// NewScatterplotBackendTableRequestAnyOfAnyOfItemRequestTypeFromValue returns a pointer to a valid ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScatterplotBackendTableRequestAnyOfAnyOfItemRequestTypeFromValue(v string) (*ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType, error) {
	ev := ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType: valid values are %v", v, allowedScatterplotBackendTableRequestAnyOfAnyOfItemRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType) IsValid() bool {
	for _, existing := range allowedScatterplotBackendTableRequestAnyOfAnyOfItemRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType value.
func (v ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType) Ptr() *ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType {
	return &v
}
