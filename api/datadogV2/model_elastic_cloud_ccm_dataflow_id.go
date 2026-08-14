// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ElasticCloudCcmDataflowId Identifier of an Elastic Cloud CCM dataflow.
type ElasticCloudCcmDataflowId string

// List of ElasticCloudCcmDataflowId.
const (
	ELASTICCLOUDCCMDATAFLOWID_COST_DATA ElasticCloudCcmDataflowId = "elastic-cloud-cost-data"
)

var allowedElasticCloudCcmDataflowIdEnumValues = []ElasticCloudCcmDataflowId{
	ELASTICCLOUDCCMDATAFLOWID_COST_DATA,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ElasticCloudCcmDataflowId) GetAllowedValues() []ElasticCloudCcmDataflowId {
	return allowedElasticCloudCcmDataflowIdEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ElasticCloudCcmDataflowId) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ElasticCloudCcmDataflowId(value)
	return nil
}

// NewElasticCloudCcmDataflowIdFromValue returns a pointer to a valid ElasticCloudCcmDataflowId
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewElasticCloudCcmDataflowIdFromValue(v string) (*ElasticCloudCcmDataflowId, error) {
	ev := ElasticCloudCcmDataflowId(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ElasticCloudCcmDataflowId: valid values are %v", v, allowedElasticCloudCcmDataflowIdEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ElasticCloudCcmDataflowId) IsValid() bool {
	for _, existing := range allowedElasticCloudCcmDataflowIdEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ElasticCloudCcmDataflowId value.
func (v ElasticCloudCcmDataflowId) Ptr() *ElasticCloudCcmDataflowId {
	return &v
}
