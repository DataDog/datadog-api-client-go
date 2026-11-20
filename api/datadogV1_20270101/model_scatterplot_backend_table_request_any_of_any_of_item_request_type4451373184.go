// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451373184 Test-only generated schema for the dashboard versioning proof of concept.
type ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451373184 string

// List of ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451373184.
const (
	SCATTERPLOTBACKENDTABLEREQUESTANYOFANYOFITEMREQUESTTYPE4451373184_ADVANCED_QUERY ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451373184 = "advanced_query"
)

var allowedScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451373184EnumValues = []ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451373184{
	SCATTERPLOTBACKENDTABLEREQUESTANYOFANYOFITEMREQUESTTYPE4451373184_ADVANCED_QUERY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451373184) GetAllowedValues() []ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451373184 {
	return allowedScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451373184EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451373184) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451373184(value)
	return nil
}

// NewScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451373184FromValue returns a pointer to a valid ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451373184
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451373184FromValue(v string) (*ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451373184, error) {
	ev := ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451373184(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451373184: valid values are %v", v, allowedScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451373184EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451373184) IsValid() bool {
	for _, existing := range allowedScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451373184EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451373184 value.
func (v ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451373184) Ptr() *ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451373184 {
	return &v
}
