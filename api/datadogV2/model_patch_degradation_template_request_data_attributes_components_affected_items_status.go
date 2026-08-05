// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PatchDegradationTemplateRequestDataAttributesComponentsAffectedItemsStatus The status of the component.
type PatchDegradationTemplateRequestDataAttributesComponentsAffectedItemsStatus string

// List of PatchDegradationTemplateRequestDataAttributesComponentsAffectedItemsStatus.
const (
	PATCHDEGRADATIONTEMPLATEREQUESTDATAATTRIBUTESCOMPONENTSAFFECTEDITEMSSTATUS_OPERATIONAL    PatchDegradationTemplateRequestDataAttributesComponentsAffectedItemsStatus = "operational"
	PATCHDEGRADATIONTEMPLATEREQUESTDATAATTRIBUTESCOMPONENTSAFFECTEDITEMSSTATUS_DEGRADED       PatchDegradationTemplateRequestDataAttributesComponentsAffectedItemsStatus = "degraded"
	PATCHDEGRADATIONTEMPLATEREQUESTDATAATTRIBUTESCOMPONENTSAFFECTEDITEMSSTATUS_PARTIAL_OUTAGE PatchDegradationTemplateRequestDataAttributesComponentsAffectedItemsStatus = "partial_outage"
	PATCHDEGRADATIONTEMPLATEREQUESTDATAATTRIBUTESCOMPONENTSAFFECTEDITEMSSTATUS_MAJOR_OUTAGE   PatchDegradationTemplateRequestDataAttributesComponentsAffectedItemsStatus = "major_outage"
)

var allowedPatchDegradationTemplateRequestDataAttributesComponentsAffectedItemsStatusEnumValues = []PatchDegradationTemplateRequestDataAttributesComponentsAffectedItemsStatus{
	PATCHDEGRADATIONTEMPLATEREQUESTDATAATTRIBUTESCOMPONENTSAFFECTEDITEMSSTATUS_OPERATIONAL,
	PATCHDEGRADATIONTEMPLATEREQUESTDATAATTRIBUTESCOMPONENTSAFFECTEDITEMSSTATUS_DEGRADED,
	PATCHDEGRADATIONTEMPLATEREQUESTDATAATTRIBUTESCOMPONENTSAFFECTEDITEMSSTATUS_PARTIAL_OUTAGE,
	PATCHDEGRADATIONTEMPLATEREQUESTDATAATTRIBUTESCOMPONENTSAFFECTEDITEMSSTATUS_MAJOR_OUTAGE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PatchDegradationTemplateRequestDataAttributesComponentsAffectedItemsStatus) GetAllowedValues() []PatchDegradationTemplateRequestDataAttributesComponentsAffectedItemsStatus {
	return allowedPatchDegradationTemplateRequestDataAttributesComponentsAffectedItemsStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PatchDegradationTemplateRequestDataAttributesComponentsAffectedItemsStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PatchDegradationTemplateRequestDataAttributesComponentsAffectedItemsStatus(value)
	return nil
}

// NewPatchDegradationTemplateRequestDataAttributesComponentsAffectedItemsStatusFromValue returns a pointer to a valid PatchDegradationTemplateRequestDataAttributesComponentsAffectedItemsStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPatchDegradationTemplateRequestDataAttributesComponentsAffectedItemsStatusFromValue(v string) (*PatchDegradationTemplateRequestDataAttributesComponentsAffectedItemsStatus, error) {
	ev := PatchDegradationTemplateRequestDataAttributesComponentsAffectedItemsStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PatchDegradationTemplateRequestDataAttributesComponentsAffectedItemsStatus: valid values are %v", v, allowedPatchDegradationTemplateRequestDataAttributesComponentsAffectedItemsStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PatchDegradationTemplateRequestDataAttributesComponentsAffectedItemsStatus) IsValid() bool {
	for _, existing := range allowedPatchDegradationTemplateRequestDataAttributesComponentsAffectedItemsStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PatchDegradationTemplateRequestDataAttributesComponentsAffectedItemsStatus value.
func (v PatchDegradationTemplateRequestDataAttributesComponentsAffectedItemsStatus) Ptr() *PatchDegradationTemplateRequestDataAttributesComponentsAffectedItemsStatus {
	return &v
}
