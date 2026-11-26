// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AlertGraphBackendDefinitionType Test-only generated schema for the dashboard versioning proof of concept.
type AlertGraphBackendDefinitionType string

// List of AlertGraphBackendDefinitionType.
const (
	ALERTGRAPHBACKENDDEFINITIONTYPE_ALERT_GRAPH AlertGraphBackendDefinitionType = "alert_graph"
)

var allowedAlertGraphBackendDefinitionTypeEnumValues = []AlertGraphBackendDefinitionType{
	ALERTGRAPHBACKENDDEFINITIONTYPE_ALERT_GRAPH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AlertGraphBackendDefinitionType) GetAllowedValues() []AlertGraphBackendDefinitionType {
	return allowedAlertGraphBackendDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AlertGraphBackendDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AlertGraphBackendDefinitionType(value)
	return nil
}

// NewAlertGraphBackendDefinitionTypeFromValue returns a pointer to a valid AlertGraphBackendDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAlertGraphBackendDefinitionTypeFromValue(v string) (*AlertGraphBackendDefinitionType, error) {
	ev := AlertGraphBackendDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AlertGraphBackendDefinitionType: valid values are %v", v, allowedAlertGraphBackendDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AlertGraphBackendDefinitionType) IsValid() bool {
	for _, existing := range allowedAlertGraphBackendDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AlertGraphBackendDefinitionType value.
func (v AlertGraphBackendDefinitionType) Ptr() *AlertGraphBackendDefinitionType {
	return &v
}
