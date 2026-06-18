// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// VercelLogSource Source of logs that Vercel forwards to Datadog.
type VercelLogSource string

// List of VercelLogSource.
const (
	VERCELLOGSOURCE_STATIC   VercelLogSource = "static"
	VERCELLOGSOURCE_LAMBDA   VercelLogSource = "lambda"
	VERCELLOGSOURCE_EDGE     VercelLogSource = "edge"
	VERCELLOGSOURCE_BUILD    VercelLogSource = "build"
	VERCELLOGSOURCE_EXTERNAL VercelLogSource = "external"
	VERCELLOGSOURCE_FIREWALL VercelLogSource = "firewall"
)

var allowedVercelLogSourceEnumValues = []VercelLogSource{
	VERCELLOGSOURCE_STATIC,
	VERCELLOGSOURCE_LAMBDA,
	VERCELLOGSOURCE_EDGE,
	VERCELLOGSOURCE_BUILD,
	VERCELLOGSOURCE_EXTERNAL,
	VERCELLOGSOURCE_FIREWALL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *VercelLogSource) GetAllowedValues() []VercelLogSource {
	return allowedVercelLogSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *VercelLogSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = VercelLogSource(value)
	return nil
}

// NewVercelLogSourceFromValue returns a pointer to a valid VercelLogSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewVercelLogSourceFromValue(v string) (*VercelLogSource, error) {
	ev := VercelLogSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for VercelLogSource: valid values are %v", v, allowedVercelLogSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v VercelLogSource) IsValid() bool {
	for _, existing := range allowedVercelLogSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VercelLogSource value.
func (v VercelLogSource) Ptr() *VercelLogSource {
	return &v
}
