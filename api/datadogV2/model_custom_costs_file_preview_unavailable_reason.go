// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomCostsFilePreviewUnavailableReason Reason the content of the Custom Costs file is not returned. Only set when `content` is omitted.
type CustomCostsFilePreviewUnavailableReason string

// List of CustomCostsFilePreviewUnavailableReason.
const (
	CUSTOMCOSTSFILEPREVIEWUNAVAILABLEREASON_CHARGE_PERIOD_OUTSIDE_RETENTION_WINDOW CustomCostsFilePreviewUnavailableReason = "CHARGE_PERIOD_OUTSIDE_RETENTION_WINDOW"
)

var allowedCustomCostsFilePreviewUnavailableReasonEnumValues = []CustomCostsFilePreviewUnavailableReason{
	CUSTOMCOSTSFILEPREVIEWUNAVAILABLEREASON_CHARGE_PERIOD_OUTSIDE_RETENTION_WINDOW,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CustomCostsFilePreviewUnavailableReason) GetAllowedValues() []CustomCostsFilePreviewUnavailableReason {
	return allowedCustomCostsFilePreviewUnavailableReasonEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CustomCostsFilePreviewUnavailableReason) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CustomCostsFilePreviewUnavailableReason(value)
	return nil
}

// NewCustomCostsFilePreviewUnavailableReasonFromValue returns a pointer to a valid CustomCostsFilePreviewUnavailableReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCustomCostsFilePreviewUnavailableReasonFromValue(v string) (*CustomCostsFilePreviewUnavailableReason, error) {
	ev := CustomCostsFilePreviewUnavailableReason(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CustomCostsFilePreviewUnavailableReason: valid values are %v", v, allowedCustomCostsFilePreviewUnavailableReasonEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CustomCostsFilePreviewUnavailableReason) IsValid() bool {
	for _, existing := range allowedCustomCostsFilePreviewUnavailableReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomCostsFilePreviewUnavailableReason value.
func (v CustomCostsFilePreviewUnavailableReason) Ptr() *CustomCostsFilePreviewUnavailableReason {
	return &v
}
