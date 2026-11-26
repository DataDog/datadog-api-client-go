// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AnomalyDetectionConfigDetectionSensitivity Test-only generated schema for the dashboard versioning proof of concept.
type AnomalyDetectionConfigDetectionSensitivity string

// List of AnomalyDetectionConfigDetectionSensitivity.
const (
	ANOMALYDETECTIONCONFIGDETECTIONSENSITIVITY_NEVER_DETECT AnomalyDetectionConfigDetectionSensitivity = "never_detect"
)

var allowedAnomalyDetectionConfigDetectionSensitivityEnumValues = []AnomalyDetectionConfigDetectionSensitivity{
	ANOMALYDETECTIONCONFIGDETECTIONSENSITIVITY_NEVER_DETECT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AnomalyDetectionConfigDetectionSensitivity) GetAllowedValues() []AnomalyDetectionConfigDetectionSensitivity {
	return allowedAnomalyDetectionConfigDetectionSensitivityEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AnomalyDetectionConfigDetectionSensitivity) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AnomalyDetectionConfigDetectionSensitivity(value)
	return nil
}

// NewAnomalyDetectionConfigDetectionSensitivityFromValue returns a pointer to a valid AnomalyDetectionConfigDetectionSensitivity
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAnomalyDetectionConfigDetectionSensitivityFromValue(v string) (*AnomalyDetectionConfigDetectionSensitivity, error) {
	ev := AnomalyDetectionConfigDetectionSensitivity(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AnomalyDetectionConfigDetectionSensitivity: valid values are %v", v, allowedAnomalyDetectionConfigDetectionSensitivityEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AnomalyDetectionConfigDetectionSensitivity) IsValid() bool {
	for _, existing := range allowedAnomalyDetectionConfigDetectionSensitivityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AnomalyDetectionConfigDetectionSensitivity value.
func (v AnomalyDetectionConfigDetectionSensitivity) Ptr() *AnomalyDetectionConfigDetectionSensitivity {
	return &v
}
