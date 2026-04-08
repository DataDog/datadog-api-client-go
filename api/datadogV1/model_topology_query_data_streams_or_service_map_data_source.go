// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TopologyQueryDataStreamsOrServiceMapDataSource Name of the data source
type TopologyQueryDataStreamsOrServiceMapDataSource string

// List of TopologyQueryDataStreamsOrServiceMapDataSource.
const (
	TOPOLOGYQUERYDATASTREAMSORSERVICEMAPDATASOURCE_DATA_STREAMS TopologyQueryDataStreamsOrServiceMapDataSource = "data_streams"
	TOPOLOGYQUERYDATASTREAMSORSERVICEMAPDATASOURCE_SERVICE_MAP  TopologyQueryDataStreamsOrServiceMapDataSource = "service_map"
)

var allowedTopologyQueryDataStreamsOrServiceMapDataSourceEnumValues = []TopologyQueryDataStreamsOrServiceMapDataSource{
	TOPOLOGYQUERYDATASTREAMSORSERVICEMAPDATASOURCE_DATA_STREAMS,
	TOPOLOGYQUERYDATASTREAMSORSERVICEMAPDATASOURCE_SERVICE_MAP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TopologyQueryDataStreamsOrServiceMapDataSource) GetAllowedValues() []TopologyQueryDataStreamsOrServiceMapDataSource {
	return allowedTopologyQueryDataStreamsOrServiceMapDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TopologyQueryDataStreamsOrServiceMapDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TopologyQueryDataStreamsOrServiceMapDataSource(value)
	return nil
}

// NewTopologyQueryDataStreamsOrServiceMapDataSourceFromValue returns a pointer to a valid TopologyQueryDataStreamsOrServiceMapDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTopologyQueryDataStreamsOrServiceMapDataSourceFromValue(v string) (*TopologyQueryDataStreamsOrServiceMapDataSource, error) {
	ev := TopologyQueryDataStreamsOrServiceMapDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TopologyQueryDataStreamsOrServiceMapDataSource: valid values are %v", v, allowedTopologyQueryDataStreamsOrServiceMapDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TopologyQueryDataStreamsOrServiceMapDataSource) IsValid() bool {
	for _, existing := range allowedTopologyQueryDataStreamsOrServiceMapDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TopologyQueryDataStreamsOrServiceMapDataSource value.
func (v TopologyQueryDataStreamsOrServiceMapDataSource) Ptr() *TopologyQueryDataStreamsOrServiceMapDataSource {
	return &v
}
