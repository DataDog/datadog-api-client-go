// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AlertGraphVizType Test-only generated schema for the dashboard versioning proof of concept.
type AlertGraphVizType string

// List of AlertGraphVizType.
const (
	ALERTGRAPHVIZTYPE_TIMESERIES AlertGraphVizType = "timeseries"
	ALERTGRAPHVIZTYPE_TOPLIST    AlertGraphVizType = "toplist"
)

var allowedAlertGraphVizTypeEnumValues = []AlertGraphVizType{
	ALERTGRAPHVIZTYPE_TIMESERIES,
	ALERTGRAPHVIZTYPE_TOPLIST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AlertGraphVizType) GetAllowedValues() []AlertGraphVizType {
	return allowedAlertGraphVizTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AlertGraphVizType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AlertGraphVizType(value)
	return nil
}

// NewAlertGraphVizTypeFromValue returns a pointer to a valid AlertGraphVizType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAlertGraphVizTypeFromValue(v string) (*AlertGraphVizType, error) {
	ev := AlertGraphVizType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AlertGraphVizType: valid values are %v", v, allowedAlertGraphVizTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AlertGraphVizType) IsValid() bool {
	for _, existing := range allowedAlertGraphVizTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AlertGraphVizType value.
func (v AlertGraphVizType) Ptr() *AlertGraphVizType {
	return &v
}
