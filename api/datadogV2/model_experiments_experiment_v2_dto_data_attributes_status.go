// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsExperimentV2DTODataAttributesStatus Current stage in the experiment lifecycle.
type ExperimentsExperimentV2DTODataAttributesStatus string

// List of ExperimentsExperimentV2DTODataAttributesStatus.
const (
	EXPERIMENTSEXPERIMENTV2DTODATAATTRIBUTESSTATUS_DRAFT              ExperimentsExperimentV2DTODataAttributesStatus = "DRAFT"
	EXPERIMENTSEXPERIMENTV2DTODATAATTRIBUTESSTATUS_SCHEDULED          ExperimentsExperimentV2DTODataAttributesStatus = "SCHEDULED"
	EXPERIMENTSEXPERIMENTV2DTODATAATTRIBUTESSTATUS_IN_PROGRESS        ExperimentsExperimentV2DTODataAttributesStatus = "IN_PROGRESS"
	EXPERIMENTSEXPERIMENTV2DTODATAATTRIBUTESSTATUS_READY_FOR_DECISION ExperimentsExperimentV2DTODataAttributesStatus = "READY_FOR_DECISION"
	EXPERIMENTSEXPERIMENTV2DTODATAATTRIBUTESSTATUS_DECISION_MADE      ExperimentsExperimentV2DTODataAttributesStatus = "DECISION_MADE"
	EXPERIMENTSEXPERIMENTV2DTODATAATTRIBUTESSTATUS_CANCELLED          ExperimentsExperimentV2DTODataAttributesStatus = "CANCELLED"
	EXPERIMENTSEXPERIMENTV2DTODATAATTRIBUTESSTATUS_UNKNOWN            ExperimentsExperimentV2DTODataAttributesStatus = "UNKNOWN"
)

var allowedExperimentsExperimentV2DTODataAttributesStatusEnumValues = []ExperimentsExperimentV2DTODataAttributesStatus{
	EXPERIMENTSEXPERIMENTV2DTODATAATTRIBUTESSTATUS_DRAFT,
	EXPERIMENTSEXPERIMENTV2DTODATAATTRIBUTESSTATUS_SCHEDULED,
	EXPERIMENTSEXPERIMENTV2DTODATAATTRIBUTESSTATUS_IN_PROGRESS,
	EXPERIMENTSEXPERIMENTV2DTODATAATTRIBUTESSTATUS_READY_FOR_DECISION,
	EXPERIMENTSEXPERIMENTV2DTODATAATTRIBUTESSTATUS_DECISION_MADE,
	EXPERIMENTSEXPERIMENTV2DTODATAATTRIBUTESSTATUS_CANCELLED,
	EXPERIMENTSEXPERIMENTV2DTODATAATTRIBUTESSTATUS_UNKNOWN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsExperimentV2DTODataAttributesStatus) GetAllowedValues() []ExperimentsExperimentV2DTODataAttributesStatus {
	return allowedExperimentsExperimentV2DTODataAttributesStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsExperimentV2DTODataAttributesStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsExperimentV2DTODataAttributesStatus(value)
	return nil
}

// NewExperimentsExperimentV2DTODataAttributesStatusFromValue returns a pointer to a valid ExperimentsExperimentV2DTODataAttributesStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsExperimentV2DTODataAttributesStatusFromValue(v string) (*ExperimentsExperimentV2DTODataAttributesStatus, error) {
	ev := ExperimentsExperimentV2DTODataAttributesStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsExperimentV2DTODataAttributesStatus: valid values are %v", v, allowedExperimentsExperimentV2DTODataAttributesStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsExperimentV2DTODataAttributesStatus) IsValid() bool {
	for _, existing := range allowedExperimentsExperimentV2DTODataAttributesStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsExperimentV2DTODataAttributesStatus value.
func (v ExperimentsExperimentV2DTODataAttributesStatus) Ptr() *ExperimentsExperimentV2DTODataAttributesStatus {
	return &v
}
