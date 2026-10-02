// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsPublicProtocolResponseDataAttributesStatus Publication status of the protocol.
type ExperimentsPublicProtocolResponseDataAttributesStatus string

// List of ExperimentsPublicProtocolResponseDataAttributesStatus.
const (
	EXPERIMENTSPUBLICPROTOCOLRESPONSEDATAATTRIBUTESSTATUS_DRAFT     ExperimentsPublicProtocolResponseDataAttributesStatus = "DRAFT"
	EXPERIMENTSPUBLICPROTOCOLRESPONSEDATAATTRIBUTESSTATUS_PUBLISHED ExperimentsPublicProtocolResponseDataAttributesStatus = "PUBLISHED"
	EXPERIMENTSPUBLICPROTOCOLRESPONSEDATAATTRIBUTESSTATUS_ARCHIVED  ExperimentsPublicProtocolResponseDataAttributesStatus = "ARCHIVED"
)

var allowedExperimentsPublicProtocolResponseDataAttributesStatusEnumValues = []ExperimentsPublicProtocolResponseDataAttributesStatus{
	EXPERIMENTSPUBLICPROTOCOLRESPONSEDATAATTRIBUTESSTATUS_DRAFT,
	EXPERIMENTSPUBLICPROTOCOLRESPONSEDATAATTRIBUTESSTATUS_PUBLISHED,
	EXPERIMENTSPUBLICPROTOCOLRESPONSEDATAATTRIBUTESSTATUS_ARCHIVED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsPublicProtocolResponseDataAttributesStatus) GetAllowedValues() []ExperimentsPublicProtocolResponseDataAttributesStatus {
	return allowedExperimentsPublicProtocolResponseDataAttributesStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsPublicProtocolResponseDataAttributesStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsPublicProtocolResponseDataAttributesStatus(value)
	return nil
}

// NewExperimentsPublicProtocolResponseDataAttributesStatusFromValue returns a pointer to a valid ExperimentsPublicProtocolResponseDataAttributesStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsPublicProtocolResponseDataAttributesStatusFromValue(v string) (*ExperimentsPublicProtocolResponseDataAttributesStatus, error) {
	ev := ExperimentsPublicProtocolResponseDataAttributesStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsPublicProtocolResponseDataAttributesStatus: valid values are %v", v, allowedExperimentsPublicProtocolResponseDataAttributesStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsPublicProtocolResponseDataAttributesStatus) IsValid() bool {
	for _, existing := range allowedExperimentsPublicProtocolResponseDataAttributesStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsPublicProtocolResponseDataAttributesStatus value.
func (v ExperimentsPublicProtocolResponseDataAttributesStatus) Ptr() *ExperimentsPublicProtocolResponseDataAttributesStatus {
	return &v
}
