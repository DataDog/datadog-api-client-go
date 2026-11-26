// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451373984 Test-only generated schema for the dashboard versioning proof of concept.
type ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451373984 string

// List of ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451373984.
const (
	SCATTERPLOTBACKENDTABLEREQUESTANYOFANYOFITEMRESPONSEFORMAT4451373984_TABULAR ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451373984 = "tabular"
)

var allowedScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451373984EnumValues = []ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451373984{
	SCATTERPLOTBACKENDTABLEREQUESTANYOFANYOFITEMRESPONSEFORMAT4451373984_TABULAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451373984) GetAllowedValues() []ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451373984 {
	return allowedScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451373984EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451373984) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451373984(value)
	return nil
}

// NewScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451373984FromValue returns a pointer to a valid ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451373984
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451373984FromValue(v string) (*ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451373984, error) {
	ev := ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451373984(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451373984: valid values are %v", v, allowedScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451373984EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451373984) IsValid() bool {
	for _, existing := range allowedScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451373984EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451373984 value.
func (v ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451373984) Ptr() *ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat4451373984 {
	return &v
}
