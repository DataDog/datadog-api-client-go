// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451375104 Test-only generated schema for the dashboard versioning proof of concept.
type ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451375104 string

// List of ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451375104.
const (
	SCATTERPLOTBACKENDTABLEREQUESTANYOFANYOFITEMREQUESTTYPE4451375104_DATA_PROJECTION ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451375104 = "data_projection"
)

var allowedScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451375104EnumValues = []ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451375104{
	SCATTERPLOTBACKENDTABLEREQUESTANYOFANYOFITEMREQUESTTYPE4451375104_DATA_PROJECTION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451375104) GetAllowedValues() []ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451375104 {
	return allowedScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451375104EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451375104) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451375104(value)
	return nil
}

// NewScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451375104FromValue returns a pointer to a valid ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451375104
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451375104FromValue(v string) (*ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451375104, error) {
	ev := ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451375104(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451375104: valid values are %v", v, allowedScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451375104EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451375104) IsValid() bool {
	for _, existing := range allowedScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451375104EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451375104 value.
func (v ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451375104) Ptr() *ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451375104 {
	return &v
}
