// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ManageStatusBackendDefinitionLastTriggeredFormat Test-only generated schema for the dashboard versioning proof of concept.
type ManageStatusBackendDefinitionLastTriggeredFormat string

// List of ManageStatusBackendDefinitionLastTriggeredFormat.
const (
	MANAGESTATUSBACKENDDEFINITIONLASTTRIGGEREDFORMAT_BOTH     ManageStatusBackendDefinitionLastTriggeredFormat = "both"
	MANAGESTATUSBACKENDDEFINITIONLASTTRIGGEREDFORMAT_RELATIVE ManageStatusBackendDefinitionLastTriggeredFormat = "relative"
	MANAGESTATUSBACKENDDEFINITIONLASTTRIGGEREDFORMAT_ABSOLUTE ManageStatusBackendDefinitionLastTriggeredFormat = "absolute"
)

var allowedManageStatusBackendDefinitionLastTriggeredFormatEnumValues = []ManageStatusBackendDefinitionLastTriggeredFormat{
	MANAGESTATUSBACKENDDEFINITIONLASTTRIGGEREDFORMAT_BOTH,
	MANAGESTATUSBACKENDDEFINITIONLASTTRIGGEREDFORMAT_RELATIVE,
	MANAGESTATUSBACKENDDEFINITIONLASTTRIGGEREDFORMAT_ABSOLUTE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ManageStatusBackendDefinitionLastTriggeredFormat) GetAllowedValues() []ManageStatusBackendDefinitionLastTriggeredFormat {
	return allowedManageStatusBackendDefinitionLastTriggeredFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ManageStatusBackendDefinitionLastTriggeredFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ManageStatusBackendDefinitionLastTriggeredFormat(value)
	return nil
}

// NewManageStatusBackendDefinitionLastTriggeredFormatFromValue returns a pointer to a valid ManageStatusBackendDefinitionLastTriggeredFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewManageStatusBackendDefinitionLastTriggeredFormatFromValue(v string) (*ManageStatusBackendDefinitionLastTriggeredFormat, error) {
	ev := ManageStatusBackendDefinitionLastTriggeredFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ManageStatusBackendDefinitionLastTriggeredFormat: valid values are %v", v, allowedManageStatusBackendDefinitionLastTriggeredFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ManageStatusBackendDefinitionLastTriggeredFormat) IsValid() bool {
	for _, existing := range allowedManageStatusBackendDefinitionLastTriggeredFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ManageStatusBackendDefinitionLastTriggeredFormat value.
func (v ManageStatusBackendDefinitionLastTriggeredFormat) Ptr() *ManageStatusBackendDefinitionLastTriggeredFormat {
	return &v
}
