// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451371744 Test-only generated schema for the dashboard versioning proof of concept.
type ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451371744 string

// List of ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451371744.
const (
	SCATTERPLOTBACKENDTABLEREQUESTANYOFANYOFITEMRESPONSEFORMAT4451371744_SCALAR ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451371744 = "scalar"
)

var allowedScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451371744EnumValues = []ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451371744{
	SCATTERPLOTBACKENDTABLEREQUESTANYOFANYOFITEMRESPONSEFORMAT4451371744_SCALAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451371744) GetAllowedValues() []ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451371744 {
	return allowedScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451371744EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451371744) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451371744(value)
	return nil
}

// NewScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451371744FromValue returns a pointer to a valid ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451371744
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451371744FromValue(v string) (*ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451371744, error) {
	ev := ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451371744(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451371744: valid values are %v", v, allowedScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451371744EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451371744) IsValid() bool {
	for _, existing := range allowedScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451371744EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451371744 value.
func (v ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451371744) Ptr() *ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451371744 {
	return &v
}
