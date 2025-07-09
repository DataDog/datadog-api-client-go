// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomDestinationResponseGoogleSecurityOperationsDestinationAuthType Type of the Google Security Operations destination authentication.
type CustomDestinationResponseGoogleSecurityOperationsDestinationAuthType string

// List of CustomDestinationResponseGoogleSecurityOperationsDestinationAuthType.
const (
	CUSTOMDESTINATIONRESPONSEGOOGLESECURITYOPERATIONSDESTINATIONAUTHTYPE_GCP_PRIVATE_KEY CustomDestinationResponseGoogleSecurityOperationsDestinationAuthType = "gcp_private_key"
)

var allowedCustomDestinationResponseGoogleSecurityOperationsDestinationAuthTypeEnumValues = []CustomDestinationResponseGoogleSecurityOperationsDestinationAuthType{
	CUSTOMDESTINATIONRESPONSEGOOGLESECURITYOPERATIONSDESTINATIONAUTHTYPE_GCP_PRIVATE_KEY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CustomDestinationResponseGoogleSecurityOperationsDestinationAuthType) GetAllowedValues() []CustomDestinationResponseGoogleSecurityOperationsDestinationAuthType {
	return allowedCustomDestinationResponseGoogleSecurityOperationsDestinationAuthTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CustomDestinationResponseGoogleSecurityOperationsDestinationAuthType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CustomDestinationResponseGoogleSecurityOperationsDestinationAuthType(value)
	return nil
}

// NewCustomDestinationResponseGoogleSecurityOperationsDestinationAuthTypeFromValue returns a pointer to a valid CustomDestinationResponseGoogleSecurityOperationsDestinationAuthType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCustomDestinationResponseGoogleSecurityOperationsDestinationAuthTypeFromValue(v string) (*CustomDestinationResponseGoogleSecurityOperationsDestinationAuthType, error) {
	ev := CustomDestinationResponseGoogleSecurityOperationsDestinationAuthType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CustomDestinationResponseGoogleSecurityOperationsDestinationAuthType: valid values are %v", v, allowedCustomDestinationResponseGoogleSecurityOperationsDestinationAuthTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CustomDestinationResponseGoogleSecurityOperationsDestinationAuthType) IsValid() bool {
	for _, existing := range allowedCustomDestinationResponseGoogleSecurityOperationsDestinationAuthTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomDestinationResponseGoogleSecurityOperationsDestinationAuthType value.
func (v CustomDestinationResponseGoogleSecurityOperationsDestinationAuthType) Ptr() *CustomDestinationResponseGoogleSecurityOperationsDestinationAuthType {
	return &v
}
