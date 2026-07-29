// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ElasticCloudMonitoringInterfaceType Interface discriminator for the Elastic Cloud monitoring interface.
type ElasticCloudMonitoringInterfaceType string

// List of ElasticCloudMonitoringInterfaceType.
const (
	ELASTICCLOUDMONITORINGINTERFACETYPE_ELASTIC_CLOUD ElasticCloudMonitoringInterfaceType = "elastic-cloud"
)

var allowedElasticCloudMonitoringInterfaceTypeEnumValues = []ElasticCloudMonitoringInterfaceType{
	ELASTICCLOUDMONITORINGINTERFACETYPE_ELASTIC_CLOUD,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ElasticCloudMonitoringInterfaceType) GetAllowedValues() []ElasticCloudMonitoringInterfaceType {
	return allowedElasticCloudMonitoringInterfaceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ElasticCloudMonitoringInterfaceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ElasticCloudMonitoringInterfaceType(value)
	return nil
}

// NewElasticCloudMonitoringInterfaceTypeFromValue returns a pointer to a valid ElasticCloudMonitoringInterfaceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewElasticCloudMonitoringInterfaceTypeFromValue(v string) (*ElasticCloudMonitoringInterfaceType, error) {
	ev := ElasticCloudMonitoringInterfaceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ElasticCloudMonitoringInterfaceType: valid values are %v", v, allowedElasticCloudMonitoringInterfaceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ElasticCloudMonitoringInterfaceType) IsValid() bool {
	for _, existing := range allowedElasticCloudMonitoringInterfaceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ElasticCloudMonitoringInterfaceType value.
func (v ElasticCloudMonitoringInterfaceType) Ptr() *ElasticCloudMonitoringInterfaceType {
	return &v
}
