// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomDestinationForwardDestinationGoogleSecurityOperationsType Type of the Google Security Operations destination.
type CustomDestinationForwardDestinationGoogleSecurityOperationsType string

// List of CustomDestinationForwardDestinationGoogleSecurityOperationsType.
const (
	CUSTOMDESTINATIONFORWARDDESTINATIONGOOGLESECURITYOPERATIONSTYPE_GOOGLE_SECURITY_OPERATIONS CustomDestinationForwardDestinationGoogleSecurityOperationsType = "google_security_operations"
)

var allowedCustomDestinationForwardDestinationGoogleSecurityOperationsTypeEnumValues = []CustomDestinationForwardDestinationGoogleSecurityOperationsType{
	CUSTOMDESTINATIONFORWARDDESTINATIONGOOGLESECURITYOPERATIONSTYPE_GOOGLE_SECURITY_OPERATIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CustomDestinationForwardDestinationGoogleSecurityOperationsType) GetAllowedValues() []CustomDestinationForwardDestinationGoogleSecurityOperationsType {
	return allowedCustomDestinationForwardDestinationGoogleSecurityOperationsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CustomDestinationForwardDestinationGoogleSecurityOperationsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CustomDestinationForwardDestinationGoogleSecurityOperationsType(value)
	return nil
}

// NewCustomDestinationForwardDestinationGoogleSecurityOperationsTypeFromValue returns a pointer to a valid CustomDestinationForwardDestinationGoogleSecurityOperationsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCustomDestinationForwardDestinationGoogleSecurityOperationsTypeFromValue(v string) (*CustomDestinationForwardDestinationGoogleSecurityOperationsType, error) {
	ev := CustomDestinationForwardDestinationGoogleSecurityOperationsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CustomDestinationForwardDestinationGoogleSecurityOperationsType: valid values are %v", v, allowedCustomDestinationForwardDestinationGoogleSecurityOperationsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CustomDestinationForwardDestinationGoogleSecurityOperationsType) IsValid() bool {
	for _, existing := range allowedCustomDestinationForwardDestinationGoogleSecurityOperationsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomDestinationForwardDestinationGoogleSecurityOperationsType value.
func (v CustomDestinationForwardDestinationGoogleSecurityOperationsType) Ptr() *CustomDestinationForwardDestinationGoogleSecurityOperationsType {
	return &v
}
