// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringAWSIdentityCenterAWSAccountsResourceType The type of the resource. The value should always be `aws_identity_center_aws_accounts`.
type SecurityMonitoringAWSIdentityCenterAWSAccountsResourceType string

// List of SecurityMonitoringAWSIdentityCenterAWSAccountsResourceType.
const (
	SECURITYMONITORINGAWSIDENTITYCENTERAWSACCOUNTSRESOURCETYPE_AWS_IDENTITY_CENTER_AWS_ACCOUNTS SecurityMonitoringAWSIdentityCenterAWSAccountsResourceType = "aws_identity_center_aws_accounts"
)

var allowedSecurityMonitoringAWSIdentityCenterAWSAccountsResourceTypeEnumValues = []SecurityMonitoringAWSIdentityCenterAWSAccountsResourceType{
	SECURITYMONITORINGAWSIDENTITYCENTERAWSACCOUNTSRESOURCETYPE_AWS_IDENTITY_CENTER_AWS_ACCOUNTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringAWSIdentityCenterAWSAccountsResourceType) GetAllowedValues() []SecurityMonitoringAWSIdentityCenterAWSAccountsResourceType {
	return allowedSecurityMonitoringAWSIdentityCenterAWSAccountsResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringAWSIdentityCenterAWSAccountsResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringAWSIdentityCenterAWSAccountsResourceType(value)
	return nil
}

// NewSecurityMonitoringAWSIdentityCenterAWSAccountsResourceTypeFromValue returns a pointer to a valid SecurityMonitoringAWSIdentityCenterAWSAccountsResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringAWSIdentityCenterAWSAccountsResourceTypeFromValue(v string) (*SecurityMonitoringAWSIdentityCenterAWSAccountsResourceType, error) {
	ev := SecurityMonitoringAWSIdentityCenterAWSAccountsResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringAWSIdentityCenterAWSAccountsResourceType: valid values are %v", v, allowedSecurityMonitoringAWSIdentityCenterAWSAccountsResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringAWSIdentityCenterAWSAccountsResourceType) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringAWSIdentityCenterAWSAccountsResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringAWSIdentityCenterAWSAccountsResourceType value.
func (v SecurityMonitoringAWSIdentityCenterAWSAccountsResourceType) Ptr() *SecurityMonitoringAWSIdentityCenterAWSAccountsResourceType {
	return &v
}
