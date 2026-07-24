// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TopologyQueryDataStreamsDataSource Name of the data source.
type TopologyQueryDataStreamsDataSource string

// List of TopologyQueryDataStreamsDataSource.
const (
	TOPOLOGYQUERYDATASTREAMSDATASOURCE_DATA_STREAMS TopologyQueryDataStreamsDataSource = "data_streams"
)

var allowedTopologyQueryDataStreamsDataSourceEnumValues = []TopologyQueryDataStreamsDataSource{
	TOPOLOGYQUERYDATASTREAMSDATASOURCE_DATA_STREAMS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TopologyQueryDataStreamsDataSource) GetAllowedValues() []TopologyQueryDataStreamsDataSource {
	return allowedTopologyQueryDataStreamsDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TopologyQueryDataStreamsDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TopologyQueryDataStreamsDataSource(value)
	return nil
}

// NewTopologyQueryDataStreamsDataSourceFromValue returns a pointer to a valid TopologyQueryDataStreamsDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTopologyQueryDataStreamsDataSourceFromValue(v string) (*TopologyQueryDataStreamsDataSource, error) {
	ev := TopologyQueryDataStreamsDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TopologyQueryDataStreamsDataSource: valid values are %v", v, allowedTopologyQueryDataStreamsDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TopologyQueryDataStreamsDataSource) IsValid() bool {
	for _, existing := range allowedTopologyQueryDataStreamsDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TopologyQueryDataStreamsDataSource value.
func (v TopologyQueryDataStreamsDataSource) Ptr() *TopologyQueryDataStreamsDataSource {
	return &v
}
