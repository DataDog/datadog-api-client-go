// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ElasticCloudIntegrationAccountBasicAuthType The authentication method type.
type ElasticCloudIntegrationAccountBasicAuthType string

// List of ElasticCloudIntegrationAccountBasicAuthType.
const (
	ELASTICCLOUDINTEGRATIONACCOUNTBASICAUTHTYPE_BASIC ElasticCloudIntegrationAccountBasicAuthType = "basic"
)

var allowedElasticCloudIntegrationAccountBasicAuthTypeEnumValues = []ElasticCloudIntegrationAccountBasicAuthType{
	ELASTICCLOUDINTEGRATIONACCOUNTBASICAUTHTYPE_BASIC,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ElasticCloudIntegrationAccountBasicAuthType) GetAllowedValues() []ElasticCloudIntegrationAccountBasicAuthType {
	return allowedElasticCloudIntegrationAccountBasicAuthTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ElasticCloudIntegrationAccountBasicAuthType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ElasticCloudIntegrationAccountBasicAuthType(value)
	return nil
}

// NewElasticCloudIntegrationAccountBasicAuthTypeFromValue returns a pointer to a valid ElasticCloudIntegrationAccountBasicAuthType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewElasticCloudIntegrationAccountBasicAuthTypeFromValue(v string) (*ElasticCloudIntegrationAccountBasicAuthType, error) {
	ev := ElasticCloudIntegrationAccountBasicAuthType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ElasticCloudIntegrationAccountBasicAuthType: valid values are %v", v, allowedElasticCloudIntegrationAccountBasicAuthTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ElasticCloudIntegrationAccountBasicAuthType) IsValid() bool {
	for _, existing := range allowedElasticCloudIntegrationAccountBasicAuthTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ElasticCloudIntegrationAccountBasicAuthType value.
func (v ElasticCloudIntegrationAccountBasicAuthType) Ptr() *ElasticCloudIntegrationAccountBasicAuthType {
	return &v
}
