// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CloudCostAccountType Type of a cloud cost account.
type CloudCostAccountType string

// List of CloudCostAccountType.
const (
	CLOUDCOSTACCOUNTTYPE_CLOUD_ACCOUNT CloudCostAccountType = "cloud_account"
)

var allowedCloudCostAccountTypeEnumValues = []CloudCostAccountType{
	CLOUDCOSTACCOUNTTYPE_CLOUD_ACCOUNT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CloudCostAccountType) GetAllowedValues() []CloudCostAccountType {
	return allowedCloudCostAccountTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CloudCostAccountType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CloudCostAccountType(value)
	return nil
}

// NewCloudCostAccountTypeFromValue returns a pointer to a valid CloudCostAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCloudCostAccountTypeFromValue(v string) (*CloudCostAccountType, error) {
	ev := CloudCostAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CloudCostAccountType: valid values are %v", v, allowedCloudCostAccountTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CloudCostAccountType) IsValid() bool {
	for _, existing := range allowedCloudCostAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CloudCostAccountType value.
func (v CloudCostAccountType) Ptr() *CloudCostAccountType {
	return &v
}
