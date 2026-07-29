// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ElasticCloudCcmTokenAuthType Authentication method discriminator.
type ElasticCloudCcmTokenAuthType string

// List of ElasticCloudCcmTokenAuthType.
const (
	ELASTICCLOUDCCMTOKENAUTHTYPE_BEARER_TOKEN ElasticCloudCcmTokenAuthType = "bearer_token"
)

var allowedElasticCloudCcmTokenAuthTypeEnumValues = []ElasticCloudCcmTokenAuthType{
	ELASTICCLOUDCCMTOKENAUTHTYPE_BEARER_TOKEN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ElasticCloudCcmTokenAuthType) GetAllowedValues() []ElasticCloudCcmTokenAuthType {
	return allowedElasticCloudCcmTokenAuthTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ElasticCloudCcmTokenAuthType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ElasticCloudCcmTokenAuthType(value)
	return nil
}

// NewElasticCloudCcmTokenAuthTypeFromValue returns a pointer to a valid ElasticCloudCcmTokenAuthType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewElasticCloudCcmTokenAuthTypeFromValue(v string) (*ElasticCloudCcmTokenAuthType, error) {
	ev := ElasticCloudCcmTokenAuthType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ElasticCloudCcmTokenAuthType: valid values are %v", v, allowedElasticCloudCcmTokenAuthTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ElasticCloudCcmTokenAuthType) IsValid() bool {
	for _, existing := range allowedElasticCloudCcmTokenAuthTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ElasticCloudCcmTokenAuthType value.
func (v ElasticCloudCcmTokenAuthType) Ptr() *ElasticCloudCcmTokenAuthType {
	return &v
}
