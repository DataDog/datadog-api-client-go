// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ElasticCloudBasicAuthType Authentication method discriminator.
type ElasticCloudBasicAuthType string

// List of ElasticCloudBasicAuthType.
const (
	ELASTICCLOUDBASICAUTHTYPE_BASIC ElasticCloudBasicAuthType = "basic"
)

var allowedElasticCloudBasicAuthTypeEnumValues = []ElasticCloudBasicAuthType{
	ELASTICCLOUDBASICAUTHTYPE_BASIC,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ElasticCloudBasicAuthType) GetAllowedValues() []ElasticCloudBasicAuthType {
	return allowedElasticCloudBasicAuthTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ElasticCloudBasicAuthType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ElasticCloudBasicAuthType(value)
	return nil
}

// NewElasticCloudBasicAuthTypeFromValue returns a pointer to a valid ElasticCloudBasicAuthType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewElasticCloudBasicAuthTypeFromValue(v string) (*ElasticCloudBasicAuthType, error) {
	ev := ElasticCloudBasicAuthType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ElasticCloudBasicAuthType: valid values are %v", v, allowedElasticCloudBasicAuthTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ElasticCloudBasicAuthType) IsValid() bool {
	for _, existing := range allowedElasticCloudBasicAuthTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ElasticCloudBasicAuthType value.
func (v ElasticCloudBasicAuthType) Ptr() *ElasticCloudBasicAuthType {
	return &v
}
