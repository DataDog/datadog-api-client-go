// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsTransactionCompute Test-only generated schema for the dashboard versioning proof of concept.
type LogsTransactionCompute struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Aggregation EventsAggregator `json:"aggregation"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Facet string `json:"facet"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewLogsTransactionCompute instantiates a new LogsTransactionCompute object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogsTransactionCompute(aggregation EventsAggregator, facet string) *LogsTransactionCompute {
	this := LogsTransactionCompute{}
	this.Aggregation = aggregation
	this.Facet = facet
	return &this
}

// NewLogsTransactionComputeWithDefaults instantiates a new LogsTransactionCompute object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogsTransactionComputeWithDefaults() *LogsTransactionCompute {
	this := LogsTransactionCompute{}
	return &this
}

// GetAggregation returns the Aggregation field value.
func (o *LogsTransactionCompute) GetAggregation() EventsAggregator {
	if o == nil {
		var ret EventsAggregator
		return ret
	}
	return o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value
// and a boolean to check if the value has been set.
func (o *LogsTransactionCompute) GetAggregationOk() (*EventsAggregator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aggregation, true
}

// SetAggregation sets field value.
func (o *LogsTransactionCompute) SetAggregation(v EventsAggregator) {
	o.Aggregation = v
}

// GetFacet returns the Facet field value.
func (o *LogsTransactionCompute) GetFacet() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Facet
}

// GetFacetOk returns a tuple with the Facet field value
// and a boolean to check if the value has been set.
func (o *LogsTransactionCompute) GetFacetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Facet, true
}

// SetFacet sets field value.
func (o *LogsTransactionCompute) SetFacet(v string) {
	o.Facet = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogsTransactionCompute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["aggregation"] = o.Aggregation
	toSerialize["facet"] = o.Facet
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LogsTransactionCompute) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregation *EventsAggregator `json:"aggregation"`
		Facet       *string           `json:"facet"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Aggregation == nil {
		return fmt.Errorf("required field aggregation missing")
	}
	if all.Facet == nil {
		return fmt.Errorf("required field facet missing")
	}
	o.Aggregation = *all.Aggregation
	o.Facet = *all.Facet

	return nil
}
