// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomDestinationResponseForwardDestinationGoogleSecurityOperationsType Type of the Google Security Operations destination.
type CustomDestinationResponseForwardDestinationGoogleSecurityOperationsType string

// List of CustomDestinationResponseForwardDestinationGoogleSecurityOperationsType.
const (
	CUSTOMDESTINATIONRESPONSEFORWARDDESTINATIONGOOGLESECURITYOPERATIONSTYPE_GOOGLE_SECURITY_OPERATIONS CustomDestinationResponseForwardDestinationGoogleSecurityOperationsType = "google_security_operations"
)

var allowedCustomDestinationResponseForwardDestinationGoogleSecurityOperationsTypeEnumValues = []CustomDestinationResponseForwardDestinationGoogleSecurityOperationsType{
	CUSTOMDESTINATIONRESPONSEFORWARDDESTINATIONGOOGLESECURITYOPERATIONSTYPE_GOOGLE_SECURITY_OPERATIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CustomDestinationResponseForwardDestinationGoogleSecurityOperationsType) GetAllowedValues() []CustomDestinationResponseForwardDestinationGoogleSecurityOperationsType {
	return allowedCustomDestinationResponseForwardDestinationGoogleSecurityOperationsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CustomDestinationResponseForwardDestinationGoogleSecurityOperationsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CustomDestinationResponseForwardDestinationGoogleSecurityOperationsType(value)
	return nil
}

// NewCustomDestinationResponseForwardDestinationGoogleSecurityOperationsTypeFromValue returns a pointer to a valid CustomDestinationResponseForwardDestinationGoogleSecurityOperationsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCustomDestinationResponseForwardDestinationGoogleSecurityOperationsTypeFromValue(v string) (*CustomDestinationResponseForwardDestinationGoogleSecurityOperationsType, error) {
	ev := CustomDestinationResponseForwardDestinationGoogleSecurityOperationsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CustomDestinationResponseForwardDestinationGoogleSecurityOperationsType: valid values are %v", v, allowedCustomDestinationResponseForwardDestinationGoogleSecurityOperationsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CustomDestinationResponseForwardDestinationGoogleSecurityOperationsType) IsValid() bool {
	for _, existing := range allowedCustomDestinationResponseForwardDestinationGoogleSecurityOperationsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomDestinationResponseForwardDestinationGoogleSecurityOperationsType value.
func (v CustomDestinationResponseForwardDestinationGoogleSecurityOperationsType) Ptr() *CustomDestinationResponseForwardDestinationGoogleSecurityOperationsType {
	return &v
}
