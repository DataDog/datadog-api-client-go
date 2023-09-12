// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/goccy/go-json"
)

// AzureFallbackDestinationType Type of the Azure archive destination.
type AzureFallbackDestinationType string

// List of AzureFallbackDestinationType.
const (
	AZUREFALLBACKDESTINATIONTYPE_AZURE AzureFallbackDestinationType = "azure"
)

var allowedAzureFallbackDestinationTypeEnumValues = []AzureFallbackDestinationType{
	AZUREFALLBACKDESTINATIONTYPE_AZURE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AzureFallbackDestinationType) GetAllowedValues() []AzureFallbackDestinationType {
	return allowedAzureFallbackDestinationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AzureFallbackDestinationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AzureFallbackDestinationType(value)
	return nil
}

// NewAzureFallbackDestinationTypeFromValue returns a pointer to a valid AzureFallbackDestinationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAzureFallbackDestinationTypeFromValue(v string) (*AzureFallbackDestinationType, error) {
	ev := AzureFallbackDestinationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AzureFallbackDestinationType: valid values are %v", v, allowedAzureFallbackDestinationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AzureFallbackDestinationType) IsValid() bool {
	for _, existing := range allowedAzureFallbackDestinationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AzureFallbackDestinationType value.
func (v AzureFallbackDestinationType) Ptr() *AzureFallbackDestinationType {
	return &v
}
