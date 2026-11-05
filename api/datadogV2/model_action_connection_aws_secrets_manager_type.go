// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ActionConnectionAWSSecretsManagerType AWS Secrets Manager secret backend.
type ActionConnectionAWSSecretsManagerType string

// List of ActionConnectionAWSSecretsManagerType.
const (
	ACTIONCONNECTIONAWSSECRETSMANAGERTYPE_AWS_SECRETS_MANAGER ActionConnectionAWSSecretsManagerType = "aws_secrets_manager"
)

var allowedActionConnectionAWSSecretsManagerTypeEnumValues = []ActionConnectionAWSSecretsManagerType{
	ACTIONCONNECTIONAWSSECRETSMANAGERTYPE_AWS_SECRETS_MANAGER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ActionConnectionAWSSecretsManagerType) GetAllowedValues() []ActionConnectionAWSSecretsManagerType {
	return allowedActionConnectionAWSSecretsManagerTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ActionConnectionAWSSecretsManagerType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ActionConnectionAWSSecretsManagerType(value)
	return nil
}

// NewActionConnectionAWSSecretsManagerTypeFromValue returns a pointer to a valid ActionConnectionAWSSecretsManagerType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewActionConnectionAWSSecretsManagerTypeFromValue(v string) (*ActionConnectionAWSSecretsManagerType, error) {
	ev := ActionConnectionAWSSecretsManagerType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ActionConnectionAWSSecretsManagerType: valid values are %v", v, allowedActionConnectionAWSSecretsManagerTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ActionConnectionAWSSecretsManagerType) IsValid() bool {
	for _, existing := range allowedActionConnectionAWSSecretsManagerTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ActionConnectionAWSSecretsManagerType value.
func (v ActionConnectionAWSSecretsManagerType) Ptr() *ActionConnectionAWSSecretsManagerType {
	return &v
}
