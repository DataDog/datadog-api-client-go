// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ElasticCloudIntegrationType Integration discriminator for Elastic Cloud.
type ElasticCloudIntegrationType string

// List of ElasticCloudIntegrationType.
const (
	ELASTICCLOUDINTEGRATIONTYPE_ELASTIC_CLOUD ElasticCloudIntegrationType = "elastic-cloud"
)

var allowedElasticCloudIntegrationTypeEnumValues = []ElasticCloudIntegrationType{
	ELASTICCLOUDINTEGRATIONTYPE_ELASTIC_CLOUD,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ElasticCloudIntegrationType) GetAllowedValues() []ElasticCloudIntegrationType {
	return allowedElasticCloudIntegrationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ElasticCloudIntegrationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ElasticCloudIntegrationType(value)
	return nil
}

// NewElasticCloudIntegrationTypeFromValue returns a pointer to a valid ElasticCloudIntegrationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewElasticCloudIntegrationTypeFromValue(v string) (*ElasticCloudIntegrationType, error) {
	ev := ElasticCloudIntegrationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ElasticCloudIntegrationType: valid values are %v", v, allowedElasticCloudIntegrationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ElasticCloudIntegrationType) IsValid() bool {
	for _, existing := range allowedElasticCloudIntegrationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ElasticCloudIntegrationType value.
func (v ElasticCloudIntegrationType) Ptr() *ElasticCloudIntegrationType {
	return &v
}
