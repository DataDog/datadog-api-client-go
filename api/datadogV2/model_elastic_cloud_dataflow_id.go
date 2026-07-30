// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ElasticCloudDataflowId Identifier of an Elastic Cloud dataflow.
type ElasticCloudDataflowId string

// List of ElasticCloudDataflowId.
const (
	ELASTICCLOUDDATAFLOWID_METRICS ElasticCloudDataflowId = "elastic-cloud-metrics"
)

var allowedElasticCloudDataflowIdEnumValues = []ElasticCloudDataflowId{
	ELASTICCLOUDDATAFLOWID_METRICS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ElasticCloudDataflowId) GetAllowedValues() []ElasticCloudDataflowId {
	return allowedElasticCloudDataflowIdEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ElasticCloudDataflowId) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ElasticCloudDataflowId(value)
	return nil
}

// NewElasticCloudDataflowIdFromValue returns a pointer to a valid ElasticCloudDataflowId
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewElasticCloudDataflowIdFromValue(v string) (*ElasticCloudDataflowId, error) {
	ev := ElasticCloudDataflowId(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ElasticCloudDataflowId: valid values are %v", v, allowedElasticCloudDataflowIdEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ElasticCloudDataflowId) IsValid() bool {
	for _, existing := range allowedElasticCloudDataflowIdEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ElasticCloudDataflowId value.
func (v ElasticCloudDataflowId) Ptr() *ElasticCloudDataflowId {
	return &v
}
