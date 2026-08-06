// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ElasticCloudInterfaceId Supported Elastic Cloud interface (source-type) ids.
type ElasticCloudInterfaceId string

// List of ElasticCloudInterfaceId.
const (
	ELASTICCLOUDINTERFACEID_ELASTIC_CLOUD     ElasticCloudInterfaceId = "elastic-cloud"
	ELASTICCLOUDINTERFACEID_ELASTIC_CLOUD_CCM ElasticCloudInterfaceId = "elastic-cloud-ccm"
)

var allowedElasticCloudInterfaceIdEnumValues = []ElasticCloudInterfaceId{
	ELASTICCLOUDINTERFACEID_ELASTIC_CLOUD,
	ELASTICCLOUDINTERFACEID_ELASTIC_CLOUD_CCM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ElasticCloudInterfaceId) GetAllowedValues() []ElasticCloudInterfaceId {
	return allowedElasticCloudInterfaceIdEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ElasticCloudInterfaceId) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ElasticCloudInterfaceId(value)
	return nil
}

// NewElasticCloudInterfaceIdFromValue returns a pointer to a valid ElasticCloudInterfaceId
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewElasticCloudInterfaceIdFromValue(v string) (*ElasticCloudInterfaceId, error) {
	ev := ElasticCloudInterfaceId(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ElasticCloudInterfaceId: valid values are %v", v, allowedElasticCloudInterfaceIdEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ElasticCloudInterfaceId) IsValid() bool {
	for _, existing := range allowedElasticCloudInterfaceIdEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ElasticCloudInterfaceId value.
func (v ElasticCloudInterfaceId) Ptr() *ElasticCloudInterfaceId {
	return &v
}
