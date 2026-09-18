// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TerraformBackendKind Backend type to synchronize.
type TerraformBackendKind string

// List of TerraformBackendKind.
const (
	TERRAFORMBACKENDKIND_TERRAFORM TerraformBackendKind = "terraform"
)

var allowedTerraformBackendKindEnumValues = []TerraformBackendKind{
	TERRAFORMBACKENDKIND_TERRAFORM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TerraformBackendKind) GetAllowedValues() []TerraformBackendKind {
	return allowedTerraformBackendKindEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TerraformBackendKind) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TerraformBackendKind(value)
	return nil
}

// NewTerraformBackendKindFromValue returns a pointer to a valid TerraformBackendKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTerraformBackendKindFromValue(v string) (*TerraformBackendKind, error) {
	ev := TerraformBackendKind(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TerraformBackendKind: valid values are %v", v, allowedTerraformBackendKindEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TerraformBackendKind) IsValid() bool {
	for _, existing := range allowedTerraformBackendKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TerraformBackendKind value.
func (v TerraformBackendKind) Ptr() *TerraformBackendKind {
	return &v
}
