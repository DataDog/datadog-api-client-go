// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ElasticCloudCcmInterfaceType Interface discriminator for the Elastic Cloud CCM interface.
type ElasticCloudCcmInterfaceType string

// List of ElasticCloudCcmInterfaceType.
const (
	ELASTICCLOUDCCMINTERFACETYPE_ELASTIC_CLOUD_CCM ElasticCloudCcmInterfaceType = "elastic-cloud-ccm"
)

var allowedElasticCloudCcmInterfaceTypeEnumValues = []ElasticCloudCcmInterfaceType{
	ELASTICCLOUDCCMINTERFACETYPE_ELASTIC_CLOUD_CCM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ElasticCloudCcmInterfaceType) GetAllowedValues() []ElasticCloudCcmInterfaceType {
	return allowedElasticCloudCcmInterfaceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ElasticCloudCcmInterfaceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ElasticCloudCcmInterfaceType(value)
	return nil
}

// NewElasticCloudCcmInterfaceTypeFromValue returns a pointer to a valid ElasticCloudCcmInterfaceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewElasticCloudCcmInterfaceTypeFromValue(v string) (*ElasticCloudCcmInterfaceType, error) {
	ev := ElasticCloudCcmInterfaceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ElasticCloudCcmInterfaceType: valid values are %v", v, allowedElasticCloudCcmInterfaceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ElasticCloudCcmInterfaceType) IsValid() bool {
	for _, existing := range allowedElasticCloudCcmInterfaceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ElasticCloudCcmInterfaceType value.
func (v ElasticCloudCcmInterfaceType) Ptr() *ElasticCloudCcmInterfaceType {
	return &v
}
