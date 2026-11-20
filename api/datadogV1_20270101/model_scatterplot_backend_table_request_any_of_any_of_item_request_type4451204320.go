// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451204320 Test-only generated schema for the dashboard versioning proof of concept.
type ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451204320 string

// List of ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451204320.
const (
	SCATTERPLOTBACKENDTABLEREQUESTANYOFANYOFITEMREQUESTTYPE4451204320_LOCAL_DATASET ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451204320 = "local_dataset"
)

var allowedScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451204320EnumValues = []ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451204320{
	SCATTERPLOTBACKENDTABLEREQUESTANYOFANYOFITEMREQUESTTYPE4451204320_LOCAL_DATASET,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451204320) GetAllowedValues() []ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451204320 {
	return allowedScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451204320EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451204320) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451204320(value)
	return nil
}

// NewScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451204320FromValue returns a pointer to a valid ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451204320
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451204320FromValue(v string) (*ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451204320, error) {
	ev := ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451204320(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451204320: valid values are %v", v, allowedScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451204320EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451204320) IsValid() bool {
	for _, existing := range allowedScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451204320EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451204320 value.
func (v ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451204320) Ptr() *ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType4451204320 {
	return &v
}
